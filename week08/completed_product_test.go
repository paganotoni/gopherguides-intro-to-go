package week08

import "testing"

func TestCompletedProductIsValid(t *testing.T) {
	p := Product{Materials: Materials{
		"A": 1,
	}}

	w := Warehouse{}
	p.Build(Employee(1), &w)

	tcases := []struct {
		name    string
		product CompletedProduct
		errors  bool
	}{
		{name: "Empty", product: CompletedProduct{}, errors: true},
		{name: "EmptyEmployee", product: CompletedProduct{Employee: Employee(1)}, errors: true},
		{name: "EmptyProduct", product: CompletedProduct{Product: p}, errors: true},
		{name: "Full", product: CompletedProduct{Employee: Employee(1), Product: p}, errors: false},
	}

	for _, tt := range tcases {
		t.Run(tt.name, func(t *testing.T) {
			if (tt.product.IsValid() != nil) == tt.errors {
				return
			}

			t.Fatalf("Expected errors: %v, got: %v", tt.errors, tt.product.IsValid())
		})
	}
}
