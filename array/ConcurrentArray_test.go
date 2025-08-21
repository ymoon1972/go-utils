package array

import (
    "fmt"
    "math/rand"
    "strconv"
    "sync"
    "sync/atomic"
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestConcurrentArrayList_Add(t *testing.T) {
    arr := createConcurrentArray(5)
    require.Equal(t, 10, arr.Size(), "ConcurrentArrayList size is not equal")

    arr.Sort(func(a int, b int) int { return a - b })
    require.ElementsMatch(t, []int{0, 1, 2, 3, 4, 10, 11, 12, 13, 14}, arr.Values(), "ConcurrentArrayList values are not equal")
    validateConcurrentArray(t, arr, 0, 0)
    validateConcurrentArray(t, arr, 9, 14)
}

func TestConcurrentArrayList_Order(t *testing.T) {
    arr := NewConcurrentArray[int]()

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
                arr.Add(int(counter))
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
                if !arr.IsEmpty() {
                    value, _ := arr.RemoveAt(0)
                    consumed = append(consumed, value)
                }
            }
        }()
    }

    wg.Wait()

    require.True(t, arr.IsEmpty(), "ConcurrentArray is not empty")
    require.Equal(t, 100, len(consumed), "ConcurrentArray consumed size is not equal")
    for i := 0; i < 100; i++ {
        require.Equal(t, i, consumed[i], "ConcurrentArray consumed item is not matched")
    }
}

func TestConcurrentArrayList_InsertAt(t *testing.T) {
    arr := NewConcurrentArray[int]()
    var wg sync.WaitGroup

    // add 0 to 4
    wg.Add(1)
    go func() {
        for i := 0; i < 5; i++ {
            time.Sleep(time.Duration(100+rand.Intn(100)) * time.Millisecond)
            arr.Add(i)
        }
        wg.Done()
    }()

    // insert 10 to 14
    wg.Add(1)
    go func() {
        iter := 0
        for iter < 5 {
            time.Sleep(time.Duration(100+rand.Intn(100)) * time.Millisecond)
            if arr.IsEmpty() {
                continue
            }

            index := rand.Intn(arr.Size())
            arr.InsertAt(index, iter+10)
            iter++
        }
        wg.Done()
    }()
    wg.Wait()

    assert.True(t, arr.Contains(3), "ConcurrentArray Contains(3) is failed")
    assert.True(t, arr.Contains(12), "ConcurrentArray Contains(12) is failed")

    arr.Sort(func(a, b int) int { return a - b })
    assert.ElementsMatch(t, []int{0, 1, 2, 3, 4, 10, 11, 12, 13, 14}, arr.Values(), "ConcurrentArray values are not equal")
}

func TestConcurrentArray_RemoveAt(t *testing.T) {
    arr := NewConcurrentArray[int]()
    var wg sync.WaitGroup

    // add 0 to 9
    wg.Add(1)
    go func() {
        for i := 0; i < 10; i++ {
            time.Sleep(time.Duration(100+rand.Intn(100)) * time.Millisecond)
            arr.Add(i)
        }
        wg.Done()
    }()

    // randomly remove 5 items
    removedValues := NewArrayList[int]()
    wg.Add(1)
    go func() {
        iter := 0
        for iter < 5 {
            time.Sleep(time.Duration(100+rand.Intn(110)) * time.Millisecond)
            if arr.IsEmpty() {
                continue
            }

            index := rand.Intn(arr.Size())
            removedValue, _ := arr.RemoveAt(index)
            removedValues.Add(removedValue)
            iter++
        }
        wg.Done()
    }()
    wg.Wait()

    assert.Equal(t, 5, arr.Size(), "ConcurrentArray size is not equal")
    assert.Equal(t, 5, removedValues.Size(), "ConcurrentArray removed values size is not equal")

    for _, value := range removedValues.Values() {
        assert.False(t, arr.Contains(value), fmt.Sprintf("ConcurrentArray contains removed value: %d", value))
    }

    arr.Merge(removedValues)
    arr.Sort(func(a, b int) int { return a - b })
    assert.ElementsMatch(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, arr.Values(), "ConcurrentArray values are not equal")
}

func TestConcurrentArray_Reverse(t *testing.T) {
    arr := createConcurrentArray(5)

    arr.Sort(func(a, b int) int { return a - b })
    arr.Reverse()
    require.ElementsMatch(t, []int{14, 13, 12, 11, 10, 4, 3, 2, 1, 0}, arr.Values(), "ConcurrentArray reverse is failed")
    validateConcurrentArray(t, arr, 0, 14)
    validateConcurrentArray(t, arr, 9, 0)
}

func TestConcurrentArray_Filter(t *testing.T) {
    arr := createConcurrentArray(5)
    arr.Sort(func(a, b int) int { return a - b })

    even := arr.Filter(func(value int) bool { return value%2 == 0 })
    require.ElementsMatch(t, []int{0, 2, 4, 10, 12, 14}, even.Values(), "ConcurrentArray filter is failed")
}

func TestConcurrentArray_Sort(t *testing.T) {
    arr := createConcurrentArray(5)
    arr.Sort(func(a, b int) int { return b - a })
    require.ElementsMatch(t, []int{14, 13, 12, 11, 10, 4, 3, 2, 1, 0}, arr.Values(), "ConcurrentArray reverse is failed")
    validateConcurrentArray(t, arr, 0, 14)
    validateConcurrentArray(t, arr, 9, 0)

    arr.Sort(func(a int, b int) int { return a - b })
    assert.ElementsMatch(t, []int{0, 1, 2, 3, 4, 10, 11, 12, 13, 14}, arr.Values(), "ConcurrentArray values are not equal")
    validateConcurrentArray(t, arr, 0, 0)
    validateConcurrentArray(t, arr, 9, 14)
}

func TestConcurrentArray_MapConcurrentArray(t *testing.T) {
    arr := createConcurrentArray(5)
    arr.Sort(func(a, b int) int { return b - a })

    mapped := MapConcurrentArray(arr, func(item int) string { return strconv.Itoa(item) })
    require.ElementsMatch(t, []string{"0", "1", "2", "3", "4", "10", "11", "12", "13", "14"}, mapped.Values(), "ConcurrentArray map is failed")
}

func TestConcurrentArray_ReduceConcurrentArray(t *testing.T) {
    arr := createConcurrentArray(5)
    arr.Sort(func(a, b int) int { return b - a })

    sum := ReduceConcurrentArray(arr, 0, func(acc int, item int) int { return acc + item })
    require.Equal(t, 70, sum, "ConcurrentArray reduce is failed")
}

func createConcurrentArray(items int) *ConcurrentArray[int] {
    arr := NewConcurrentArray[int]()
    var wg sync.WaitGroup

    // add 0 to 4
    wg.Add(1)
    go func() {
        for i := 0; i < items; i++ {
            time.Sleep(time.Duration(100+rand.Intn(100)) * time.Millisecond)
            arr.Add(i)
        }
        wg.Done()
    }()

    // add 10 to 14
    wg.Add(1)
    go func() {
        for i := 0; i < items; i++ {
            time.Sleep(time.Duration(100+rand.Intn(100)) * time.Millisecond)
            arr.Add(i + 10)
        }
        wg.Done()
    }()
    wg.Wait()

    return arr
}

func validateConcurrentArray(t *testing.T, arr *ConcurrentArray[int], index, expectedValue int) {
    value, err := arr.Get(index)
    require.Nil(t, err, "ConcurrentArray get is failed")
    require.Equal(t, expectedValue, value, "ConcurrentArray get item is not matched")
}
