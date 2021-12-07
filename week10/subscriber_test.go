package week10_test

import (
	"bytes"
	"encoding/json"
	"sync"

	"github.com/paganotoni/gopherguides-intro-to-go/week10"
)

type tSubscriber struct {
	id   string
	out  *bytes.Buffer
	news []week10.News

	sync.RWMutex
}

func (ts *tSubscriber) Identifier() string {
	return ts.id
}

func (ts *tSubscriber) Receive(news week10.News) {
	ts.Lock()
	defer ts.Unlock()

	bb, err := json.Marshal(news)
	if err != nil {
		return
	}

	ts.out.Write(bb)
	ts.news = append(ts.news, news)
}

func (ts *tSubscriber) News() []week10.News {
	ts.RLock()
	defer ts.RUnlock()

	return ts.news
}
