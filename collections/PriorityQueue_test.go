package collections

import (
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestPriorityQueue_MinQueue(t *testing.T) {
    queue := NewPriorityQueue[int](func(a, b int) int { return a - b })
    queue.Offer(3)
    queue.Offer(1)
    queue.Offer(2)

    assert.Equal(t, 3, queue.Size(), "PriorityQueue size is not equal")
    assert.True(t, queue.Contains(3), "PriorityQueue contains is failed")
    validatePriorityQueuePeek(t, queue, 1)
    validatePriorityQueuePoll(t, queue, 1)

    validatePriorityQueuePeek(t, queue, 2)
    validatePriorityQueuePoll(t, queue, 2)

    // add in the middle
    queue.OfferValues([]int{5, 4})
    assert.True(t, queue.Contains(5), "PriorityQueue contains is failed")

    validatePriorityQueuePeek(t, queue, 3)
    validatePriorityQueuePoll(t, queue, 3)

    validatePriorityQueuePeek(t, queue, 4)
    validatePriorityQueuePoll(t, queue, 4)

    validatePriorityQueuePeek(t, queue, 5)
    validatePriorityQueuePoll(t, queue, 5)
    assert.Equal(t, 0, queue.Size(), "PriorityQueue size is not equal")
    assert.True(t, queue.IsEmpty(), "PriorityQueue is not empty")
}

func TestPriorityQueue_MaxQueue(t *testing.T) {
    queue := NewPriorityQueue[int](func(a, b int) int { return b - a })
    queue.Offer(3)
    queue.Offer(1)
    queue.Offer(2)

    assert.Equal(t, 3, queue.Size(), "PriorityQueue size is not equal")
    assert.True(t, queue.Contains(3), "PriorityQueue contains is failed")
    validatePriorityQueuePeek(t, queue, 3)
    validatePriorityQueuePoll(t, queue, 3)

    validatePriorityQueuePeek(t, queue, 2)
    validatePriorityQueuePoll(t, queue, 2)

    // add in the middle
    queue.OfferValues([]int{5, 4})
    assert.True(t, queue.Contains(5), "PriorityQueue contains is failed")

    validatePriorityQueuePeek(t, queue, 5)
    validatePriorityQueuePoll(t, queue, 5)

    validatePriorityQueuePeek(t, queue, 4)
    validatePriorityQueuePoll(t, queue, 4)

    validatePriorityQueuePeek(t, queue, 1)
    validatePriorityQueuePoll(t, queue, 1)
    assert.Equal(t, 0, queue.Size(), "PriorityQueue size is not equal")
    assert.True(t, queue.IsEmpty(), "PriorityQueue is not empty")
}

func validatePriorityQueuePoll(t *testing.T, queue *PriorityQueue[int], expectedValue int) {
    value, err := queue.Poll()
    require.Nil(t, err, "PriorityQueue poll is failed")
    require.Equal(t, expectedValue, value, "PriorityQueue poll item is not matched")
}

func validatePriorityQueuePeek(t *testing.T, queue *PriorityQueue[int], expectedValue int) {
    value, err := queue.Peek()
    require.Nil(t, err, "PriorityQueue peek is failed")
    require.Equal(t, expectedValue, value, "PriorityQueue peek item is not matched")
}
