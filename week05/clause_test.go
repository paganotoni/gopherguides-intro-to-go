package demo

import (
	"fmt"
	"testing"
)

func TestClausesString(t *testing.T) {
	t.Parallel()

	tcases := []struct {
		name    string
		clauses Clauses
		want    string
	}{
		{"empty", Clauses{}, ""},
		{"one", Clauses{"A": 1}, fmt.Sprintf(`"A" = %q`, 1)},
		{"two", Clauses{"A": 1, "B": "aaa"}, fmt.Sprintf(`"A" = %q and "B" = "aaa"`, 1)},
		{"more", Clauses{"A": 1, "B": "aaa", "C": true}, fmt.Sprintf(`"A" = %q and "B" = "aaa" and "C" = %%!q(bool=true)`, 1)},
	}

	for _, tt := range tcases {
		t.Run(tt.name, func(t *testing.T) {

			c := tt.clauses
			got := c.String()
			if got != tt.want {
				t.Errorf("Clauses.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClausesMatch(t *testing.T) {
	t.Parallel()

	tcases := []struct {
		name    string
		clauses Clauses
		model   Model
		want    bool
	}{
		{"empty-match", Clauses{}, Model{}, true},
		{"dispair-match", Clauses{}, Model{"a": "b"}, true},
		{"dispair-no-match", Clauses{"a": "b"}, Model{}, false},
		{"multiple-match", Clauses{"a": "b", "c": "d"}, Model{"a": "b", "c": "d"}, true},
		{"multiple-dispair-match", Clauses{"a": "b", "c": "d"}, Model{"a": "b", "c": "d", "d": "c"}, true},
		{"multiple-no-match", Clauses{"a": "b", "d": "c"}, Model{"a": "b", "c": "d"}, false},
	}

	for _, tt := range tcases {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.clauses
			got := c.Match(tt.model)
			if got != tt.want {
				t.Errorf("Clauses.Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
