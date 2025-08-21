package stack

import (
    "math/rand"
    "sort"
    "sync"
    "sync/atomic"
    "testing"
    "time"

    "github.com/stretchr/testify/require"
)

func TestConcurrentStack(t *testing.T) {
    stack := NewConcurrentStack[int]()
    require.True(t, stack.IsEmpty(), "ConcurrentStack is not empty")

    var wg sync.WaitGroup

    // generate value from 0 to 99
    for p := 0; p < 10; p++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for i := 0; i < 10; i++ {
                time.Sleep(time.Duration(100+rand.Intn(100)) * time.Millisecond)
                stack.Push(i + p*10)
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
                if !stack.IsEmpty() {
                    value, _ := stack.Pop()
                    consumed = append(consumed, value)
                }
            }
        }()
    }

    wg.Wait()

    require.True(t, stack.IsEmpty(), "ConcurrentStack is not empty")
    require.Equal(t, 100, len(consumed), "ConcurrentStack consumed size is not equal")

    sort.Ints(consumed)
    for i := 0; i < 100; i++ {
        require.Equal(t, i, consumed[i], "ConcurrentStack consumed item is not matched")
    }
}

func TestConcurrentStack_Order(t *testing.T) {
    stack := NewConcurrentStack[int]()
    require.True(t, stack.IsEmpty(), "ConcurrentStack is not empty")

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
                stack.Push(int(counter))
                atomic.AddInt32(&counter, 1)
                mutex.Unlock()
            }
        }()
    }
    wg.Wait()

    // consume 100 items
    consumed := make([]int, 0, 100)
    for !stack.IsEmpty() {
        value, _ := stack.Pop()
        consumed = append(consumed, value)
    }
    require.True(t, stack.IsEmpty(), "ConcurrentStack is not empty")
    require.Equal(t, 100, len(consumed), "ConcurrentStack consumed size is not equal")
    for i := 0; i < 100; i++ {
        require.Equal(t, 99-i, consumed[i], "ConcurrentStack consumed item is not matched")
    }
}
