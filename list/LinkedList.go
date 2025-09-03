package list

import (
	"errors"
	"go-utils/array"
)

type doubleNode[T comparable] struct {
	value T
	prev  *doubleNode[T]
	next  *doubleNode[T]
}

type LinkedList[T comparable] struct {
	head *doubleNode[T]
	tail *doubleNode[T]
	size int
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	var zero T
	head := &doubleNode[T]{value: zero}
	tail := &doubleNode[T]{value: zero}
	head.next = tail
	tail.prev = head
	return &LinkedList[T]{
		head: head,
		tail: tail,
		size: 0,
	}
}

func (s *LinkedList[T]) Size() int {
	return s.size
}

func (s *LinkedList[T]) IsEmpty() bool {
	return s.size == 0
}

func (s *LinkedList[T]) Clear() {
	s.head.next = s.tail
	s.tail.prev = s.head
	s.size = 0
}

func (s *LinkedList[T]) Values() []T {
	values := make([]T, s.size)
	current := s.head.next
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
	newNode := &doubleNode[T]{value: value, prev: s.head, next: s.head.next}
	s.attachHeadNode(newNode)
}

func (s *LinkedList[T]) AddTail(value T) {
	newNode := &doubleNode[T]{value: value, prev: s.tail.prev, next: s.tail}
	s.attachTailNode(newNode)
}

func (s *LinkedList[T]) InsertAt(index int, value T) {
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

func (s *LinkedList[T]) GetHead() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("List is empty")
	}
	return s.head.next.value, nil
}

func (s *LinkedList[T]) GetTail() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("List is empty")
	}
	return s.tail.prev.value, nil
}

func (s *LinkedList[T]) GetAt(index int) (T, error) {
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

func (s *LinkedList[T]) Contains(value T) bool {
	current := s.head.next
	for current.next != nil {
		if current.value == value {
			return true
		}
		current = current.next
	}
	return false
}

func (s *LinkedList[T]) Clone() *LinkedList[T] {
	list := NewLinkedList[T]()

	current := s.head.next
	for current.next != nil {
		list.Add(current.value)
		current = current.next
	}
	return list
}

func (s *LinkedList[T]) Merge(list *LinkedList[T]) {
	current := list.head.next
	for current.next != nil {
		s.Add(current.value)
		current = current.next
	}
}

func (s *LinkedList[T]) MergeArray(arr *array.Array[T]) {
	for _, value := range arr.Values() {
		s.Add(value)
	}
}

func (s *LinkedList[T]) MergeList(list *LinkedList[T]) {
	current := list.head
	for current != nil {
		s.Add(current.value)
		current = current.next
	}
}

func (s *LinkedList[T]) Reverse() {
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

func (s *LinkedList[T]) Sort(comparator func(T, T) int) {
	if s.IsEmpty() {
		return
	}

	s.tail.prev.next = nil
	s.head.next = mergeSortList(s.head.next, comparator)

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

func (s *LinkedList[T]) attachHeadNode(node *doubleNode[T]) {
	s.head.next.prev = node
	s.head.next = node
	s.size++
}

func (s *LinkedList[T]) attachTailNode(node *doubleNode[T]) {
	s.tail.prev.next = node
	s.tail.prev = node
	s.size++
}

func (s *LinkedList[T]) detachHeadNode() *doubleNode[T] {
	node := s.head.next
	s.head.next = s.head.next.next
	s.head.next.prev = s.head
	s.size--
	return node
}

func (s *LinkedList[T]) detachTailNode() *doubleNode[T] {
	node := s.tail.prev
	s.tail.prev = s.tail.prev.prev
	s.tail.prev.next = s.tail
	s.size--
	return node
}

func mergeSortList[T comparable](head *doubleNode[T], comparator func(T, T) int) *doubleNode[T] {
	if head == nil || head.next == nil {
		return head
	}

	left, right := splitList(head)
	left = mergeSortList(left, comparator)
	right = mergeSortList(right, comparator)
	return mergeList(left, right, comparator)
}

func splitList[T comparable](head *doubleNode[T]) (*doubleNode[T], *doubleNode[T]) {
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

func mergeList[T comparable](left, right *doubleNode[T], comparator func(T, T) int) *doubleNode[T] {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	var head *doubleNode[T]
	if comparator(left.value, right.value) <= 0 {
		head = left
		head.next = mergeList(left.next, right, comparator)
	} else {
		head = right
		head.next = mergeList(left, right.next, comparator)
	}
	return head
}
