package queue

import "sync"

type ConcurrentPriorityQueue[T comparable] struct {
    mu    sync.RWMutex
    queue *PriorityQueue[T]
}

func NewConcurrentPriorityQueue[T comparable](comparator func(a, b T) int) *ConcurrentPriorityQueue[T] {
    return &ConcurrentPriorityQueue[T]{queue: NewPriorityQueue[T](comparator)}
}

func (s *ConcurrentPriorityQueue[T]) Size() int {
    return s.queue.Size()
}

func (s *ConcurrentPriorityQueue[T]) IsEmpty() bool {
    return s.queue.IsEmpty()
}

func (s *ConcurrentPriorityQueue[T]) Clear() {
    s.mu.Lock()
    defer s.mu.Unlock()

    s.queue.Clear()
}

func (s *ConcurrentPriorityQueue[T]) Values() []T {
    s.mu.RLock()
    defer s.mu.RUnlock()

    return s.queue.Values()
}

func (s *ConcurrentPriorityQueue[T]) Offer(value T) {
    s.mu.Lock()
    defer s.mu.Unlock()

    s.queue.Offer(value)
}

func (s *ConcurrentPriorityQueue[T]) OfferValues(values []T) {
    s.mu.Lock()
    defer s.mu.Unlock()

    s.queue.OfferValues(values)
}

func (s *ConcurrentPriorityQueue[T]) Poll() (T, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    return s.queue.Poll()
}

func (s *ConcurrentPriorityQueue[T]) Peek() (T, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()

    return s.queue.Peek()
}
