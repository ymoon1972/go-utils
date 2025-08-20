package collections

type Queue[T comparable] struct {
    *LinkedList[T]
}

func NewQueue[T comparable]() *Queue[T] {
    return &Queue[T]{NewLinkedList[T]()}
}

func (s *Queue[T]) Offer(value T) {
    s.AddTail(value)
}

func (s *Queue[T]) OfferValues(values []T) {
    s.AddAll(values)
}

func (s *Queue[T]) Poll() (T, error) {
    return s.RemoveHead()
}

func (s *Queue[T]) Peek() (T, error) {
    return s.GetHead()
}
