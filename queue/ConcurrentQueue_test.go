package queue

import (
    "math/rand"
    "sync"
    "sync/atomic"
    "testing"
    "time"

    "github.com/stretchr/testify/require"
)

func TestConcurrentQueue(t *testing.T) {
    queue := NewConcurrentQueue[int]()
    require.True(t, queue.IsEmpty(), "ConcurrentQueue is not empty")

    var wg sync.WaitGroup
    var mutex sync.Mutex

    // generate value from 0 to 99
    var counter int32
    for p := 0; p < 10; p++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for i := 0; i < 10; i++ {
                time.Sleep(time.Duration(100+rand.Intn(100)) * time.Millisecond)

                mutex.Lock()
                queue.Offer(int(counter))
                atomic.AddInt32(&counter, 1)
                mutex.Unlock()
            }
        }()
    }

    // consume 100 items
    consumed := make([]int, 0, 100)
    for c := 0; c < 10; c++ {
        wg.Add(1)
        go func() {
            defer wg.Done()

            for len(consumed) < 100 {
                time.Sleep(time.Duration(100+rand.Intn(100)) * time.Millisecond)
                if !queue.IsEmpty() {
                    value, _ := queue.Poll()
                    consumed = append(consumed, value)
                }
            }
        }()
    }

    wg.Wait()

    require.True(t, queue.IsEmpty(), "ConcurrentQueue is not empty")
    require.Equal(t, 100, len(consumed), "ConcurrentQueue consumed size is not equal")
    for i := 0; i < 100; i++ {
        require.Equal(t, i, consumed[i], "ConcurrentQueue consumed item is not matched")
    }
}
