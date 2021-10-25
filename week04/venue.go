package week04

import (
	"fmt"
	"io"
)

type Venue struct {
	Audience int
	Log      io.Writer
}

func (v *Venue) Entertain(audience int, entertainers []Entertainer) {
	for _, e := range entertainers {
		if st, ok := e.(Setuper); ok {
			st.Setup(*v)

			sm := fmt.Sprintf("%s has completed setup.\n", e.Name())
			v.Log.Write([]byte(sm))
		}

		e.Perform(*v)
		sm := fmt.Sprintf("%s has performed for %d people .\n", e.Name(), audience)
		v.Log.Write([]byte(sm))

		if st, ok := e.(Teardowner); ok {
			err := st.Teardown(*v)
			if err != nil {
				sm := fmt.Sprintf("teardown error: %s", err.Error())
				v.Log.Write([]byte(sm))
			}

			sm := fmt.Sprintf("%s has completed teardown.\n", e.Name())
			v.Log.Write([]byte(sm))
		}
	}
}
