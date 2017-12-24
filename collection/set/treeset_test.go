// This module contains tests for TreeSet.go
//
// Note:
//  These tests are not ordered by reliance.
//  Some possible concurrency issues are not covered by this test suite.

package set

import (
    "sync" // for sync.WaitGroup
    "testing"
    "github.com/michalpiszczek/nonstdlib/util/math"
)

type S struct {
    v string
}

func (m S) Compare(o interface{}) int {
    oc, _ := o.(S)

    if m.v < oc.v {
        return -1
    } else if m.v > oc.v {
        return 1
    } else {
        return 0
    }
}

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

type Car struct {
    x int
    y int
}

func (m Car) Compare(o interface{}) int {
    oc, _ := o.(Car)

    if m.x == oc.x {
        return math.Signum(m.y - oc.y)
    } else {
        return math.Signum(m.x - oc.x)
    }
}

// Vrooom...
type raceCar struct {
    num  int
    fast bool
}

func (m raceCar) Compare(o interface{}) int {
    oc, _ := o.(raceCar)

    return math.Signum(m.num - oc.num)
}

func TestNewEmptyTreeSet(t *testing.T) {
    s := NewTreeSet()

    if s.Size() != 0 || s.Empty() != true {
        t.Error("NewTreeSet with 0 args does not create an empty set!")
    }
}

func TestNewSimpleTreeSet(t *testing.T) {
    s := NewTreeSet(S{"Hello"}, S{"Cruel"}, S{"World"})

    if s.Size() != 3 || s.Empty() != false {
        t.Error("3 unique args does not create a set with 3 items!")
    }

    if s.Contains(S{"Hello"}, S{"Cruel"}, S{"World"}) == false {
        t.Error("Items missing!")
    }
}

func TestNewUniqueTreeSet(t *testing.T) {
    s := NewTreeSet(S{"Hello"}, S{"Cruel"}, S{"World"}, S{"World"})

    if s.Size() != 3 || s.Empty() != false {
        t.Error("3 unique args does not create a set with 3 items!")
    }

    if s.Contains(S{"Hello"}, S{"Cruel"}, S{"World"}) == false {
        t.Error("Items missing!")
    }
}



func TestNewUniqueStructTreeSet(t *testing.T) {

    s := NewTreeSet(Car{1, 2}, Car{2, 1}, Car{1, 2})

    if s.Size() != 2 || s.Empty() != false {
        t.Error("3 unique args does not create a set with 3 items!")
    }

    if s.Contains(Car{1, 2}, Car{2, 1}) == false {
        t.Error("Items missing!")
    }
}

func BenchmarkNewTreeSet(t *testing.B) {
    _ = NewTreeSet(S{"Hello"}, S{"Cruel"}, S{"World"})

}

func TestInitEmptyTreeSet(t *testing.T) {
    s := &TreeSet{}

    s.Init()

    if s.Size() != 0 || s.Empty() != true {
        t.Error("NewTreeSet with 0 args does not create an empty set!")
    }
}

func TestInitSimpleTreeSet(t *testing.T) {
    s := &TreeSet{}
    s.Init()
    s.Insert(S{"Hello"}, S{"Cruel"}, S{"World"})

    if s.Size() != 3 || s.Empty() != false {
        t.Error("3 unique args does not create a set with 3 items!")
    }

    if s.Contains(S{"Hello"}, S{"Cruel"}, S{"World"}) == false {
        t.Error("Items missing!")
    }
}

func TestInitUniqueTreeSet(t *testing.T) {
    s := &TreeSet{}
    s.Init()
    s.Insert(S{"Hello"}, S{"Cruel"}, S{"World"}, S{"World"})

    if s.Size() != 3 || s.Empty() != false {
        t.Error("3 unique args does not create a set with 3 items!")
    }

    if s.Contains(S{"Hello"}, S{"Cruel"}, S{"World"}) == false {
        t.Error("Items missing!")
    }
}

