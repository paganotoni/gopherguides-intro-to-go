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
	ctx, cf := context.WithCancel(ctx)

	w.Lock()
	w.cancel = cf
	w.Unlock()

	return ctx
}

// Stop the warehouse
func (w *Warehouse) Stop() {
	w.RLock()
	if w.cancel != nil {
		w.cancel()
	}
	w.RUnlock()
}

// Retrieve quantity of material from the warehouse
func (w *Warehouse) Retrieve(m Material, q int) (Material, error) {
	ctx := w.fill(m)

	// wait for the materials to become available
	<-ctx.Done()

	// remove the materials from the warehouse
	w.Lock()
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

		w.RLock()
		ca := w.cap
		mt := w.materials
		w.RUnlock()

		if ca <= 0 {
			ca = 10
		}

		w.Lock()
		if mt == nil {
			mt = Materials{}
		}

		// until the warehouse is full of
		// the material create the material and
		// fill the warehouse
		q := mt[m]
		for q < ca {
			time.Sleep(m.Duration())
			mt[m]++
			q = mt[m]
		}

		w.materials = mt
		w.Unlock()
	}()

	return ctx
}
