package stack

import "sync"

type ConcurrentStack[T comparable] struct {
    mu    sync.RWMutex
    stack *Stack[T]
}

func NewConcurrentStack[T comparable]() *ConcurrentStack[T] {
    return &ConcurrentStack[T]{stack: NewStack[T]()}
}

func (s *ConcurrentStack[T]) Size() int {
    return s.stack.Size()
}

func (s *ConcurrentStack[T]) IsEmpty() bool {
    return s.stack.IsEmpty()
}

func (s *ConcurrentStack[T]) Clear() {
    s.mu.Lock()
    defer s.mu.Unlock()

    s.stack.Clear()
}

func (s *ConcurrentStack[T]) Values() []T {
    s.mu.RLock()
    defer s.mu.RUnlock()

    return s.stack.Values()
}

func (s *ConcurrentStack[T]) Push(value T) {
    s.mu.Lock()
    defer s.mu.Unlock()

    s.stack.Push(value)
}

func (s *ConcurrentStack[T]) PushValues(values []T) {
    s.mu.Lock()
    defer s.mu.Unlock()

    s.stack.PushValues(values)
}

func (s *ConcurrentStack[T]) Pop() (T, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    return s.stack.Pop()
}

func (s *ConcurrentStack[T]) Peek() (T, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()

    return s.stack.Peek()
}
