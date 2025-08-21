package queue

import (
    "go-utils/array"
    "math/rand"
    "sync"
    "testing"
    "time"

    "github.com/stretchr/testify/require"
)

func TestConcurrentPriorityQueue_MinQueue(t *testing.T) {
    comparator := func(a, b int) int { return a - b }
    queue := NewConcurrentPriorityQueue[int](comparator)

    var wg sync.WaitGroup

    // generate value from 0 to 99
    for p := 0; p < 10; p++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for i := 0; i < 10; i++ {
                time.Sleep(time.Duration(10+rand.Intn(10)) * time.Millisecond)
                queue.Offer(i + p*10)
            }
        }()
    }
    wg.Wait()

    for c := 0; c < 1; c++ {
        wg.Add(1)
        go func() {
            defer wg.Done()

            consumed := array.NewArrayList[int]()
            for !queue.IsEmpty() {
                time.Sleep(time.Duration(10+rand.Intn(10)) * time.Millisecond)
                value, err := queue.Poll()
                if err == nil {
                    consumed.Add(value)
                }
            }

            for i := 1; i < consumed.Size(); i++ {
                require.True(t, consumed.Compare(i-1, i, comparator) <= 0, "ConcurrentPriorityQueue is not min queue")
            }
        }()
    }
    wg.Wait()
    require.True(t, queue.IsEmpty(), "ConcurrentPriorityQueue is not empty")
}

func TestConcurrentPriorityQueue_MaxQueue(t *testing.T) {
    comparator := func(a, b int) int { return b - a }
    queue := NewConcurrentPriorityQueue[int](comparator)

    var wg sync.WaitGroup

    // generate value from 0 to 99
    for p := 0; p < 10; p++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for i := 0; i < 10; i++ {
                time.Sleep(time.Duration(10+rand.Intn(10)) * time.Millisecond)
                queue.Offer(i + p*10)
            }
        }()
    }
    wg.Wait()

    for c := 0; c < 1; c++ {
        wg.Add(1)
        go func() {
            defer wg.Done()

            consumed := array.NewArrayList[int]()
            for !queue.IsEmpty() {
                time.Sleep(time.Duration(10+rand.Intn(10)) * time.Millisecond)
                value, err := queue.Poll()
                if err == nil {
                    consumed.Add(value)
                }
            }

            for i := 1; i < consumed.Size(); i++ {
                require.True(t, consumed.Compare(i-1, i, comparator) <= 0, "ConcurrentPriorityQueue is not max queue")
            }
        }()
    }
    wg.Wait()
    require.True(t, queue.IsEmpty(), "ConcurrentPriorityQueue is not empty")
}
