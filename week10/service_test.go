package week10_test

import (
	"bytes"
	"fmt"
	"strings"
	"sync"
	"testing"

	"github.com/paganotoni/gopherguides-intro-to-go/week09"
)

func TestServiceSubscribeMany(t *testing.T) {
	service := &week09.Service{}

	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wgx *sync.WaitGroup, index int) {
			sub := tSubscriber{
				id:  fmt.Sprintf("%d", index),
				out: bytes.NewBufferString(""),
			}

			service.Subscribe(&sub, []string{"sports"})
			wgx.Done()
		}(wg, i)
	}

	wg.Wait()

	if len(service.Subscribers()) != 10 {
		t.Fatalf("Expected 10 subscriber, got %d", len(service.Subscribers()))
	}
}

func TestServiceSubscribeRepeated(t *testing.T) {
	service := &week09.Service{}

	sub := tSubscriber{
		id:  "1",
		out: bytes.NewBufferString(""),
	}

	sub2 := tSubscriber{
		id:  "1",
		out: bytes.NewBufferString(""),
	}

	err := service.Subscribe(&sub, []string{"sports"})
	if err != nil {
		t.Fatal("err should be nil, got %w", err)
	}

	err = service.Subscribe(&sub2, []string{"sports"})
	if err == nil {
		t.Fatal("err should not nil")
	}
}

func TestServiceUnsubscribe(t *testing.T) {
	service := &week09.Service{}

	for i := 0; i < 10; i++ {
		sub := tSubscriber{
			id:  fmt.Sprintf("%d", i),
			out: bytes.NewBufferString(""),
		}

		service.Subscribe(sub, []string{"sports"})
	}

	if len(service.Subscribers()) != 10 {
		t.Fatalf("Expected 0 subscriber, got %d", len(service.Subscribers()))
	}

	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			service.Unsubscribe(fmt.Sprintf("%d", id))
			wg.Done()
		}(i)
	}
	wg.Wait()

	if len(service.Subscribers()) != 0 {
		t.Fatalf("Expected 0 subscriber, got %d", len(service.Subscribers()))
	}
}

func TestReceive(t *testing.T) {
	service := &week09.Service{}

	sub := &tSubscriber{
		id:  fmt.Sprintf("1"),
		out: bytes.NewBufferString(""),
	}

	err := service.Subscribe(sub, []string{"sports"})
	if err != nil {
		t.Fatal("err should be nil, got %w", err)
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)

	crazyNews := week09.News{
		Title:   "redsox win the super-bowl",
		Content: "the redsox win on the ice",
		Author:  "Mark Bates",

		Categories: []string{"sports"},
	}

	go func() {
		service.Receive(crazyNews)
		wg.Done()
	}()

	wg.Add(1)

	smartNews := week09.News{
		Title:   "super interesting article",
		Content: "this is a crazy interesting topic that no-one reads.",
		Author:  "Phd Dr KnowItAll",

		Categories: []string{"smartythings"},
	}

	go func() {
		service.Receive(smartNews)
		wg.Done()
	}()

	wg.Wait()

	if !strings.Contains(sub.out.String(), crazyNews.Title) {
		t.Fatalf("expected %s to contain %s", sub.out.String(), crazyNews.Title)
	}

	if strings.Contains(sub.out.String(), smartNews.Title) {
		t.Fatalf("expected %s not to contain %s", sub.out.String(), smartNews.Title)
	}
}
