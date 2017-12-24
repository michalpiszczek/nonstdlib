// This module contains tests for hashset.go
//
// Note:
// 	These tests are not ordered by reliance.
// 	Some possible concurrency issues are not covered by this test suite.

package set

import (
	"sync" // for sync.WaitGroup
	"testing"
)

func TestNewEmptyHashSet(t *testing.T) {
	s := NewHashSet()

	if s.Size() != 0 || s.Empty() != true {
		t.Error("NewHashSet with 0 args does not create an empty set!")
	}
}

func TestNewSimpleHashSet(t *testing.T) {
	s := NewHashSet("Hello", "Cruel", "World")

	if s.Size() != 3 || s.Empty() != false {
		t.Error("3 unique args does not create a set with 3 items!")
	}

	if s.Contains("Hello", "Cruel", "World") == false {
		t.Error("Items missing!")
	}
}

func TestNewUniqueHashSet(t *testing.T) {
	s := NewHashSet("Hello", "Cruel", "World", "World")

	if s.Size() != 3 || s.Empty() != false {
		t.Error("3 unique args does not create a set with 3 items!")
	}

	if s.Contains("Hello", "Cruel", "World") == false {
		t.Error("Items missing!")
	}
}

func TestNewUniqueStructHashSet(t *testing.T) {

    type Car struct {
        x int
        y int
    }

    s := NewHashSet(Car{1, 2}, Car{2, 1}, Car{1, 2})

    if s.Size() != 2 || s.Empty() != false {
        t.Error("3 unique args does not create a set with 3 items!")
    }

    if s.Contains(Car{1, 2}, Car{2, 1}) == false {
        t.Error("Items missing!")
    }
}

func BenchmarkNewHashSet(t *testing.B) {
	_ = NewHashSet("Hello", "Cruel", "World")

}

func TestInitEmptyHashSet(t *testing.T) {
	s := &HashSet{}

	s.Init()

	if s.Size() != 0 || s.Empty() != true {
		t.Error("NewHashSet with 0 args does not create an empty set!")
	}
}

func TestInitSimpleHashSet(t *testing.T) {
	s := &HashSet{}
	s.Init()
	s.Insert("Hello", "Cruel", "World")

	if s.Size() != 3 || s.Empty() != false {
		t.Error("3 unique args does not create a set with 3 items!")
	}

	if s.Contains("Hello", "Cruel", "World") == false {
		t.Error("Items missing!")
	}
}

func TestInitUniqueHashSet(t *testing.T) {
	s := &HashSet{}
	s.Init()
	s.Insert("Hello", "Cruel", "World", "World")

	if s.Size() != 3 || s.Empty() != false {
		t.Error("3 unique args does not create a set with 3 items!")
	}

	if s.Contains("Hello", "Cruel", "World") == false {
		t.Error("Items missing!")
	}
}

func BenchmarkInitHashSet(t *testing.B) {
	s := &HashSet{}
	s.Init()
	s.Insert("Hello", "Cruel", "World")
}

func TestInsertEmptyHashSet(t *testing.T) {
	s := NewHashSet()
	s.Insert()

	if !s.Empty() {
		t.Error("Inserting nothing to an empty set should leave it empty.")
	}
}

func TestInsertSimpleHashSet(t *testing.T) {
	s := NewHashSet()
	s.Insert(1, 2, 3, 3)

	if !s.Contains(1, 2, 3) && s.Size() != 3 {
		t.Error("Insert {1, 2, 3, 3} to empty should result in {1, 2, 3}"+" not: ", s)
	}
}

func TestInsertUniqueStructHashSet(t *testing.T) {

    type Car struct {
        x int
        y int
    }

    s := NewHashSet()
    s.Insert(Car{1, 2}, Car{2, 1}, Car{1, 2})

    if s.Size() != 2 || s.Empty() != false {
        t.Error("3 unique args does not create a set with 3 items!")
    }

    if s.Contains(Car{1, 2}, Car{2, 1}) == false {
        t.Error("Items missing!")
    }
}

func TestRemoveHashSet(t *testing.T) {
	s := NewHashSet(1, 2, 3)
	s.Remove(1, 2)

	if s.Contains(1, 2) && !s.Contains(3) && s.Size() != 1 {
		t.Error("{1, 2, 3}.Remove(1, 2) should be {3}"+" not: ", s)
	}
}

