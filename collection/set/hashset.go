// This module implements a HashSet, conforming to set.Interface.

package set

import (
	"fmt" // To help with String().
	"github.com/michalpiszczek/nonstdlib/collection"
)

// A HashSet implements set.Interface.
//
// Behavior unspecified if a HashSet is not created using NewHashSet() or
// if HashSet.Init() is not first called on a new &HashSet{}.
//
// By unspecified, I mean I will Fatal() you with a helpful message :)
//
type HashSet struct {
	collection.Base
	m map[interface{}]struct{}
}

// Alias for ease of use.
var present = struct{}{}

// Returns a pointer to a new HashSet containing the given items.
func NewHashSet(items ...interface{}) *HashSet {
	s := &HashSet{}
	s.Init()
	s.Insert(items...)
	return s
}

// Returns a pointer to a new unsafe HashSet containing the given items.
func NewHashSetUnsafe(items ...interface{}) *HashSet {
	s := &HashSet{}
	s.InitUnsafe()
	s.Insert(items...)
	return s
}

func (s *HashSet) Init() {
	s.InitBase()

	s.m = make(map[interface{}]struct{})
}

func (s *HashSet) InitUnsafe() {
	s.InitBaseUnsafe()

	s.m = make(map[interface{}]struct{})
}

func (s *HashSet) Insert(items ...interface{}) {
	s.CheckInit()

	if len(items) == 0 {
		return
	}
	if s.Threadsafe() {
		s.Lock()
		defer s.Unlock()
	}

	for _, item := range items {
        if _, ok := s.m[item]; !ok {
            s.Sizeb += 1
        }
        s.m[item] = present
	}
}

func (s *HashSet) Remove(items ...interface{}) {
	s.CheckInit()

	if len(items) == 0 {
		return
	}

	s.Lockb.Lock()
	defer s.Lockb.Unlock()

	for _, item := range items {
        if _, ok := s.m[item]; ok {
            delete(s.m, item)
            s.Sizeb -= 1
        }
	}
}

func (s *HashSet) Contains(items ...interface{}) bool {
	s.CheckInit()

	if len(items) == 0 {
		return false
	}

	s.Lockb.RLock()
	defer s.RUnlock()

	for _, item := range items {
		if _, ok := s.m[item]; !ok {
			return false
		}
	}
	return true
}

// Returns a pointer to a new Set containing all the items in either this
// Set or the given Set.
func (s *HashSet) Union(o Set) Set {
	s.CheckInit()

	result := s.Copy()

	o.Map(func(item interface{}) bool {
		result.Insert(item)
		return true
	})

	return result
}

// Returns a pointer to a new Set containing all the items in both this Set
// and the given Set.
func (s *HashSet) Intersection(o Set) Set {
	s.CheckInit()

	result := NewHashSet()

	var iter Set = s
	var itee Set = o
	if s.Size() > o.Size() {
		iter = o
		itee = s
	}

	iter.Map(func(item interface{}) bool {
		if itee.Contains(item) {
			result.Insert(item)
		}
		return true
	})

	return result
}

// Returns a pointer to a new Set containing all the items in this Set that
// are not in the given Set.
func (s *HashSet) Difference(o Set) Set {
	s.CheckInit()

	result := s.Copy()

	o.Map(func(item interface{}) bool {
		result.Remove(item)
		return true
	})

	return result
}

// Returns true if this Set and the given Set contain exactly the same items,
// false otherwise.
func (s *HashSet) Equal(o Set) bool {
	s.CheckInit()

	// Use len() instead of Size to avoid deadlock
	if s.Size() != o.Size() {
		return false
	}

	s.RLock()
	defer s.RUnlock()
	equal := true
	o.Map(func(item interface{}) bool {
		_, equal = s.m[item]
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
func (s *HashSet) Subset(o Set) (subset bool, proper bool) {
	s.CheckInit()

	if s.Size() != o.Size() {
		proper = true
	} else {
		proper = false
	}

	subset = true
	s.Map(func(item interface{}) bool {
		subset = o.Contains(item)
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
func (s *HashSet) Superset(o Set) (superset bool, proper bool) {
	s.CheckInit()

	if s.Size() != o.Size() {
		proper = true
	} else {
		proper = false
	}

	superset = true
	o.Map(func(item interface{}) bool {
		superset = s.Contains(item)
		return superset
	})

	return
}

// Returns a pointer to a new Set that is a copy of this Set.
func (s *HashSet) Copy() Set {
	s.CheckInit()

    if s.Threadsafe() {
        return NewHashSet(*s.Slice()...)
    } else {
        return NewHashSetUnsafe(*s.Slice()...)
    }
}

// Attempts to apply the given function to every items in this Set.
// Stops once all elements have been processed, or once the function
// returns false, whichever occurs first.
func (s *HashSet) Map(f func(item interface{}) bool) bool {
	s.CheckInit()

	s.RLock()
	defer s.RUnlock()

	ok := true
	for item := range s.m {
		if ok = f(item); !ok {
			break
		}
	}
	return ok
}

// Returns a slice of all the items in this Set in no particular order.
func (s *HashSet) Slice() *[]interface{} {
	s.CheckInit()

	s.RLock()
	defer s.RUnlock()

	slice := make([]interface{}, 0, len(s.m))

	for item := range s.m {
		slice = append(slice, item)
	}

	return &slice
}

// Removes all items from this Set.
func (s *HashSet) Clear() {
	s.CheckInit()

	s.Lock()
	defer s.Unlock()

	s.m = make(map[interface{}]struct{})
    s.Sizeb = 0
}

func (s *HashSet) String() string {
	return fmt.Sprintf("%v", s.Slice())
}
