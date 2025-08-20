package collections

import (
    "testing"

    "github.com/stretchr/testify/require"
)

func TestArrayList_NewArrayList(t *testing.T) {
    arr := NewArrayList[int]()
    require.NotNil(t, arr, "NewArrayList() is failed")
}

func TestArrayList_Add(t *testing.T) {
    arr := NewArrayList[int]()
    require.Equal(t, 0, arr.Size(), "ArrayList size is not equal")
    require.Equal(t, true, arr.IsEmpty(), "ArrayList is not empty")

    arr.Add(1)
    arr.Add(2)
    arr.Add(3)
    require.Equal(t, 3, arr.Size(), "ArrayList size is not equal")
    require.ElementsMatch(t, []int{1, 2, 3}, arr.Values(), "ArrayList values are not equal")
    value, err := arr.Get(1)
    require.Nil(t, err, "ArrayList get is failed")
    require.Equal(t, 2, value, "ArrayList get(1) is not matched")
    require.Equal(t, false, arr.IsEmpty(), "ArrayList is empty")

    // get out of range
    value, err = arr.Get(10)
    require.NotNil(t, err, "ArrayList get is failed")
    require.Equal(t, 0, value, "ArrayList get(10) is not matched")
}

func TestArrayList_AddAll(t *testing.T) {
    arr := NewArrayList[int]()
    arr.Add(1)
    arr.Add(2)
    arr.Add(3)
    arr.AddAll([]int{4, 5, 6})
    require.Equal(t, 6, arr.Size(), "ArrayList size is not equal")
    require.ElementsMatch(t, []int{1, 2, 3, 4, 5, 6}, arr.Values(), "ArrayList values are not equal")
    value, err := arr.Get(4)
    require.Nil(t, err, "ArrayList get is failed")
    require.Equal(t, 5, value, "ArrayList get(4) is not matched")
}

func TestArrayList_Insert(t *testing.T) {
    arr := NewArrayList[int]()
    arr.Add(1)
    arr.Add(2)
    arr.Add(3)

    err := arr.InsertAt(2, 10)
    require.Nil(t, err, "ArrayList insert is failed")
    require.Equal(t, 4, arr.Size(), "ArrayList size is not equal")
    require.ElementsMatch(t, []int{1, 2, 10, 3}, arr.Values(), "ArrayList values are not equal")

    value, err := arr.Get(2)
    require.Nil(t, err, "ArrayList get is failed")
    require.Equal(t, 10, value, "ArrayList get(2) is not matched")

    // insert at out of range
    err = arr.InsertAt(10, 10)
    require.NotNil(t, err, "ArrayList insert is failed")
    require.Equal(t, 4, arr.Size(), "ArrayList size is not equal")
}

func TestArrayList_Remove(t *testing.T) {
    arr := NewArrayList[int]()
    arr.Add(1)
    arr.Add(2)
    arr.Add(3)
    arr.Add(4)

    err := arr.RemoveAt(2)
    require.Nil(t, err, "ArrayList removeAt is failed")
    require.Equal(t, 3, arr.Size(), "ArrayList size is not equal")
    require.ElementsMatch(t, []int{1, 2, 4}, arr.Values(), "ArrayList values are not equal")

    // remove at out of range
    err = arr.RemoveAt(10)
    require.NotNil(t, err, "ArrayList removeAt is failed")
    require.Equal(t, 3, arr.Size(), "ArrayList size is not equal")
}

func TestArrayList_Clear(t *testing.T) {
    arr := NewArrayList[int]()
    arr.Add(1)
    arr.Add(2)
    arr.Add(3)
    arr.Add(4)
    require.Equal(t, 4, arr.Size(), "ArrayList size is not equal")
    require.ElementsMatch(t, []int{1, 2, 3, 4}, arr.Values(), "ArrayList values are not equal")

    arr.Clear()
    require.Equal(t, 0, arr.Size(), "ArrayList size is not equal")
    require.Equal(t, true, arr.IsEmpty(), "ArrayList is not empty")
    require.ElementsMatch(t, []int{}, arr.Values(), "ArrayList values are not equal")
}

func TestArrayList_Clone(t *testing.T) {
    arr := NewArrayList[int]()
    arr.AddAll([]int{1, 2, 3})

    second := arr.Clone()
    require.ElementsMatch(t, []int{1, 2, 3}, second.Values(), "ArrayList clone is failed")
}

func TestArrayList_Merge(t *testing.T) {
    arr := NewArrayList[int]()
    arr.AddAll([]int{1, 2, 3})

    second := arr.Clone()
    require.ElementsMatch(t, []int{1, 2, 3}, second.Values(), "ArrayList clone is failed")

    second.Merge(arr)
    require.ElementsMatch(t, []int{1, 2, 3, 1, 2, 3}, second.Values(), "ArrayList merge is failed")
}

func TestArrayList_Reverse(t *testing.T) {
    arr := NewArrayList[int]()
    arr.AddAll([]int{1, 2, 3})

    arr.Reverse()
    require.ElementsMatch(t, []int{3, 2, 1}, arr.Values(), "ArrayList reverse is failed")
}

func TestArrayList_Filter(t *testing.T) {
    arr := NewArrayList[int]()
    arr.AddAll([]int{1, 2, 3, 4, 5, 6})

    even := arr.Filter(func(value int) bool { return value%2 == 0 })
    require.ElementsMatch(t, []int{2, 4, 6}, even.Values(), "ArrayList filter is failed")
}

func TestArrayList_Sort(t *testing.T) {
    arr := NewArrayList[int]()
    arr.AddAll([]int{1, 3, 4, 6, 5, 2})

    arr.Sort(func(a int, b int) bool { return a < b })
    require.ElementsMatch(t, []int{1, 2, 3, 4, 5, 6}, arr.Values(), "ArrayList sort is failed")

    arr.Sort(func(a int, b int) bool { return a > b })
    require.ElementsMatch(t, []int{6, 5, 4, 3, 2, 1}, arr.Values(), "ArrayList reverse sort is failed")
}

func TestArrayList_Contains(t *testing.T) {
    arr := NewArrayList[int]()
    arr.AddAll([]int{1, 3, 4, 6, 5, 2})

    require.True(t, arr.Contains(1), "ArrayList contains is failed")
    require.True(t, arr.Contains(3), "ArrayList contains is failed")
    require.False(t, arr.Contains(7), "ArrayList contains is failed")
}

func TestArrayList_MapArrayList(t *testing.T) {
    arr := NewArrayList[int]()
    arr.AddAll([]int{1, 2, 3, 4, 5, 6})

    even := MapArrayList(arr, func(value int) bool { return value%2 == 0 })
    require.ElementsMatch(t, []bool{false, true, false, true, false, true}, even.Values(), "ArrayList map is failed")
}

func TestArrayList_ReduceArrayList(t *testing.T) {
    arr := NewArrayList[int]()
    arr.AddAll([]int{1, 2, 3, 4, 5, 6})

    sum := ReduceArrayList(arr, 0, func(acc int, value int) int { return acc + value })
    require.Equal(t, 21, sum, "ArrayList reduce is failed")
}
