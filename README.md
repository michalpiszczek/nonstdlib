# nonstdlib

## About

`nonstdlib` is a non-idiomatic library of thread-safe data structures for Go. 

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

First make sure you have [Go](http://golang.org) installed, then:

```go
go get github.com/michalpiszczek/nonstdlib
```

## Usage

Import and create an instance of a Collection,

```golang
import "github.com/michalpiszczek/libmp/collection/set"

func main() {
    myNewHashSet := set.NewHashSet()
    myNewUnsafeHashSet := set.NewHashSetUnsafe()
    ...
}
```

All Collections provide these convenience functions,

```
my[Collection] := [collection].New[Collection]() 

myUnsafe[Collection] := [collection].New[Collection]Unsafe()
```

### Summary

This a summary of the operations supported by each Collection, but does not fully capture all their behavior. Please see the respective interfaces for complete details.

Operations supported by all Collections (details: `collection/collection.go`):

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

Additional operations supported by all WorkLists (details: `collection/worklist/worklist.go`):

```go 
    var w WorkList
    
    w.Push(work)         // pushes the given work onto w
    w.Pop(work)          // pops the next item of work off w
```

Additional operations supported by all Dictionaries (details: `collection/dictionary/dictionary.go`):

    Note: All keys for TreeMaps must implement collections.Comparer.

```go 
    var d Dictionary
    
    d.Insert(key, val)    // adds val to d, associated with key
    d.Locate(key)         // returns the value associated with key in d
    d.Remove(key)         // removes the value associated wit key from d
    d.Contains(...keys)   // returns true if d contains values for all keys
```

Additional operations supported by all Sets (details: `collection/set/set.go`):

    Note: All keys for TreeSets must implement collections.Comparer.

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

I'm fairly new to Go, but I'm loving it so far! I've got another project Going, and this library is sort of a quick detour from that, and as such, it hasn't had too much thought put into it. 

Here are some things I'm happy with:

    - Struct composition! Embeding collections.Base in all the Collections saved a lot of work. 
      I like this language feature!
    - Similarly, interface inheritance/compositon! 
    - lock.Lock(); defer lock.Unlock() is super useful!
    - Testing. I wrote a quick JUnit-esque wrapper for goes `testing`, check it out in `nonstdlib/util/test`.

Here are some things I'd like to address at some point:
    
    - Embeding collections.Base and using collections.Base required making some internal fields public
    - There's some issues with covariance, or lack thereof, that Cause problems for Copy()
    - Some of the interfaces could be more consistent
    - The Comparer interface is pretty hacky
    - HashMap just wrap's Go's map (I doubt I can do better, but it'd be fun to try)
    - The thread-safety is kind of weighty
    - Everythings probably slower than it could be

## Todo
    - cleanup interfaces
    - complete test coverage
    
    - Skiplist
    - Bloomfilter
    - x-fast trie


