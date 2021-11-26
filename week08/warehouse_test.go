package week08

import (
	"context"
	"testing"
)

func TestWarehouseStart(t *testing.T) {
	w := &Warehouse{}
	ctx := context.Background()
	_ = w.Start(ctx)

	t.Cleanup(func() {
		w.Stop()
	})

	if w.cancel == nil {
		t.Error("warehouse cancel should not be nil")
	}
}

func TestWarehouseFill(t *testing.T) {
	w := &Warehouse{cap: 100}
	ctx := context.Background()
	_ = w.Start(ctx)

	t.Cleanup(func() {
		w.Stop()
	})

	m, err := w.Retrieve(Material("apple"), 2)
	if err != nil {
		t.Errorf("error should be nil, got %v", err)
	}

	if m != Material("apple") {
		t.Errorf("material should be apple, got %v", m)
	}
}
