# nonstdlib/collection

collection.go defines the interface Collection, implemented by all Collections.

The best way of instantiating a new Collection is through the New[CollectionName] functions. For example:

```golang
go get github.com/rakyll/magicmime
```

'''golang
    import "github.com/michalpiszczek/libmp/collection/set"

    func main() {
        myNewHashSet := set.NewHashSet()
        ...
    }
'''

The package collection/set defines the Set interface. Sets are Collections that store unique items.

    - set/hashset, stores items unordered.

    - set/treeset, stores items ordered.

The package collection/map defines the Map interface. Maps are Collections that associated keys with items.

    - map/hashmap, associates keys with items in an unordered way.

    - map/treemap, associates keys with items in an ordered way.
        -

The package collection/list defines the List interface. Lists are Collections that store items in order.

    - list/queue, stores items in a FIFO order.

    - list/stack, stores items in a LIFO order.

