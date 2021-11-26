package week09_test

import (
	"bytes"
	"encoding/json"
	"gopherguides-intro-to-go/week09"
)

type tSubscriber struct {
	id  string
	out *bytes.Buffer
}

func (ts tSubscriber) Identifier() string {
	return ts.id
}

func (ts tSubscriber) Receive(news week09.News) {
	bb, err := json.Marshal(news)
	if err != nil {
		return
	}

	ts.out.Write(bb)
}
