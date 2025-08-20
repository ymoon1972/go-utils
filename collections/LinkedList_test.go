package collections

import (
    "strconv"
    "testing"

    "github.com/stretchr/testify/require"
)

func TestLinkedList_Add(t *testing.T) {
    list := NewLinkedList[int]()
    require.Equal(t, 0, list.Size(), "LinkedList size is not equal")
    require.Equal(t, true, list.IsEmpty(), "LinkedList is not empty")

    list.Add(1)
    list.Add(2)
    list.Add(3)
    require.Equal(t, 3, list.Size(), "LinkedList size is not equal")
    require.Equal(t, false, list.IsEmpty(), "LinkedList is empty")
    require.ElementsMatch(t, []int{1, 2, 3}, list.Values(), "LinkedList values are not equal")
    validateHead(t, list, 1)
    validateTail(t, list, 3)
}

func TestLinkedList_Insert(t *testing.T) {
    list := NewLinkedList[int]()
    require.Equal(t, 0, list.Size(), "LinkedList size is not equal")
    require.Equal(t, true, list.IsEmpty(), "LinkedList is not empty")

    list.InsertAt(1, 1)
    require.Equal(t, 1, list.Size(), "LinkedList size is not equal")
    require.ElementsMatch(t, []int{1}, list.Values(), "LinkedList values are not equal")
    validateHead(t, list, 1)
    validateTail(t, list, 1)
    validateGetAt(t, list, 0, 1)

    list.InsertAt(0, 0)
    require.Equal(t, 2, list.Size(), "LinkedList size is not equal")
    require.ElementsMatch(t, []int{0, 1}, list.Values(), "LinkedList values are not equal")
    validateHead(t, list, 0)
    validateTail(t, list, 1)
    validateGetAt(t, list, 0, 0)
    validateGetAt(t, list, 1, 1)

    list.InsertAt(100, 2)
    require.Equal(t, 3, list.Size(), "LinkedList size is not equal")
    require.ElementsMatch(t, []int{0, 1, 2}, list.Values(), "LinkedList values are not equal")
    validateHead(t, list, 0)
    validateTail(t, list, 2)
    validateGetAt(t, list, 0, 0)
    validateGetAt(t, list, 1, 1)
    validateGetAt(t, list, 2, 2)

    _, err := list.GetAt(100)
    require.NotNil(t, err, "LinkedList getAt is failed")
}

func TestLinkedList_Remove(t *testing.T) {
    list := NewLinkedList[int]()
    require.Equal(t, 0, list.Size(), "LinkedList size is not equal")
    require.Equal(t, true, list.IsEmpty(), "LinkedList is not empty")

    _, err := list.RemoveAt(1)
    require.NotNil(t, err, "LinkedList removeAt is failed")

    list.AddAll([]int{1, 2, 3})
    validateRemoveHead(t, list, 1)
    validateRemoveTail(t, list, 3)
    validateRemoveAt(t, list, 0, 2)
    require.Equal(t, 0, list.Size(), "LinkedList size is not equal")
    require.Equal(t, true, list.IsEmpty(), "LinkedList is not empty")

    list.AddAll([]int{4, 5, 6, 7, 8})
    validateRemoveAt(t, list, 1, 5)
    validateHead(t, list, 4)
    validateTail(t, list, 8)
    require.ElementsMatch(t, []int{4, 6, 7, 8}, list.Values(), "LinkedList values are not equal")

    validateRemoveAt(t, list, 1, 6)
    validateHead(t, list, 4)
    validateTail(t, list, 8)
    require.ElementsMatch(t, []int{4, 7, 8}, list.Values(), "LinkedList values are not equal")

    validateRemoveAt(t, list, 2, 8)
    validateHead(t, list, 4)
    validateTail(t, list, 7)
    require.ElementsMatch(t, []int{4, 7}, list.Values(), "LinkedList values are not equal")
}

func TestLinkedList_Clone(t *testing.T) {
    list := NewLinkedList[int]()
    list.AddAll([]int{1, 2, 3})

    cloned := list.Clone()
    require.ElementsMatch(t, []int{1, 2, 3}, cloned.Values(), "LinkedList clone is failed")
}

