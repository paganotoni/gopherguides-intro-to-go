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
		{name: "empty", want: "table not found ", table: ""},
		{name: "provided", want: "table not found cars", table: "cars"},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			err := ErrTableNotFound{
				table: tc.table,
			}

			got := err.Error()
			if got != tc.want {
				t.Fatalf("expected %q, got %q", tc.want, got)
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
		{name: "empty", want: "", table: ""},
		{name: "provided", want: "cars", table: "cars"},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			err := ErrTableNotFound{
				table: tc.table,
			}

			got := err.TableNotFound()
			if got != tc.want {
				t.Fatalf("expected %q, got %q", tc.want, got)
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
		{name: "matches", want: true, terr: ErrTableNotFound{}},
		{name: "other", want: false, terr: fmt.Errorf("Some other error")},
		{name: "wrapped", want: false, terr: fmt.Errorf("Some other error wrapping: %w", ErrTableNotFound{})},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			e := ErrTableNotFound{}

			got := e.Is(tc.terr)
			if got != tc.want {
				t.Fatalf("expected %t, got %t", tc.want, got)
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
		{name: "match", err: ErrTableNotFound{}, want: true},
		{name: "other", err: fmt.Errorf("Sother error"), want: false},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			got := IsErrTableNotFound(tc.err)
			if got != tc.want {
				t.Fatalf("expected %t, got %t", tc.want, got)
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
		{name: "empty", want: "[] no rows found\nquery: ", err: errNoRows{}},
		{
			name: "with data",
			want: `[cars] no rows found` + "\n" + `query: "A" = "B"`,
			err: errNoRows{
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
				t.Fatalf("expected %q, got %q", tc.want, got)
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
		{name: "empty", want: Clauses{}, err: errNoRows{}},
		{
			name: "with",
			want: Clauses{"A": "B"},
			err: errNoRows{
				clauses: Clauses{"A": "B"},
			},
		},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.err.Clauses()
			if got.String() != tc.want.String() {
				t.Fatalf("expected %q, got %q", tc.want, got)
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
		{name: "empty", err: errNoRows{}, table: "", clauses: Clauses{}},
		{name: "empty", err: errNoRows{table: "cars", clauses: Clauses{"A": "B"}}, table: "cars", clauses: Clauses{"A": "B"}},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			gtable, gclauses := tc.err.RowNotFound()
			if gtable != tc.table {
				t.Fatalf("expected %q for tables, got %q", tc.table, gtable)
			}

			if !gclauses.Match(Model(tc.clauses)) {
				t.Fatalf("expected %q for clauses, got %q", tc.clauses, gclauses)
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
		{name: "matches", err: &errNoRows{}, want: true},
		{name: "no match", err: fmt.Errorf("some error"), want: false},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			e := &errNoRows{}

			got := e.Is(tc.err)
			if got != tc.want {
				t.Fatalf("expected %t, got %t", tc.want, got)
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
		{name: "matches", err: &errNoRows{}, want: true},
		{name: "no match", err: fmt.Errorf("some error"), want: false},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			got := IsErrNoRows(tc.err)
			if got != tc.want {
				t.Fatalf("expected %t, got %t", tc.want, got)
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
		{name: "matches", err: &errNoRows{}, ok: true},
		{name: "no match", err: fmt.Errorf("some error"), ok: false},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			_, got := AsErrNoRows(tc.err)
			if got != tc.ok {
				t.Fatalf("expected %t, got %t", tc.ok, got)
			}
		})
	}
}
