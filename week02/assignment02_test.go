package week02

import "testing"

func TestArray(t *testing.T) {
	exp := [4]string{"max", "tuto", "bombom", "tony"}
	act := make([]string, 4)

	// iterating through the array and appending
	// to act
	for i := range exp {
		act[i] = exp[i]
	}

	for i, v := range act {
		if v != exp[i] {
			t.Errorf("Expected %v, got %v", exp[i], v)
		}
	}

	// For array I'm not sure this test makes much sense
	if len(exp) != len(act) {
		t.Errorf("Expected %v, got %v", len(exp), len(act))
	}
}

func TestSlice(t *testing.T) {
	exp := []string{"max", "tuto", "bombom", "tony"}
	act := []string{}

	// iterating through the array and appending
	// to act
	for i := range exp {
		act = append(act, exp[i])
	}

	for i, v := range act {
		if v != exp[i] {
			t.Errorf("Expected %v, got %v", exp[i], v)
		}
	}

	// For array I'm not sure this test makes much sense
	if len(exp) != len(act) {
		t.Errorf("Expected %v, got %v", len(exp), len(act))
	}
}

func TestMap(t *testing.T) {
	exp := map[int]string{
		1: "max",
		2: "tuto",
		3: "bombom",
		4: "tony",
	}

	act := make(map[int]string)

	for k, v := range exp {
		act[k] = v
	}

	for k := range exp {
		v, ok := act[k]
		if !ok {
			t.Errorf("Expected %v, got %v", exp[k], v)
		}
	}
}