func BenchmarkInitTreeSet(t *testing.B) {
    s := &TreeSet{}
    s.Init()
    s.Insert(S{"Hello"}, S{"Cruel"}, S{"World"})
}

func TestInsertEmptyTreeSet(t *testing.T) {
    s := NewTreeSet()
    s.Insert()

    if !s.Empty() {
        t.Error("Inserting nothing to an empty set should leave it empty.")
    }
}

func TestInsertSimpleTreeSet(t *testing.T) {
    s := NewTreeSet()
    s.Insert(compInt{1}, compInt{2}, compInt{3}, compInt{3})

    if !s.Contains(compInt{1}, compInt{2}, compInt{3}) && s.Size() != 3 {
        t.Error("Insert {1, 2, 3, 3} to empty should result in {1, 2, 3}"+" not: ", s)
    }
}

func TestInsertUniqueStructTreeSet(t *testing.T) {

    s := NewTreeSet()
    s.Insert(Car{1, 2}, Car{2, 1}, Car{1, 2})

    if s.Size() != 2 || s.Empty() != false {
        t.Error("3 unique args does not create a set with 3 items!")
    }

    if s.Contains(Car{1, 2}, Car{2, 1}) == false {
        t.Error("Items missing!")
    }
}

func TestRemoveTreeSet(t *testing.T) {
    s := NewTreeSet(compInt{1}, compInt{2}, compInt{3})
    s.Remove(compInt{1}, compInt{2})

    if s.Contains(compInt{1}, compInt{2}) && !s.Contains(compInt{3}) && s.Size() != 1 {
        t.Error("{1, 2, 3}.Remove(1, 2) should be {3}"+" not: ", s)
    }
}

// Kind of implicitly tested by a lot of other tests...
func TestContainsPresentTreeSet(t *testing.T) {
    s := NewTreeSet(compInt{1}, compInt{2}, compInt{3})

    if !s.Contains(compInt{1}, compInt{2}) {
        t.Error("{1, 2} should countain {1, 2}")
    }
}

func TestContainsAbsentTreeSet(t *testing.T) {
    s := NewTreeSet(compInt{1}, compInt{2}, compInt{3})

    if s.Contains(compInt{4}) {
        t.Error("{1, 2} should not countain {4}")
    }
}

func TestSizeTreeSet(t *testing.T) {
    s := NewTreeSet(compInt{1}, compInt{2}, compInt{3})

    if s.Size() != 3 {
        t.Error("{1, 2, 3} should have Size 3")
    }
}

func TestEmptyTrueTreeSet(t *testing.T) {
    s := NewTreeSet()

    if !s.Empty() {
        t.Error("{ } should be empty.")
    }
}

func TestEmptyFalseTreeSet(t *testing.T) {
    s := NewTreeSet(compInt{1})

    if s.Empty() {
        t.Error("{1} should not be empty")
    }

}

func TestUnionEmptyTreeSet(t *testing.T) {
    s1 := NewTreeSet()
    s2 := NewTreeSet()

    u := s1.Union(s2)

    if !u.Empty() {
        t.Error("The union of two empty sets should be empty"+" not: ", u)
    }
}

func TestUnionIdenticalTreeSet(t *testing.T) {
    s1 := NewTreeSet(compInt{1}, compInt{2}, compInt{3})
    s2 := NewTreeSet(compInt{1}, compInt{2}, compInt{3})

    u := s1.Union(s2)

    if !u.Equal(s1) || !u.Equal(s2) {
        t.Error("The union of {1, 2, 3} and {1, 2, 3} should be {1, 2, 3}"+
            " not: ", u)
    }
}

func TestUnionSimpleTreeSet(t *testing.T) {
    s1 := NewTreeSet(compInt{1}, compInt{2}, compInt{3})
    s2 := NewTreeSet(compInt{3}, compInt{4}, compInt{5})

    u := s1.Union(s2)

    if !u.Equal(NewTreeSet(compInt{1}, compInt{2}, compInt{3}, compInt{4}, compInt{5})) {
        t.Error("The union of {1, 2, 3} and {3, 4, 5} should be {1, 2, 3, 4, 5}"+
            " not: ", u)
    }
}

