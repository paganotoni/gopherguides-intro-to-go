package week08

import (
	"context"
	"testing"
)

func TestManagerStartError(t *testing.T) {
	m := Manager{}
	_, err := m.Start(context.Background(), -1)
	if err == nil {
		t.Fatal("manager should err when calling start with less than 1 worker")
	}
}

func TestManagerStartValid(t *testing.T) {
	m := Manager{}
	_, err := m.Start(context.Background(), 1)
	t.Cleanup(func() {
		m.Stop()
	})

	if err != nil {
		t.Fatal("manager should not err when calling start with 1 worker")
	}
}

func TestManagerAssignStopped(t *testing.T) {
	m := Manager{}
	m.Start(context.Background(), 1)
	t.Cleanup(func() {
		m.Stop()
	})

	err := m.Assign(&Product{
		Materials: Materials{},
	})

	if err == nil {
		t.Fatal("manager should err when calling assign on a stopped manager")
	}
}

func TestManagerAssignInvalid(t *testing.T) {
	m := Manager{}
	m.Start(context.Background(), 1)
	m.Stop()

	err := m.Assign(ProductA)
	if err == nil {
		t.Fatal("manager should err when calling assign on a stopped manager")
	}
}

func TestManagerAssignOk(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	m := &Manager{}
	ctx, err := m.Start(ctx, 5)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for i := 0; i < 25; i++ {
		go m.Assign(&Product{
			Materials: Materials{Material("iron"): 1},
		})

		go m.Assign(&Product{
			Materials: Materials{Material("copper"): 1},
		})
	}

	var completed []CompletedProduct

	go func() {
		for cp := range m.Completed() {
			completed = append(completed, cp)

			if len(completed) >= 50 {
				m.Stop()
			}
		}
	}()

	<-ctx.Done()
	if len(completed) != 20 {
		t.Errorf("expected 20 products, got %d", len(completed))
	}
}
