package verify

import (
	"reflect"
	"testing"
)

func Equals(t *testing.T, a, b interface{}, msg ...interface{}) {
	t.Helper()

	if reflect.DeepEqual(a, b) {
		return
	}

	t.Errorf("%s is not equal to %s", a, b)
}

func Equalsf(t *testing.T, a, b interface{}, msg string) {
	t.Helper()

	if reflect.DeepEqual(a, b) {
		return
	}

	t.Errorf(msg, a, b)
}
