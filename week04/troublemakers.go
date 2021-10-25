package week04

import "fmt"

var (
	_ Entertainer = &Troublemakers{}
	_ Teardowner  = &Troublemakers{}
)

// Troublemakers is a type of entertainer
// which does not require a lot of work from the
// venue to start, but causes a lot of trouble
// at the end of the show and requires some work
// to teardown the event.
type Troublemakers struct{}

func (spo Troublemakers) Name() string {
	return "Troublemakers"
}

func (spo Troublemakers) Teardown(v Venue) error {
	v.Log.Write([]byte("Troublemakers: Ran Away after the concert"))
	v.Log.Write([]byte("Troublemakers: Police pursue them for the chair-trowing event"))
	v.Log.Write([]byte("Troublemakers: Detained"))
	v.Log.Write([]byte("Troublemakers: Venue has to pay fine"))

	return fmt.Errorf("could not solve Troublemakers issues, legal department still working")
}

func (spo Troublemakers) Perform(v Venue) {
	v.Log.Write([]byte("Troublemakers: Entering the stage\n"))
	v.Log.Write([]byte("Troublemakers: Performs its first song\n"))
	v.Log.Write([]byte("Troublemakers: Breaks a Camera with the guitar\n"))
	v.Log.Write([]byte("Troublemakers: Breaks a Guitar, pieces hit an attendant\n"))
	v.Log.Write([]byte("Troublemakers: Throw a chair to the crowd\n"))
	v.Log.Write([]byte("Troublemakers: Leaving the stage\n"))
}
