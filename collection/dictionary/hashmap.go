// This module implements a HashMap, conforming to the Dictionary interface,
// with the additional guarantee that last item to be Pushed,
// will be the first item to be Popped from this WorkList (LIFO).

package dictionary

import (
	"fmt" // To help with String().
	"github.com/michalpiszczek/nonstdlib/collection"
	"log"
)

// A HashMap implements Dictionary.
//
// Behavior unspecified if a HashMap is not created using NewHashMap(), NewHashMapUnsafe()
// or if HashMap.Init() / HashMap.InitUnsafe(), is not first called on a new &HashMap{}.
//
type HashMap struct {
	collection.Base
	m map[interface{}]interface{} // We'll just wrap this for now
}

// Returns a pointer to a new HashMap.
func NewHashMap() *HashMap {
	s := &HashMap{}
	s.Init()
	return s
}

// Returns a pointer to a new unsafe HashMap.
func NewHashMapUnsafe() *HashMap {
	s := &HashMap{}
	s.InitUnsafe()
	return s
}

func (s *HashMap) Init() {
	s.InitBase()

	s.m = make(map[interface{}]interface{})
}

func (s *HashMap) InitUnsafe() {
	s.InitBaseUnsafe()

	s.m = make(map[interface{}]interface{})
}

func (s *HashMap) Insert(key interface{}, value interface{}) interface{} {
	s.CheckInit()
	if key == nil {
		log.Panic("Nil key.")
	}
	if s.Threadsafe() {
		s.Lockb.Lock()
		defer s.Lockb.Unlock()
	}

	old, ok := s.m[key]
	s.m[key] = value
	s.Sizeb += 1

	if !ok {
		return nil
	}
	return old
}

func (s *HashMap) Locate(key interface{}) interface{} {
	s.CheckInit()
	if key == nil {
		log.Panic("Nil key.")
	}
	if s.Threadsafe() {
		s.Lockb.RLock()
		defer s.Lockb.RUnlock()
	}

	value, ok := s.m[key]

	if !ok {
		return nil
	}

	return value
}

func (s *HashMap) Remove(key interface{}) interface{} {
	s.CheckInit()
	if key == nil {
		log.Panic("Nil key.")
	}
	if s.Threadsafe() {
		s.Lockb.Lock()
		defer s.Lockb.Unlock()
	}

	value, ok := s.m[key]

	if !ok {
		return nil
	}

	delete(s.m, key)
	s.Sizeb -= 1
	return value
}

func (s *HashMap) Contains(keys ...interface{}) bool {
	s.CheckInit()

	if s.Threadsafe() {
		s.Lockb.RLock()
		defer s.Lockb.RUnlock()
	}

	ok := true
	for _, key := range keys {
		if key == nil {
			log.Panic("Nil key.")
		}
		_, ok = s.m[key]
		if !ok {
			break
		}
	}

	return ok
}

func (s *HashMap) Copy() Dictionary {
	s.CheckInit()

    var c *HashMap
    if s.Threadsafe() {
        c = NewHashMap()
    } else {
        c = NewHashMapUnsafe()
    }

	// maybe?
	if s.Size() == 0 {
		return c
	}

	if s.Threadsafe() {
		s.Lockb.RLock()
		defer s.Lockb.RUnlock()
	}

	for k, v := range s.m {
		c.m[k] = v
	}
	c.Sizeb = s.Sizeb
	return c
}

// Maps over KeyValues
func (s *HashMap) Map(f func(interface{}) bool) bool {
	s.CheckInit()

	if s.Threadsafe() {
		s.Lockb.RLock()
		defer s.Lockb.RUnlock()
	}

	ok := true
	for k, v := range s.m {
		if ok = f(&KeyValue{k, v}); !ok {
			break
		}
	}
	return ok
}

// Returns a slice of pointers to KeyValue structs.
func (s *HashMap) Slice() *[]interface{} {
	s.CheckInit()

	if s.Threadsafe() {
		s.Lockb.RLock()
		defer s.Lockb.RUnlock()
	}

	slice := make([]interface{}, 0, len(s.m))

	for k, v := range s.m {
		slice = append(slice, &KeyValue{k, v})
	}

	return &slice
}

func (s *HashMap) Clear() {
	s.CheckInit()

	if s.Threadsafe() {
		s.Lockb.Lock()
		defer s.Lockb.Unlock()
	}

	s.m = make(map[interface{}]interface{})
	s.Sizeb = 0
}

func (s *HashMap) String() string {
	return fmt.Sprintf("%v", s.Slice())
}
