package collections

import (
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestBinaryTree(t *testing.T) {
    tree := NewBinaryTree[int](func(a, b int) int { return a - b })
    tree.Offer(7)
    tree.OfferAll([]int{2, 4, 8})

    assert.Equal(t, 4, tree.Size(), "BinaryTree size is not equal")
    assert.ElementsMatch(t, []int{2, 4, 7, 8}, tree.Values(), "BinaryTree values are not equal")
    assert.True(t, tree.Contains(4), "BinaryTree contains(4) is not matched")

    tree.Offer(1)
    tree.Offer(3)
    tree.Offer(9)
    assert.ElementsMatch(t, []int{1, 2, 3, 4, 7, 8, 9}, tree.Values(), "BinaryTree values are not equal")
    assert.True(t, tree.Contains(1), "BinaryTree contains(1) is not matched")

    // remove node only with the right child
    removed := tree.Remove(8)
    assert.True(t, removed, "BinaryTree remove(8) is not matched")
    assert.ElementsMatch(t, []int{1, 2, 3, 4, 7, 9}, tree.Values(), "BinaryTree values are not equal")
    assert.False(t, tree.Contains(8), "BinaryTree contains(8) is not matched")

    // remove node with left child
    removed = tree.Remove(2)
    assert.True(t, removed, "BinaryTree remove(2) is not matched")
    assert.ElementsMatch(t, []int{1, 3, 4, 7, 9}, tree.Values(), "BinaryTree values are not equal")
    assert.False(t, tree.Contains(2), "BinaryTree contains(2) is not matched")

    // remove leaf node
    removed = tree.Remove(4)
    assert.True(t, removed, "BinaryTree remove(4) is not matched")
    assert.ElementsMatch(t, []int{1, 3, 7, 9}, tree.Values(), "BinaryTree values are not equal")
    assert.False(t, tree.Contains(4), "BinaryTree contains(4) is not matched")

    // remove root node
    removed = tree.Remove(7)
    assert.True(t, removed, "BinaryTree remove(7) is not matched")
    assert.ElementsMatch(t, []int{1, 3, 9}, tree.Values(), "BinaryTree values are not equal")
    assert.False(t, tree.Contains(7), "BinaryTree contains(7) is not matched")

    tree.Clear()
    assert.Equal(t, 0, tree.Size(), "BinaryTree size is not equal")
    assert.True(t, tree.IsEmpty(), "BinaryTree is not empty")
    assert.ElementsMatch(t, []int{}, tree.Values(), "BinaryTree values are not equal")
}

func TestBinaryTree_MinQueue(t *testing.T) {
    tree := NewBinaryTree[int](func(a, b int) int { return a - b })
    tree.Offer(7)
    tree.OfferAll([]int{2, 4, 8})

    validateTreePeek(t, tree, 2)
    validateTreePoll(t, tree, 2)
    assert.Equal(t, 3, tree.Size(), "BinaryTree size is not equal")
    assert.ElementsMatch(t, []int{4, 7, 8}, tree.Values(), "BinaryTree values are not equal")
    assert.False(t, tree.Contains(2), "BinaryTree contains(2) is not matched")

    validateTreePeek(t, tree, 4)
    validateTreePoll(t, tree, 4)
    assert.Equal(t, 2, tree.Size(), "BinaryTree size is not equal")
    assert.ElementsMatch(t, []int{7, 8}, tree.Values(), "BinaryTree values are not equal")
    assert.False(t, tree.Contains(4), "BinaryTree contains(4) is not matched")

    tree.Offer(1)
    assert.ElementsMatch(t, []int{1, 7, 8}, tree.Values(), "BinaryTree values are not equal")
    assert.True(t, tree.Contains(1), "BinaryTree contains(1) is not matched")

    validateTreePeek(t, tree, 1)
    validateTreePoll(t, tree, 1)
    validateTreePeek(t, tree, 7)
    validateTreePoll(t, tree, 7)
    validateTreePeek(t, tree, 8)
    validateTreePoll(t, tree, 8)
    assert.Equal(t, 0, tree.Size(), "BinaryTree size is not equal")
    assert.True(t, tree.IsEmpty(), "BinaryTree is not empty")
}

func TestBinaryTree_MaxQueue(t *testing.T) {
    tree := NewBinaryTree[int](func(a, b int) int { return b - a })
    tree.Offer(7)
    tree.OfferAll([]int{2, 4, 8})
    assert.Equal(t, 4, tree.Size(), "BinaryTree size is not equal")
    assert.ElementsMatch(t, []int{8, 7, 4, 2}, tree.Values(), "BinaryTree values are not equal")

    validateTreePeek(t, tree, 8)
    validateTreePoll(t, tree, 8)
    assert.Equal(t, 3, tree.Size(), "BinaryTree size is not equal")
    assert.ElementsMatch(t, []int{7, 4, 2}, tree.Values(), "BinaryTree values are not equal")
    assert.False(t, tree.Contains(8), "BinaryTree contains(8) is not matched")

    validateTreePeek(t, tree, 7)
    validateTreePoll(t, tree, 7)
    assert.Equal(t, 2, tree.Size(), "BinaryTree size is not equal")
    assert.ElementsMatch(t, []int{4, 2}, tree.Values(), "BinaryTree values are not equal")
    assert.False(t, tree.Contains(7), "BinaryTree contains(7) is not matched")

    tree.Offer(1)
    assert.ElementsMatch(t, []int{4, 2, 1}, tree.Values(), "BinaryTree values are not equal")
    assert.True(t, tree.Contains(1), "BinaryTree contains(1) is not matched")

    validateTreePeek(t, tree, 4)
    validateTreePoll(t, tree, 4)
    validateTreePeek(t, tree, 2)
    validateTreePoll(t, tree, 2)
    validateTreePeek(t, tree, 1)
    validateTreePoll(t, tree, 1)
    assert.Equal(t, 0, tree.Size(), "BinaryTree size is not equal")
    assert.True(t, tree.IsEmpty(), "BinaryTree is not empty")
}

func TestBinaryTree_Clone(t *testing.T) {
    tree := NewBinaryTree[int](func(a, b int) int { return a - b })
    tree.Offer(7)
    tree.OfferAll([]int{2, 4, 8})

    clone := tree.Clone()
    assert.Equal(t, 4, clone.Size(), "BinaryTree size is not equal")
    assert.ElementsMatch(t, []int{2, 4, 7, 8}, clone.Values(), "BinaryTree values are not equal")
    assert.True(t, clone.Contains(2), "BinaryTree contains(2) is not matched")

    validateTreePeek(t, clone, 2)
    validateTreePoll(t, clone, 2)
    assert.Equal(t, 3, clone.Size(), "BinaryTree size is not equal")
    assert.ElementsMatch(t, []int{4, 7, 8}, clone.Values(), "BinaryTree values are not equal")
    assert.False(t, clone.Contains(2), "BinaryTree contains(2) is not matched")

    clone.Offer(1)
    assert.ElementsMatch(t, []int{1, 4, 7, 8}, clone.Values(), "BinaryTree values are not equal")
    assert.True(t, clone.Contains(1), "BinaryTree contains(1) is not matched")
}

func validateTreePoll(t *testing.T, tree *BinaryTree[int], expectedValue int) {
    value, err := tree.Poll()
    require.Nil(t, err, "Tree poll is failed")
    require.Equal(t, expectedValue, value, "Tree poll item is not matched")
}

func validateTreePeek(t *testing.T, tree *BinaryTree[int], expectedValue int) {
    value, err := tree.Peek()
    require.Nil(t, err, "Tree peek is failed")
    require.Equal(t, expectedValue, value, "Tree peek item is not matched")
}
