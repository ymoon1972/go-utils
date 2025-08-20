package collections

import (
    "strconv"
    "testing"

    "github.com/stretchr/testify/require"
)

func TestNewDoubleLinkedList_Add(t *testing.T) {
    list := NewDoubleLinkedList[int]()
    require.Equal(t, 0, list.Size(), "DoubleLinkedList size is not equal")
    require.Equal(t, true, list.IsEmpty(), "DoubleLinkedList is not empty")

    list.Add(1)
    list.Add(2)
    list.Add(3)
    require.Equal(t, 3, list.Size(), "DoubleLinkedList size is not equal")
    require.Equal(t, false, list.IsEmpty(), "DoubleLinkedList is empty")
    require.ElementsMatch(t, []int{1, 2, 3}, list.Values(), "DoubleLinkedList values are not equal")
    validateDLinkedListHead(t, list, 1)
    validateDLinkedListTail(t, list, 3)
}

func TestNewDoubleLinkedList_Insert(t *testing.T) {
    list := NewDoubleLinkedList[int]()
    require.Equal(t, 0, list.Size(), "DoubleLinkedList size is not equal")
    require.Equal(t, true, list.IsEmpty(), "DoubleLinkedList is not empty")

    list.InsertAt(1, 1)
    require.Equal(t, 1, list.Size(), "DoubleLinkedList size is not equal")
    require.ElementsMatch(t, []int{1}, list.Values(), "DoubleLinkedList values are not equal")
    validateDLinkedListHead(t, list, 1)
    validateDLinkedListTail(t, list, 1)
    validateDLinkedListGetAt(t, list, 0, 1)

    list.InsertAt(0, 0)
    require.Equal(t, 2, list.Size(), "DoubleLinkedList size is not equal")
    require.ElementsMatch(t, []int{0, 1}, list.Values(), "DoubleLinkedList values are not equal")
    validateDLinkedListHead(t, list, 0)
    validateDLinkedListTail(t, list, 1)
    validateDLinkedListGetAt(t, list, 0, 0)
    validateDLinkedListGetAt(t, list, 1, 1)

    list.InsertAt(100, 2)
    require.Equal(t, 3, list.Size(), "DoubleLinkedList size is not equal")
    require.ElementsMatch(t, []int{0, 1, 2}, list.Values(), "DoubleLinkedList values are not equal")
    validateDLinkedListHead(t, list, 0)
    validateDLinkedListTail(t, list, 2)
    validateDLinkedListGetAt(t, list, 0, 0)
    validateDLinkedListGetAt(t, list, 1, 1)
    validateDLinkedListGetAt(t, list, 2, 2)

    _, err := list.GetAt(100)
    require.NotNil(t, err, "DoubleLinkedList getAt is failed")
}

func TestNewDoubleLinkedList_Remove(t *testing.T) {
    list := NewDoubleLinkedList[int]()
    require.Equal(t, 0, list.Size(), "DoubleLinkedList size is not equal")
    require.Equal(t, true, list.IsEmpty(), "DoubleLinkedList is not empty")

    _, err := list.RemoveAt(1)
    require.NotNil(t, err, "DoubleLinkedList removeAt is failed")

    list.AddAll([]int{1, 2, 3})
    validateDLinkedListRemoveHead(t, list, 1)
    validateDLinkedListRemoveTail(t, list, 3)
    validateDLinkedListRemoveAt(t, list, 0, 2)
    require.Equal(t, 0, list.Size(), "DoubleLinkedList size is not equal")
    require.Equal(t, true, list.IsEmpty(), "DoubleLinkedList is not empty")

    list.AddAll([]int{4, 5, 6, 7, 8})
    validateDLinkedListRemoveAt(t, list, 1, 5)
    validateDLinkedListHead(t, list, 4)
    validateDLinkedListTail(t, list, 8)
    require.ElementsMatch(t, []int{4, 6, 7, 8}, list.Values(), "DoubleLinkedList values are not equal")

    validateDLinkedListRemoveAt(t, list, 1, 6)
    validateDLinkedListHead(t, list, 4)
    validateDLinkedListTail(t, list, 8)
    require.ElementsMatch(t, []int{4, 7, 8}, list.Values(), "DoubleLinkedList values are not equal")

    validateDLinkedListRemoveAt(t, list, 2, 8)
    validateDLinkedListHead(t, list, 4)
    validateDLinkedListTail(t, list, 7)
    require.ElementsMatch(t, []int{4, 7}, list.Values(), "DoubleLinkedList values are not equal")
}

func TestNewDoubleLinkedList_Clone(t *testing.T) {
    list := NewDoubleLinkedList[int]()
    list.AddAll([]int{1, 2, 3})

    cloned := list.Clone()
    require.ElementsMatch(t, []int{1, 2, 3}, cloned.Values(), "DoubleLinkedList clone is failed")
}

