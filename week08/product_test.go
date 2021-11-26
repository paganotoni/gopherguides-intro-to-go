package week08

import "testing"

func TestProductIsValid(t *testing.T) {
	tcases := []struct {
		name    string
		product Product
		errors  bool
	}{
		{name: "Invalid Product", product: Product{Materials: Materials{}}, errors: true},
		{name: "Valid Product", product: Product{Materials: Materials{"A": 1}}, errors: false},
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
	p := Product{Materials: Materials{"A": 1}}
	if p.BuiltBy() != Employee(0) {
		t.Error("Expected employee 0 to be built by")
	}

	w := &Warehouse{}

	(&p).Build(Employee(1), w)
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
		{name: "Invalid Product", product: Product{Materials: Materials{}}, employee: Employee(1), errors: true},
		{name: "Invalid Employee", product: Product{Materials: Materials{"A": 1}}, employee: Employee(0), errors: true},
		{name: "Build valid", product: Product{Materials: Materials{"A": 1}}, employee: Employee(1), errors: false},
	}

	for _, tcase := range tcases {
		t.Run(tcase.name, func(t *testing.T) {
			w := &Warehouse{}
			terr := tcase.product.Build(tcase.employee, w)
			terrors := terr != nil
			if terrors == tcase.errors {
				return
			}

			t.Fatalf("Expected errors: %v, got: %v", tcase.errors, terrors)
		})
	}
}

func TestProductIsBuilt(t *testing.T) {
	p := Product{Materials: Materials{"A": 1}}
	w := &Warehouse{}
	(&p).Build(Employee(1), w)

	tcases := []struct {
		name    string
		product Product
		errors  bool
	}{
		{name: "Invalid Product", product: Product{Materials: Materials{"A": 1}}, errors: true},
		{name: "Not Built", product: Product{Materials: Materials{"A": 1}}, errors: true},
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
