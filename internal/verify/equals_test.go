package verify

import "testing"

func TestEquals(t *testing.T) {
	tcases := []struct {
		a, b interface{}
	}{
		{1, 1},
		{"1", "1"},
		{[]string{"A"}, []string{"A"}},
		// Arrays must be in the same order!
		{[]string{"A", "B"}, []string{"A", "B"}},
		{map[string]string{"A": "B"}, map[string]string{"A": "B"}},
		{map[string]string{"A": "B", "C": "D"}, map[string]string{"C": "D", "A": "B"}},
	}

	for _, tcase := range tcases {
		Equals(t, tcase.a, tcase.b)
	}

	for _, tcase := range tcases {
		Equalsf(t, tcase.a, tcase.b, "%s is not equal to %s")
	}
}
