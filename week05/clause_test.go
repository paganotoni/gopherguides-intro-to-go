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
		{name: "empty", clauses: Clauses{}, want: ""},
		{name: "one", clauses: Clauses{"A": 1}, want: fmt.Sprintf(`"A" = %q`, 1)},
		{name: "two", clauses: Clauses{"A": 1, "B": "aaa"}, want: fmt.Sprintf(`"A" = %q and "B" = "aaa"`, 1)},
		{name: "more", clauses: Clauses{"A": 1, "B": "aaa", "C": true}, want: fmt.Sprintf(`"A" = %q and "B" = "aaa" and "C" = %%!q(bool=true)`, 1)},
	}

	for _, tt := range tcases {
		t.Run(tt.name, func(t *testing.T) {

			c := tt.clauses
			got := c.String()
			if got != tt.want {
				t.Fatalf("expected %v, got %v", tt.want, got)
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
		{name: "empty match", clauses: Clauses{}, model: Model{}, want: true},
		{name: "dispair match", clauses: Clauses{}, model: Model{"a": "b"}, want: true},
		{name: "dispair no match", clauses: Clauses{"a": "b"}, model: Model{}, want: false},
		{name: "multiple match", clauses: Clauses{"a": "b", "c": "d"}, model: Model{"a": "b", "c": "d"}, want: true},
		{name: "multiple dispair match", clauses: Clauses{"a": "b", "c": "d"}, model: Model{"a": "b", "c": "d", "d": "c"}, want: true},
		{name: "multiple no-match", clauses: Clauses{"a": "b", "d": "c"}, model: Model{"a": "b", "c": "d"}, want: false},
	}

	for _, tt := range tcases {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.clauses
			got := c.Match(tt.model)
			if got != tt.want {
				t.Fatalf("expected %v, got %v", tt.want, got)
			}
		})
	}
}
