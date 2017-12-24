// This module implements a Queue, conforming to the WorkList interface,
// with the additional guarantee that last item to be Pushed,
// will be the first item to be Popped from this WorkList (FIFO).

package worklist

import (
	"fmt" // To help with String().
	"github.com/michalpiszczek/nonstdlib/collection"
)

// A singly-linked list should suffice.
type qnode struct {
	work interface{}
	next *qnode
}

// A Queue implements WorkList, with the additional guarantee of
// being a FIFO WorkList.
//
// Behavior unspecified if a Queue is not created using NewQueue(), NewQueueUnsafe()
// or if Queue.Init() / Queue.InitUnsafe() is not first called on a new &Queue{}.
//
type Queue struct {
	collection.Base
	front *qnode // last in
	back  *qnode // first in
}

// Returns a pointer to a new Queue.
func NewQueue() *Queue {
	s := &Queue{}
	s.Init()
	return s
}

// Returns a pointer to a new Unsafe Queue.
func NewQueueUnsafe() *Queue {
	s := &Queue{}
	s.InitUnsafe()
	return s
}

func (s *Queue) Init() {
	s.InitBase()
}

func (s *Queue) InitUnsafe() {
	s.InitBaseUnsafe()
}

func (s *Queue) Push(work interface{}) {
	s.CheckInit()

	if s.Threadsafe() {
		s.Lockb.Lock()
		defer s.Lockb.Unlock()
	}

	next := &qnode{work: work}

	// todo: this
	if s.front == nil {
		s.front = next
		s.back = next
	} else {
		s.front.next = next
		s.front = next
	}
	s.Sizeb += 1
}

func (s *Queue) Pop() interface{} {
	s.CheckInit()

	if s.Empty() {
		return nil
	}

	if s.Threadsafe() {
		s.Lockb.Lock()
		defer s.Lockb.Unlock()
	}

	work := s.back.work
	s.back = s.back.next

	s.Sizeb -= 1
	return work
}

func (s *Queue) Copy() WorkList {
	s.CheckInit()

    var c *Queue
	if s.Threadsafe() {
        c = NewQueue()
    } else {
        c = NewQueueUnsafe()
    }

	if s.Empty() {
		return c
	}

	if s.Threadsafe() {
		s.Lockb.RLock()
		defer s.Lockb.RUnlock()
	}

	c.back = &qnode{s.back.work, s.back.next}
	temp := c.back

	curr := s.front.next

	for curr != nil {
		temp.next = &qnode{curr.work, curr.next}
		temp = temp.next
		curr = curr.next
	}

	c.Sizeb = s.Sizeb
	return c
}

// Applies first in -> last in
func (s *Queue) Map(f func(interface{}) bool) bool {
	s.CheckInit()

	if s.Threadsafe() {
		s.Lockb.RLock()
		defer s.Lockb.RUnlock()
	}

	ok := true
	curr := s.back
	for curr != nil && ok {
		ok = f(curr.work)
		curr = curr.next
	}
	return ok
}

// The first item in will be the first item in the slice.
func (s *Queue) Slice() *[]interface{} {
	s.CheckInit()

	if s.Threadsafe() {
		s.Lockb.RLock()
		defer s.Lockb.RUnlock()
	}

	slice := make([]interface{}, 0, s.Size())

	curr := s.back
	for curr != nil {
		slice = append(slice, curr.work)
		curr = curr.next
	}

	return &slice
}

func (s *Queue) Clear() {
	s.CheckInit()

	if s.Threadsafe() {
		s.Lockb.Lock()
		defer s.Lockb.Unlock()
	}

	s.front = nil
	s.back = nil
	s.Sizeb = 0
}

// Top item first.
func (s *Queue) String() string {
	return fmt.Sprintf("%v", s.Slice())
}
