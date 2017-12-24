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

Operations supported by all Collections (for details, consult `collections.go`):

```go
    
    Init()                      // initalizes the Collection. It will manage its own thread-safety.
    InitUnsafe()                // initalizes the Collection. It will not manage its own thread-safety.
    Size()                      // returns the number of items in the Collection
    Empty()                     // returns true if there are no items in the Collection, false otherwise.
    Map()                       // applies a given function to each item in the Collection.
    Slice()                     // returns a slice of the Collection
    Clear()                     // removes all items from the Collection
    Copy()                      // returns a copy of the Collection
    Threadsafe()                // returns true if the Collection is thread-safe, false otherwise
    Lock()                      
    Unlock()
    RLock()
    RUnlock()
    String()                    // returns a string representation of the Collection
    
```

## Todo
    - Skiplist
    - Bloomfilter
    - x-fast trie


