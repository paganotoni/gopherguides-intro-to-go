package week04

import (
	"bytes"
	"testing"
)

func TestSpoiledBand(t *testing.T) {
	log := bytes.NewBuffer([]byte{})
	venue := &Venue{
		Log: log,
	}

	venue.Entertain(1_000, []Entertainer{
		&SpoiledBand{},
	})

	expectedContent := []string{
		"Spoiled Band: Preparing for the show\n",
		"Spoiled Band: Artists want flowers",
		"Spoiled Band: Wants a Carl Sagan Picture",
		"Spoiled Band: They want a 73F degree room, working on it",
		"Spoiled Band: Finding an air-hockey table for the band",
		"Spoiled Band: Adding 3 Mangos, 3 Hawaiian Papaya, 6 Bananas and 3 Peaches",
		"Spoiled Band: Performing 3 only songs they have\n",
		"Spoiled Band: Singing first song\n",
		"Spoiled Band: Singing seccond song\n",
		"Spoiled Band: Singing (with performance difficulties) third song\n",
		`Spoiled Band has completed setup`,
		`Spoiled Band has performed for 1000 people`,
	}

	for _, content := range expectedContent {
		if bytes.Contains(log.Bytes(), []byte(content)) {
			continue
		}

		t.Fatalf("expected log to contain %s", content)
	}
}
