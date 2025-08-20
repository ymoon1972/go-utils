package collections

import (
    "testing"

    "github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
    stack := NewStack[int]()
    require.True(t, stack.IsEmpty(), "Stack is not empty")

    stack.PushValues([]int{1, 2, 3, 4})
    require.False(t, stack.IsEmpty(), "Stack is empty")
    require.Equal(t, 4, stack.Size(), "Stack size is not equal")
    validateStackPeek(t, stack, 4)
    validateStackPop(t, stack, 4)
    validateStackPeek(t, stack, 3)
    validateStackPop(t, stack, 3)

    // add in the middle
    stack.Push(10)
    validateStackPeek(t, stack, 10)
    validateStackPop(t, stack, 10)

    // continue pop
    validateStackPeek(t, stack, 2)
    validateStackPop(t, stack, 2)
    validateStackPeek(t, stack, 1)
    validateStackPop(t, stack, 1)
    require.True(t, stack.IsEmpty(), "Stack is not empty")

    stack.Push(10)
    require.Equal(t, 1, stack.Size(), "Stack size is not equal")
    validateStackPeek(t, stack, 10)
    validateStackPop(t, stack, 10)
}

func validateStackPop(t *testing.T, stack *Stack[int], expectedValue int) {
    value, err := stack.Pop()
    require.Nil(t, err, "Stack pop is failed")
    require.Equal(t, expectedValue, value, "Stack pop item is not matched")
}

func validateStackPeek(t *testing.T, stack *Stack[int], expectedValue int) {
    value, err := stack.Peek()
    require.Nil(t, err, "Stack peek is failed")
    require.Equal(t, expectedValue, value, "Stack peek item is not matched")
}
