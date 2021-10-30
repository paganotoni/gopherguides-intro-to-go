package demo

import (
	"fmt"
	"testing"
)

func TestErrTableNotFoundError(t *testing.T) {
	t.Parallel()

	tcases := []struct {
		name  string
		want  string
		table string
	}{
		{"empty", "table not found ", ""},
		{"provided", "table not found cars", "cars"},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			err := ErrTableNotFound{
				table: tc.table,
			}

			got := err.Error()
			if got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestErrTableNotFoundTableNotFound(t *testing.T) {
	t.Parallel()

	tcases := []struct {
		name  string
		want  string
		table string
	}{
		{"empty", "", ""},
		{"provided", "cars", "cars"},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			err := ErrTableNotFound{
				table: tc.table,
			}

			got := err.TableNotFound()
			if got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestErrTableNotFoundIs(t *testing.T) {
	t.Parallel()

	tcases := []struct {
		name string
		want bool
		terr error
	}{
		{"matches", true, ErrTableNotFound{}},
		{"other", false, fmt.Errorf("Some other error")},
		{"wrapped", false, fmt.Errorf("Some other error wrapping: %w", ErrTableNotFound{})},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			e := ErrTableNotFound{}

			got := e.Is(tc.terr)
			if got != tc.want {
				t.Errorf("got %t, want %t", got, tc.want)
			}
		})
	}
}

func TestIsErrTableNotFound(t *testing.T) {
	t.Parallel()

	tcases := []struct {
		name string
		err  error
		want bool
	}{
		{"matches", ErrTableNotFound{}, true},
		{"other", fmt.Errorf("Some other error"), false},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			got := IsErrTableNotFound(tc.err)
			if got != tc.want {
				t.Errorf("got %t, want %t", got, tc.want)
			}
		})
	}
}

func TestErrNoRowsError(t *testing.T) {
	t.Parallel()

	tcases := []struct {
		name string
		want string
		err  errNoRows
	}{
		{"empty", "[] no rows found\nquery: ", errNoRows{}},
		{
			"withdata",
			`[cars] no rows found` + "\n" + `query: "A" = "B"`,
			errNoRows{
				table: "cars",
				clauses: Clauses{
					"A": "B",
				},
			},
		},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.err.Error()
			if got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestErrNoRowsErrorClauses(t *testing.T) {
	t.Parallel()

	tcases := []struct {
		name string
		want Clauses
		err  errNoRows
	}{
		{"empty", Clauses{}, errNoRows{}},
		{
			"with",
			Clauses{"A": "B"},
			errNoRows{
				clauses: Clauses{"A": "B"},
			},
		},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.err.Clauses()
			if got.String() != tc.want.String() {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestErrNoRowsErrorRowsNotFound(t *testing.T) {
	t.Parallel()

	tcases := []struct {
		name string
		err  errNoRows

		table   string
		clauses Clauses
	}{
		{"empty", errNoRows{}, "", Clauses{}},
		{"empty", errNoRows{table: "cars", clauses: Clauses{"A": "B"}}, "cars", Clauses{"A": "B"}},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			gtable, gclauses := tc.err.RowNotFound()
			if gtable != tc.table {
				t.Errorf("got %q for tables, wanted %q", gtable, tc.table)
			}

			if !gclauses.Match(Model(tc.clauses)) {
				t.Errorf("got %q for clauses, wanted %q", gclauses, tc.clauses)
			}
		})
	}
}

func TestErrNoRowsErrorRowsIs(t *testing.T) {
	t.Parallel()

	tcases := []struct {
		name string
		err  error
		want bool
	}{
		{"matches", &errNoRows{}, true},
		{"no-match", fmt.Errorf("some error"), false},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			e := &errNoRows{}

			got := e.Is(tc.err)
			if got != tc.want {
				t.Errorf("got %t, want %t", got, tc.want)
			}
		})
	}
}

func TestIsErrNoRowsErrorRows(t *testing.T) {
	t.Parallel()

	tcases := []struct {
		name string
		err  error
		want bool
	}{
		{"matches", &errNoRows{}, true},
		{"no-match", fmt.Errorf("some error"), false},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			got := IsErrNoRows(tc.err)
			if got != tc.want {
				t.Errorf("got %t, want %t", got, tc.want)
			}
		})
	}
}

func TestAsErrNoRows(t *testing.T) {
	t.Parallel()

	tcases := []struct {
		name string
		err  error
		ok   bool
	}{
		{"matches", &errNoRows{}, true},
		{"no-match", fmt.Errorf("some error"), false},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			_, got := AsErrNoRows(tc.err)
			if got != tc.ok {
				t.Errorf("got %t, want %t", got, tc.ok)
			}
		})
	}
}
