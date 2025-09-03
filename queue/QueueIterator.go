package queue

type Iterator[T comparable] struct {
	queue *Queue[T]
	index int
}

func (s *Queue[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{queue: s, index: 0}
}

func (it *Iterator[T]) HasNext() bool {
	return it.index < it.queue.Size()
}

func (it *Iterator[T]) Next() T {
	value, _ := it.queue.GetAt(it.index)
	it.index++
	return value
}

func (it *Iterator[T]) Each(action func(T)) {
	for it.HasNext() {
		action(it.Next())
	}
}

func Map[T comparable, V comparable](it *Iterator[T], mapper func(T) V) *Queue[V] {
	result := NewQueue[V]()
	for it.HasNext() {
		result.Add(mapper(it.Next()))
	}
	return result
}

func Reduce[T comparable, V comparable](it *Iterator[T], initial V, reducer func(V, T) V) V {
	acc := initial
	for it.HasNext() {
		acc = reducer(acc, it.Next())
	}
	return acc
}

func Filter[T comparable](it *Iterator[T], predicate func(T) bool) *Queue[T] {
	filtered := NewQueue[T]()
	for it.HasNext() {
		value := it.Next()
		if predicate(value) {
			filtered.Add(value)
		}
	}
	return filtered
}
