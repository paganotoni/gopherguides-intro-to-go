package week04

import (
	"bytes"
	"testing"
)

func TestEntertain(t *testing.T) {
	venue := &Venue{}

	t.Run("SetuperKindOfBand", func(t *testing.T) {
		log := bytes.NewBuffer([]byte(``))
		venue.Log = log

		venue.Entertain(20_000, []Entertainer{
			&SpoiledBand{},
		})

		expectedContent := []string{
			`Spoiled Band has completed setup`,
			`Spoiled Band has performed for 20000 people`,
		}

		for _, content := range expectedContent {
			if bytes.Contains(log.Bytes(), []byte(content)) {
				continue
			}

			t.Fatalf("expected log to contain %s", content)
		}
	})

	t.Run("TeardownKindOfBand", func(t *testing.T) {
		log := bytes.NewBuffer([]byte(``))
		venue.Log = log

		venue.Entertain(20_000, []Entertainer{
			&Troublemakers{},
		})

		expectedContent := []string{
			`Troublemakers has performed for 20000 people`,
			`Troublemakers has completed teardown`,
		}

		for _, content := range expectedContent {
			if bytes.Contains(log.Bytes(), []byte(content)) {
				continue
			}

			t.Fatalf("expected log to contain %s", content)
		}
	})

}
