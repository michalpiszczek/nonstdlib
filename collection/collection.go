// This module defines the Collection interface, and convenience functions
// usable on all types implementing Collection.
//

// Package collection defines the Collection interface. Collections
// store items and are thread-safe by default.
//
package collection

import (
	"fmt" // For Stringer.
	"log"
	"sync"
)

// Defines the Comparer interface. Types implementing this interface
// must be able to be ordered, as if supporting <, > and =.
//
// Right now TreeMaps use this.
type Comparer interface {

	// Returns -1 if this Comparer is less than the given Comparer,
	// 0 if they are equal, and 1 otherwise.
	//
	// Fatal()s if the given object is not of the same type as This.
	Compare(o interface{}) int
}

// Defines the Collection interface. Collections store items, and
// implement the operations defined below.
//
// Collections are internally thread-safe by default, but allow for
// direct management of their thread-safety should that be desired.
// See InitUnsafe() below.
//
// Collections must be initialized through their Init() or InitUnsafe()
// method before use, otherwise, all other methods will Panic.
//
type Collection interface {

	// Initializes this Collection. This Collection will be thread-safe,
	// meaning it will manage its own thread-safety.
	//
	// Panics if this Collection has already been initialized.
	Init()

	// Initializes this Collection. This Collection will be thread-unsafe,
	// meaning it will require manual management of its thread-safety using
	// the exposed Lock(), Unlock(), RLock() and RUnlock() methods below.
	//
	// Panics if this Collection has already been initialized.
	InitUnsafe()

	// Returns the number of items in this Collection.
	//
	// Panics if this Collection has not been initialized.
	Size() int

	// Returns true if this Collection holds no items.
	//
	// Panics if this Collection has not been initialized.
	Empty() bool

	// Attempts to apply the given function to every item in this Collection.
	// Stops once all items have been processed, or once the given function
	// returns false, whichever occurs first.
	//
	// Returns true if the function was applied to every item, false otherwise.
	//
	// Should items be mutated, no guarantees are given about the behavior
	// of this Collection thereafter.
	//
	// Panics if this Collection has not been initialized.
	Map(f func(interface{}) bool) bool

	// Returns a pointer to a slice of all the items in this Collection.
	//
	// Panics if this Collection has not been initialized.
	Slice() *[]interface{}

	//
	// TODO: reason about lack of covariance, and decide if this really
	// is something you want to support.
	//

	// // Returns a new, initialized Collection, that is a copy of this
	// // Collection.
	// //
	// // Not necessarily a deep copy.
	// //
	// // Panics if this Collection has not been initialized.
	// Copy() Collection

	// Clears all items from this Collection, making it equivalent to a newly
	// initialized Collection of the same type.
	//
	// Panics if this Collection has not been initialized.
	Clear()

	// Returns true if this Collection is thread-safe. A thread-safe
	// Collection manages its own resources to ensure thread-safety.
	// This is the default behavior for all Collections.
	//
	// Returns false if this Collection is thread-unsafe. A thread-
	// unsafe Collection does not manage its own resources to ensure
	// thread-safety, and instead expects clients to do so though the
	// exposed Lock(), Unlock(), RLock() and RUnlock() methods.
	//
	// Panics if this Collection has not been initialized.
	Threadsafe() bool

	// Attempts to acquire the lock on this Collection for both reading and
	// writing, blocking to do so.
	//
	// Lock() should only be called directly if this Collection is not managing
	// its own thread-safety, that is, when this Collection.ThreadSafe() returns
	// false. Calls to Lock() while this Collection.ThreadSafe() returns true
	// will Panic.
	//
	// Panics if this Collection has not been initialized.
	Lock()

	// Releases the lock on this Collection for both reading and writing.
	//
	// Unlock() should only be called directly if this Collection is not managing
	// its own thread-safety, that is, when this Collection.ThreadSafe() returns
	// false. Calls to Unlock() while this Collection.ThreadSafe() returns true
	// will Panic.
	//
	// Panics if this Collection has not been initialized.
	Unlock()

	// Attempts to acquire the lock on this Collection for reading, blocking
	// to do so.
	//
	// RLock() should only be called directly if this Collection is not managing
	// its own thread-safety, that is, when this Collection.ThreadSafe() returns
	// false. Calls to RLock() while this Collection.ThreadSafe() returns true
	// will Panic.
	//
	// Panics if this Collection has not been initialized.
	RLock()

	// Releases the lock on this Collection for reading.
	//
	// RUnlock() should only be called directly if this Collection is not managing
	// its own thread-safety, that is, when this Collection.ThreadSafe() returns
	// false. Calls to RUnlock() while this Collection.ThreadSafe() returns true
	// will Panic.
	//
	// Panics if this Collection has not been initialized.
	RUnlock()

	// Collections implement Stringer, meaning, this Collection.String() should
	// return a meaningful string representation of this Collection.
	fmt.Stringer
}

