package week06

import "testing"

func TestEmployeeValid(t *testing.T) {
	tcases := []struct {
		name     string
		employee Employee
		errors   bool
	}{
		{name: "Zero", employee: Employee(0), errors: true},
		{name: "One", employee: Employee(1), errors: false},
		{name: "One", employee: Employee(-1), errors: true},
	}

	for _, tt := range tcases {
		terrors := tt.employee.IsValid() != nil
		if terrors == tt.errors {
			continue
		}

		t.Errorf("%s: expected %v, got %v", tt.name, tt.errors, tt.employee.IsValid())
	}
}
