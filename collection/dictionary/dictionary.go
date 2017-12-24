// Package dictionary defines the interface for Dictionaries. A Dictionary
// is a Collection of unique (key, value) pairs.
//
package dictionary

import (
	"github.com/michalpiszczek/nonstdlib/collection"
)

// Dictionary.Map() maps over pointers to KeyValue structs.
//
// Dictionary.Slice() returns a slice of pointers to KeyValue structs.
//
type KeyValue struct {
	Key   interface{}
	Value interface{}
}

// Defines the interface for Dictionaries. Dictionaries are Collections
// of unique (key, value) pairs, supporting the operations given here.
//
type Dictionary interface {

	collection.Collection

	// Inserts the given value associated with the given key into this
	// Dictionary. Returns the previous value associated with that key,
	// or nil, if none existed.
	//
	// Panics if the given key or value are nil.
	// Panics if this Dictionary has not been initialized.
	Insert(key interface{}, value interface{}) interface{}

	// Returns the value associated with the given key, If there is no
	// value associated with the given key, returns nil.
	//
	// Panics if the given key is nil.
	// Panics if this Dictionary has not been initialized.
	Locate(key interface{}) interface{}

	// Removes and returns the value associated with the given key from
	// this Dictionary. If there is no value associated with the given key,
	// returns nil.
	//
	// Panics if the given key is nil.
	// Panics if this Dictionary has not been initialized.
	Remove(key interface{}) interface{}

	// Returns true if all the given keys have entries in this Dictionary,
	// false otherwise.
	//
	// Panics if any of the given keys are nil.
	// Panics if this Dictionary has not been initialized.
	Contains(keys ...interface{}) bool

	// Returns a new, initialized Dictionary, that contains the same items
	// as this Dictionary.
	//
	// Not necessarily a deep copy.
	//
	// Panics if this Dictionary has not been initialized.
	Copy() Dictionary
}

// ****************************************************************************
//
//  Convenience functions to be used as so: dictionary = function(dictionary).
//
// ****************************************************************************

// Inserts the given value associated with the given key into this
// Dictionary. Returns the previous value associated with that key,
// or nil, if none existed.
//
// Panics if the given key or value are nil.
// Panics if this Dictionary has not been initialized.
func Insert(d Dictionary, k interface{}, v interface{}) (this Dictionary, old interface{}) {
	this = d
	old = d.Insert(k, v)
	return
}

// Returns the value associated with the given key, If there is no
// value associated with the given key, returns nil.
//
// Panics if the given key is nil.
// Panics if this Dictionary has not been initialized.
func Locate(d Dictionary, k interface{}) (this Dictionary, value interface{}) {
	this = d
	value = d.Locate(k)
	return
}

// Removes and returns the value associated with the given key from
// this Dictionary. If there is no value associated with the given key,
// returns nil.
//
// Panics if the given key is nil.
// Panics if this Dictionary has not been initialized.
func Remove(d Dictionary, k interface{}) (this Dictionary, value interface{}) {
	this = d
	value = d.Remove(k)
	return
}

// Returns true if all the given keys have entries in this Dictionary,
// false otherwise.
//
// Panics if any of the given keys are nil.
// Panics if this Dictionary has not been initialized.
func Contains(d Dictionary, ks ...interface{}) (this Dictionary, contains bool) {
	this = d
	contains = d.Contains(ks...)
	return
}

// Returns a new, initialized Dictionary, that contains the same items
// as this Dictionary.
//
// Not necessarily a deep copy.
//
// Panics if this Dictionary has not been initialized.
func Copy(d Dictionary) (this Dictionary, cpy Dictionary) {
	this = d
	cpy = d.Copy()
	return
}