// ****************************************************************************
//
//	collections.Base. All Collections should embed this struct.
//
// ****************************************************************************

type Base struct {
	init  bool
	Sizeb int

	threadsafe bool
	Lockb      sync.RWMutex
}

func (b *Base) InitBase() {
	if b.init {
		log.Panic("Cannot initialize an already initialized Collection.")
	}
	b.Sizeb = 0
	b.init = true
}

func (b *Base) InitBaseUnsafe() {
	if b.init {
		log.Panic("Cannot initialize an already initialized Collection.")
	}
	b.Sizeb = 0
	b.threadsafe = false
	b.init = true
}

func (b *Base) CheckInit() {
	if !b.init {
		log.Panic("Collection not initialized!")
	}
}

func (b *Base) Size() int {
	if !b.init {
		log.Panic("Cannot call Size() on an uninitialized Collection.")
	}

	if b.threadsafe {
		b.Lockb.RLock()
		defer b.Lockb.RUnlock()
	}

	return b.Sizeb
}

func (b *Base) Empty() bool {
	if !b.init {
		log.Panic("Cannot call Empty() on an uninitialized Collection.")
	}

	if b.threadsafe {
		b.Lockb.RLock()
		defer b.Lockb.RUnlock()
	}

	return b.Sizeb == 0
}

func (b *Base) Threadsafe() bool {
	if !b.init {
		log.Panic("Cannot call ThreadSafe() on an uninitialized Collection.")
	}

	return b.threadsafe
}

func (b *Base) Lock() {
	if !b.init {
		log.Panic("Cannot call Lock() on an uninitialized Collection.")
	}

	if b.threadsafe {
		log.Panic("Cannot call Lock() on a thread-safe Collection.")
	}

	b.Lockb.Lock()
}

func (b *Base) Unlock() {
	if !b.init {
		log.Panic("Cannot call Unlock() on an uninitialized Collection.")
	}

	if b.threadsafe {
		log.Panic("Cannot call UnLock() on a thread-safe Collection.")
	}

	b.Lockb.Unlock()
}

func (b *Base) RLock() {
	if !b.init {
		log.Panic("Cannot call RLock() on an uninitialized Collection.")
	}

	if b.threadsafe {
		log.Panic("Cannot call RLock() on a thread-safe Collection.")
	}

	b.Lockb.RLock()
}

func (b *Base) RUnlock() {
	if !b.init {
		log.Panic("Cannot call RUnlock() on an uninitialized Collection.")
	}

	if b.threadsafe {
		log.Panic("Cannot call RUnlock() on a thread-safe Collection.")
	}

	b.Lockb.RUnlock()
}

// ****************************************************************************
//
//	Convenience functions to be used as so: collection = function(collection).
//
// ****************************************************************************

// Returns the given Collection, and:
//
// Initializes this Collection. The given Collection will be thread-safe,
// meaning it will manage its own thread-safety.
//
// Will result in a Panic if the given Collection has been initialized.
func Init(c Collection) (this Collection) {
	this = c
	c.Init()
	return c
}

// Returns the given Collection, and:
//
// Initializes the given Collection. The given Collection will be thread-unsafe,
// meaning it will require manual management of its thread-safety using
// the exposed Lock(), Unlock(), RLock() and RUnlock() methods below.
//
// Will result in a Panic if the given Collection has been initialized.
func InitUnsafe(c Collection) (this Collection) {
	this = c
	c.InitUnsafe()
	return c
}

// Returns the given Collection, and:
//
// Returns an int representing the number of items in the given Collection.
//
// Will result in a Panic if the given Collection has not been initialized.
func Size(c Collection) (this Collection, size int) {
	this = c
	size = c.Size()
	return
}

