package stack

import "go-utils/list"

type Stack[T comparable] struct {
    *list.LinkedList[T]
}

func NewStack[T comparable]() *Stack[T] {
    return &Stack[T]{list.NewLinkedList[T]()}
}

func (s *Stack[T]) Push(value T) {
    s.AddHead(value)
}

func (s *Stack[T]) PushValues(values []T) {
    for _, value := range values {
        s.Push(value)
    }
}

func (s *Stack[T]) Pop() (T, error) {
    return s.RemoveHead()
}

func (s *Stack[T]) Peek() (T, error) {
    return s.GetHead()
}
