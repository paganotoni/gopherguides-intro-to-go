package week10

import (
	"sync"
	"time"
)

// NewMockSource returns a new mock source with the configured
// interval set.
func NewMockSource(interval time.Duration) Source {
	return &mockSource{interval: interval}
}

// mockSource sends mock news at a configured rate.
type mockSource struct {
	stopped  bool
	interval time.Duration

	sync.RWMutex
}

func (ms *mockSource) Stop() {
	ms.Lock()
	defer ms.Unlock()

	ms.stopped = true
}

// PublishTo starts a goroutine that sends mock news at the configured
// interval through the given news channel.
func (ms *mockSource) PublishTo(ch chan News) {
	go func() {
		t := time.Tick(ms.interval)
		for {
			select {
			case <-t:
				ms.RLock()
				if ms.stopped {
					ms.RUnlock()
					break
				}
				ms.RUnlock()

				ch <- News{
					ID:         int(time.Now().Unix()),
					Title:      "mock news",
					Content:    "mock news body",
					Categories: []string{"mock"},
				}
			}
		}
	}()
}
