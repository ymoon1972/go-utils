package list

import (
    "go-utils/array"
    "sync"
)

type ConcurrentList[T comparable] struct {
    mu   sync.RWMutex
    list *DoubleLinkedList[T]
}

func NewConcurrentList[T comparable]() *ConcurrentList[T] {
    return &ConcurrentList[T]{list: NewDoubleLinkedList[T]()}
}

func (s *ConcurrentList[T]) Size() int {
    return s.list.Size()
}

func (s *ConcurrentList[T]) IsEmpty() bool {
    return s.list.IsEmpty()
}

func (s *ConcurrentList[T]) Clear() {
    s.mu.Lock()
    defer s.mu.Unlock()

    s.list.Clear()
}

func (s *ConcurrentList[T]) Values() []T {
    s.mu.RLock()
    defer s.mu.RUnlock()

    return s.list.Values()
}

func (s *ConcurrentList[T]) Add(value T) {
    s.mu.Lock()
    defer s.mu.Unlock()

    s.list.Add(value)
}

func (s *ConcurrentList[T]) AddAll(values []T) {
    s.mu.Lock()
    defer s.mu.Unlock()

    for _, value := range values {
        s.list.Add(value)
    }
}

func (s *ConcurrentList[T]) AddHead(value T) {
    s.mu.Lock()
    defer s.mu.Unlock()

    s.list.AddHead(value)
}

func (s *ConcurrentList[T]) AddTail(value T) {
    s.mu.Lock()
    defer s.mu.Unlock()

    s.list.AddTail(value)
}

func (s *ConcurrentList[T]) InsertAt(index int, value T) {
    s.mu.Lock()
    defer s.mu.Unlock()

    s.list.InsertAt(index, value)
}

func (s *ConcurrentList[T]) GetHead() (T, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()

    return s.list.GetHead()
}

func (s *ConcurrentList[T]) GetTail() (T, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()

    return s.list.GetTail()
}

func (s *ConcurrentList[T]) GetAt(index int) (T, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()

    return s.list.GetAt(index)
}

func (s *ConcurrentList[T]) RemoveHead() (T, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    return s.list.RemoveHead()
}

func (s *ConcurrentList[T]) RemoveTail() (T, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    return s.list.RemoveTail()
}

func (s *ConcurrentList[T]) RemoveAt(index int) (T, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    return s.list.RemoveAt(index)
}

func (s *ConcurrentList[T]) Contains(value T) bool {
    s.mu.RLock()
    defer s.mu.RUnlock()

    return s.list.Contains(value)
}

func (s *ConcurrentList[T]) Clone() *ConcurrentList[T] {
    s.mu.RLock()
    defer s.mu.RUnlock()

    newList := s.list.Clone()
    return &ConcurrentList[T]{list: newList}
}

func (s *ConcurrentList[T]) Merge(concurrentList *ConcurrentList[T]) {
    s.mu.Lock()
    defer s.mu.Unlock()

    concurrentList.mu.RLock()
    defer concurrentList.mu.RUnlock()

    s.list.Merge(concurrentList.list)
}

func (s *ConcurrentList[T]) MergeArray(arr *array.List[T]) {
    s.mu.Lock()
    defer s.mu.Unlock()

    s.list.MergeArray(arr)
}

func (s *ConcurrentList[T]) MergeList(list *LinkedList[T]) {
    s.mu.Lock()
    defer s.mu.Unlock()

    s.list.MergeList(list)
}

func (s *ConcurrentList[T]) Reverse() {
    s.mu.Lock()
    defer s.mu.Unlock()

    s.list.Reverse()
}

func (s *ConcurrentList[T]) Filter(predicate func(T) bool) *DoubleLinkedList[T] {
    s.mu.RLock()
    defer s.mu.RUnlock()

    return s.list.Filter(predicate)
}

func (s *ConcurrentList[T]) Sort(comparator func(T, T) int) {
    s.mu.Lock()
    defer s.mu.Unlock()

    s.list.Sort(comparator)
}

func MapConcurrentList[T, V comparable](s *ConcurrentList[T], mapper func(T) V) *DoubleLinkedList[V] {
    s.mu.RLock()
    defer s.mu.RUnlock()

    return MapDoubleLinkedList[T, V](s.list, mapper)
}

func ReduceConcurrentList[T comparable, V any](s *ConcurrentList[T], initial V, reducer func(V, T) V) V {
    s.mu.RLock()
    defer s.mu.RUnlock()

    return ReduceDoubleLinkedList[T, V](s.list, initial, reducer)
}
