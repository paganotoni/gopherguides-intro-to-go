package week07

import "testing"

func TestErrors(t *testing.T) {
	tcases := []struct {
		err error
		msg string
	}{
		{ErrInvalidQuantity(-1), "quantity must be greater than 0, got -1"},
		{ErrProductNotBuilt("product not built"), "product not built"},
		{ErrInvalidEmployee(-1), "invalid employee number: -1"},
		{ErrInvalidEmployeeCount(-1), "invalid employee count: -1"},
		{ErrManagerStopped{}, "manager is stopped"},
	}

	for _, tc := range tcases {
		if tc.err.Error() != tc.msg {
			t.Errorf("expected %q, got %q", tc.msg, tc.err.Error())
		}
	}
}
