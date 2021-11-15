package week06

import (
	"sync"
	"testing"
)

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
