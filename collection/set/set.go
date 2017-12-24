// This module defines the Set interface, and convenience functions
// useable on all types implementing Set.

// Package collection defines the Set interface. Sets are Collections
// of unique items.
//
package set

import (
	"github.com/michalpiszczek/nonstdlib/collection"
)

// Defines the Set interface. Sets store unique items, and
// implement the operations defined below.
//
// Sets are also Collections, and implement the behavior
// expected of all Collections.
//
type Set interface {
	collection.Collection

	// Inserts the given items into this Set.
	//
	// Panics if this Set has not been initialized.
	Insert(items ...interface{})

	// Removes the given items from this Set.
	//
	// Panics if this Set has not been initialized.
	Remove(items ...interface{})

	// Returns true if all the given items are in the Set, false otherwise.
	//
	// Panics if this Set has not been initialized.
	Contains(items ...interface{}) bool

	// Returns a new Set containing all the items that are in this Set or the
	// other given Set.
	//
	// Panics if this Set or the given other Set have not been initialized.
	Union(o Set) Set

	// Returns a new Set containing all the items in this Set and the other
	// given Set.
	//
	// Panics if this Set or the given other Set have not been initialized.
	Intersection(o Set) Set

	// Returns a new Set containing all the items in this Set that are not
	// in the other given Set.
	//
	// Panics if this Set or the given other Set have not been initialized.
	Difference(o Set) Set

	// Returns true if this Set and the other given Set have the exact same
	// contents, false otherwise.
	//
	// Panics if this Set or the given other Set have not been initialized.
	Equal(o Set) bool

	// The first bool returned is true if this Set is a subset of the given Set,
	// false otherwise. If the first returned bool is true, then the second bool
	// will be false if these two sets are equal, true otherwise.
	//
	// true, true -> s is a proper subset of o
	// true, false -> s is equal to o
	// false, true -> s is not a subset of o
	// false, false -> s is not a subset of o
	//
	// Panics if this Set or the given other Set have not been initialized.
	Subset(o Set) (subset bool, proper bool)

	// The first bool returned is true if this Set is a superset of the given
	// Set, false otherwise. If the first returned bool is true, then the
	// second bool will be false if these two sets are equal, true otherwise.
	//
	// true, true -> s is a proper superset of o
	// true, false -> s is equal to o
	// false, true -> s is not a superset of o
	// false, false -> s is not a superset of o
	//
	// Panics if this Set or the given other Set have not been initialized.
	Superset(o Set) (superset bool, proper bool)

	// Returns a new, initialized Set, that contains the same items
	// as this Set.
	//
	// Not necessarily a deep copy.
	//
	// Panics if this Set has not been initialized.
	Copy() Set
}

// ****************************************************************************
//
//  Convenience functions to be used as so: set = function(set).
//
// ****************************************************************************

// Returns the given Set, and:
//
// Inserts the given items into this Set.
//
// Panics if the given Set has not been initialized.
func Insert(s Set, items ...interface{}) (this Set) {
	this = s
	s.Insert(items...)
	return
}

// Returns the given Set, and:
//
// Removes the given items from this Set.
//
// Panics if the given Set has not been initialized.
func Remove(s Set, items ...interface{}) (this Set) {
	this = s
	s.Remove(items...)
	return
}

// Returns the given Set, and:
//
// Returns true if all the given items are in the Set, false otherwise.
//
// Panics if the given Set has not been initialized.
func Contains(s Set, items ...interface{}) (this Set, contains bool) {
	this = s
	contains = s.Contains(items...)
	return
}

// Returns the given Sets, and:
//
// Returns a new Set containing all the items that are in either of the
// given Sets.
//
// Panics if either given Set has not been initialized.
func Union(s Set, o Set) (this Set, other Set, union Set) {
	this = s
	other = o
	union = s.Union(o)
	return

}

// Returns the given Sets, and:
//
// Returns a new Set containing the items that are in both of the of
// the given Sets.
//
// Panics if either given Set has not been initialized.
func Intersection(s Set, o Set) (this Set, other Set, intersection Set) {
	this = s
	other = o
	intersection = s.Intersection(o)
	return

}

// Returns the given Sets, and:
//
// Returns a new Set containing all the items in the first given Set, but
// not in the other given Set.
//
// Panics if either given Set has not been initialized.
func Difference(s Set, o Set) (this Set, other Set, difference Set) {
	this = s
	other = o
	difference = s.Difference(o)
	return
}

// Returns the given Sets, and:
//
// Returns true if this Set and the other given Set have the exact same
// contents, false otherwise.
//
// Panics if either given Set has not been initialized.
func Equal(s Set, o Set) (this Set, other Set, equal bool) {
	this = s
	other = o
	equal = s.Equal(o)
	return

}

// Returns the given Sets, and:
//
// The first bool returned is true if the first given Set is a subset of the second
// given Set, false otherwise. If the first returned bool is true, then the second bool
// will be false if the two given sets are equal, true otherwise.
//
// true, true   -> s is a proper subset of o
// true, false  -> s is equal to o
// false, true  -> s is not a subset of o
// false, false -> s is not a subset of o
//
// Panics if either given Set has not been initialized.
func Subset(s Set, o Set) (this Set, other Set, subset bool, proper bool) {
	this = s
	other = o
	subset, proper = s.Subset(o)
	return
}

// Returns the given Sets, and:
//
// The first bool returned is true if the first given Set is a superset of
// the second given Set, false otherwise. If the first returned bool is true,
// then the second bool will be false if the two given sets are equal,
// true otherwise.
//
// true, true   -> s is a proper superset of o
// true, false  -> s is equal to o
// false, true  -> s is not a superset of o
// false, false -> s is not a superset of o
//
// Panics if either given Set has not been initialized.
func Superset(s Set, o Set) (this Set, other Set, superset bool, proper bool) {
	this = s
	other = o
	superset, proper = s.Superset(o)
	return
}

// Returns the given Set, and:
//
// Returns to a new Set that contains the same items
// as the given Set.
//
// Will result in a Panic if the given Set has not been initialized.
func Copy(c Set) (this Set, copy Set) {
	this = c
	copy = c.Copy()
	return
}
