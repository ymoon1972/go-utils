# go-utils

Java collection-like utilities for Go (generic collections inspired by Java's Collections Framework).

This repository provides a small set of generic data structures with familiar APIs:
- ArrayList: dynamic array-backed list with utilities such as InsertAt, Contains, Sort, Filter, Map, Reduce
- LinkedList: singly linked list with head/tail operations and index-based access
- DoubleLinkedList: doubly linked list with bidirectional traversal operations
- Stack: LIFO stack backed by LinkedList (Push, Pop, Peek)
- Queue: FIFO queue backed by LinkedList (Offer, Poll, Peek)
- PriorityQueue: binary-heap priority queue with a user-supplied comparator (min-/max-heap behavior by comparator)
- BinaryTree: binary search tree with comparator-defined ordering (Offer, OfferAll, Remove, Contains, in-order Values)

All collections are implemented using Go generics (type parameters), with methods designed to be easy to use and test.

## Module

Module name (from go.mod): `go-utils`

Note: The module path is local (`go-utils`). If you plan to use this as a dependency from another project, update the module path to your VCS hosting path (e.g., `github.com/<you>/go-utils`) and run `go mod tidy`. For local usage within this repository, the current module name works as-is.

## Installation

- As a local module: clone this repository.
- As a dependency: set the module path to your VCS location and `go get <your-path>/go-utils`.

## Usage

Below are small examples showing how to use the collections.

### ArrayList
```go
package main

import (
    "fmt"
    "go-utils/array"
)

func main() {
    arr := array.NewArrayList[int]()
    arr.Add(1)
    arr.AddAll([]int{2, 3})
    _ = arr.InsertAt(2, 13) // [1, 2, 13, 3]

    fmt.Println("values:", arr.Values())

    // Utilities
    fmt.Println("contains 2?", arr.Contains(2))
    arr.Reverse()
    arr.Sort(func(a, b int) int { return a - b })

    evens := arr.Filter(func(x int) bool { return x%2 == 0 })
    doubled := array.MapArrayList(evens, func(x int) int { return x * 2 })
    sum := array.ReduceArrayList(doubled, 0, func(acc, x int) int { return acc + x })

    fmt.Println("evens:", evens.Values())
    fmt.Println("doubled:", doubled.Values())
    fmt.Println("sum:", sum)
}
```

### LinkedList
```go
import "go-utils/list"

ll := list.NewLinkedList[string]()
ll.Add("a")
ll.AddHead("z")         // [z, a]
_ = ll.InsertAt(1, "b") // [z, b, a]
head, _ := ll.GetHead()
_ = head
vals := ll.Values() // []string{"z","b","a"}
```

### Stack
```go
import "go-utils/stack"

s := stack.NewStack[int]()
s.Push(10)
s.Push(20)
top, _ := s.Peek() // 20
v, _ := s.Pop()     // 20, stack now has 10
_ = v
```

### Queue
```go
import "go-utils/queue"

q := queue.NewQueue[int]()
q.Offer(1)
q.OfferValues([]int{2, 3})
front, _ := q.Peek() // 1
v, _ := q.Poll()     // 1, queue now has 2,3
_ = v
```

### BinaryTree
```go
import "go-utils/tree"

// Comparator returns negative if a<b, zero if equal, positive if a>b
cmp := func(a, b int) int { return a - b }

bt := tree.NewBinaryTree[int](cmp)
bt.Offer(7)
bt.OfferAll([]int{2, 4, 8})

// In-order values are sorted by comparator
vals := bt.Values() // []int{2, 4, 7, 8}
_ = vals

contains := bt.Contains(4) // true
_ = contains

// Remove a value
_ = bt.Remove(8)
_ = bt.Remove(2)

// Size/empty
size := bt.Size()     // current node count
empty := bt.IsEmpty() // false unless cleared
_ = size
_ = empty

bt.Clear()
_ = bt.IsEmpty() // true
```

### PriorityQueue
```go
import "go-utils/queue"

// Comparator returns negative if a<b, zero if equal, positive if a>b
cmp := func(a, b int) int { return a - b } // min-heap
pq := queue.NewPriorityQueue[int](cmp)

pq.Offer(5)
pq.OfferValues([]int{3, 8, 1})

peek, _ := pq.Peek() // 1 (smallest)
val, _ := pq.Poll()  // 1, then 3 will be next
_ = val
_ = peek
```

Note: The comparator controls heap ordering. For a max-heap, invert the comparison (e.g., return b - a for ints).

## Testing

This project uses `testify` for assertions. To run all tests:

```
go test ./...
```

## License

This project is licensed under the terms of the LICENSE file included in this repository.
