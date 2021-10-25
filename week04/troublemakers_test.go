package week04

import (
	"bytes"
	"testing"
)

func TestTroubleMakers(t *testing.T) {
	log := bytes.NewBuffer([]byte{})
	venue := &Venue{
		Log: log,
	}

	venue.Entertain(100_000, []Entertainer{
		&Troublemakers{},
	})

	expectedContent := []string{
		"Troublemakers: Entering the stage",
		"Troublemakers: Performs its first song",
		"Troublemakers: Breaks a Camera with the guitar",
		"Troublemakers: Breaks a Guitar, pieces hit an attendant",
		"Troublemakers: Throw a chair to the crowd",
		"Troublemakers: Leaving the stage",
		"Troublemakers: Ran Away after the concert",
		"Troublemakers: Police pursue them for the chair-trowing event",
		"Troublemakers: Detained",
		"Troublemakers: Venue has to pay fine",
		"teardown error: could not solve Troublemakers issues, legal department still working",
	}

	for _, content := range expectedContent {
		if bytes.Contains(log.Bytes(), []byte(content)) {
			continue
		}

		t.Errorf("expected log to contain %s", content)
	}

}
