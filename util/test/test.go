// Test provides some basic test utilities on top, to be used
// on top of Go's existing "testing" package.
package test

import (
	"testing"
)

// Compares a and b for equality. If they are not equal, will Errorf() using
// the given testing.T and print a and b, followed by the given msg.
//
// b should be the expected value, a the actual value.
func AssertEqual(t *testing.T, a interface{}, b interface{}, msg string) {
	if a != b {
		t.Errorf("Expected %#v, got %#v. " + msg, b, a)
	}
}

// Compares a and b for equality. If they are equal, will Errorf() using
// the given testing.T and print a and b, followed by the given msg.
//
// b should be the expected value, a the actual value.
func AssertNotEqual(t *testing.T, a interface{}, b interface{}, msg string) {
	if a == b {
		t.Errorf("Expected %#v, got %#v. " + msg, b, a)
	}
}

// If a is false, will Errorf() using the given *testing.T and print the
// given message.
func AssertTrue(t *testing.T, a bool, msg string) {
	if !a {
		t.Errorf(msg)
	}
}

// If a is true, will Errorf() using the given *testing.T and print the
// given message.
func AssertFalse(t *testing.T, a bool, msg string) {
	if a {
		t.Errorf(msg)
	}
}

// If a is non nil, will Errorf() using the given *testing.T, printing
// a and the given message.
func AssertNil(t *testing.T, a interface{}, msg string) {
	if a != nil {
		t.Errorf("Expected nil, got %#v. " + msg, a)
	}
}

// If a is nil, will Errorf() will Errorf() using the given *testing.T,
// and print the given message.
func AssertNonNil(t *testing.T, a interface{}, msg string) {
	if a == nil {
		t.Errorf("Expected non-nil. " + msg)
	}
}
