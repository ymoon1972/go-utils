package queue

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {
	queue := NewQueue[int]()
	require.True(t, queue.IsEmpty(), "Queue is not empty")

	queue.OfferValues([]int{1, 2, 3, 4})
	require.False(t, queue.IsEmpty(), "Queue is empty")
	require.Equal(t, 4, queue.Size(), "Queue size is not equal")
	validateQueuePeek(t, queue, 1)
	validateQueuePoll(t, queue, 1)
	validateQueuePeek(t, queue, 2)
	validateQueuePoll(t, queue, 2)

	// add in the middle
	queue.Offer(10)

	// continue poll
	validateQueuePeek(t, queue, 3)
	validateQueuePoll(t, queue, 3)
	validateQueuePeek(t, queue, 4)
	validateQueuePoll(t, queue, 4)
	validateQueuePeek(t, queue, 10)
	validateQueuePoll(t, queue, 10)
	require.True(t, queue.IsEmpty(), "Queue is not empty")

	queue.Offer(10)
	require.Equal(t, 1, queue.Size(), "Queue size is not equal")
	validateQueuePeek(t, queue, 10)
	validateQueuePoll(t, queue, 10)
}

func TestQueue_Each(t *testing.T) {
	queue := NewQueue[int]()
	queue.OfferValues([]int{1, 2, 3, 4})

	cloned := NewQueue[int]()
	queue.Iterator().Each(func(value int) {
		cloned.Offer(value)
	})
	require.False(t, cloned.IsEmpty(), "Queue is empty")
	require.Equal(t, 4, cloned.Size(), "Queue size is not equal")
	validateQueuePoll(t, cloned, 1)
	validateQueuePoll(t, cloned, 2)
	validateQueuePoll(t, cloned, 3)
	validateQueuePoll(t, cloned, 4)
}

func TestQueue_Map(t *testing.T) {
	queue := NewQueue[int]()
	queue.OfferValues([]int{1, 2, 3, 4})

	mapped := Map[int, string](queue.Iterator(), func(value int) string { return strconv.Itoa(value) })
	require.Equal(t, 4, mapped.Size(), "Queue size is not equal")
	validateQueuePoll(t, mapped, "1")
	validateQueuePoll(t, mapped, "2")
	validateQueuePoll(t, mapped, "3")
	validateQueuePoll(t, mapped, "4")
}

func TestQueue_Reduce(t *testing.T) {
	queue := NewQueue[int]()
	queue.OfferValues([]int{1, 2, 3, 4})

	result := Reduce[string, string](
		Map[int, string](queue.Iterator(), func(value int) string { return strconv.Itoa(value) }).Iterator(),
		"",
		func(acc string, value string) string {
			if len(acc) == 0 {
				return value
			}

			return acc + "-" + value
		},
	)
	require.Equal(t, "1-2-3-4", result, "Reduce result is not equal")
}

func TestQueue_Filter(t *testing.T) {
	queue := NewQueue[int]()
	queue.OfferValues([]int{1, 2, 3, 4})

	filtered := Filter[int](queue.Iterator(), func(value int) bool { return value%2 == 0 })
	require.False(t, filtered.IsEmpty(), "Queue is empty")
	require.Equal(t, 2, filtered.Size(), "Queue size is not equal")
	validateQueuePoll(t, filtered, 2)
	validateQueuePoll(t, filtered, 4)
}

func validateQueuePoll[T comparable](t *testing.T, queue *Queue[T], expectedValue T) {
	value, err := queue.Poll()
	require.Nil(t, err, "Queue poll is failed")
	require.Equal(t, expectedValue, value, "Queue poll item is not matched")
}

func validateQueuePeek[T comparable](t *testing.T, queue *Queue[T], expectedValue T) {
	value, err := queue.Peek()
	require.Nil(t, err, "Queue peek is failed")
	require.Equal(t, expectedValue, value, "Queue peek item is not matched")
}
