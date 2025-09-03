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
- ConcurrentArray: thread-safe array-backed list (Add, AddAll, InsertAt, RemoveAt, Contains, Sort, Filter, Map, Reduce)
- ConcurrentList: concurrent list with thread-safe operations (Add, AddAll, Remove, Contains, Values)
- ConcurrentQueue: thread-safe FIFO queue (Offer, OfferValues, Poll, Peek)
- ConcurrentStack: thread-safe LIFO stack (Push, PushValues, Pop, Peek)
- ConcurrentPriorityQueue: thread-safe binary-heap priority queue with comparator (Offer, OfferValues, Poll, Peek)

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
  _ = arr.InsertAt(2, 12)
  fmt.Println("values:", arr.Values()) // [1, 2, 12, 3]

  // Utilities
  fmt.Println("contains 2?", arr.Contains(2)) // true
  arr.Reverse()
  fmt.Println("values:", arr.Values()) // [3, 12, 2, 1]

  arr.Sort(func(a, b int) int { return a - b })
  fmt.Println("values:", arr.Values()) // [1, 2, 3, 12]

  evens := Filter(arr.Iterator(), func(x int) bool { return x%2 == 0 })
  doubled := Map(evens.Iterator(), func(x int) int { return x * 2 })
  sum := Reduce(doubled.Iterator(), 0, func(acc, x int) int { return acc + x })

  fmt.Println("evens:", evens.Values())     // [2, 12]
  fmt.Println("doubled:", doubled.Values()) // [4, 24]
  fmt.Println("sum:", sum)                  // 28
}
```

### LinkedList
```go
import "go-utils/list"

func main() {
  ll := list.NewLinkedList[int]()
  ll.Add(1)
  ll.AddTail(3)
  ll.AddHead(0) // [0, 1, 3]
  ll.InsertAt(2, 2) // [0, 1, 2, 3]

  head, _ := ll.GetHead()
  tail, _ := ll.GetTail()
  vals := ll.Values()
  fmt.Println("head:", head) // 0
  fmt.Println("tail:", tail) // 3
  fmt.Println("vals:", vals) // []int{0, 1, 2, 3}

  evens := Filter(ll.Iterator(), func (x int) bool { return x%2 == 0 })
  doubled := Map(evens.Iterator(), func (x int) int { return x * 2 })
  sum := Reduce(doubled.Iterator(), 0, func (acc, x int) int { return acc + x })
  fmt.Println("evens:", evens.Values()) // [0, 2]
  fmt.Println("doubled:", doubled.Values()) // [0, 4]
  fmt.Println("sum:", sum) // 4

  v, _ := ll.RemoveTail() // 3, now values are [0, 1, 2]
  fmt.Println("Removed:", v) // 3
  fmt.Println("Values:", ll.Values()) // [0, 1, 2]
}
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

### ConcurrentList
```go
import "go-utils/list"

cl := list.NewConcurrentList[int]()
cl.Add(1)
cl.Add(2)
_ = cl.InsertAt(1, 10) // [1, 10, 2]
cl.Sort(func(a, b int) int { return a - b })
vals := cl.Values() // []int{1, 2, 10}
_ = vals
```

### ConcurrentArray
```go
import "go-utils/array"

ca := array.NewConcurrentArray[int]()
ca.Add(1)
ca.AddAll([]int{2, 3})
_ = ca.InsertAt(1, 10) // [1, 10, 2, 3]
ca.Sort(func(a, b int) int { return a - b })
valsA := ca.Values() // []int{1, 2, 3, 10}
_ = valsA
```

### ConcurrentQueue
```go
import "go-utils/queue"

cq := queue.NewConcurrentQueue[int]()
cq.Offer(1)
cq.OfferValues([]int{2, 3})
frontQ, _ := cq.Peek() // 1
valQ, _ := cq.Poll()   // 1, queue now has 2,3
_, _ = frontQ, valQ
```

### ConcurrentStack
```go
import "go-utils/stack"

cs := stack.NewConcurrentStack[int]()
cs.Push(10)
cs.Push(20)
topS, _ := cs.Peek() // 20
valS, _ := cs.Pop()   // 20, stack now has 10
_, _ = topS, valS
```

### ConcurrentPriorityQueue
```go
import "go-utils/queue"

// Comparator returns negative if a<b, zero if equal, positive if a>b
cmp := func(a, b int) int { return a - b } // min-heap
cpq := queue.NewConcurrentPriorityQueue[int](cmp)

cpq.Offer(5)
cpq.OfferValues([]int{3, 8, 1})

peekCPQ, _ := cpq.Peek() // 1 (smallest)
valCPQ, _ := cpq.Poll()  // 1, then 3 will be next
_, _ = valCPQ, peekCPQ
```

Note: The comparator controls heap ordering. For a max-heap, invert the comparison (e.g., return b - a for ints).

## Testing

This project uses `testify` for assertions. To run all tests:

```
go test ./...
```

## Requirements
- Go 1.18+ (uses Go generics)

## License
This project is licensed under the terms of the LICENSE file included in this repository.
