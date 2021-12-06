package week10_test

import (
	"bytes"
	"encoding/json"

	"github.com/paganotoni/gopherguides-intro-to-go/week10"
)

type tSubscriber struct {
	id  string
	out *bytes.Buffer
}

func (ts tSubscriber) Identifier() string {
	return ts.id
}

func (ts tSubscriber) Receive(news week10.News) {
	bb, err := json.Marshal(news)
	if err != nil {
		return
	}

	ts.out.Write(bb)
}
