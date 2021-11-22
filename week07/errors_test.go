package week07

import "testing"

func TestErrors(t *testing.T) {
	tcases := []struct {
		err error
		msg string
	}{
		{err: ErrInvalidQuantity(-1), msg: "quantity must be greater than 0, got -1"},
		{err: ErrProductNotBuilt("product not built"), msg: "product not built"},
		{err: ErrInvalidEmployee(-1), msg: "invalid employee number: -1"},
		{err: ErrInvalidEmployeeCount(-1), msg: "invalid employee count: -1"},
		{err: ErrManagerStopped{}, msg: "manager is stopped"},
	}

	for _, tc := range tcases {
		if tc.err.Error() != tc.msg {
			t.Errorf("expected %q, got %q", tc.msg, tc.err.Error())
		}
	}
}
