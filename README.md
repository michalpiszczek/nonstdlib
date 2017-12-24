# nonstdlib

## About

`nonstdlib` is a library of thread-safe data structures for Go. 

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
    ...
}
```

Operations supported by all Collections (for details see `collection/collection.go`):

```go
    var c Collection
    
    c.Init()             // initalizes the Collection. It will manage its own thread-safety.
    c.InitUnsafe()       // initalizes the Collection. It will not manage its own thread-safety.
    c.Size()             // returns the number of items in the Collection
    c.Empty()            // returns true if there are no items in the Collection, false otherwise.
    c.Map()              // applies a given function to each item in the Collection.
    c.Slice()            // returns a slice of the Collection
    c.Clear()            // removes all items from the Collection
    c.Copy()             // returns a copy of the Collection
    c.Threadsafe()       // returns true if the Collection is thread-safe, false otherwise
    c.Lock()             // attempts to acquire the lock on the Collection   
    c.Unlock()           // releases the lock on the Collection
    c.RLock()            // attempts to acquire the lock on the Collection  (for reading)
    c.RUnlock()          // releases the lock on the Collection (for reading)
    c.String()           // returns a string representation of the Collection
```

Additional operations supported by all WorkLists (see `collection/worklist/worklist.go`):

```go 
    var w WorkList
    
    w.Push()             // pushes the given work onto the WorkList
    w.Pop()              // pops the next item of work off the WorkList
```

Additional operations supported by all Dictionaries (see `collection/worklist/dictionary.go`):

```go 
    var w WorkList
    
    w.Push()             // pushes the given work onto the WorkList
    w.Pop()              // pops the next item of work off the WorkList
```

## Todo
    - Skiplist
    - Bloomfilter
    - x-fast trie


