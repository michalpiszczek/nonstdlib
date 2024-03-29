# nonstdlib

`nonstdlib` is a non-idiomatic library of thread-safe collections for Go. 

Currently, these Collections are avaiable:
    
    - Worklist 
        - Queue 
        - Stack
    - Dictionary
        - HashMap (backed by Go's `map`)
        - TreeMap (AVL backed)
    - Set
        - HashSet 
        - TreeSet
       

## Installation

Make sure you have [Go](http://golang.org) installed. Then,

```go
go get github.com/michalpiszczek/nonstdlib
```

## Usage

Create an instance of a Collection:

```golang
import "github.com/michalpiszczek/libmp/collection/set"

func main() {
    myNewHashSet := set.NewHashSet()
    myNewUnsafeHashSet := set.NewHashSetUnsafe()
    ...
}
```

All Collections provide these convenience functions:

```
my[Collection] := [collection].New[Collection]() 

myUnsafe[Collection] := [collection].New[Collection]Unsafe()
```

### Method Summary

Please see the interfaces for complete details!

#### Collections (details: `collection/collection.go`):

```go
    var c Collection
    
    c.Init()             // initalizes c. It will manage its own thread-safety
    c.InitUnsafe()       // initalizes c. It will not manage its own thread-safety
    c.Size()             // returns the number of items in c
    c.Empty()            // returns true if there are no items in c, false otherwise
    c.Map(func())        // applies a given function to each item in c
    c.Slice()            // returns a slice of c
    c.Clear()            // removes all items from c
    c.Copy()             // returns a copy of c
    c.Threadsafe()       // returns true if c is thread-safe, false otherwise
    c.Lock()             // attempts to acquire the lock on c
    c.Unlock()           // releases the lock on c
    c.RLock()            // attempts to acquire the lock on c (for reading)
    c.RUnlock()          // releases the lock on c (for reading)
    c.String()           // returns a string representation of c
```

#### WorkLists (details: `collection/worklist/worklist.go`):

```go 
    var w WorkList
    
    w.Push(work)         // pushes the given work onto w
    w.Pop(work)          // pops the next item of work off w
```

#### Dictionaries (details: `collection/dictionary/dictionary.go`):

**Note**: *All keys for TreeMaps must implement collections.Comparer.*

```go 
    var d Dictionary
    
    d.Insert(key, val)    // adds val to d, associated with key
    d.Locate(key)         // returns the value associated with key in d
    d.Remove(key)         // removes the value associated with key from d
    d.Contains(...keys)   // returns true if d contains values for all keys
```

#### Sets (details: `collection/set/set.go`):

**Note**: *All keys for TreeSets must implement collections.Comparer.*

```go 
    var s Set
    
    s.Insert(...items)    // adds all items to s
    s.Remove(...items)    // removes all items from s
    s.Contains(...items)  // returns true if all items are in s
    s.Union(Set)          // returns the union of s and the given Set
    s.Intersection(Set)   // returns the intersection of s and the given Set
    s.Difference(Sety)    // returns the difference of s and the given Set
    s.Equal(Sety)         // returns true if s and the given Set are equal, false otherwise
    s.Subset(Set)         // returns true if s is a Subset of the given Set, false otherwise
    s.Superset(Set)       // returns true if s is a Superset of the given Set, false otherwise

```
## Implementation

I'm fairly new to Go, and I'm loving it so far! 

This library has been a fun distraction from another project that I've got **Go**ing.

I am happy with:

    - Struct composition! Embeding collections.Base saved a lot of work! 
    - Similarly, interface inheritance/composition! 
    - lock.Lock(); defer lock.Unlock() is super useful!
    - I wrote a quick JUnit-esque wrapper for Go's `testing`, check it out in `nonstdlib/util/test`.
    - Testing in general. I think I've managed decent coverage so far.
    - AVL-Tree's Remove. First time implementing that. A little trickier than Insert, but not too bad.

I would like to address:
    
    - Embeding collections.Base required making some internal fields public
    - Issues with covariance prevent Copy() from being defined in the Collection interface
    - Some of the interfaces could be more consistent
    - The Comparer interface is pretty hacky and not well docced
    - HashMap just wrap's Go's map (I doubt I can do better, but it'd be fun to try)
    - The thread-safety is kind of weighty
    - Everything's probably slower than it could be
    - Benchmarking
    - Improve docs

It would be cool to add:

    - skiplist
    - bloomfilter
    - x-fast trie

## Cheers!
