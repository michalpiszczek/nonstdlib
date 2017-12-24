// This module implements a TreeSet, conforming to set.Interface.

package set

import (
    "fmt" // To help with String().
    "github.com/michalpiszczek/nonstdlib/collection"
    "github.com/michalpiszczek/nonstdlib/collection/dictionary"
    "log"
)

// A TreeSet implements set.Interface.
//
// Behavior unspecified if a TreeSet is not created using NewTreeSet() or
// if TreeSet.Init() is not first called on a new &TreeSet{}.
//
// By unspecified, I mean I will Fatal() you with a helpful message :)
//
type TreeSet struct {
    collection.Base
    m *dictionary.TreeMap
}

// Returns a pointer to a new TreeSet containing the given items.
func NewTreeSet(items ...interface{}) *TreeSet {
    s := &TreeSet{}
    s.Init()
    s.Insert(items...)
    return s
}

// Returns a pointer to a new unsafe TreeSet containing the given items.
func NewTreeSetUnsafe(items ...interface{}) *TreeSet {
    s := &TreeSet{}
    s.InitUnsafe()
    s.Insert(items...)
    return s
}

func (s *TreeSet) Init() {
    s.InitBase()

    s.m = dictionary.NewTreeMapUnsafe()
}

func (s *TreeSet) InitUnsafe() {
    s.InitBaseUnsafe()

    s.m = dictionary.NewTreeMapUnsafe();
}

func (s *TreeSet) Insert(items ...interface{}) {
    s.CheckInit()

    if len(items) == 0 {
        return
    }
    if s.Threadsafe() {
        s.Lock()
        defer s.Unlock()
    }

    for _, item := range items {
        itemc, ok := item.(collection.Comparer)
        if !ok {
            log.Fatal("Item doesn't implement collection.Comparer.")
        }
        old := s.m.Insert(itemc, true)
        if old == nil {
            s.Sizeb += 1
        }
    }
}

// non-nil apparently...
func (s *TreeSet) Remove(items ...interface{}) {
    s.CheckInit()

    if len(items) == 0 {
        return
    }

    s.Lockb.Lock()
    defer s.Lockb.Unlock()

    for _, item := range items {
        itemc, ok := item.(collection.Comparer)
        if !ok {
            log.Fatal("Item doesn't implement collection.Comparer.")
        }
        old := s.m.Remove(itemc)
        if old != nil {
            s.Sizeb -= 1
        }
    }
}

func (s *TreeSet) Contains(items ...interface{}) bool {
    s.CheckInit()

    if len(items) == 0 {
        return false
    }

    s.Lockb.RLock()
    defer s.RUnlock()

    return s.m.Contains(items...)
}

// Returns a pointer to a new Set containing all the items in either this
// Set or the given Set.
func (s *TreeSet) Union(o Set) Set {
    s.CheckInit()

    result := s.Copy()

    o.Map(func(item interface{}) bool {
        itemc, _  := item.(collection.Comparer)
        result.Insert(itemc)
        return true
    })

    return result
}

// Returns a pointer to a new Set containing all the items in both this Set
// and the given Set.
func (s *TreeSet) Intersection(o Set) Set {
    s.CheckInit()

    result := NewTreeSet()

    var iter Set = s
    var itee Set = o
    if s.Size() > o.Size() {
        iter = o
        itee = s
    }

    iter.Map(func(item interface{}) bool {
        itemc, _  := item.(collection.Comparer)

        if itee.Contains(itemc) {
            result.Insert(itemc)
        }
        return true
    })

    return result
}

// Returns a pointer to a new Set containing all the items in this Set that
// are not in the given Set.
func (s *TreeSet) Difference(o Set) Set {
    s.CheckInit()

    result := s.Copy()

    o.Map(func(item interface{}) bool {
        itemc, _  := item.(collection.Comparer)
        result.Remove(itemc)
        return true
    })

    return result
}

// Returns true if this Set and the given Set contain exactly the same items,
// false otherwise.
func (s *TreeSet) Equal(o Set) bool {
    s.CheckInit()

    // Use len() instead of Size to avoid deadlock
    if s.Size() != o.Size() {
        return false
    }

    s.RLock()
    defer s.RUnlock()
    equal := true
    o.Map(func(item interface{}) bool {
        itemc, _  := item.(collection.Comparer)
        equal = s.Contains(itemc)
        return equal
    })

    return equal
}

// The first bool returned is true if this Set is a subset of the given Set,
// false otherwise. If the first returned bool is true, then the second bool
// will be false if these two sets are equal, true otherwise.
//
// true, true -> s is a proper subset of o
// true, false -> s is equal to o
// false, true -> s is not a subset of o
// false, false -> s is not a subset of o
func (s *TreeSet) Subset(o Set) (subset bool, proper bool) {
    s.CheckInit()

    if s.Size() != o.Size() {
        proper = true
    } else {
        proper = false
    }

    subset = true
    s.Map(func(item interface{}) bool {
        itemc, _  := item.(collection.Comparer)
        subset = o.Contains(itemc)
        return subset
    })

    return
}

// The first bool returned is true if this Set is a superset of the given
// Set, false otherwise. If the first returned bool is true, then the
// second bool will be false if these two sets are equal, true otherwise.
//
// true, true -> s is a proper superset of o
// true, false -> s is equal to o
// false, true -> s is not a superset of o
// false, false -> s is not a superset of o
func (s *TreeSet) Superset(o Set) (superset bool, proper bool) {
    s.CheckInit()

    if s.Size() != o.Size() {
        proper = true
    } else {
        proper = false
    }

    superset = true
    o.Map(func(item interface{}) bool {
        itemc, _  := item.(collection.Comparer)
        superset = s.Contains(itemc)
        return superset
    })

    return
}

// Returns a pointer to a new Set that is a copy of this Set.
func (s *TreeSet) Copy() Set {
    s.CheckInit()

    if s.Threadsafe() {
        return NewTreeSet(*s.Slice()...)
    } else {
        return NewTreeSetUnsafe(*s.Slice()...)
    }
}

// Attempts to apply the given function to every items in this Set.
// Stops once all elements have been processed, or once the function
// returns false, whichever occurs first.
func (s *TreeSet) Map(f func(item interface{}) bool) bool {
    s.CheckInit()

    s.RLock()
    defer s.RUnlock()

    ok := true
    s.m.Map(func(kv interface{}) bool {

        kvc, _ := kv.(*dictionary.KeyValue)
        vc, ok := kvc.Key.(collection.Comparer)
        if !ok {
            log.Fatal("Item doesn't implement collection.Comparer.")
        }
        ok = f(vc)
        return ok
    })

    return ok
}

// Returns a slice of all the items in this Set in no particular order.
func (s *TreeSet) Slice() *[]interface{} {
    s.CheckInit()

    s.RLock()
    defer s.RUnlock()

    slice := make([]interface{}, 0, s.Sizeb)

    s.m.Map(func(kv interface{}) bool {
        kvc, _ := kv.(*dictionary.KeyValue)
        vc, ok := kvc.Key.(collection.Comparer)
        if !ok {
            log.Fatal("Item doesn't implement collection.Comparer.")
        }
        slice = append(slice, vc)
        return true
    })

    return &slice
}

// Removes all items from this Set.
func (s *TreeSet) Clear() {
    s.CheckInit()

    s.Lock()
    defer s.Unlock()

    s.m = dictionary.NewTreeMapUnsafe()
    s.Sizeb = 0
}

func (s *TreeSet) String() string {
    return fmt.Sprintf("%v", s.Slice())
}
