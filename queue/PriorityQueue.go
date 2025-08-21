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
        top := (pos - 1) / 2
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
    last, _ := s.RemoveAt(s.Size() - 1)
    if !s.IsEmpty() {
        s.SetAt(0, last)

        length := s.Size()
        index := 0
        for {
            left := index*2 + 1
            right := index*2 + 2
            smallest := index

            if left < length && s.Compare(left, smallest, s.comparator) <= 0 {
                smallest = left
            }
            if right < length && s.Compare(right, smallest, s.comparator) <= 0 {
                smallest = right
            }

            if smallest == index {
                break
            }

            s.Swap(smallest, index)
            index = smallest
        }
    }

    return value, nil
}
