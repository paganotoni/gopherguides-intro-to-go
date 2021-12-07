package week10_test

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/paganotoni/gopherguides-intro-to-go/week10"
)

func TestMockSource(t *testing.T) {
	t.Cleanup(week10.Clean)

	sub := &tSubscriber{
		out: bytes.NewBuffer(nil),
	}

	week10.Subscribe(sub, "sports")

	ms := week10.NewMockSource(time.Millisecond*10, "sports")
	week10.AddSource(ms)

	week10.Start(context.Background())

	time.Sleep(time.Millisecond * 105)
	week10.Stop()

	if len(sub.News()) != 10 {
		t.Errorf("Expected 10 news, got %d", len(sub.news))
	}
}
