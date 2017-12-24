// This module contains tests for Queue.go
//
// Note:
// 	These tests are not ordered by reliance.
// 	Some possible concurrency issues are not covered.
//  This test suite was written for a different implementation and interface,
//      and does not cover all current behaviors.

package worklist

import (
	"github.com/michalpiszczek/nonstdlib/util/test"
	"testing"
)

func TestNewEmptyQueue(t *testing.T) {
	q := NewQueue()

	test.AssertEqual(t, q.Size(), 0, "A new Queue should have size 0.")
	test.AssertTrue(t, q.Empty(), "A new Queue should be empty.")
	test.AssertNonNil(t, q, "A new Queue should not be nil.")
}

func TestInitEmptyQueue(t *testing.T) {
	q := &Queue{}
	q.Init()

	if q.Size() != 0 || q.Empty() != true {
		t.Error("NewQueue with 0 args does not create an empty worklist")
	}
}

func TestPushSimpleQueue(t *testing.T) {
	q := NewQueue()

	for i := 0; i < 10; i++ {
		q.Push(i)
	}

	if q.Size() != 10 {
		t.Error("Pushed 10 items but size is not 10")
	}
}

func TestPushAboveInitCapQueue(t *testing.T) {
	q := NewQueue()

	for i := 0; i < 100; i++ {
		q.Push(i)
	}

	if q.Size() != 100 {
		t.Error("Pushed 100 items but size is not 100")
	}
}

func TestPopEmptyQueue(t *testing.T) {
	q := NewQueue()

	if w := q.Pop(); w != nil {
		t.Error("Popping off an empty Queue did not return nil!")
	}
}

func TestPopManyQueue(t *testing.T) {
	q := NewQueue()

	for i := 0; i < 100; i++ {
		q.Push(i)
	}

	i := 0
	for !q.Empty() {
		w := q.Pop()
		if w != i {
			t.Error("Popped: ", w, "Expected ", i)
		}
		i += 1
	}

	if q.Size() != 0 {
		t.Error("After popping of all elements, Size() should be 0.")
	}
}

func TestSizeEmptyQueue(t *testing.T) {
	q := NewQueue()

	if q.Size() != 0 {
		t.Error("Size of an empty WorkList should be 0.")
	}
}

func TestSizeSimpleQueue(t *testing.T) {
	q := NewQueue()

	for i := 0; i < 100; i++ {
		q.Push(i)
	}

	if q.Size() != 100 {
		t.Error("Pushed 100 items but size is not 100")
	}
}

func TestEmptyPresentQueue(t *testing.T) {
	q := NewQueue()

	for i := 0; i < 100; i++ {
		q.Push(i)
	}

	if q.Size() != 100 {
		t.Error("A Queue with 100 items should not be Empty()")
	}
}

func TestEmptyAbsentQueue(t *testing.T) {
	q := NewQueue()

	if !q.Empty() {
		t.Error("An new Queue should be empty")
	}
}

func TestCopyQueue(t *testing.T) {
	q1 := NewQueue()

	for i := 0; i < 100; i++ {
		q1.Push(i)
	}

	q2 := q1.Copy()

	for !q1.Empty() {
		if q1.Pop() != q2.Pop() {
			t.Error("Popping off the original and a copy should yeild same work")
		}
	}
}

func TestClearQueue(t *testing.T) {
	q := NewQueue()

	for i := 0; i < 100; i++ {
		q.Push(i)
	}

	q.Clear()

	if !q.Empty() {
		t.Error("A cleared Queue should be empty")
	}
}
