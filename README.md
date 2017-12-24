# nonstdlib

## About

`nonstdlib` is a library of thread-safe data structures for Go. Currently, these structures are avaiable:
    
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


```golang
import "github.com/michalpiszczek/libmp/collection/set"

func main() {
    myNewHashSet := set.NewHashSet()
    ...
}
```

## Todo
    - Skiplist
    - Bloomfilter
    - x-fast trie