func TestIntersectionEmptyTreeSet(t *testing.T) {
    s1 := NewTreeSet()
    s2 := NewTreeSet()

    i := s1.Intersection(s2)

    if !i.Empty() {
        t.Error("The Intersection of two empty sets should be empty"+
            " not: ", i)
    }
}

func TestIntersectionIdenticalTreeSet(t *testing.T) {
    s1 := NewTreeSet(compInt{1}, compInt{2}, compInt{3})
    s2 := NewTreeSet(compInt{1}, compInt{2}, compInt{3})

    i := s1.Intersection(s2)

    if !i.Equal(s1) || !i.Equal(s2) {
        t.Error("The Intersection of {1, 2, 3} and {1, 2, 3} should be {1, 2, 3}"+
            " not: ", i)
    }
}

func TestIntersectionSimpleTreeSet(t *testing.T) {
    s1 := NewTreeSet(compInt{1}, compInt{2}, compInt{3})
    s2 := NewTreeSet(compInt{3}, compInt{4}, compInt{5})

    i := s1.Intersection(s2)

    if !i.Equal(NewTreeSet(compInt{3})) {
        t.Error("The Intersection of {1, 2, 3} and {3, 4, 5} should be {3}"+
            " not: ", i)
    }
}

func TestDifferenceEmptyTreeSet(t *testing.T) {
    s1 := NewTreeSet()
    s2 := NewTreeSet()

    d := s1.Difference(s2)

    if !d.Empty() {
        t.Error("The Difference of two empty sets should be empty"+
            " not: ", d)
    }
}

func TestDifferenceIdenticalTreeSet(t *testing.T) {
    s1 := NewTreeSet(compInt{1}, compInt{2}, compInt{3})
    s2 := NewTreeSet(compInt{1}, compInt{2}, compInt{3})

    d := s1.Difference(s2)

    if !d.Empty() {
        t.Error("The Difference of {1, 2, 3} and {1, 2, 3} should be { }"+
            " not: ", d)
    }
}

func TestDifferenceSimple1TreeSet(t *testing.T) {
    s1 := NewTreeSet(compInt{1}, compInt{2}, compInt{3})
    s2 := NewTreeSet(compInt{3}, compInt{4}, compInt{5})

    d := s1.Difference(s2)

    if !d.Equal(NewTreeSet(compInt{1}, compInt{2})) {
        t.Error("The Difference of {1, 2, 3} and {3, 4, 5} should be {1, 2}"+
            " not: ", d)
    }
}

func TestDifferenceSimple2TreeSet(t *testing.T) {
    s1 := NewTreeSet(compInt{1}, compInt{2}, compInt{3})
    s2 := NewTreeSet(compInt{3}, compInt{4}, compInt{5})

    d := s2.Difference(s1)

    if !d.Equal(NewTreeSet(compInt{4}, compInt{5})) {
        t.Error("The Difference of {3, 4, 5} and {1, 2, 3} should be {4, 5}"+
            " not: ", d)
    }
}

func TestEqualEmptyTreeSet(t *testing.T) {
    s1 := NewTreeSet()
    s2 := NewTreeSet()

    if !s1.Equal(s2) {
        t.Error("Two empty sets should be equal")
    }
}

func TestEqualReflexiveTreeSet(t *testing.T) {
    s1 := NewTreeSet(compInt{1})

    if !s1.Equal(s1) {
        t.Error("A set should be equal to itself")
    }
}

func TestEqualSymmetricTreeSet(t *testing.T) {
    s1 := NewTreeSet(compInt{1}, compInt{2}, compInt{3})
    s2 := NewTreeSet(compInt{1}, compInt{2}, compInt{3})

    if !s1.Equal(s2) || !s2.Equal(s1) {
        t.Error("Equality should be symmetric")
    }
}

