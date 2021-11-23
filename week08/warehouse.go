package week08

import (
	"context"
	"sync"
	"time"
)

// Warehouse is where the materials are stored
// and where the materials are retrieved from
type Warehouse struct {
	cancel    context.CancelFunc // cancels the warehouse
	cap       int                // capacity of the warehouse
	materials Materials          // materials in the warehouse

	sync.RWMutex
}

// Start the warehouse
func (w *Warehouse) Start(ctx context.Context) context.Context {
	ctx, c := context.WithCancel(ctx)

	w.cancel = c

	return ctx
}

// Stop the warehouse
func (w *Warehouse) Stop() {
	w.cancel()
}

// Retrieve quantity of material from the warehouse
func (w *Warehouse) Retrieve(m Material, q int) (Material, error) {
	ctx := w.fill(m)

	// wait for the materials to become available
	<-ctx.Done()

	w.Lock()
	// remove the materials from the warehouse
	w.materials[m] -= q
	w.Unlock()

	return m, nil
}

// fill the warehouse with the material until it is full
func (w *Warehouse) fill(m Material) context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	// lauch a goroutine to fill the warehouse
	// until it is full
	// context is cancelled when the warehouse is full
	go func() {
		defer cancel()

		w.Lock()
		if w.cap <= 0 {
			w.cap = 10
		}

		if w.materials == nil {
			w.materials = Materials{}
		}
		w.Unlock()

		cap := w.cap
		mats := w.materials

		w.Lock()
		// until the warehouse is full of
		// the material create the material and
		// fill the warehouse
		q := mats[m]
		for q < cap {
			time.Sleep(m.Duration())
			mats[m]++
			q = mats[m]
		}
		w.Unlock()
	}()

	return ctx
}
