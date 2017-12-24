// This module implements a Stack, conforming to the WorkList interface,
// with the additional guarantee that last item to be Pushed,
// will be the first item to be Popped from this WorkList (LIFO).

package worklist

import (
	"fmt" // To help with String().
	"github.com/michalpiszczek/nonstdlib/collection"
)

type snode struct {
	work interface{}
	prev *snode
}

// A Stack implements WorkList with the additional guarantee of
// being a LIFO WorkList.
//
// Behavior unspecified if a Stack is not created using NewStack(), NewStackUnsafe()
// or if Stack.Init() / Stack.InitUnsafe() is not first called on a new &Stack{}.
//
type Stack struct {
	collection.Base
	front *snode
}

// Returns a pointer to a new Stack.
func NewStack() *Stack {
	s := &Stack{}
	s.Init()
	return s
}

// Returns a pointer to a new Stack.
func NewStackUnsafe() *Stack {
	s := &Stack{}
	s.InitUnsafe()
	return s
}

func (s *Stack) Init() {
	s.InitBase()
}

func (s *Stack) InitUnsafe() {
	s.InitBaseUnsafe()
}

func (s *Stack) Push(work interface{}) {
	s.CheckInit()

	if s.Threadsafe() {
		s.Lockb.Lock()
		defer s.Lockb.Unlock()
	}

	s.front = &snode{work: work, prev: s.front}
	s.Sizeb += 1
}

func (s *Stack) Pop() interface{} {
	s.CheckInit()

	if s.Empty() {
		return nil
	}

	if s.Threadsafe() {
		s.Lockb.Lock()
		defer s.Lockb.Unlock()
	}

	work := s.front.work
	s.front = s.front.prev

	s.Sizeb -= 1
	return work
}

func (s *Stack) Copy() WorkList {
	s.CheckInit()

    var c *Stack
	if s.Threadsafe() {
        c = NewStack()
    } else {
        c = NewStackUnsafe()
    }

	if s.Size() == 0 {
		return c
	}

	if s.Threadsafe() {
		s.Lockb.RLock()
		defer s.Lockb.RUnlock()
	}

	c.front = &snode{s.front.work, s.front.prev}
	temp := c.front

	curr := s.front.prev

	for curr != nil {
		temp.prev = &snode{curr.work, curr.prev}
		temp = temp.prev
		curr = curr.prev
	}

	c.Sizeb = s.Sizeb
	return c
}

// Applies top -> bottom
func (s *Stack) Map(f func(interface{}) bool) bool {
	s.CheckInit()

	if s.Threadsafe() {
		s.Lockb.RLock()
		defer s.Lockb.RUnlock()
	}

	ok := true
	curr := s.front
	for curr != nil && ok {
		ok = f(curr.work)
		curr = curr.prev
	}
	return ok
}

// The top item in the worklist will be the first in the slice.
func (s *Stack) Slice() *[]interface{} {
	s.CheckInit()

	if s.Threadsafe() {
		s.Lockb.RLock()
		defer s.Lockb.RUnlock()
	}

	slice := make([]interface{}, 0, s.Size())

	curr := s.front
	for curr != nil {
		slice = append(slice, curr.work)
		curr = curr.prev
	}

	return &slice
}

func (s *Stack) Clear() {
	s.CheckInit()

	if s.Threadsafe() {
		s.Lockb.Lock()
		defer s.Lockb.Unlock()
	}

	s.front = nil
	s.Sizeb = 0
}

// Top item first.
func (s *Stack) String() string {
	return fmt.Sprintf("%v", s.Slice())
}
