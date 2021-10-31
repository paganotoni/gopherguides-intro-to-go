package demo

import (
	"reflect"
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
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
			if !reflect.DeepEqual(terr, tt.err) {
				t.Errorf("got %v, want %v", terr, tt.err)
			}

			if !reflect.DeepEqual(tmod, tt.models) {
				t.Errorf("got %v, want %v", tmod, tt.models)
			}
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
				t.Errorf("got %v, want %v", count, tt.count)
			}

			if !reflect.DeepEqual(terr, tt.err) {
				t.Errorf("got %v, want %v", terr, tt.err)
			}

		})
	}
}

func TestStoreInsert(t *testing.T) {
	t.Parallel()

	store := Store{}
	store.Insert("car", Model{"A": "B"})

	len, _ := store.Len("car")
	if len != 1 {
		t.Errorf("got len %v, want %v", len, 1)
	}

	mods, err := store.All("car")
	if err != nil {
		t.Errorf("got err %v, want nil", err)
	}

	if !reflect.DeepEqual(mods, Models{{"A": "B"}}) {
		t.Errorf("got %v, want %v", mods, Models{{"A": "B"}})
	}

	store.Insert("car", Model{"C": "D"})

	len, _ = store.Len("car")
	if len != 2 {
		t.Errorf("got len %v, want %v", len, 1)
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
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("got err %v, want %v", err, tt.err)
			}

			if len(got) != tt.len {
				t.Errorf("got len %v, want %v", len(got), tt.len)
			}
		})
	}
}
