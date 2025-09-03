package list

import (
	"errors"
	"go-utils/array"
)

type singleNode[T comparable] struct {
	value T
	next  *singleNode[T]
}

type LinkedList[T comparable] struct {
	head *singleNode[T]
	tail *singleNode[T]
	size int
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (s *LinkedList[T]) Size() int {
	return s.size
}

func (s *LinkedList[T]) IsEmpty() bool {
	return s.size == 0
}

func (s *LinkedList[T]) Clear() {
	s.head = nil
	s.tail = nil
	s.size = 0
}

func (s *LinkedList[T]) Values() []T {
	values := make([]T, s.size)

	current := s.head
	for i := 0; i < s.size; i++ {
		values[i] = current.value
		current = current.next
	}
	return values
}

func (s *LinkedList[T]) Add(value T) {
	if s.IsEmpty() {
		s.AddHead(value)
	} else {
		s.AddTail(value)
	}
}

func (s *LinkedList[T]) AddAll(values []T) {
	for _, value := range values {
		s.Add(value)
	}
}

func (s *LinkedList[T]) AddHead(value T) {
	if s.IsEmpty() {
		s.attachHeadNode(&singleNode[T]{value: value})
	} else {
		s.attachHeadNode(&singleNode[T]{value: value, next: s.head})
	}
}

func (s *LinkedList[T]) AddTail(value T) {
	if s.IsEmpty() {
		s.AddHead(value)
	} else {
		s.attachTailNode(&singleNode[T]{value: value})
	}
}

func (s *LinkedList[T]) InsertAt(index int, value T) {
	if s.IsEmpty() || index <= 0 {
		s.AddHead(value)
	} else if index >= s.size {
		s.AddTail(value)
	} else {
		current := s.head
		for i := 0; i < index; i++ {
			current = current.next
		}

		if current == s.tail {
			s.tail.next = &singleNode[T]{value: value}
			s.tail = s.tail.next
		} else {
			next := current.next
			current.next = &singleNode[T]{value: value, next: next}
		}
		s.size++
	}
}

func (s *LinkedList[T]) GetHead() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("List is empty")
	}

	return s.head.value, nil
}

func (s *LinkedList[T]) GetTail() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("List is empty")
	}

	return s.tail.value, nil
}

func (s *LinkedList[T]) GetAt(index int) (T, error) {
	if s.IsEmpty() || index < 0 || index >= s.size {
		var zero T
		return zero, errors.New("Index out of range")
	}

	current := s.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.value, nil
}

func (s *LinkedList[T]) RemoveHead() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("List is empty")
	}

	return s.detachHeadNode().value, nil
}

func (s *LinkedList[T]) RemoveTail() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("List is empty")
	}

	return s.detachTailNode().value, nil
}

func (s *LinkedList[T]) RemoveAt(index int) (T, error) {
	if s.IsEmpty() || index < 0 || index >= s.size {
		var zero T
		return zero, errors.New("Index out of range")
	}

	if index == 0 || s.Size() == 1 {
		return s.RemoveHead()
	}

	if index == s.size-1 {
		return s.RemoveTail()
	}

	prev := s.head
	current := s.head
	for i := 0; i < index; i++ {
		prev = current
		current = current.next
	}

	prev.next = current.next
	s.size--
	return current.value, nil
}

func (s *LinkedList[T]) Contains(value T) bool {
	current := s.head
	for current != nil {
		if current.value == value {
			return true
		}
		current = current.next
	}
	return false
}

func (s *LinkedList[T]) Clone() *LinkedList[T] {
	list := NewLinkedList[T]()
	current := s.head
	for current != nil {
		list.AddTail(current.value)
		current = current.next
	}
	return list
}

func (s *LinkedList[T]) Merge(list *LinkedList[T]) {
	if list.IsEmpty() {
		return
	}

	current := list.head
	for current != nil {
		s.AddTail(current.value)
		current = current.next
	}
}

func (s *LinkedList[T]) MergeArray(arr *array.Array[T]) {
	for _, value := range arr.Values() {
		s.AddTail(value)
	}
}

func (s *LinkedList[T]) Reverse() {
	if s.IsEmpty() || s.size == 1 {
		return
	}

	var prev *singleNode[T]
	current := s.head
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	s.head, s.tail = s.tail, s.head
}

func (s *LinkedList[T]) Filter(predicate func(T) bool) *LinkedList[T] {
	filtered := NewLinkedList[T]()
	current := s.head
	for current != nil {
		if predicate(current.value) {
			filtered.AddTail(current.value)
		}
		current = current.next
	}
	return filtered
}

func (s *LinkedList[T]) Sort(comparator func(T, T) int) {
	if s.IsEmpty() {
		return
	}

	s.head = mergeSort(s.head, comparator)
	s.tail = s.head
	for s.tail.next != nil {
		s.tail = s.tail.next
	}
}

func MapLinkedList[T comparable, V comparable](s *LinkedList[T], mapper func(T) V) *LinkedList[V] {
	if s.IsEmpty() {
		return NewLinkedList[V]()
	}

	result := NewLinkedList[V]()
	current := s.head
	for current != nil {
		result.AddTail(mapper(current.value))
		current = current.next
	}
	return result
}

func ReduceLinkedList[T comparable, V any](s *LinkedList[T], initial V, reducer func(acc V, item T) V) V {
	acc := initial
	current := s.head
	for current != nil {
		acc = reducer(acc, current.value)
		current = current.next
	}

	return acc
}

func (s *LinkedList[T]) attachHeadNode(node *singleNode[T]) {
	if s.IsEmpty() {
		s.head = node
		s.tail = s.head
	} else {
		newNode := node
		s.head = newNode
	}
	s.size++
}

func (s *LinkedList[T]) attachTailNode(node *singleNode[T]) {
	s.tail.next = node
	s.tail = node
	s.tail.next = nil
	s.size++
}

func (s *LinkedList[T]) detachHeadNode() *singleNode[T] {
	node := s.head
	if s.head == s.tail {
		s.Clear()
	} else {
		s.head = s.head.next
		s.size--
	}
	return node
}

func (s *LinkedList[T]) detachTailNode() *singleNode[T] {
	node := s.tail
	if s.head == s.tail {
		s.Clear()
	} else {
		prev := s.head
		current := s.head
		for current.next != nil {
			prev = current
			current = current.next
		}
		prev.next = nil
		s.tail = prev
		s.size--
	}

	return node
}

func mergeSort[T comparable](head *singleNode[T], comparator func(T, T) int) *singleNode[T] {
	if head == nil || head.next == nil {
		return head
	}

	left, right := split(head)
	left = mergeSort(left, comparator)
	right = mergeSort(right, comparator)
	return merge(left, right, comparator)
}

func split[T comparable](head *singleNode[T]) (*singleNode[T], *singleNode[T]) {
	if head == nil || head.next == nil {
		return head, nil
	}

	slow, fast := head, head.next
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	middle := slow.next
	slow.next = nil
	return head, middle
}

func merge[T comparable](left, right *singleNode[T], comparator func(T, T) int) *singleNode[T] {
	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	var head *singleNode[T]
	if comparator(left.value, right.value) <= 0 {
		head = left
		head.next = merge(left.next, right, comparator)
	} else {
		head = right
		head.next = merge(left, right.next, comparator)
	}

	return head
}
