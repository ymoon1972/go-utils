package list

type Iterator[T comparable] struct {
	node *doubleNode[T]
	end  *doubleNode[T]
}

func (s *LinkedList[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{node: s.head.next, end: s.tail}
}

func (it *Iterator[T]) HasNext() bool {
	return it.node != it.end
}

func (it *Iterator[T]) Next() T {
	value := it.node.value
	it.node = it.node.next
	return value
}

func (it *Iterator[T]) Each(action func(T)) {
	for it.HasNext() {
		action(it.Next())
	}
}

func Map[T comparable, V comparable](it *Iterator[T], mapper func(T) V) *LinkedList[V] {
	result := NewLinkedList[V]()
	for it.HasNext() {
		result.Add(mapper(it.Next()))
	}
	return result
}

func Reduce[T comparable, V any](it *Iterator[T], initial V, reducer func(V, T) V) V {
	acc := initial
	for it.HasNext() {
		acc = reducer(acc, it.Next())
	}
	return acc
}

func Filter[T comparable](it *Iterator[T], predicate func(T) bool) *LinkedList[T] {
	filtered := NewLinkedList[T]()
	for it.HasNext() {
		value := it.Next()
		if predicate(value) {
			filtered.Add(value)
		}
	}
	return filtered
}
