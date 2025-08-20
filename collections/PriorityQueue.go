package collections

import (
    "errors"
)

type PriorityQueue[T comparable] struct {
    *ArrayList[T]
    comparator func(a, b T) int
}

func NewPriorityQueue[T comparable](comparator func(a, b T) int) *PriorityQueue[T] {
    return &PriorityQueue[T]{
        NewArrayList[T](),
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
        if s.comparator(s.items[top], s.items[pos]) <= 0 {
            break
        }

        // swap
        s.items[top], s.items[pos] = s.items[pos], s.items[top]
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

    return s.items[0], nil
}

func (s *PriorityQueue[T]) Poll() (T, error) {
    if s.IsEmpty() {
        var zero T
        return zero, errors.New("queue is empty")
    }

    value := s.items[0]
    s.items[0] = s.items[s.Size()-1]
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
            if s.comparator(s.items[top], s.items[left]) > 0 {
                s.items[top], s.items[left] = s.items[left], s.items[top]
            }
            break
        }

        if s.comparator(s.items[left], s.items[right]) <= 0 {
            // take the left child
            if s.comparator(s.items[top], s.items[left]) > 0 {
                s.items[top], s.items[left] = s.items[left], s.items[top]
                top = left
            }
            continue
        }

        // take the right child
        if s.comparator(s.items[top], s.items[right]) <= 0 {
            break
        }
        s.items[top], s.items[right] = s.items[right], s.items[top]
        top = right
    }

    return value, nil
}