// Returns the given Collection, and:
//
// Returns true if it holds no items, or false if it does.
//
// Will result in a Panic if the given Collection has not been initialized.
func Empty(c Collection) (this Collection, empty bool) {
	this = c
	empty = c.Empty()
	return
}

// Returns the given Collection, and:
//
// Attempts to apply the given function to every item in the given Collection.
// Stops once all items have been processed, or once the given function returns
// false, whichever occurs first.
//
// Returns true if the function was applied to every item, false otherwise.
//
// Should items be mutated, no guarantees are given about the behavior
// of the given Collection thereafter.
//
// Will result in a Panic if the given Collection has not been initialized.
func Map(c Collection, f func(interface{}) bool) (this Collection) {
	this = c
	c.Map(f)
	return
}

// Returns the given Collection, and:
//
// Returns a slice of all the items in the given Collection.
//
// Will result in a Panic if the given Collection has not been initialized.
func Slice(c Collection) (this Collection, slice *[]interface{}) {
	this = c
	slice = c.Slice()
	return
}

// // Returns the given Collection, and:
// //
// // Returns to a new Collection that is a copy of the given
// // Collection.
// //
// // Will result in a Panic if the given Collection has not been initialized.
// func Copy(c Collection) (this Collection, copy Collection) {
// 	this = c
// 	copy = c.Copy()
// 	return
// }

// Returns the given Collection, and:
//
// Clears all items from the given Collection, making it empty.
func Clear(c Collection) (this Collection) {
	this = c
	c.Clear()
	return
}

// Returns the given Collection, and:
//
// Returns true if the given Collection is thread-safe. A thread-safe
// Collection manages its own resources to ensure thread-safety.
// This is the default behavior for all Collections.
//
// Returns false if the given Collection is thread-unsafe. A thread-
// unsafe Collection does not manage its own resources to ensure
// thread-safety, and instead expects clients to do so though the
// exposed Lock(), Unlock(), RLock() and RUnlock() methods.
//
// Will result in a Panic if the given Collection has not been initialized.
func Threadsafe(c Collection) (this Collection, threadsafe bool) {
	this = c
	threadsafe = c.Threadsafe()
	return
}

// Returns the given Collection, and:
//
// Attempts to acquire the lock on the given Collection for both reading and
// writing, blocking to do so.
//
// Lock() should only be called directly if the given Collection is not managing
// its own thread-safety, that is, when the given Collection.ThreadSafe() returns
// false. Calls to Lock() while the given Collection.ThreadSafe() returns true
// will Panic.
//
// Will result in a Panic if the given Collection has not been initialized.
func Lock(c Collection) (this Collection) {
	this = c
	c.Lock()
	return
}

// Returns the given Collection, and:
//
// Releases the lock on the given Collection for both reading and writing.
//
// Unlock() should only be called directly if the given Collection is not managing
// its own thread-safety, that is, when the given Collection.ThreadSafe() returns
// false. Calls to Unlock() while the given Collection.ThreadSafe() returns true
// will Panic.
//
// Will result in a Panic if the given Collection has not been initialized.
func Unlock(c Collection) (this Collection) {
	this = c
	c.Unlock()
	return
}

// Returns the given Collection, and:
//
// Attempts to acquire the lock on the given Collection for reading, blocking
// to do so.
//
// RLock() should only be called directly if the given Collection is not managing
// its own thread-safety, that is, when the given Collection.ThreadSafe() returns
// false. Calls to RLock() while the given Collection.ThreadSafe() returns true
// will Panic.
//
// Will result in a Panic if the given Collection has not been initialized.
func RLock(c Collection) (this Collection) {
	this = c
	c.RLock()
	return
}

// Returns the given Collection, and:
//
// Releases the lock on the given Collection for reading.
//
// RUnlock() should only be called directly if the given Collection is not managing
// its own thread-safety, that is, when the given Collection.ThreadSafe() returns
// false. Calls to RUnlock() while the given Collection.ThreadSafe() returns true
// will Panic.
//
// Will result in a Panic if the given Collection has not been initialized.
func RUnlock(c Collection) (this Collection) {
	this = c
	c.RUnlock()
	return
}
