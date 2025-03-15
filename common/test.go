package common

import (
	"reflect"
	"testing"
)

func AssertInt(t *testing.T, expected, actual int, message string) {
	if expected != actual {
		t.Errorf("%s (expected: %v, actual %v)", message, expected, actual)
	}
}

func AssertArray(t *testing.T, expected, actual []int, message string) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%s (expected: %v, actual: %v)", message, expected, actual)
	}
}