// Kind of implicitly tested by a lot of other tests...
func TestContainsPresentHashSet(t *testing.T) {
	s := NewHashSet(1, 2, 3)

	if !s.Contains(1, 2) {
		t.Error("{1, 2} should countain {1, 2}")
	}
}

func TestContainsAbsentHashSet(t *testing.T) {
	s := NewHashSet(1, 2, 3)

	if s.Contains(4) {
		t.Error("{1, 2} should not countain {4}")
	}
}

func TestSizeHashSet(t *testing.T) {
	s := NewHashSet(1, 2, 3)

	if s.Size() != 3 {
		t.Error("{1, 2, 3} should have Size 3")
	}
}

func TestEmptyTrueHashSet(t *testing.T) {
	s := NewHashSet()

	if !s.Empty() {
		t.Error("{ } should be empty.")
	}
}

func TestEmptyFalseHashSet(t *testing.T) {
	s := NewHashSet(1)

	if s.Empty() {
		t.Error("{1} should not be empty")
	}

}

func TestUnionEmptyHashSet(t *testing.T) {
	s1 := NewHashSet()
	s2 := NewHashSet()

	u := s1.Union(s2)

	if !u.Empty() {
		t.Error("The union of two empty sets should be empty"+" not: ", u)
	}
}

func TestUnionIdenticalHashSet(t *testing.T) {
	s1 := NewHashSet(1, 2, 3)
	s2 := NewHashSet(1, 2, 3)

	u := s1.Union(s2)

	if !u.Equal(s1) || !u.Equal(s2) {
		t.Error("The union of {1, 2, 3} and {1, 2, 3} should be {1, 2, 3}"+
			" not: ", u)
	}
}

func TestUnionSimpleHashSet(t *testing.T) {
	s1 := NewHashSet(1, 2, 3)
	s2 := NewHashSet(3, 4, 5)

	u := s1.Union(s2)

	if !u.Equal(NewHashSet(1, 2, 3, 4, 5)) {
		t.Error("The union of {1, 2, 3} and {3, 4, 5} should be {1, 2, 3, 4, 5}"+
			" not: ", u)
	}
}

func TestIntersectionEmptyHashSet(t *testing.T) {
	s1 := NewHashSet()
	s2 := NewHashSet()

	i := s1.Intersection(s2)

	if !i.Empty() {
		t.Error("The Intersection of two empty sets should be empty"+
			" not: ", i)
	}
}

func TestIntersectionIdenticalHashSet(t *testing.T) {
	s1 := NewHashSet(1, 2, 3)
	s2 := NewHashSet(1, 2, 3)

	i := s1.Intersection(s2)

	if !i.Equal(s1) || !i.Equal(s2) {
		t.Error("The Intersection of {1, 2, 3} and {1, 2, 3} should be {1, 2, 3}"+
			" not: ", i)
	}
}

func TestIntersectionSimpleHashSet(t *testing.T) {
	s1 := NewHashSet(1, 2, 3)
	s2 := NewHashSet(3, 4, 5)

	i := s1.Intersection(s2)

	if !i.Equal(NewHashSet(3)) {
		t.Error("The Intersection of {1, 2, 3} and {3, 4, 5} should be {3}"+
			" not: ", i)
	}
}

func TestDifferenceEmptyHashSet(t *testing.T) {
	s1 := NewHashSet()
	s2 := NewHashSet()

	d := s1.Difference(s2)

	if !d.Empty() {
		t.Error("The Difference of two empty sets should be empty"+
			" not: ", d)
	}
}

func TestDifferenceIdenticalHashSet(t *testing.T) {
	s1 := NewHashSet(1, 2, 3)
	s2 := NewHashSet(1, 2, 3)

	d := s1.Difference(s2)

	if !d.Empty() {
		t.Error("The Difference of {1, 2, 3} and {1, 2, 3} should be { }"+
			" not: ", d)
	}
}

func TestDifferenceSimple1HashSet(t *testing.T) {
	s1 := NewHashSet(1, 2, 3)
	s2 := NewHashSet(3, 4, 5)

	d := s1.Difference(s2)

	if !d.Equal(NewHashSet(1, 2)) {
		t.Error("The Difference of {1, 2, 3} and {3, 4, 5} should be {1, 2}"+
			" not: ", d)
	}
}

