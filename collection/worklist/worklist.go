// Package worklist defines the interface for WorkLists. A WorkList
// is an ordered Collection of work.
//
package worklist

import (
	"github.com/michalpiszczek/nonstdlib/collection"
)

// Defines the interface for WorkLists. WorkLists are ordered Collections
// of work, supporting the additional operations given here.
//
type WorkList interface {
	collection.Collection

	// Pushes the given item of work to this WorkList.
	//
	// Panics if this WorkList has not been initialized.
	Push(work interface{})

	// Pops and returns the next item of work in this WorkList.
	// Returns nil if there is no work remaining in this WorkList.
	//
	// Panics if this WorkList has not been initialized.
	Pop() interface{}

	// Returns a new, initialized WorkList, that contains the same items
	// as this WorkList.
	//
	// Not necessarily a deep copy.
	//
	// Panics if this WorkList has not been initialized.
	Copy() WorkList
}

// ****************************************************************************
//
//	Convenience functions to be used as so: worklist = function(worklist).
//
// ****************************************************************************

// Returns the given WorkList, and:
//
// Adds the given item of work to the given WorkList.
//
// Panics if the given WorkList has not been initialized.
func Push(w WorkList, work interface{}) (this WorkList) {
	this = w
	w.Push(work)
	return
}

// Returns the given WorkList, and:
//
// Removes and returns the next item of work in the given WorkList.
// Returns nil if there is no work remaining in the given WorkList.
//
// Panics if the given WorkList has not been initialized.
func Pop(w WorkList) (this WorkList, work interface{}) {
	this = w
	work = w.Pop()
	return
}

// Returns the given WorkList, and:
//
// Returns to a new WorkList that contains the same items
// as the given WorkList.
//
// Will result in a Panic if the given WorkList has not been initialized.
func Copy(c WorkList) (this WorkList, cpy WorkList) {
	this = c
	cpy = c.Copy()
	return
}