func TestNewDoubleLinkedList_Merge(t *testing.T) {
    list := NewDoubleLinkedList[int]()
    list.AddAll([]int{1, 2, 3})

    cloned := list.Clone()
    cloned.Merge(list)
    require.ElementsMatch(t, []int{1, 2, 3, 1, 2, 3}, cloned.Values(), "DoubleLinkedList merge is failed")
}

func TestNewDoubleLinkedList_Reverse(t *testing.T) {
    list := NewDoubleLinkedList[int]()
    list.AddAll([]int{1, 2, 3})
    list.Reverse()
    require.ElementsMatch(t, []int{3, 2, 1}, list.Values(), "DoubleLinkedList reverse is failed")
    validateDLinkedListHead(t, list, 3)
    validateDLinkedListTail(t, list, 1)
}

func TestNewDoubleLinkedList_Filter(t *testing.T) {
    list := NewDoubleLinkedList[int]()
    list.AddAll([]int{1, 2, 3, 4, 5, 6})

    even := list.Filter(func(value int) bool { return value%2 == 0 })
    require.ElementsMatch(t, []int{2, 4, 6}, even.Values(), "DoubleLinkedList filter is failed")
}

func TestNewDoubleLinkedList_Sort(t *testing.T) {
    list := NewDoubleLinkedList[int]()
    list.AddAll([]int{6, 4, 3, 5, 1, 2})

    list.Sort(func(a int, b int) int { return a - b })
    require.ElementsMatch(t, []int{1, 2, 3, 4, 5, 6}, list.Values(), "DoubleLinkedList sort is failed")

    list.Sort(func(a int, b int) int { return b - a })
    require.ElementsMatch(t, []int{6, 5, 4, 3, 2, 1}, list.Values(), "DoubleLinkedList reverse sort is failed")
    validateDLinkedListHead(t, list, 6)
    validateDLinkedListTail(t, list, 1)
}

func TestNewDoubleLinkedList_MapLinkedList(t *testing.T) {
    list := NewDoubleLinkedList[int]()
    list.AddAll([]int{1, 2, 3, 4, 5, 6})

    mapped := MapDoubleLinkedList(list, func(item int) string { return strconv.Itoa(item) })
    require.ElementsMatch(t, []string{"1", "2", "3", "4", "5", "6"}, mapped.Values(), "DoubleLinkedList map is failed")
}

func TestNewDoubleLinkedList_ReduceLinkedList(t *testing.T) {
    list := NewDoubleLinkedList[int]()
    list.AddAll([]int{1, 2, 3, 4, 5, 6})

    mapped := MapDoubleLinkedList(list, func(item int) string { return strconv.Itoa(item) })
    aggregated := ReduceDoubleLinkedList(mapped, "", func(acc string, item string) string {
        if len(acc) == 0 {
            return item
        }
        return acc + "-" + item
    })
    require.Equal(t, "1-2-3-4-5-6", aggregated, "DoubleLinkedList reduce is failed")
}

func validateDLinkedListHead(t *testing.T, list *DoubleLinkedList[int], expectedValue int) {
    value, err := list.GetHead()
    require.Nil(t, err, "DoubleLinkedList getHead is failed")
    require.Equal(t, expectedValue, value, "DoubleLinkedList head item is not matched")
}

func validateDLinkedListTail(t *testing.T, list *DoubleLinkedList[int], expectedValue int) {
    value, err := list.GetTail()
    require.Nil(t, err, "DoubleLinkedList getTail is failed")
    require.Equal(t, expectedValue, value, "DoubleLinkedList tail item is not matched")
}

func validateDLinkedListGetAt(t *testing.T, list *DoubleLinkedList[int], index int, expectedValue int) {
    value, err := list.GetAt(index)
    require.Nil(t, err, "DoubleLinkedList getTail is failed")
    require.Equal(t, expectedValue, value, "DoubleLinkedList tail item is not matched")
}

func validateDLinkedListRemoveHead(t *testing.T, list *DoubleLinkedList[int], expectedValue int) {
    value, err := list.RemoveHead()
    require.Nil(t, err, "DoubleLinkedList removeHead is failed")
    require.Equal(t, expectedValue, value, "DoubleLinkedList removeHead value is not matched")
}

func validateDLinkedListRemoveTail(t *testing.T, list *DoubleLinkedList[int], expectedValue int) {
    value, err := list.RemoveTail()
    require.Nil(t, err, "DoubleLinkedList removeTail is failed")
    require.Equal(t, expectedValue, value, "DoubleLinkedList removeTail item is not matched")
}

func validateDLinkedListRemoveAt(t *testing.T, list *DoubleLinkedList[int], index, expectedValue int) {
    value, err := list.RemoveAt(index)
    require.Nil(t, err, "DoubleLinkedList removeAt is failed")
    require.Equal(t, expectedValue, value, "DoubleLinkedList removeAt item is not matched")
}
