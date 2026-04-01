package main

import (
	"testing"
)

func TestNodeCreation(t *testing.T) {
	left := NewTerminalNode(2)
	right := NewTerminalNode(3)
	root := NewNode(1, left, right)

	if root.Value != 1 {
		t.Errorf("Expected root value 1, got %d", root.Value)
	}

	if root.left != left || root.right != right {
		t.Error("Expected left and right to be set")
	}

	if left.parent != root || right.parent != root {
		t.Error("Expected parent to be set")
	}
}

func TestTerminalNodeCreation(t *testing.T) {
	leaf := NewTerminalNode(5)

	if leaf.Value != 5 {
		t.Errorf("Expected value 5, got %d", leaf.Value)
	}

	if leaf.left != nil || leaf.right != nil {
		t.Error("Expected terminal node to have no children")
	}
}

func TestInOrderIteratorCreation(t *testing.T) {
	left := NewTerminalNode(2)
	right := NewTerminalNode(3)
	root := NewNode(1, left, right)

	iterator := NewInOrderIterator(root)

	if iterator.root != root {
		t.Error("Expected iterator root to be set")
	}

	if iterator.Current.Value != 2 {
		t.Errorf("Expected first current to be node 2 (leftmost), got %d", iterator.Current.Value)
	}
}

func TestInOrderTraversal(t *testing.T) {
	// Tree:
	//   1
	//  / \
	// 2   3
	// In-order: 2, 1, 3

	left := NewTerminalNode(2)
	right := NewTerminalNode(3)
	root := NewNode(1, left, right)

	iterator := NewInOrderIterator(root)

	values := []int{}
	for iterator.MoveNext() {
		values = append(values, iterator.Current.Value)
	}

	if len(values) != 3 {
		t.Errorf("Expected 3 values, got %d", len(values))
	}

	if values[0] != 2 || values[1] != 1 || values[2] != 3 {
		t.Errorf("Expected [2, 1, 3], got %v", values)
	}
}

func TestBinaryTreeCreation(t *testing.T) {
	left := NewTerminalNode(2)
	right := NewTerminalNode(3)
	root := NewNode(1, left, right)

	tree := NewBinaryTree(root)

	if tree.root != root {
		t.Error("Expected tree root to be set")
	}
}

func TestBinaryTreeInOrder(t *testing.T) {
	left := NewTerminalNode(2)
	right := NewTerminalNode(3)
	root := NewNode(1, left, right)

	tree := NewBinaryTree(root)
	iterator := tree.InOrder()

	values := []int{}
	for iterator.MoveNext() {
		values = append(values, iterator.Current.Value)
	}

	if len(values) != 3 {
		t.Errorf("Expected 3 values, got %d", len(values))
	}
}

func TestIteratorReset(t *testing.T) {
	left := NewTerminalNode(2)
	right := NewTerminalNode(3)
	root := NewNode(1, left, right)

	iterator := NewInOrderIterator(root)

	// Traverse once
	values1 := []int{}
	for iterator.MoveNext() {
		values1 = append(values1, iterator.Current.Value)
	}

	// Create a new iterator for second traversal
	iterator2 := NewInOrderIterator(root)

	// Traverse again with new iterator
	values2 := []int{}
	for iterator2.MoveNext() {
		values2 = append(values2, iterator2.Current.Value)
	}

	if len(values1) != len(values2) {
		t.Errorf("Expected same traversal with new iterator, got %d vs %d", len(values1), len(values2))
	}
}

func TestComplexTreeTraversal(t *testing.T) {
	// Tree:
	//       4
	//      / \
	//     2   6
	//    / \ / \
	//   1  3 5  7

	one := NewTerminalNode(1)
	three := NewTerminalNode(3)
	five := NewTerminalNode(5)
	seven := NewTerminalNode(7)

	two := NewNode(2, one, three)
	six := NewNode(6, five, seven)
	root := NewNode(4, two, six)

	iterator := NewInOrderIterator(root)

	values := []int{}
	for iterator.MoveNext() {
		values = append(values, iterator.Current.Value)
	}

	// In-order: 1, 2, 3, 4, 5, 6, 7
	expected := []int{1, 2, 3, 4, 5, 6, 7}
	for i, v := range expected {
		if i >= len(values) || values[i] != v {
			t.Errorf("In-order traversal mismatch at position %d, expected %d, got %v", i, v, values)
			break
		}
	}
}
