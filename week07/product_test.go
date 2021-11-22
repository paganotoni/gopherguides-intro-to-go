package week07

import "testing"

func TestProductIsValid(t *testing.T) {
	tcases := []struct {
		name    string
		product Product
		errors  bool
	}{
		{name: "Invalid Product", product: Product{Quantity: 0}, errors: true},
		{name: "Valid Product", product: Product{Quantity: 1}, errors: false},
	}

	for _, tcase := range tcases {
		t.Run(tcase.name, func(t *testing.T) {
			terr := tcase.product.IsValid()
			terrors := terr != nil
			if terrors == tcase.errors {
				return
			}

			t.Fatalf("Expected errors: %v, got: %v", tcase.errors, terrors)
		})
	}
}

func TestProductBuiltBy(t *testing.T) {
	p := Product{Quantity: 2}
	if p.BuiltBy() != Employee(0) {
		t.Error("Expected employee 0 to be built by")
	}

	(&p).Build(Employee(1))
	if p.BuiltBy() != Employee(1) {
		t.Error("Expected employee 0 to be built by")
	}
}

func TestProductBuild(t *testing.T) {
	tcases := []struct {
		name     string
		product  Product
		employee Employee
		errors   bool
	}{
		{name: "Invalid Product", product: Product{Quantity: 0}, employee: Employee(1), errors: true},
		{name: "Invalid Employee", product: Product{Quantity: 1}, employee: Employee(0), errors: true},
		{name: "Build valid", product: Product{Quantity: 1}, employee: Employee(1), errors: false},
	}

	for _, tcase := range tcases {
		t.Run(tcase.name, func(t *testing.T) {
			terr := tcase.product.Build(tcase.employee)
			terrors := terr != nil
			if terrors == tcase.errors {
				return
			}

			t.Fatalf("Expected errors: %v, got: %v", tcase.errors, terrors)
		})
	}
}

func TestProductIsBuilt(t *testing.T) {
	p := Product{Quantity: 2}
	(&p).Build(Employee(1))

	tcases := []struct {
		name    string
		product Product
		errors  bool
	}{
		{name: "Invalid Product", product: Product{Quantity: 0}, errors: true},
		{name: "Not Built", product: Product{Quantity: 1}, errors: true},
		{name: "Built", product: p, errors: false},
	}

	for _, tcase := range tcases {
		t.Run(tcase.name, func(t *testing.T) {
			err := tcase.product.IsBuilt()
			errs := err != nil
			if errs == tcase.errors {
				return
			}

			t.Fatalf("Expected errors: %v, got: %v", tcase.errors, errs)
		})
	}
}
