package queue

import (
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

func validateQueuePoll(t *testing.T, queue *Queue[int], expectedValue int) {
    value, err := queue.Poll()
    require.Nil(t, err, "Queue poll is failed")
    require.Equal(t, expectedValue, value, "Queue poll item is not matched")
}

func validateQueuePeek(t *testing.T, queue *Queue[int], expectedValue int) {
    value, err := queue.Peek()
    require.Nil(t, err, "Queue peek is failed")
    require.Equal(t, expectedValue, value, "Queue peek item is not matched")
}
