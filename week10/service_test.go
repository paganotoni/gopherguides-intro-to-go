package week10_test

import (
	"bytes"
	"fmt"
	"sync"
	"testing"

	"github.com/paganotoni/gopherguides-intro-to-go/week10"
)

func TestServiceSubscribeMany(t *testing.T) {
	t.Cleanup(week10.Clean)

	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wgx *sync.WaitGroup, index int) {
			sub := tSubscriber{
				id:  fmt.Sprintf("%d", index),
				out: bytes.NewBufferString(""),
			}

			week10.Subscribe(&sub, []string{"sports"})
			wgx.Done()
		}(wg, i)
	}

	wg.Wait()

	subs := len(week10.Subscribers())
	if subs != 10 {
		t.Fatalf("Expected 10 subscriber, got %d", subs)
	}
}

func TestServiceSubscribeRepeated(t *testing.T) {
	t.Cleanup(week10.Clean)

	sub := tSubscriber{
		id:  "1",
		out: bytes.NewBufferString(""),
	}

	sub2 := tSubscriber{
		id:  "1",
		out: bytes.NewBufferString(""),
	}

	err := week10.Subscribe(&sub, []string{"sports"})
	if err != nil {
		t.Fatalf("err should be nil, got %v", err)
	}

	err = week10.Subscribe(&sub2, []string{"sports"})
	if err == nil {
		t.Fatal("err should not nil")
	}
}

func TestServiceUnsubscribe(t *testing.T) {
	t.Cleanup(week10.Clean)

	for i := 0; i < 10; i++ {
		sub := &tSubscriber{
			id:  fmt.Sprintf("%d", i),
			out: bytes.NewBufferString(""),
		}

		week10.Subscribe(sub, []string{"sports"})
	}

	if subs := len(week10.Subscribers()); subs != 10 {
		t.Fatalf("Expected 0 subscriber, got %d", subs)
	}

	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			week10.Unsubscribe(&tSubscriber{id: fmt.Sprintf("%d", id)})
			wg.Done()
		}(i)
	}
	wg.Wait()

	if subs := len(week10.Subscribers()); subs != 0 {
		t.Fatalf("Expected 0 subscriber, got %d", subs)
	}
}

// func TestReceive(t *testing.T) {
// 	sub := &tSubscriber{
// 		id:  fmt.Sprintf("1"),
// 		out: bytes.NewBufferString(""),
// 	}

// 	week10.Subscribe(sub, []string{"sports"})

// 	wg := new(sync.WaitGroup)
// 	wg.Add(1)

// 	crazyNews := week09.News{
// 		Title:   "redsox win the super-bowl",
// 		Content: "the redsox win on the ice",
// 		Author:  "Mark Bates",

// 		Categories: []string{"sports"},
// 	}

// 	go func() {
// 		week10.Receive(crazyNews)
// 		wg.Done()
// 	}()

// 	wg.Add(1)

// 	smartNews := week09.News{
// 		Title:   "super interesting article",
// 		Content: "this is a crazy interesting topic that no-one reads.",
// 		Author:  "Phd Dr KnowItAll",

// 		Categories: []string{"smartythings"},
// 	}

// 	go func() {
// 		service.Receive(smartNews)
// 		wg.Done()
// 	}()

// 	wg.Wait()

// 	if !strings.Contains(sub.out.String(), crazyNews.Title) {
// 		t.Fatalf("expected %s to contain %s", sub.out.String(), crazyNews.Title)
// 	}

// 	if strings.Contains(sub.out.String(), smartNews.Title) {
// 		t.Fatalf("expected %s not to contain %s", sub.out.String(), smartNews.Title)
// 	}
// }
