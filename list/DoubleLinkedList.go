package list

import (
    "errors"
)

type doubleNode[T comparable] struct {
    value T
    prev  *doubleNode[T]
    next  *doubleNode[T]
}

type DoubleLinkedList[T comparable] struct {
    head *doubleNode[T]
    tail *doubleNode[T]
    size int
}

func NewDoubleLinkedList[T comparable]() *DoubleLinkedList[T] {
    var zero T
    head := &doubleNode[T]{value: zero}
    tail := &doubleNode[T]{value: zero}
    head.next = tail
    tail.prev = head
    return &DoubleLinkedList[T]{
        head: head,
        tail: tail,
        size: 0,
    }
}

func (s *DoubleLinkedList[T]) Size() int {
    return s.size
}

func (s *DoubleLinkedList[T]) IsEmpty() bool {
    return s.size == 0
}

func (s *DoubleLinkedList[T]) Clear() {
    s.head.next = s.tail
    s.tail.prev = s.head
    s.size = 0
}

func (s *DoubleLinkedList[T]) Values() []T {
    values := make([]T, s.size)
    current := s.head.next
    for i := 0; i < s.size; i++ {
        values[i] = current.value
        current = current.next
    }
    return values
}

func (s *DoubleLinkedList[T]) Add(value T) {
    if s.IsEmpty() {
        s.AddHead(value)
    } else {
        s.AddTail(value)
    }
}

func (s *DoubleLinkedList[T]) AddAll(values []T) {
    for _, value := range values {
        s.Add(value)
    }
}

func (s *DoubleLinkedList[T]) AddHead(value T) {
    newNode := &doubleNode[T]{value: value, prev: s.head, next: s.head.next}
    s.attachHeadNode(newNode)
}

func (s *DoubleLinkedList[T]) AddTail(value T) {
    newNode := &doubleNode[T]{value: value, prev: s.tail.prev, next: s.tail}
    s.attachTailNode(newNode)
}

func (s *DoubleLinkedList[T]) InsertAt(index int, value T) {
    if s.IsEmpty() || index <= 0 {
        s.AddHead(value)
    } else if index >= s.size {
        s.AddTail(value)
    } else {
        current := s.head.next
        for i := 0; i < index; i++ {
            current = current.next
        }

        newNode := &doubleNode[T]{value: value, prev: current.prev, next: current}
        current.prev.next = newNode
        current.prev = newNode
        s.size++
    }
}

func (s *DoubleLinkedList[T]) GetHead() (T, error) {
    if s.IsEmpty() {
        var zero T
        return zero, errors.New("List is empty")
    }
    return s.head.next.value, nil
}

func (s *DoubleLinkedList[T]) GetTail() (T, error) {
    if s.IsEmpty() {
        var zero T
        return zero, errors.New("List is empty")
    }
    return s.tail.prev.value, nil
}

func (s *DoubleLinkedList[T]) GetAt(index int) (T, error) {
    if s.IsEmpty() || index < 0 || index >= s.size {
        var zero T
        return zero, errors.New("Index out of range")
    }

    current := s.head.next
    for i := 0; i < index; i++ {
        current = current.next
    }
    return current.value, nil
}

func (s *DoubleLinkedList[T]) RemoveHead() (T, error) {
    if s.IsEmpty() {
        var zero T
        return zero, errors.New("List is empty")
    }

    return s.detachHeadNode().value, nil
}

func (s *DoubleLinkedList[T]) RemoveTail() (T, error) {
    if s.IsEmpty() {
        var zero T
        return zero, errors.New("List is empty")
    }

    return s.detachTailNode().value, nil
}

func (s *DoubleLinkedList[T]) RemoveAt(index int) (T, error) {
    if s.IsEmpty() || index < 0 || index >= s.size {
        var zero T
        return zero, errors.New("Index out of range")
    }

    current := s.head.next
    for i := 0; i < index; i++ {
        current = current.next
    }
    value := current.value
    current.prev.next = current.next
    current.next.prev = current.prev
    s.size--
    return value, nil
}

func (s *DoubleLinkedList[T]) Contains(value T) bool {
    current := s.head.next
    for current != nil {
        if current.value == value {
            return true
        }
        current = current.next
    }
    return false
}

func (s *DoubleLinkedList[T]) Clone() *DoubleLinkedList[T] {
    list := NewDoubleLinkedList[T]()

    current := s.head.next
    for current.next != nil {
        list.Add(current.value)
        current = current.next
    }
    return list
}

