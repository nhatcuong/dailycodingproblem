// This problem was asked by Dropbox.
// Given the root to a binary search tree, find the second largest node in the tree.

package main

import "fmt"

func main() {
	root := SearchBinaryTree{10, nil, nil}
	root.addValue(5)
	root.addValue(3)
	root.addValue(7)
	var secondMax int
	secondMax, _ = root.findSecondMax()
	fmt.Printf("expecting %v true value %v\n", 7, secondMax)
	root.addValue(15)
	root.addValue(12)
	root.addValue(17)
	secondMax, _ = root.findSecondMax()
	fmt.Printf("expecting %v true value %v\n", 15, secondMax)
	root.addValue(19)
	secondMax, _ = root.findSecondMax()
	fmt.Printf("expecting %v true value %v\n", 17, secondMax)
	root.addValue(18)
	secondMax, _ = root.findSecondMax()
	fmt.Printf("expecting %v true value %v\n", 18, secondMax)
	// case 1: no right tree, but left tree with children
	// case 2: w right tree with right child that has right child and left child
	// case 3: like above, but the right child has right child
	// case 4: like above, but the far right child has left child
}

type SearchBinaryTree struct {
	value int
	left  *SearchBinaryTree
	right *SearchBinaryTree
}

func (t *SearchBinaryTree) addValue(new int) {
	switch {
	case t.value > new && t.left == nil:
		left := SearchBinaryTree{new, nil, nil}
		t.left = &left
	case t.value > new && t.left != nil:
		t.left.addValue(new)
	case t.value < new && t.right == nil:
		right := SearchBinaryTree{new, nil, nil}
		t.right = &right
	case t.value < new && t.right != nil:
		t.right.addValue(new)
	}
}

func (t *SearchBinaryTree) findMax() int {
	if t.right == nil {
		return t.value
	}
	return t.right.findMax()
}

func (t *SearchBinaryTree) findSecondMax() (int, bool) {
	if t.right != nil {
		if t.right.left == nil && t.right.right == nil {
			return t.value, true
		}
		return t.right.findSecondMax()
	}
	if t.left != nil {
		return t.left.findMax(), true
	}
	return 0, false
}
