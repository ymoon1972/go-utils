package array

type Iterator[T comparable] struct {
	array *Array[T]
	index int
}

func (s *Array[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{array: s, index: 0}
}

func (it *Iterator[T]) HasNext() bool {
	return it.index < it.array.Size()
}

func (it *Iterator[T]) Next() T {
	value := it.array.items[it.index]
	it.index++
	return value
}

func (it *Iterator[T]) Each(action func(T)) {
	for it.HasNext() {
		action(it.Next())
	}
}

func Map[T comparable, V comparable](it *Iterator[T], mapper func(T) V) *Array[V] {
	result := NewArrayList[V]()
	for it.HasNext() {
		result.Add(mapper(it.Next()))
	}
	return result
}

func Reduce[T comparable, V any](it *Iterator[T], initial V, reducer func(acc V, item T) V) V {
	acc := initial
	for it.HasNext() {
		acc = reducer(acc, it.Next())
	}

	return acc
}

func Filter[T comparable](it *Iterator[T], predicate func(T) bool) *Array[T] {
	filtered := NewArrayList[T]()
	for it.HasNext() {
		item := it.Next()
		if predicate(item) {
			filtered.Add(item)
		}
	}
	return filtered
}
