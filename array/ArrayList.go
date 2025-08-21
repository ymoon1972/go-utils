package array

import (
    "errors"
    "fmt"
    "sort"
)

type List[T comparable] struct {
    items []T
}

func NewArrayList[T comparable]() *List[T] {
    return &List[T]{}
}

func (s *List[T]) Size() int {
    return len(s.items)
}

func (s *List[T]) IsEmpty() bool {
    return s.Size() == 0
}

func (s *List[T]) Values() []T {
    return s.items
}

func (s *List[T]) Add(value T) {
    s.items = append(s.items, value)
}

func (s *List[T]) AddAll(values []T) {
    s.items = append(s.items, values...)
}

func (s *List[T]) InsertAt(index int, value T) error {
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

func (s *List[T]) Get(index int) (T, error) {
    if index < 0 || index >= s.Size() {
        var zero T
        return zero, errors.New(fmt.Sprintf("Index %d is out of range with size %d", index, s.Size()))
    }

    return s.items[index], nil
}

func (s *List[T]) RemoveAt(index int) (T, error) {
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

func (s *List[T]) Clone() *List[T] {
    itemsCopy := make([]T, s.Size())
    copy(itemsCopy, s.items)
    return &List[T]{items: itemsCopy}
}

func (s *List[T]) Merge(list *List[T]) {
    s.items = append(s.items, list.items...)
}

func (s *List[T]) Reverse() {
    for i, j := 0, s.Size()-1; i < j; i, j = i+1, j-1 {
        s.items[i], s.items[j] = s.items[j], s.items[i]
    }
}

func (s *List[T]) Compare(left, right int, comparator func(a, b T) int) int {
    return comparator(s.items[left], s.items[right])
}

func (s *List[T]) Swap(left, right int) {
    s.items[left], s.items[right] = s.items[right], s.items[left]
}

func (s *List[T]) Filter(predicate func(T) bool) *List[T] {
    filtered := NewArrayList[T]()
    for _, item := range s.items {
        if predicate(item) {
            filtered.Add(item)
        }
    }
    return filtered
}

func (s *List[T]) Sort(comparator func(T, T) int) {
    sort.Slice(s.items, func(i, j int) bool {
        return comparator(s.items[i], s.items[j]) <= 0
    })
}

func (s *List[T]) Contains(value T) bool {
    for _, item := range s.items {
        if item == value {
            return true
        }
    }

    return false
}

func MapArrayList[T comparable, V comparable](s *List[T], mapper func(T) V) *List[V] {
    if s.IsEmpty() {
        return NewArrayList[V]()
    }

    result := NewArrayList[V]()
    for _, item := range s.items {
        result.Add(mapper(item))
    }
    return result
}

func ReduceArrayList[T comparable, V any](s *List[T], initial V, reducer func(acc V, item T) V) V {
    acc := initial
    for _, item := range s.items {
        acc = reducer(acc, item)
    }
    return acc
}

func (s *List[T]) Clear() {
    s.items = []T{}
}
