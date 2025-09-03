package array

import "sync"

type ConcurrentArray[T comparable] struct {
	mu  sync.RWMutex
	arr *Array[T]
}

func NewConcurrentArray[T comparable]() *ConcurrentArray[T] {
	return &ConcurrentArray[T]{arr: NewArrayList[T]()}
}

func (s *ConcurrentArray[T]) Size() int {
	return s.arr.Size()
}

func (s *ConcurrentArray[T]) IsEmpty() bool {
	return s.arr.IsEmpty()
}

func (s *ConcurrentArray[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.arr.Clear()
}

func (s *ConcurrentArray[T]) Values() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.arr.Values()
}

func (s *ConcurrentArray[T]) Add(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.arr.Add(value)
}

func (s *ConcurrentArray[T]) AddAll(values []T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.arr.AddAll(values)
}

func (s *ConcurrentArray[T]) InsertAt(index int, value T) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.arr.InsertAt(index, value)
}

func (s *ConcurrentArray[T]) RemoveAt(index int) (T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.arr.RemoveAt(index)
}

func (s *ConcurrentArray[T]) Get(index int) (T, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.arr.Get(index)
}

func (s *ConcurrentArray[T]) Contains(value T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.arr.Contains(value)
}

func (s *ConcurrentArray[T]) Clone() *ConcurrentArray[T] {
	s.mu.RLock()
	defer s.mu.RUnlock()

	newList := s.arr.Clone()
	return &ConcurrentArray[T]{arr: newList}
}

func (s *ConcurrentArray[T]) Merge(list *Array[T]) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.arr.Merge(list)
}

func (s *ConcurrentArray[T]) Reverse() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.arr.Reverse()
}

func (s *ConcurrentArray[T]) Compare(left, right int, comparator func(a, b T) int) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.arr.Compare(left, right, comparator)
}

func (s *ConcurrentArray[T]) Swap(left, right int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.arr.Swap(left, right)
}

func (s *ConcurrentArray[T]) Filter(predicate func(T) bool) *Array[T] {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.arr.Filter(predicate)
}

func (s *ConcurrentArray[T]) Sort(comparator func(T, T) int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.arr.Sort(comparator)
}

func MapConcurrentArray[T comparable, V comparable](s *ConcurrentArray[T], mapper func(T) V) *Array[V] {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return MapArray(s.arr, mapper)
}

func ReduceConcurrentArray[T comparable, V any](s *ConcurrentArray[T], initial V, reducer func(acc V, item T) V) V {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return ReduceArray(s.arr, initial, reducer)
}
