// This module implements an AVL Tree backed Dictionary, conforming to
// Dictionary, with the additional stipulation that all
// all Keys must implement collection.Comparer, defined herein.

// Normally, the Dictionary should be an abstraction on top of the backing
// AVL Tree, but we're rolling it all at once here.

package dictionary

import (
    "fmt" // To help with String().
	"github.com/michalpiszczek/nonstdlib/collection"
	"github.com/michalpiszczek/nonstdlib/util/math"
    "github.com/michalpiszczek/nonstdlib/collection/worklist"
	"log"
)

// * * * * * * * * * * * * * * * * * * * * * * * * * *
//
// node definition, and associated helper functions.
//
// * * * * * * * * * * * * * * * * * * * * * * * * * *

// The node struct for the AVL Tree. Again, stipulates that keys implement
// collection.Comparer.
type node struct {
	K collection.Comparer
	V interface{}
	H int
	C []*node
}

// Returns a pointer to a new node with the given key, value
// and height.
func newNode(k collection.Comparer, v interface{}, h int) *node {
	n := &node{K: k, V: v, H: h, C: make([]*node, 2)}
	return n
}

// Returns the height of the given node. The height of a node is defined
// as the maximum number of edges from it to a leaf. Thus, height(leaf) = 0,
// and the height of a leaf's children (which should be nil) = -1.
func height(n *node) int {
	if n == nil {
		return -1
	} else {
		return n.H
	}
}

// Sets the height of the given node to one greater than the max
// height of its children.
func updateHeight(n *node) {
	n.H = math.Max(height(child(n, 0)), height(child(n, 1))+1)
}

// Returns a pointer to the 0th or 1st child of the given node.
//
// Dangerous.
func child(n *node, c int) *node {
	return n.C[c]
}

// Returns 0 if k1 <= k2, 1 if k1 > k2.
//
// TODO: a better version of this could be elegantly reused in Insert()
// to help clean up some of the garbage there.
func direction(k1 collection.Comparer, k2 interface{}) int {
	return math.Signum(math.Signum(k1.Compare(k2) + 1))
}

// Rotates the given node based on the two given directions, and then
// returns a pointer to the new parent node of the resulting subtree.
func rotate(p *node, dir1 int, dir2 int) *node {
	if dir1 != dir2 {
		p.C[dir1] = rotate(child(p, dir1), dir2, dir2)
	}

	temp := p
	p = child(p, dir1)

	temp.C[dir1] = p.C[1-dir1]
	p.C[1-dir1] = temp

	updateHeight(temp)
	updateHeight(p)

	return p
}

// Returns either the left most (if dir == 0), or right most
// if (dir == 1) successor of n, and removes it from its parent
// node. Don't pass in a leaf as n.
func successor(n *node) *node {
	if child(n, 0) == nil && child(n, 1) == nil {
		return nil
	}

	dir := tallestDir(n)

	parent := n
	current := child(n, dir)

	if current == nil {
		parent.C[dir] = nil
		return current
	}

	dir = math.Abs(dir - 1)

	for child(current, dir) != nil {
		parent = current
		current = child(current, dir)
	}
	parent.C[dir] = nil
	return current
}

// Returns 0 of height(child(n, 0)) > height(child(n, 1)), or 1 otherwise.
func tallestDir(n *node) int {
	return math.Signum(math.Signum(height(child(n, 1))-height(child(n, 0))) + 1)
}

// Returns true if the difference in height between the given node's children
// is no greater than 1. False otherwise.
func balanced(n *node) bool {
	return math.Abs(height(child(n, 0))-height(child(n, 1))) > 1
}

// * * * * * * * * * * * * * * * * * * * * * * * * * *
//
// end node stuff
//
// * * * * * * * * * * * * * * * * * * * * * * * * * *

// An TreeMap implements Dictionary, with the additional guarantee
// of storing its KeyValues in sorted order, as defined by the Key's Compare()
// method (Keys should implement collection.Comparer, defined above).
//
// Behavior unspecified if a HashMap is not created using NewTreeMap(), NewTreeMapUnsafe()
// or if TreeMap.Init() / TreeMap.InitUnsafe(), is not first called on a new &TreeMap{}.
//
type TreeMap struct {
	collection.Base
	root *node
}

// Returns a pointer to a new HashMap.
func NewTreeMap() *TreeMap {
	s := &TreeMap{}
	s.Init()
	return s
}

// Returns a pointer to a new unsafe HashMap.
func NewTreeMapUnsafe() *TreeMap {
	s := &TreeMap{}
	s.InitUnsafe()
	return s
}

func (s *TreeMap) Init() {
	s.InitBase()
}

func (s *TreeMap) InitUnsafe() {
	s.InitBaseUnsafe()
}

