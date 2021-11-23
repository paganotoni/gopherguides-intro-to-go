package week08

import (
	"testing"
	"time"
)

func TestMaterialDuration(t *testing.T) {
	tcases := []struct {
		material Material
		duration time.Duration
	}{
		{Material("iron"), 4 * time.Second},
		{Material("steel"), 5 * time.Second},
		{Material("wood"), 4 * time.Second},
		{Material("cotton"), 6 * time.Second},
	}

	for _, tcase := range tcases {
		if tcase.material.Duration() != tcase.duration {
			t.Errorf("%s.Duration() = %d, want %d", tcase.material, tcase.material.Duration(), tcase.duration)
		}
	}
}

func TestMaterialsDuration(t *testing.T) {
	tcases := []struct {
		material Materials
		duration time.Duration
	}{
		{Materials{"iron": 1, "steel": 1, "wood": 1, "cotton": 1}, 19000000},
		{Materials{"iron": 2, "steel": 1, "wood": 1, "cotton": 1}, 23000000},
		{Materials{}, 0},
	}

	for _, tcase := range tcases {
		if tcase.material.Duration() != tcase.duration {
			t.Errorf("%s.Duration() = %d, want %d", tcase.material, tcase.material.Duration(), tcase.duration)
		}
	}
}
