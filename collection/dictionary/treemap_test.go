// This module contains tests for TreeMap.go
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

// A wrapper for int that will implement Comparer.
type compInt struct {
    i int
}

// Implement Comparer on compInt.
func (m compInt) Compare(o interface{}) int {
    oc, _ := o.(compInt)

    if m.i < oc.i {
        return -1
    } else if m.i > oc.i {
        return 1
    } else {
        return 0
    }
}

func TestOrdering(t *testing.T) {
    s := NewTreeMap()

    list := make([]compInt, 0)

    list = append(list, compInt{1})
    list = append(list, compInt{3})
    list = append(list, compInt{7})
    list = append(list, compInt{2})
    list = append(list, compInt{5})

    for _, num := range list {
        s.Insert(num, 0)
    }

    keys := []int{1, 2, 3, 5, 7}
    for i, kv := range *s.Slice() {
        k, _ := kv.(*KeyValue).Key.(compInt)
        test.AssertEqual(t, k.i, keys[i], "Items returned out of order!")
    }

}

func TestNewEmptyTreeMap(t *testing.T) {
    s := NewTreeMap()

    test.AssertEqual(t, s.Size(), 0, "New dictionary has size != 0.")
    test.AssertTrue(t, s.Empty(), "New dictionary is not empty.")
}

func TestInitEmptyTreeMap(t *testing.T) {
    s := &TreeMap{}
    s.Init()

    test.AssertEqual(t, s.Size(), 0, "New dictionary has size != 0.")
    test.AssertTrue(t, s.Empty(), "New dictionary is not empty.")
    test.AssertNonNil(t, s, "New dictionary is nil.")
}

func TestInsertEmptyTreeMap(t *testing.T) {
    s := NewTreeMap()

    x := compInt{1}

    old1 := s.Insert(x, "hello")

    test.AssertNil(t, old1, "")
    test.AssertEqual(t, s.Size(), 1, "Wrong size after inserting.")
    test.AssertTrue(t, s.Contains(x), "Inserted element missing")
}

func TestInsertSimpleTreeMap(t *testing.T) {
    s := NewTreeMap()

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

func TestRemoveChainInOrderTreeMap(t *testing.T) {
    s := NewTreeMap()

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

func TestRemoveChainInReverseTreeMap(t *testing.T) {
    s := NewTreeMap()

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

func TestLargeRandomLoadTreeMap(t *testing.T) {
    s := NewTreeMap()

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

func BenchmarkLargeRandomLoadTreeMap(t *testing.B) {
    s := NewTreeMap()

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