func TestDifferenceSimple2HashSet(t *testing.T) {
	s1 := NewHashSet(1, 2, 3)
	s2 := NewHashSet(3, 4, 5)

	d := s2.Difference(s1)

	if !d.Equal(NewHashSet(4, 5)) {
		t.Error("The Difference of {3, 4, 5} and {1, 2, 3} should be {4, 5}"+
			" not: ", d)
	}
}

func TestEqualEmptyHashSet(t *testing.T) {
	s1 := NewHashSet()
	s2 := NewHashSet()

	if !s1.Equal(s2) {
		t.Error("Two empty sets should be equal")
	}
}

func TestEqualReflexiveHashSet(t *testing.T) {
	s1 := NewHashSet(1)

	if !s1.Equal(s1) {
		t.Error("A set should be equal to itself")
	}
}

func TestEqualSymmetricHashSet(t *testing.T) {
	s1 := NewHashSet(1, 2, 3)
	s2 := NewHashSet(1, 2, 3)

	if !s1.Equal(s2) || !s2.Equal(s1) {
		t.Error("Equality should be symmetric")
	}
}

func TestEqualNonEmptyHashSet(t *testing.T) {
	s1 := NewHashSet(1)
	s2 := NewHashSet(1)

	if !s1.Equal(s2) {
		t.Error("{1} should equal {1}")
	}
}

func TestEqualModifiedHashSet(t *testing.T) {
	s1 := NewHashSet(1)
	s2 := NewHashSet(1)

	if !s1.Equal(s2) {
		t.Error("{1} should equal {1}")
	}

	s2.Insert(2)

	if s1.Equal(s2) {
		t.Error("{1} should equal {1, 2}")
	}
}

func TestEqualSubsetHashSet(t *testing.T) {
	s1 := NewHashSet(1)
	s2 := NewHashSet(1, 2)

	if s1.Equal(s2) || s2.Equal(s1) {
		t.Error("{1} should not equal {1, 2}")
	}
}

