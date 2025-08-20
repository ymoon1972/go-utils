package collections

import "errors"

type treeNode[T comparable] struct {
    value T
    left  *treeNode[T]
    right *treeNode[T]
}

type BinaryTree[T comparable] struct {
    head       *treeNode[T]
    size       int
    comparator func(a, b T) int
}

func NewBinaryTree[T comparable](comparator func(a, b T) int) *BinaryTree[T] {
    return &BinaryTree[T]{comparator: comparator}
}

func (s *BinaryTree[T]) Size() int {
    return s.size
}

func (s *BinaryTree[T]) IsEmpty() bool {
    return s.size == 0
}

func (s *BinaryTree[T]) Values() []T {
    if s.IsEmpty() {
        return []T{}
    }

    values := NewQueue[T]()
    collectValues(s.head, values)
    return values.Values()
}

func (s *BinaryTree[T]) Offer(value T) {
    s.head = addNode(s.head, value, s.comparator)
    s.size++
}

func (s *BinaryTree[T]) OfferAll(values []T) {
    for _, value := range values {
        s.Offer(value)
    }
}

func (s *BinaryTree[T]) Contains(value T) bool {
    if s.IsEmpty() {
        return false
    }

    return findNode(s.head, value, s.comparator)
}

func (s *BinaryTree[T]) Remove(value T) bool {
    if s.IsEmpty() {
        return false
    }

    head, removed := removeNode(s.head, value, s.comparator)
    s.head = head
    if removed {
        s.size--
    }
    return removed
}

func (s *BinaryTree[T]) Clear() {
    s.head = nil
    s.size = 0
}

func (s *BinaryTree[T]) Peek() (T, error) {
    if s.IsEmpty() {
        var zero T
        return zero, errors.New("tree is empty")
    }

    return findLeftMostNode(s.head).value, nil
}

func (s *BinaryTree[T]) Poll() (T, error) {
    if s.IsEmpty() {
        var zero T
        return zero, errors.New("tree is empty")
    }

    top, value := pollNode(s.head)
    s.head = top
    s.size--
    return value, nil
}

func pollNode[T comparable](node *treeNode[T]) (*treeNode[T], T) {
    if node.left == nil {
        return node.right, node.value
    }

    left, value := pollNode(node.left)
    node.left = left
    return node, value
}

func (s *BinaryTree[T]) Clone() *BinaryTree[T] {
    if s.IsEmpty() {
        return NewBinaryTree[T](s.comparator)
    }

    src := NewQueue[*treeNode[T]]()
    src.Offer(s.head)

    dest := NewQueue[*treeNode[T]]()
    destHead := &treeNode[T]{value: s.head.value}
    dest.Offer(destHead)

    for !src.IsEmpty() {
        length := src.Size()
        for i := 0; i < length; i++ {
            node, _ := src.Poll()
            target, _ := dest.Poll()

            if node.left != nil {
                target.left = &treeNode[T]{value: node.left.value}
                src.Offer(node.left)
                dest.Offer(target.left)
            }

            if node.right != nil {
                target.right = &treeNode[T]{value: node.right.value}
                src.Offer(node.right)
                dest.Offer(target.right)
            }
        }
    }

    return &BinaryTree[T]{
        head:       destHead,
        size:       s.size,
        comparator: s.comparator,
    }
}

func collectValues[T comparable](node *treeNode[T], values *Queue[T]) {
    if node == nil {
        return
    }

    collectValues(node.left, values)
    values.Offer(node.value)
    collectValues(node.right, values)
}

func addNode[T comparable](node *treeNode[T], value T, comparator func(a, b T) int) *treeNode[T] {
    if node == nil {
        return &treeNode[T]{value: value}
    }

    if comparator(value, node.value) <= 0 {
        node.left = addNode(node.left, value, comparator)
    } else {
        node.right = addNode(node.right, value, comparator)
    }

    return node
}

func findNode[T comparable](node *treeNode[T], value T, comparator func(a, b T) int) bool {
    if node == nil {
        return false
    }

    result := comparator(value, node.value)
    if result == 0 {
        return true
    } else if result < 0 {
        return findNode(node.left, value, comparator)
    } else {
        return findNode(node.right, value, comparator)
    }
}

func removeNode[T comparable](node *treeNode[T], value T, comparator func(a, b T) int) (*treeNode[T], bool) {
    if node == nil {
        return nil, false
    }

    result := comparator(value, node.value)
    if result == 0 {
        if node.right == nil {
            return node.left, true
        }

        right, leftMost := detachLeftMostNode(node.right)
        leftMost.left = node.left
        leftMost.right = right
        return leftMost, true
    }

    if result < 0 {
        left, removed := removeNode(node.left, value, comparator)
        node.left = left
        return node, removed
    }

    right, removed := removeNode(node.right, value, comparator)
    node.right = right
    return node, removed
}

func findLeftMostNode[T comparable](node *treeNode[T]) *treeNode[T] {
    for node.left != nil {
        node = node.left
    }
    return node
}

func detachLeftMostNode[T comparable](node *treeNode[T]) (*treeNode[T], *treeNode[T]) {
    if node.left == nil {
        return node.right, node
    }

    left, end := detachLeftMostNode(node.left)
    node.left = left
    return node, end
}
