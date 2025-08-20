package collections

import (
    "errors"
    "fmt"
    "sort"
)

type ArrayList[T comparable] struct {
    items []T
}

func NewArrayList[T comparable]() *ArrayList[T] {
    return &ArrayList[T]{}
}

func (s *ArrayList[T]) Size() int {
    return len(s.items)
}

func (s *ArrayList[T]) IsEmpty() bool {
    return s.Size() == 0
}

func (s *ArrayList[T]) Values() []T {
    return s.items
}

func (s *ArrayList[T]) Add(value T) {
    s.items = append(s.items, value)
}

func (s *ArrayList[T]) AddAll(values []T) {
    s.items = append(s.items, values...)
}

func (s *ArrayList[T]) InsertAt(index int, value T) error {
    if index < 0 || index > s.Size() {
        return errors.New(fmt.Sprintf("Index %d is out of range with size %d", index, s.Size()))
    }

    if index == 0 {
        // add to the start
        s.items = append([]T{value}, s.items...)
        return nil
    }

    // Grow the slice by one.
    s.items = append(s.items, value)

    // Shift the elements to the right.
    copy(s.items[index+1:], s.items[index:])

    // Insert the value at the given index.
    s.items[index] = value
    return nil
}

func (s *ArrayList[T]) Get(index int) (T, error) {
    if index < 0 || index >= s.Size() {
        var zero T
        return zero, errors.New(fmt.Sprintf("Index %d is out of range with size %d", index, s.Size()))
    }

    return s.items[index], nil
}

func (s *ArrayList[T]) RemoveAt(index int) (T, error) {
    if index < 0 || index >= s.Size() {
        var zero T
        return zero, errors.New(fmt.Sprintf("Index %d is out of range with size %d", index, s.Size()))
    }

    value := s.items[index]
    if index == 0 {
        // remove the start
        s.items = s.items[1:]
    } else if index == s.Size()-1 {
        // remove the end
        s.items = s.items[:s.Size()-1]
    } else {
        // remove in the middle
        s.items = append(s.items[:index], s.items[index+1:]...)
    }
    return value, nil
}

func (s *ArrayList[T]) Clone() *ArrayList[T] {
    itemsCopy := make([]T, s.Size())
    copy(itemsCopy, s.items)
    return &ArrayList[T]{items: itemsCopy}
}

func (s *ArrayList[T]) Merge(list *ArrayList[T]) {
    s.items = append(s.items, list.items...)
}

func (s *ArrayList[T]) Reverse() {
    for i, j := 0, s.Size()-1; i < j; i, j = i+1, j-1 {
        s.items[i], s.items[j] = s.items[j], s.items[i]
    }
}

func (s *ArrayList[T]) Filter(predicate func(T) bool) *ArrayList[T] {
    filtered := NewArrayList[T]()
    for _, item := range s.items {
        if predicate(item) {
            filtered.Add(item)
        }
    }
    return filtered
}

func (s *ArrayList[T]) Sort(comparator func(T, T) int) {
    sort.Slice(s.items, func(i, j int) bool {
        return comparator(s.items[i], s.items[j]) <= 0
    })
}

func (s *ArrayList[T]) Contains(value T) bool {
    for _, item := range s.items {
        if item == value {
            return true
        }
    }

    return false
}

func MapArrayList[T comparable, V comparable](s *ArrayList[T], mapper func(T) V) *ArrayList[V] {
    if s.IsEmpty() {
        return NewArrayList[V]()
    }

    result := NewArrayList[V]()
    for _, item := range s.items {
        result.Add(mapper(item))
    }
    return result
}

func ReduceArrayList[T comparable, V any](s *ArrayList[T], initial V, reducer func(acc V, item T) V) V {
    acc := initial
    for _, item := range s.items {
        acc = reducer(acc, item)
    }
    return acc
}

func (s *ArrayList[T]) Clear() {
    s.items = []T{}
}
