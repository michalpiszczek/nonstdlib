// This module contains tests for stack.go
//
// Note:
// 	These tests are not ordered by reliance.
// 	Some possible concurrency issues are not covered by this test suite.

package worklist

import (
	"testing"
)

func TestNewEmptyStack(t *testing.T) {
	s := NewStack()

	if s.Size() != 0 || s.Empty() != true {
		t.Error("NewStack with 0 args does not create an empty worklist")
	}
}

func TestInitEmptyStack(t *testing.T) {
	s := &Stack{}
	s.Init()

	if s.Size() != 0 || s.Empty() != true {
		t.Error("NewStack with 0 args does not create an empty worklist")
	}
}

func TestPushSimpleStack(t *testing.T) {
	s := NewStack()

	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	if s.Size() != 10 {
		t.Error("Pushed 10 items but size is not 10")
	}
}

func TestPushAboveInitCapStack(t *testing.T) {
	s := NewStack()

	for i := 0; i < 100; i++ {
		s.Push(i)
	}

	if s.Size() != 100 {
		t.Error("Pushed 100 items but size is not 100")
	}
}

func TestPopEmptyStack(t *testing.T) {
	s := NewStack()

	if w := s.Pop(); w != nil {
		t.Error("Popping off an empty stack did not return nil!")
	}
}

func TestPopManyStack(t *testing.T) {
	s := NewStack()

	for i := 0; i < 100; i++ {
		s.Push(i)
	}

	i := 99
	for !s.Empty() {
		w := s.Pop()
		if w != i {
			t.Error("Popped: ", w, "Expected ", i)
		}
		i -= 1
	}

	if s.Size() != 0 {
		t.Error("After popping of all elements, Size() should be 0.")
	}
}

func TestSizeEmptyStack(t *testing.T) {
	s := NewStack()

	if s.Size() != 0 {
		t.Error("Size of an empty WorkList should be 0.")
	}
}

func TestSizeSimpleStack(t *testing.T) {
	s := NewStack()

	for i := 0; i < 100; i++ {
		s.Push(i)
	}

	if s.Size() != 100 {
		t.Error("Pushed 100 items but size is not 100")
	}
}

func TestEmptyPresentStack(t *testing.T) {
	s := NewStack()

	for i := 0; i < 100; i++ {
		s.Push(i)
	}

	if s.Size() != 100 {
		t.Error("A Stack with 100 items should not be Empty()")
	}
}

func TestEmptyAbsentStack(t *testing.T) {
	s := NewStack()

	if !s.Empty() {
		t.Error("An new Stack should be empty")
	}
}

func TestCopyStack(t *testing.T) {
	s1 := NewStack()

	for i := 0; i < 100; i++ {
		s1.Push(i)
	}

	s2 := s1.Copy()

	for !s1.Empty() {
		if s1.Pop() != s2.Pop() {
			t.Error("Popping off the original and a copy should yeild same work")
		}
	}
}

func TestClearStack(t *testing.T) {
	s := NewStack()

	for i := 0; i < 100; i++ {
		s.Push(i)
	}

	s.Clear()

	if !s.Empty() {
		t.Error("A cleared stack should be empty")
	}
}