// Panic if key doesn't implement collection.Comparer
func (s *TreeMap) Insert(key interface{}, value interface{}) interface{} {
	s.CheckInit()
	if key == nil {
		log.Panic("Nil key.")
	}
	// Use kc from here on
	kc, ok := key.(collection.Comparer)
	if !ok {
		log.Fatal("Key doesn't implement collection.Comparer.")
	}

	if s.Threadsafe() {
		s.Lockb.Lock()
		defer s.Lockb.Unlock()
	}

	current := s.root
	visited := worklist.NewStackUnsafe()

	for current != nil && kc.Compare(current.K) != 0 {
		visited.Push(current)
		direction := direction(kc, current.K)
		current = child(current, direction)
	}

	if current == nil {
		current = newNode(kc, value, 0)
		for !visited.Empty() {
			parent, _ := visited.Pop().(*node)
			dir1 := direction(kc, parent.K)
			parent.C[dir1] = current
			if balanced(current) {
				dir2 := tallestDir(current)
				parent = rotate(parent, dir1, dir2)
			}
			updateHeight(parent)
			current = parent
		}
		s.Sizeb += 1
		s.root = current
	} else if kc.Compare(current.K) == 0 {
		old := current.V
		current.V = value
		return old
	}
	return nil
}

func (s *TreeMap) Locate(key interface{}) interface{} {
	s.CheckInit()
	if key == nil {
		log.Panic("Nil key.")
	}
	// Use kc from here on
	kc, ok := key.(collection.Comparer)
	if !ok {
		log.Fatal("Key doesn't implement collection.Comparer.")
	}

	if s.Threadsafe() {
		s.Lockb.RLock()
		defer s.Lockb.RUnlock()
	}

	current := s.root

	for current != nil && kc.Compare(current.K) != 0 {
		direction := direction(kc, current.K)
		current = child(current, direction)
	}

	if current == nil {
		return nil
	} else if kc.Compare(current.K) == 0 {
		return current.V
	}
	return nil
}

func (s *TreeMap) Remove(key interface{}) interface{} {
	s.CheckInit()
	if key == nil {
		log.Panic("Nil key.")
	}
	// Use kc from here on
	kc, ok := key.(collection.Comparer)
	if !ok {
		log.Fatal("Key doesn't implement collection.Comparer.")
	}

	if s.Threadsafe() {
		s.Lockb.Lock()
		defer s.Lockb.Unlock()
	}

	current := s.root
	visited := worklist.NewStackUnsafe()

	for current != nil && kc.Compare(current.K) != 0 {
		visited.Push(current)
		direction := direction(kc, current.K)
		current = child(current, direction)
	}

	if current == nil {
		return nil
	} else if kc.Compare(current.K) == 0 {
		oldV := current.V
		current = successor(current)
		for !visited.Empty() {
			parent, _ := visited.Pop().(*node)
			dir1 := direction(kc, parent.K)
			parent.C[dir1] = current
			if current != nil && balanced(current) {
				dir2 := tallestDir(current)
				parent = rotate(parent, dir1, dir2)
			}
			updateHeight(parent)
			current = parent
		}

		s.Sizeb -= 1
		s.root = current

		return oldV
	}
	return nil
}

// Will also Fatal() if any key doesn't implement collection.Comparer.
func (s *TreeMap) Contains(keys ...interface{}) bool {

	ok := true
	for _, k := range keys {
		if k == nil {
			log.Fatal("Cannot check Containment with nil key(s).")
		}

		if ok = s.Locate(k) != nil; !ok {
			break
		}
	}
	return ok
}

func (s *TreeMap) Copy() Dictionary {
    s.CheckInit()

    var c *TreeMap
    if s.Threadsafe() {
        c = NewTreeMap()
    } else {
        c = NewTreeMapUnsafe()
    }

    // maybe?
    if s.Size() == 0 {
        return c
    }

    if s.Threadsafe() {
        s.Lockb.RLock()
        defer s.Lockb.RUnlock()
    }

    s.Map(func(kv interface{}) bool {
        kvc, _ := kv.(*KeyValue)
        c.Insert(kvc.Key, kvc.Value)
        return true
    })

    return c
}

// Maps over KeyValues
func (s *TreeMap) Map(f func(item interface{}) bool) bool {
	s.CheckInit()

	if s.Threadsafe() {
		s.Lockb.RLock()
		defer s.Lockb.RUnlock()
	}

    q := worklist.NewStackUnsafe()
    ok := true

    if s.root != nil {
        q.Push(s.root)
    }

    for !q.Empty() {
        curr := q.Pop()

        currc, _ := curr.(*node)

        if right := child(currc, 1); right != nil {
            q.Push(right)
        }

        if left := child(currc, 0); left != nil {
            q.Push(left)
        }

        if ok = f(&KeyValue{currc.K, currc.V}); !ok {
            break
        }
    }

    return ok
}

// Returns a slice of pointers to KeyValue structs.
func (s *TreeMap) Slice() *[]interface{} {
	s.CheckInit()

	if s.Threadsafe() {
		s.Lockb.RLock()
		defer s.Lockb.RUnlock()
	}

	slice := make([]interface{}, 0, s.Size())

    s.Map(func(kv interface{}) bool {
        kvc, _ := kv.(*KeyValue)
        slice = append(slice, kvc)
        return true
    })
	return &slice
}

func (s *TreeMap) Clear() {
	s.CheckInit()

	if s.Threadsafe() {
		s.Lockb.Lock()
		defer s.Lockb.Unlock()
	}

	s.root = nil
	s.Sizeb = 0
}

func (s *TreeMap) String() string {
	return fmt.Sprintf("%v", s.Slice())
}
