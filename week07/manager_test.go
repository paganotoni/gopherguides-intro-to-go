package week07

import (
	"context"
	"os/signal"
	"sync"
	"syscall"
	"testing"
	"time"
)

const TEST_SIGNAL = syscall.SIGUSR2

// snippet: example
func Test_Run(t *testing.T) {
	t.Parallel()

	t.Run("timeout after 5 seconds if nothing happens", func(t *testing.T) {
		t.Parallel()

		ctx, cnFn := context.WithTimeout(context.Background(), 5*time.Second)
		defer cnFn()

		products := []*Product{
			// 6_000 milliseconds = 6 seconds
			&Product{Quantity: 6_000},
		}

		cprod, err := Run(ctx, 1, products...)
		if err != nil {
			t.Errorf("got error %v, want nil", err)
		}

		if len(cprod) != 0 {
			t.Errorf("got %v, want 0", len(cprod))
		}
	})

	t.Run("interruption by a signal", func(t *testing.T) {
		t.Parallel()

		ctx, cnFn := signal.NotifyContext(context.Background(), TEST_SIGNAL)
		defer cnFn()

		go func() {
			time.Sleep(time.Second)

			t.Log("sending test signal")

			// send the TEST_SIGNAL to the system
			syscall.Kill(syscall.Getpid(), TEST_SIGNAL)
		}()

		products := []*Product{
			// 1_500 milliseconds = 1.5 seconds
			&Product{Quantity: 1_500},
		}

		cprod, err := Run(ctx, 1, products...)
		if err != nil {
			t.Errorf("got error %v, want nil", err)
		}

		if len(cprod) != 0 {
			t.Errorf("got %v, want 0", len(cprod))
		}
	})

	t.Run("products completed", func(t *testing.T) {
		t.Parallel()

		ctx, cnFn := context.WithTimeout(context.Background(), 5*time.Second)
		defer cnFn()

		products := []*Product{
			// 1_500 milliseconds = 1.5 seconds
			&Product{Quantity: 1_500},
			// 2_000 milliseconds = 2 seconds
			&Product{Quantity: 2_000},
			// 3_000 milliseconds = 3 seconds
			&Product{Quantity: 3_000},
		}

		cprod, err := Run(ctx, 3, products...)
		if err != nil {
			t.Fatalf("got error %v, want nil", err)
		}

		if len(cprod) != 3 {
			t.Fatalf("got %v, want 3", len(cprod))
		}

		for _, v := range cprod {
			if v.Employee != Employee(0) {
				continue
			}

			t.Fatalf("Employee should not be %v", v.Employee)
		}

	})
}

func TestManagerStartError(t *testing.T) {
	m := NewManager()

	err := m.Start(-1)
	if err == nil {
		t.Fatal("manager should err when calling start with less than 1 worker")
	}
}

func TestManagerStart(t *testing.T) {
	m := NewManager()
	t.Cleanup(func() {
		m.Stop()
	})

	err := m.Start(1)
	if err != nil {
		t.Fatal("manager should not err when calling start with 1 worker")
	}
}

func TestManagerAssign(t *testing.T) {

	tcases := []struct {
		name      string
		product   *Product
		managerFn func() *Manager
		errors    bool
	}{
		{
			name:    "stopped manager",
			errors:  true,
			product: &Product{Quantity: 1},
			managerFn: func() *Manager {
				m := NewManager()
				m.Stop()

				return m
			},
		},

		{
			name:    "invalid product",
			errors:  true,
			product: &Product{Quantity: 0},
			managerFn: func() *Manager {
				m := NewManager()

				return m
			},
		},

		{
			name:    "correct",
			errors:  false,
			product: &Product{Quantity: 1},
			managerFn: func() *Manager {
				m := NewManager()
				m.Start(1)

				return m
			},
		},
	}

	for _, tcase := range tcases {
		t.Run(tcase.name, func(t *testing.T) {
			m := tcase.managerFn()
			t.Cleanup(func() {
				m.Stop()
			})

			err := m.Assign(tcase.product)
			if tcase.errors && err == nil {
				t.Fatal("manager should err when calling assign")
			}

			if !tcase.errors && err != nil {
				t.Fatal("manager should not err when calling assign")
			}
		})
	}
}

func TestManagerComplete(t *testing.T) {
	m := NewManager()

	p := &Product{Quantity: 1}
	p.Build(Employee(1))

	tcases := []struct {
		name     string
		employee Employee
		product  *Product
		errors   bool
	}{
		{name: "Invalid employee", product: &Product{Quantity: 0}, errors: true},
		{name: "Invalid product", employee: Employee(1), product: &Product{Quantity: 1}, errors: true},
		{name: "All Good", employee: Employee(1), product: p, errors: false},
	}

	// Go routine to avoid locking on the success case.
	var wg sync.WaitGroup
	var cp CompletedProduct
	wg.Add(1)

	go func() {
		cp = <-m.Completed()
		wg.Done()
	}()

	for _, tcase := range tcases {
		t.Run(tcase.name, func(t *testing.T) {
			err := m.Complete(tcase.employee, tcase.product)
			if tcase.errors && err == nil {
				t.Fatal("manager should err when calling complete")
			}

			if !tcase.errors && err != nil {
				t.Fatal("manager should not err when calling complete")
			}
		})
	}

	wg.Wait()
	if cp.Employee != Employee(1) || cp.Product.Quantity != 1 {
		t.Fatal("manager should send completed product")
	}

}

func TestManagerAsssignBuilt(t *testing.T) {
	m := NewManager()
	m.Start(2)

	t.Cleanup(func() {
		m.Stop()
	})

	p := &Product{Quantity: -1}
	m.Jobs() <- p

	err := <-m.Errors()
	if err == nil {
		t.Fatal("manager should err when calling assign with invalid product")
	}
}