// Intended to reveal potential deadlock issues...
func TestEqualConcurrencyHashSet(t *testing.T) {
	s1 := NewHashSet(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		16, 17, 18, 19, 20)
	s2 := NewHashSet(5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18)

	results := make(chan bool)

	var wg sync.WaitGroup

	// Uncertain of a better way to do this.
	wg.Add(200000)

	go func() {
		for i := 0; i < 100000; i++ {
			go func() {
				defer wg.Done()
				e := s2.Equal(s1)
				results <- !e
			}()

			go func() {
				defer wg.Done()
				e := s1.Equal(s2)
				results <- !e
			}()
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	ok := true
	for x := range results {
		ok = x && ok
	}

	if ok == false {
		t.Error("Problem with concurrent Equal calls!")
	}
}

func TestSubsetSimpleHashSet(t *testing.T) {
	s1 := NewHashSet(1, 2, 3)
	s2 := NewHashSet(2)

	if s, _ := s2.Subset(s1); !s {
		t.Error("{2} should be a subset of {1, 2, 3}")
	}
}

func TestSubsetReflexiveHashSet(t *testing.T) {
	s1 := NewHashSet(1, 2, 3)

	if s, p := s1.Subset(s1); !s || p {
		t.Error("{1, 2, 3} should be a non-proper subset of itself")
	}
}

func TestSubsetProperHashSet(t *testing.T) {
	s1 := NewHashSet(1, 2, 3)
	s2 := NewHashSet(2)

	if s, p := s2.Subset(s1); !s || !p {
		t.Error("{2} should be a proper subset of {1, 2, 3}")
	}

	s3 := s1.Copy()

	if _, p := s1.Subset(s3); p {
		t.Error("{1, 2, 3} should not be a proper subset of {1, 2, 3}")
	}
}

// Intended to reveal deadlocks...
func TestSubsetConcurrencyHashSet(t *testing.T) {
	s1 := NewHashSet(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		16, 17, 18, 19, 20)
	s2 := NewHashSet(5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18)

	results := make(chan bool)

	var wg sync.WaitGroup

	// Uncertain of a better way to do this.
	wg.Add(200000)

	go func() {
		for i := 0; i < 100000; i++ {
			go func() {
				defer wg.Done()
				s, p := s2.Subset(s1)
				results <- s && p
			}()

			go func() {
				defer wg.Done()
				s, _ := s1.Subset(s2)
				results <- !s
			}()
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	ok := true
	for x := range results {
		ok = x && ok
	}

	if ok == false {
		t.Error("Problem with concurrent Subset calls!")
	}
}

func TestSupersetSimpleHashSet(t *testing.T) {
	s1 := NewHashSet(1, 2, 3)
	s2 := NewHashSet(2)

	if s, _ := s2.Superset(s1); s {
		t.Error("{2} should not be a Superset of {1, 2, 3}")
	}
}

func TestSupersetReflexiveHashSet(t *testing.T) {
	s1 := NewHashSet(1, 2, 3)

	if s, p := s1.Superset(s1); !s || p {
		t.Error("{1, 2, 3} should be a non-proper Superset of itself")
	}
}

func TestSupersetProperHashSet(t *testing.T) {
	s1 := NewHashSet(1, 2, 3)
	s2 := NewHashSet(2)

	if s, _ := s2.Superset(s1); s {
		t.Error("{2} should not be a proper Superset of {1, 2, 3}")
	}

	s3 := s1.Copy()

	if _, p := s1.Superset(s3); p {
		t.Error("{1, 2, 3} should not be a proper Superset of {1, 2, 3}")
	}
}

// Intended to reveal deadlocks...
func TestSupersetConcurrencyHashSet(t *testing.T) {
	s1 := NewHashSet(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		16, 17, 18, 19, 20)
	s2 := NewHashSet(5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18)

	results := make(chan bool)

	var wg sync.WaitGroup

	// Uncertain of a better way to do this.
	wg.Add(200000)

	go func() {
		for i := 0; i < 100000; i++ {
			go func() {
				defer wg.Done()
				s, p := s1.Superset(s2)
				results <- s && p
			}()

			go func() {
				defer wg.Done()
				s, _ := s2.Superset(s1)
				results <- !s
			}()
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	ok := true
	for x := range results {
		ok = x && ok
	}

	if ok == false {
		t.Error("Problem with concurrent Superset calls!")
	}
}

func TestMapRaceCarHashSet(t *testing.T) {

	// Vrooom...
	type raceCar struct {
		num  int
		fast bool
	}

	s := NewHashSet()

	for i := 0; i < 100; i++ {
		s.Insert(&raceCar{num: i, fast: false})
	}

	s.Map(func(item interface{}) bool {
		rc, _ := item.(*raceCar)
		rc.fast = true
		return true
	})

	s.Map(func(item interface{}) bool {
		rc, _ := item.(*raceCar)
		if rc.fast != true {
			t.Error("At least one race car isn't fast!")
		}
		return true
	})
}

func TestSliceLengthHashSet(t *testing.T) {
	s := NewHashSet(1, 2, 3)

	if len(*s.Slice()) != 3 {
		t.Error("The Slice of {1, 2, 3} should have len() 3")
	}
}

func TestSliceContainsHashSet(t *testing.T) {
	s := NewHashSet(1, 2, 3)

	for _, i := range *s.Slice() {
		if !s.Contains(i) {
			t.Error("A Set should countain all the elements in its Slice!")
			// Bar parallel cases...
		}
	}
}

func TestSliceConstructionEqualityHashSet(t *testing.T) {
	s1 := NewHashSet(1, 2, 3)
	s2 := NewHashSet(*s1.Slice()...)

	if !s1.Equal(s2) {
		t.Error("A Set should equal a Set made from its Slice()")
	}
}

func TestCopyEqualOriginalHashSet(t *testing.T) {
	s1 := NewHashSet(1, 2, 3)
	s2 := s1.Copy()

	if !s1.Equal(s2) {
		t.Error("A copy should equal the original")
	}
}

func TestCopyModifyHashSet(t *testing.T) {
	s1 := NewHashSet(1, 2, 3)
	s2 := s1.Copy()

	s2.Remove(2)

	if !s1.Contains(1, 2, 3) {
		t.Error("Remove from a copy should not modify the original Set")
	}

	if !s2.Contains(1, 3) {
		t.Error("Removing {2} from a copy of {1, 2, 3} should yield {2, 3}")
	}
}

func TestClearEmptyHashSet(t *testing.T) {
	s := NewHashSet()
	s.Clear()

	if !s.Empty() {
		t.Error("A cleared Empty Set should be empty")
	}
}

func TestClearToEmptyHashSet(t *testing.T) {
	s := NewHashSet(1, 2, 3)
	s.Clear()

	if !s.Empty() {
		t.Error("{1, 2, 3}.Clear() should yield {}")
	}
}
