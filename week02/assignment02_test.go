package week02

import "testing"

func TestArray(t *testing.T) {
	exp := [4]string{"max", "tuto", "bombom", "tony"}
	act := make([]string, 0)

	// iterating through the array and appending
	// to act
	for i := range exp {
		new := make([]string, len(act)+1)
		copy(new, act)

		act = new
		act[len(act)-1] = exp[i]
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

func TestSlicePointer(t *testing.T) {
	orig := []string{"A", "B", "C"}
	lor := len(orig)

	modif := func(dat []string) {
		dat[0] = "X"
		dat = append(dat, "D")
	}

	modif(orig)
	// Length still the same
	if lor != len(orig) {
		t.Errorf("Expected %v, got %v", lor, len(orig))
	}

	if orig[0] != "X" {
		t.Errorf("Expected %v, got %v", "X", orig[0])
	}

	modip := func(dat *[]string) {
		(*dat)[0] = "W"
		*dat = append(*dat, "D")
	}

	modip(&orig)
	// Length not the same
	if lor == len(orig) {
		t.Errorf("Expected length to be %[1]v, got %[1]v", len(orig))
	}

	if orig[0] != "W" {
		t.Errorf("Expected %v, got %v", "W", orig[0])
	}

}