func (s *DoubleLinkedList[T]) Merge(list *DoubleLinkedList[T]) {
    current := list.head.next
    for current.next != nil {
        s.Add(current.value)
        current = current.next
    }
}

func (s *DoubleLinkedList[T]) Reverse() {
    if s.IsEmpty() || s.size == 1 {
        return
    }

    var temp *doubleNode[T]
    node := s.head.next
    current := s.head.next
    for current.next != nil {
        temp = current.prev
        current.prev = current.next
        current.next = temp

        current = current.prev
    }

    if temp != nil {
        temp.prev.prev = s.head
        s.head.next = temp.prev

        s.tail.prev = node
        node.next = s.tail
    }
}

func (s *DoubleLinkedList[T]) Filter(predicate func(T) bool) *DoubleLinkedList[T] {
    filtered := NewDoubleLinkedList[T]()
    current := s.head.next
    for current.next != nil {
        if predicate(current.value) {
            filtered.Add(current.value)
        }
        current = current.next
    }
    return filtered
}

func (s *DoubleLinkedList[T]) Sort(comparator func(T, T) int) {
    if s.IsEmpty() {
        return
    }

    s.tail.prev.next = nil
    s.head.next = mergeSortDoubleLinkedList(s.head.next, comparator)

    // Fix the tail pointer
    prev := s.head
    curr := s.head.next
    for curr != nil {
        curr.prev = prev
        prev = curr
        curr = curr.next
    }

    prev.next = s.tail
    s.tail.prev = prev
}

func MapDoubleLinkedList[T, V comparable](s *DoubleLinkedList[T], mapper func(T) V) *DoubleLinkedList[V] {
    if s.IsEmpty() {
        return NewDoubleLinkedList[V]()
    }

    result := NewDoubleLinkedList[V]()
    current := s.head.next
    for current.next != nil {
        result.Add(mapper(current.value))
        current = current.next
    }
    return result
}

func ReduceDoubleLinkedList[T comparable, V any](s *DoubleLinkedList[T], initial V, reducer func(V, T) V) V {
    if s.IsEmpty() {
        return initial
    }

    acc := initial
    current := s.head.next
    for current.next != nil {
        acc = reducer(acc, current.value)
        current = current.next
    }

    return acc
}

func (s *DoubleLinkedList[T]) attachHeadNode(node *doubleNode[T]) {
    s.head.next.prev = node
    s.head.next = node
    s.size++
}

func (s *DoubleLinkedList[T]) attachTailNode(node *doubleNode[T]) {
    s.tail.prev.next = node
    s.tail.prev = node
    s.size++
}

func (s *DoubleLinkedList[T]) detachHeadNode() *doubleNode[T] {
    node := s.head.next
    s.head.next = s.head.next.next
    s.head.next.prev = s.head
    s.size--
    return node
}

func (s *DoubleLinkedList[T]) detachTailNode() *doubleNode[T] {
    node := s.tail.prev
    s.tail.prev = s.tail.prev.prev
    s.tail.prev.next = s.tail
    s.size--
    return node
}

func mergeSortDoubleLinkedList[T comparable](head *doubleNode[T], comparator func(T, T) int) *doubleNode[T] {
    if head == nil || head.next == nil {
        return head
    }

    left, right := splitDoubleLinkedList(head)
    left = mergeSortDoubleLinkedList(left, comparator)
    right = mergeSortDoubleLinkedList(right, comparator)
    return mergeDoubleLinkedList(left, right, comparator)
}

func splitDoubleLinkedList[T comparable](head *doubleNode[T]) (*doubleNode[T], *doubleNode[T]) {
    if head == nil || head.next == nil {
        return head, nil
    }

    slow, fast := head, head.next
    for fast != nil && fast.next != nil {
        slow = slow.next
        fast = fast.next.next
    }

    mid := slow.next
    slow.next = nil
    return head, mid
}

func mergeDoubleLinkedList[T comparable](left, right *doubleNode[T], comparator func(T, T) int) *doubleNode[T] {
    if left == nil {
        return right
    }
    if right == nil {
        return left
    }

    var head *doubleNode[T]
    if comparator(left.value, right.value) <= 0 {
        head = left
        head.next = mergeDoubleLinkedList(left.next, right, comparator)
    } else {
        head = right
        head.next = mergeDoubleLinkedList(left, right.next, comparator)
    }
    return head
}