func TestEqualNonEmptyTreeSet(t *testing.T) {
    s1 := NewTreeSet(compInt{1})
    s2 := NewTreeSet(compInt{1})

    if !s1.Equal(s2) {
        t.Error("{1} should equal {1}")
    }
}

func TestEqualModifiedTreeSet(t *testing.T) {
    s1 := NewTreeSet(compInt{1})
    s2 := NewTreeSet(compInt{1})

    if !s1.Equal(s2) {
        t.Error("{1} should equal {1}")
    }

    s2.Insert(compInt{2})

    if s1.Equal(s2) {
        t.Error("{1} should equal {1, 2}")
    }
}

func TestEqualSubsetTreeSet(t *testing.T) {
    s1 := NewTreeSet(compInt{1})
    s2 := NewTreeSet(compInt{1}, compInt{2})

    if s1.Equal(s2) || s2.Equal(s1) {
        t.Error("{1} should not equal {1, 2}")
    }
}

// Intended to reveal potential deadlock issues...
func TestEqualConcurrencyTreeSet(t *testing.T) {
    s1 := NewTreeSet(compInt{1}, compInt{2}, compInt{3}, compInt{4}, compInt{5}, compInt{6}, compInt{7}, compInt{8},
        compInt{9}, compInt{10}, compInt{11}, compInt{12}, compInt{13}, compInt{14}, compInt{15},
        compInt{16}, compInt{17}, compInt{18}, compInt{19}, compInt{20})
    s2 := NewTreeSet(compInt{5}, compInt{6}, compInt{7}, compInt{8},
        compInt{9}, compInt{10}, compInt{11}, compInt{12}, compInt{13}, compInt{14}, compInt{15},
        compInt{16}, compInt{17}, compInt{18})

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

func TestSubsetSimpleTreeSet(t *testing.T) {
    s1 := NewTreeSet(compInt{1}, compInt{2}, compInt{3})
    s2 := NewTreeSet(compInt{2})

    if s, _ := s2.Subset(s1); !s {
        t.Error("{2} should be a subset of {1, 2, 3}")
    }
}

func TestSubsetReflexiveTreeSet(t *testing.T) {
    s1 := NewTreeSet(compInt{1}, compInt{2}, compInt{3})

    if s, p := s1.Subset(s1); !s || p {
        t.Error("{1, 2, 3} should be a non-proper subset of itself")
    }
}

func TestSubsetProperTreeSet(t *testing.T) {
    s1 := NewTreeSet(compInt{1}, compInt{2}, compInt{3})
    s2 := NewTreeSet(compInt{2})

    if s, p := s2.Subset(s1); !s || !p {
        t.Error("{2} should be a proper subset of {1, 2, 3}")
    }

    s3 := s1.Copy()

    if _, p := s1.Subset(s3); p {
        t.Error("{1, 2, 3} should not be a proper subset of {1, 2, 3}")
    }
}

// Intended to reveal deadlocks...
func TestSubsetConcurrencyTreeSet(t *testing.T) {
    s1 := NewTreeSet(compInt{1}, compInt{2}, compInt{3}, compInt{4}, compInt{5}, compInt{6}, compInt{7}, compInt{8},
        compInt{9}, compInt{10}, compInt{11}, compInt{12}, compInt{13}, compInt{14}, compInt{15},
        compInt{16}, compInt{17}, compInt{18}, compInt{19}, compInt{20})
    s2 := NewTreeSet(compInt{5}, compInt{6}, compInt{7}, compInt{8},
        compInt{9}, compInt{10}, compInt{11}, compInt{12}, compInt{13}, compInt{14}, compInt{15},
        compInt{16}, compInt{17}, compInt{18})
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

func TestSupersetSimpleTreeSet(t *testing.T) {
    s1 := NewTreeSet(compInt{1}, compInt{2}, compInt{3})
    s2 := NewTreeSet(compInt{2})

    if s, _ := s2.Superset(s1); s {
        t.Error("{2} should not be a Superset of {1, 2, 3}")
    }
}

func TestSupersetReflexiveTreeSet(t *testing.T) {
    s1 := NewTreeSet(compInt{1}, compInt{2}, compInt{3})

    if s, p := s1.Superset(s1); !s || p {
        t.Error("{1, 2, 3} should be a non-proper Superset of itself")
    }
}

func TestSupersetProperTreeSet(t *testing.T) {
    s1 := NewTreeSet(compInt{1}, compInt{2}, compInt{3})
    s2 := NewTreeSet(compInt{2})

    if s, _ := s2.Superset(s1); s {
        t.Error("{2} should not be a proper Superset of {1, 2, 3}")
    }

    s3 := s1.Copy()

    if _, p := s1.Superset(s3); p {
        t.Error("{1, 2, 3} should not be a proper Superset of {1, 2, 3}")
    }
}

// Intended to reveal deadlocks...
func TestSupersetConcurrencyTreeSet(t *testing.T) {
    s1 := NewTreeSet(compInt{1}, compInt{2}, compInt{3}, compInt{4}, compInt{5}, compInt{6}, compInt{7}, compInt{8},
        compInt{9}, compInt{10}, compInt{11}, compInt{12}, compInt{13}, compInt{14}, compInt{15},
        compInt{16}, compInt{17}, compInt{18}, compInt{19}, compInt{20})
    s2 := NewTreeSet(compInt{5}, compInt{6}, compInt{7}, compInt{8},
        compInt{9}, compInt{10}, compInt{11}, compInt{12}, compInt{13}, compInt{14}, compInt{15},
        compInt{16}, compInt{17}, compInt{18})

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



func TestMapRaceCarTreeSet(t *testing.T) {



    s := NewTreeSet()

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

func TestSliceLengthTreeSet(t *testing.T) {
    s := NewTreeSet(compInt{1}, compInt{2}, compInt{3})

    if len(*s.Slice()) != 3 {
        t.Error("The Slice of {1, 2, 3} should have len() 3")
    }
}

func TestSliceContainsTreeSet(t *testing.T) {
    s := NewTreeSet(compInt{1}, compInt{2}, compInt{3})

    for _, i := range *s.Slice() {
        if !s.Contains(i) {
            t.Error("A Set should countain all the elements in its Slice!")
            // Bar parallel cases...
        }
    }
}

func TestSliceConstructionEqualityTreeSet(t *testing.T) {
    s1 := NewTreeSet(compInt{1}, compInt{2}, compInt{3})
    s2 := NewTreeSet(*s1.Slice()...)

    if !s1.Equal(s2) {
        t.Error("A Set should equal a Set made from its Slice()")
    }
}

func TestCopyEqualOriginalTreeSet(t *testing.T) {
    s1 := NewTreeSet(compInt{1}, compInt{2}, compInt{3})
    s2 := s1.Copy()

    if !s1.Equal(s2) {
        t.Error("A copy should equal the original")
    }
}

func TestCopyModifyTreeSet(t *testing.T) {
    s1 := NewTreeSet(compInt{1}, compInt{2}, compInt{3})
    s2 := s1.Copy()

    s2.Remove(compInt{2})

    if !s1.Contains(compInt{1}, compInt{2}, compInt{3}) {
        t.Error("Remove from a copy should not modify the original Set")
    }

    if !s2.Contains(compInt{1}, compInt{3}) {
        t.Error("Removing {2} from a copy of {1, 2, 3} should yield {2, 3}")
    }
}

func TestClearEmptyTreeSet(t *testing.T) {
    s := NewTreeSet()
    s.Clear()

    if !s.Empty() {
        t.Error("A cleared Empty Set should be empty")
    }
}

func TestClearToEmptyTreeSet(t *testing.T) {
    s := NewTreeSet(compInt{1}, compInt{2}, compInt{3})
    s.Clear()

    if !s.Empty() {
        t.Error("{1, 2, 3}.Clear() should yield {}")
    }
}