func TestLinkedList_Merge(t *testing.T) {
    list := NewLinkedList[int]()
    list.AddAll([]int{1, 2, 3})

    cloned := list.Clone()
    cloned.Merge(list)
    require.ElementsMatch(t, []int{1, 2, 3, 1, 2, 3}, cloned.Values(), "LinkedList merge is failed")
}

func TestLinkedList_Reverse(t *testing.T) {
    list := NewLinkedList[int]()
    list.AddAll([]int{1, 2, 3})

    list.Reverse()
    require.ElementsMatch(t, []int{3, 2, 1}, list.Values(), "LinkedList reverse is failed")
    validateHead(t, list, 3)
    validateTail(t, list, 1)
}

func TestLinkedList_Filter(t *testing.T) {
    list := NewLinkedList[int]()
    list.AddAll([]int{1, 2, 3, 4, 5, 6})

    even := list.Filter(func(value int) bool { return value%2 == 0 })
    require.ElementsMatch(t, []int{2, 4, 6}, even.Values(), "LinkedList filter is failed")
}

func TestLinkedList_Sort(t *testing.T) {
    list := NewLinkedList[int]()
    list.AddAll([]int{6, 4, 3, 5, 1, 2})

    list.Sort(func(a int, b int) bool { return a < b })
    require.ElementsMatch(t, []int{1, 2, 3, 4, 5, 6}, list.Values(), "LinkedList sort is failed")

    list.Sort(func(a int, b int) bool { return a > b })
    require.ElementsMatch(t, []int{6, 5, 4, 3, 2, 1}, list.Values(), "LinkedList reverse sort is failed")
    validateHead(t, list, 6)
    validateTail(t, list, 1)
}

func TestLinkedList_MapLinkedList(t *testing.T) {
    list := NewLinkedList[int]()
    list.AddAll([]int{1, 2, 3, 4, 5, 6})

    mapped := MapLinkedList(list, func(item int) string { return strconv.Itoa(item) })
    require.ElementsMatch(t, []string{"1", "2", "3", "4", "5", "6"}, mapped.Values(), "LinkedList map is failed")
}

func TestLinkedList_ReduceLinkedList(t *testing.T) {
    list := NewLinkedList[int]()
    list.AddAll([]int{1, 2, 3, 4, 5, 6})

    sum := ReduceLinkedList(list, 0, func(acc int, item int) int { return acc + item })
    require.Equal(t, 21, sum, "LinkedList reduce is failed")
}

func validateHead(t *testing.T, list *LinkedList[int], expectedValue int) {
    value, err := list.GetHead()
    require.Nil(t, err, "LinkedList getHead is failed")
    require.Equal(t, expectedValue, value, "LinkedList head item is not matched")
}

func validateTail(t *testing.T, list *LinkedList[int], expectedValue int) {
    value, err := list.GetTail()
    require.Nil(t, err, "LinkedList getTail is failed")
    require.Equal(t, expectedValue, value, "LinkedList tail item is not matched")
}

func validateGetAt(t *testing.T, list *LinkedList[int], index int, expectedValue int) {
    value, err := list.GetAt(index)
    require.Nil(t, err, "LinkedList getTail is failed")
    require.Equal(t, expectedValue, value, "LinkedList tail item is not matched")
}

func validateRemoveHead(t *testing.T, list *LinkedList[int], expectedValue int) {
    value, err := list.RemoveHead()
    require.Nil(t, err, "LinkedList removeHead is failed")
    require.Equal(t, expectedValue, value, "LinkedList removeHead value is not matched")
}

func validateRemoveTail(t *testing.T, list *LinkedList[int], expectedValue int) {
    value, err := list.RemoveTail()
    require.Nil(t, err, "LinkedList removeTail is failed")
    require.Equal(t, expectedValue, value, "LinkedList removeTail item is not matched")
}

func validateRemoveAt(t *testing.T, list *LinkedList[int], index int, expectedValue int) {
    value, err := list.RemoveAt(index)
    require.Nil(t, err, "LinkedList removeAt is failed")
    require.Equal(t, expectedValue, value, "LinkedList removeAt item is not matched")
}
