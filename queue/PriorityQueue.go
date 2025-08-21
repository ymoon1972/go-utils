package queue

import (
    "errors"
    "go-utils/array"
)

type PriorityQueue[T comparable] struct {
    *array.List[T]
    comparator func(a, b T) int
}

func NewPriorityQueue[T comparable](comparator func(a, b T) int) *PriorityQueue[T] {
    return &PriorityQueue[T]{
        array.NewArrayList[T](),
        comparator,
    }
}

func (s *PriorityQueue[T]) Offer(value T) {
    // add to the tail
    s.Add(value)

    // blow up
    pos := s.Size() - 1
    for pos > 0 {
        top := pos / 2
        if s.Compare(top, pos, s.comparator) <= 0 {
            break
        }

        // swap
        s.Swap(top, pos)
        pos = top
    }
}

func (s *PriorityQueue[T]) OfferValues(values []T) {
    for _, value := range values {
        s.Offer(value)
    }
}

func (s *PriorityQueue[T]) Peek() (T, error) {
    if s.IsEmpty() {
        var zero T
        return zero, errors.New("queue is empty")
    }

    return s.Get(0)
}

func (s *PriorityQueue[T]) Poll() (T, error) {
    if s.IsEmpty() {
        var zero T
        return zero, errors.New("queue is empty")
    }

    value, _ := s.Get(0)
    s.Swap(0, s.Size()-1)
    _, _ = s.RemoveAt(s.Size() - 1)

    // bubble down the values
    top := 0
    for top < s.Size() {
        left := top*2 + 1
        right := top*2 + 2
        if left >= s.Size() {
            // out of range
            break
        }

        if right >= s.Size() {
            // only a left child exists
            if s.Compare(top, left, s.comparator) > 0 {
                s.Swap(top, left)
            }
            break
        }

        if s.Compare(left, right, s.comparator) <= 0 {
            // take the left child
            if s.Compare(top, left, s.comparator) > 0 {
                s.Swap(top, left)
                top = left
            }
            continue
        }

        // take the right child
        if s.Compare(top, right, s.comparator) <= 0 {
            break
        }
        s.Swap(top, right)
        top = right
    }

    return value, nil
}
