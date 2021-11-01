package demo

import (
	"errors"
	"testing"
)

func TestStoreDb(t *testing.T) {
	t.Parallel()

	tcases := []struct {
		name  string
		store Store
		want  data
	}{
		{name: "empty", store: Store{}, want: data{}},
		{name: "empty", store: Store{data: data{"cars": Models{{"A": "B"}}}}, want: data{"cars": Models{{"A": "B"}}}},
	}

	for _, tt := range tcases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.store.db()
			checkDataMatches(t, tt.want, got)
		})
	}
}

func TestStoreAll(t *testing.T) {
	t.Parallel()

	tcases := []struct {
		name   string
		store  Store
		table  string
		models Models
		err    error
	}{
		{name: "empty", store: Store{}, table: "cars", models: nil, err: &ErrTableNotFound{table: "cars"}},
		{name: "with data", store: Store{data: data{"cars": Models{{"A": "B"}}}}, table: "cars", models: Models{{"A": "B"}}, err: nil},
	}

	for _, tt := range tcases {
		t.Run(tt.name, func(t *testing.T) {
			tmod, terr := tt.store.All(tt.table)

			checkErrorsMatch(t, tt.err, terr)
			checkModelsMatch(t, tt.models, tmod)
		})
	}
}

func TestStoreLen(t *testing.T) {
	t.Parallel()

	tcases := []struct {
		name  string
		store Store

		table string
		count int
		err   error
	}{
		{name: "empty", store: Store{}, table: "cars", count: 0, err: &ErrTableNotFound{table: "cars"}},
		{name: "with data", store: Store{data: data{"cars": Models{{"A": "B"}}}}, table: "cars", count: 1, err: nil},
		{name: "with data multiple models", store: Store{data: data{"cars": Models{{"A": "B"}}, "books": Models{{"C": "D"}}}}, table: "cars", count: 1, err: nil},
	}

	for _, tt := range tcases {
		t.Run(tt.name, func(t *testing.T) {
			count, terr := tt.store.Len(tt.table)
			if count != tt.count {
				t.Fatalf("expected %v, got %v", tt.count, count)
			}

			checkErrorsMatch(t, tt.err, terr)
		})
	}
}

func TestStoreInsert(t *testing.T) {
	t.Parallel()

	store := Store{}
	store.Insert("car", Model{"A": "B"})

	len, _ := store.Len("car")
	if len != 1 {
		t.Fatalf("expected len %v, got %v", 1, len)
	}

	mods, err := store.All("car")
	if err != nil {
		t.Fatalf("expected err %v, got nil", err)
	}

	checkModelsMatch(t, Models{{"A": "B"}}, mods)
	store.Insert("car", Model{"C": "D"})

	len, _ = store.Len("car")
	if len != 2 {
		t.Fatalf("expected len %v, want %v", len, 1)
	}
}

func TestStoreSelect(t *testing.T) {
	t.Parallel()

	tcases := []struct {
		name    string
		err     error
		len     int
		clauses Clauses
		store   Store
	}{
		{name: "empty", err: &ErrTableNotFound{table: "cars"}, len: 0, clauses: Clauses{}, store: Store{}},
		{
			name:    "no rows",
			err:     &errNoRows{table: "cars", clauses: Clauses{"A": "C"}},
			len:     0,
			clauses: Clauses{"A": "C"},
			store: Store{
				data{
					"cars": Models{{"A": "B"}},
				},
			},
		},
		{
			name:    "no clauses",
			err:     nil,
			len:     1,
			clauses: Clauses{},
			store: Store{
				data{
					"cars": Models{{"A": "B"}},
				},
			},
		},

		{
			name:    "matching",
			err:     nil,
			len:     1,
			clauses: Clauses{"A": "B"},
			store: Store{
				data{
					"cars": Models{{"A": "B"}},
				},
			},
		},
	}

	for _, tt := range tcases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.store.Select("cars", tt.clauses)
			checkErrorsMatch(t, err, tt.err)

			if len(got) != tt.len {
				t.Fatalf("expected len %v, got %v", tt.len, len(got))
			}
		})
	}
}

// Helper to check if errors are matching.
func checkErrorsMatch(t *testing.T, err, err2 error) {
	t.Helper()

	if err == nil && err2 == nil {
		return
	}

	if errors.Is(err, err) {
		return
	}

	t.Fatalf("expected err %v, got %v", err, err2)
}

func checkModelsMatch(t *testing.T, models, models2 Models) {
	t.Helper()
	for i, v := range models {
		if Clauses(v).Match(models2[i]) {
			continue
		}

		t.Fatalf("expected %v, got %v", models, models2)
	}
}

func checkDataMatches(t *testing.T, exp, act data) {
	t.Helper()
	for table, v := range exp {
		x, ok := act[table]
		if !ok {
			t.Fatalf("expected %v, got %v", exp, act)

			return
		}

		for i, el := range x {
			// Check if the element matches
			if !Clauses(el).Match(v[i]) {
				t.Fatalf("expected %v, got %v", exp, act)

				return
			}
		}
	}
}
