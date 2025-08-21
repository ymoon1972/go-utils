package list

import (
    "fmt"
    "go-utils/array"
    "math/rand"
    "strconv"
    "sync"
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestConcurrentList_Add(t *testing.T) {
    list := createConcurrentList(5)

    assert.True(t, list.Contains(3), "ConcurrentList Contains(3) is failed")
    assert.True(t, list.Contains(12), "ConcurrentList Contains(12) is failed")

    list.Sort(func(a, b int) int { return a - b })
    assert.ElementsMatch(t, []int{0, 1, 2, 3, 4, 10, 11, 12, 13, 14}, list.Values(), "ConcurrentList values are not equal")
}

func TestConcurrentList_InsertAt(t *testing.T) {
    list := NewConcurrentList[int]()
    var wg sync.WaitGroup

    // add 0 to 4
    wg.Add(1)
    go func() {
        for i := 0; i < 5; i++ {
            time.Sleep(time.Duration(100+rand.Intn(100)) * time.Millisecond)
            list.Add(i)
        }
        wg.Done()
    }()

    // insert 10 to 14
    wg.Add(1)
    go func() {
        iter := 0
        for iter < 5 {
            time.Sleep(time.Duration(100+rand.Intn(100)) * time.Millisecond)
            if list.IsEmpty() {
                continue
            }

            index := rand.Intn(list.Size())
            list.InsertAt(index, iter+10)
            iter++
        }
        wg.Done()
    }()
    wg.Wait()

    assert.True(t, list.Contains(3), "ConcurrentList Contains(3) is failed")
    assert.True(t, list.Contains(12), "ConcurrentList Contains(12) is failed")

    list.Sort(func(a, b int) int { return a - b })
    assert.ElementsMatch(t, []int{0, 1, 2, 3, 4, 10, 11, 12, 13, 14}, list.Values(), "ConcurrentList values are not equal")
}

func TestConcurrentList_RemoveAt(t *testing.T) {
    list := NewConcurrentList[int]()
    var wg sync.WaitGroup

    // add 0 to 9
    wg.Add(1)
    go func() {
        for i := 0; i < 10; i++ {
            time.Sleep(time.Duration(100+rand.Intn(100)) * time.Millisecond)
            list.Add(i)
        }
        wg.Done()
    }()

    // randomly remove 5 items
    removedValues := array.NewArrayList[int]()
    wg.Add(1)
    go func() {
        iter := 0
        for iter < 5 {
            time.Sleep(time.Duration(100+rand.Intn(110)) * time.Millisecond)
            if list.IsEmpty() {
                continue
            }

            index := rand.Intn(list.Size())
            removedValue, _ := list.RemoveAt(index)
            removedValues.Add(removedValue)
            iter++
        }
        wg.Done()
    }()
    wg.Wait()

    assert.Equal(t, 5, list.Size(), "ConcurrentList size is not equal")
    assert.Equal(t, 5, removedValues.Size(), "ConcurrentList removed values size is not equal")

    for _, value := range removedValues.Values() {
        assert.False(t, list.Contains(value), fmt.Sprintf("ConcurrentList contains removed value: %d", value))
    }

    list.MergeArray(removedValues)
    list.Sort(func(a, b int) int { return a - b })
    assert.ElementsMatch(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, list.Values(), "ConcurrentList values are not equal")
}

func TestConcurrentList_Reverse(t *testing.T) {
    list := createConcurrentList(5)

    list.Sort(func(a, b int) int { return a - b })
    list.Reverse()
    require.ElementsMatch(t, []int{14, 13, 12, 11, 10, 4, 3, 2, 1, 0}, list.Values(), "ConcurrentList reverse is failed")
    validateConcurrentListHead(t, list, 14)
    validateConcurrentListTail(t, list, 0)
}

func TestConcurrentList_Filter(t *testing.T) {
    list := createConcurrentList(5)
    list.Sort(func(a, b int) int { return a - b })

    even := list.Filter(func(value int) bool { return value%2 == 0 })
    require.ElementsMatch(t, []int{0, 2, 4, 10, 12, 14}, even.Values(), "ConcurrentList filter is failed")
}

func TestConcurrentList_Sort(t *testing.T) {
    list := createConcurrentList(5)
    list.Sort(func(a, b int) int { return b - a })
    require.ElementsMatch(t, []int{14, 13, 12, 11, 10, 4, 3, 2, 1, 0}, list.Values(), "ConcurrentList reverse is failed")
    validateConcurrentListHead(t, list, 14)
    validateConcurrentListTail(t, list, 0)

    list.Sort(func(a int, b int) int { return a - b })
    assert.ElementsMatch(t, []int{0, 1, 2, 3, 4, 10, 11, 12, 13, 14}, list.Values(), "ConcurrentList values are not equal")
    validateConcurrentListHead(t, list, 0)
    validateConcurrentListTail(t, list, 14)
}

func TestConcurrentList_MapConcurrentList(t *testing.T) {
    list := createConcurrentList(5)
    list.Sort(func(a, b int) int { return b - a })

    mapped := MapConcurrentList(list, func(item int) string { return strconv.Itoa(item) })
    require.ElementsMatch(t, []string{"0", "1", "2", "3", "4", "10", "11", "12", "13", "14"}, mapped.Values(), "ConcurrentList map is failed")
}

func TestConcurrentList_ReduceConcurrentList(t *testing.T) {
    list := createConcurrentList(5)
    list.Sort(func(a, b int) int { return b - a })

    sum := ReduceConcurrentList(list, 0, func(acc int, item int) int { return acc + item })
    require.Equal(t, 70, sum, "ConcurrentList reduce is failed")
}

func createConcurrentList(items int) *ConcurrentList[int] {
    list := NewConcurrentList[int]()
    var wg sync.WaitGroup

    // add 0 to 4
    wg.Add(1)
    go func() {
        for i := 0; i < items; i++ {
            time.Sleep(time.Duration(100+rand.Intn(100)) * time.Millisecond)
            list.Add(i)
        }
        wg.Done()
    }()

    // add 10 to 14
    wg.Add(1)
    go func() {
        for i := 0; i < items; i++ {
            time.Sleep(time.Duration(100+rand.Intn(100)) * time.Millisecond)
            list.Add(i + 10)
        }
        wg.Done()
    }()
    wg.Wait()

    return list
}

func validateConcurrentListHead(t *testing.T, list *ConcurrentList[int], expectedValue int) {
    value, err := list.GetHead()
    require.Nil(t, err, "ConcurrentList getHead is failed")
    require.Equal(t, expectedValue, value, "ConcurrentList head item is not matched")
}

func validateConcurrentListTail(t *testing.T, list *ConcurrentList[int], expectedValue int) {
    value, err := list.GetTail()
    require.Nil(t, err, "ConcurrentList getTail is failed")
    require.Equal(t, expectedValue, value, "ConcurrentList tail item is not matched")
}
