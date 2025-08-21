package queue

import "sync"

type ConcurrentQueue[T comparable] struct {
    mu    sync.RWMutex
    queue *Queue[T]
}

func NewConcurrentQueue[T comparable]() *ConcurrentQueue[T] {
    return &ConcurrentQueue[T]{queue: NewQueue[T]()}
}

func (q *ConcurrentQueue[T]) Size() int {
    return q.queue.Size()
}

func (q *ConcurrentQueue[T]) IsEmpty() bool {
    return q.queue.IsEmpty()
}

func (q *ConcurrentQueue[T]) Clear() {
    q.mu.Lock()
    defer q.mu.Unlock()

    q.queue.Clear()
}

func (q *ConcurrentQueue[T]) Values() []T {
    q.mu.RLock()
    defer q.mu.RUnlock()

    return q.queue.Values()
}

func (q *ConcurrentQueue[T]) Offer(value T) {
    q.mu.Lock()
    defer q.mu.Unlock()

    q.queue.Offer(value)
}

func (q *ConcurrentQueue[T]) OfferValues(values []T) {
    q.mu.Lock()
    defer q.mu.Unlock()

    q.queue.AddAll(values)
}

func (q *ConcurrentQueue[T]) Poll() (T, error) {
    q.mu.Lock()
    defer q.mu.Unlock()

    return q.queue.Poll()
}

func (q *ConcurrentQueue[T]) Peek() (T, error) {
    q.mu.RLock()
    defer q.mu.RUnlock()

    return q.queue.Peek()
}
