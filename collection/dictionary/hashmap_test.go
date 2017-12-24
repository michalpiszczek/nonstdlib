// This module contains tests for hashmap.go
//
// Note:
//  These tests are not ordered by reliance.
//  Some possible concurrency issues are not covered by this test suite.
//  TODO: cover new behavior

package dictionary

import (
	"github.com/michalpiszczek/nonstdlib/util/test"
	"math/rand"
	"testing"
)


func TestNewEmptyHashMap(t *testing.T) {
	s := NewHashMap()

	test.AssertEqual(t, s.Size(), 0, "New dictionary has size != 0.")
	test.AssertTrue(t, s.Empty(), "New dictionary is not empty.")
}

func TestInitEmptyHashMap(t *testing.T) {
	s := &HashMap{}
	s.Init()

	test.AssertEqual(t, s.Size(), 0, "New dictionary has size != 0.")
	test.AssertTrue(t, s.Empty(), "New dictionary is not empty.")
	test.AssertNonNil(t, s, "New dictionary is nil.")
}

func TestInsertEmptyHashMap(t *testing.T) {
	s := NewHashMap()

	x := compInt{1}

	old1 := s.Insert(x, "hello")

	test.AssertNil(t, old1, "")
	test.AssertEqual(t, s.Size(), 1, "Wrong size after inserting.")
	test.AssertTrue(t, s.Contains(x), "Inserted element missing")
}

func TestInsertSimpleHashMap(t *testing.T) {
	s := NewHashMap()

	x := compInt{5}
	y := compInt{10}
	z := compInt{15}

	old1 := s.Insert(x, "hello")
	old2 := s.Insert(y, "cruel")
	old3 := s.Insert(z, "world")

	test.AssertNil(t, old1, "")
	test.AssertNil(t, old2, "")
	test.AssertNil(t, old3, "")

	test.AssertEqual(t, s.Size(), 3, "Wrong size after inserting.")

	test.AssertTrue(t, s.Contains(x), "Inserted element missing")
	test.AssertTrue(t, s.Contains(y), "Inserted element missing")
	test.AssertTrue(t, s.Contains(z), "Inserted element missing")
}

func TestRemoveChainInOrderHashMap(t *testing.T) {
	s := NewHashMap()

	x := compInt{5}
	y := compInt{10}
	z := compInt{15}

	s.Insert(x, "hello")
	s.Insert(y, "cruel")
	s.Insert(z, "world")

	old1 := s.Remove(x)
	old2 := s.Remove(y)
	old3 := s.Remove(z)

	test.AssertEqual(t, old1, "hello", "")
	test.AssertEqual(t, old2, "cruel", "")
	test.AssertEqual(t, old3, "world", "")

	test.AssertEqual(t, s.Size(), 0, "Wrong size after deleting.")

	test.AssertFalse(t, s.Contains(x), "Removed element present.")
	test.AssertFalse(t, s.Contains(y), "Removed element present.")
	test.AssertFalse(t, s.Contains(z), "Removed element present.")
}

func TestRemoveChainInReverseHashMap(t *testing.T) {
	s := NewHashMap()

	x := compInt{5}
	y := compInt{10}
	z := compInt{15}

	s.Insert(x, "hello")
	s.Insert(y, "cruel")
	s.Insert(z, "world")

	old1 := s.Remove(z)
	old2 := s.Remove(y)
	old3 := s.Remove(x)

	test.AssertEqual(t, old3, "hello", "")
	test.AssertEqual(t, old2, "cruel", "")
	test.AssertEqual(t, old1, "world", "")

	test.AssertEqual(t, s.Size(), 0, "Wrong size after deleting.")

	test.AssertFalse(t, s.Contains(x), "Removed element present.")
	test.AssertFalse(t, s.Contains(y), "Removed element present.")
	test.AssertFalse(t, s.Contains(z), "Removed element present.")
}

func TestLargeRandomLoadHashMap(t *testing.T) {

    s := NewHashMap()

	kvs := make(map[int]int)

	for i := 0; i < 10000; i++ {
		k := rand.Intn(10000)
		temp := compInt{k}
		val := i
		s.Insert(temp, val)
		kvs[k] = val;
	}

	for k, v := range kvs {
		r := s.Locate(compInt{k})
		test.AssertEqual(t, r, v, "Retrieved wrong value.")
		if r != v {
			t.FailNow()
		}
	}
}

func BenchmarkLargeRandomLoadHashMap(t *testing.B) {
    s := NewHashMap()

    kvs := make(map[int]int)

    for i := 0; i < 100000; i++ {
        k := rand.Intn(100000)
        temp := compInt{k}
        val := i
        s.Insert(temp, val)
        kvs[k] = val;
    }

    for k, v := range kvs {
        r := s.Locate(compInt{k})
        if r != v {
            t.FailNow()
        }
    }
}
