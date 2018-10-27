// Given pre-order and in-order traversals of a binary tree, write a function to reconstruct the tree.
// For example, given the following preorder traversal:
// [a, b, d, e, c, f, g]
// And the following inorder traversal:
// [d, b, e, a, f, c, g]
// You should return the following tree:
//     a
//    / \
//   b   c
//  / \ / \
// d  e f  g

package main

import "fmt"

func main() {
	b := reconstructTreeFromTraversals(
		[]string{"a", "b", "d", "e", "c", "f", "g"},
		[]string{"d", "b", "e", "a", "f", "c", "g"},
	)
	fmt.Println(b)
}

type BinaryTree struct {
	value string
	left  *BinaryTree
	right *BinaryTree
}

func reconstructTreeFromTraversals(preOrder, inOrder []string) *BinaryTree {
	if len(preOrder) == 0 {
		return nil
	}
	if len(preOrder) == 1 {
		return &BinaryTree{preOrder[0], nil, nil}
	}

	//get root from preOrder
	rootValue := preOrder[0]
	inOrderRootIndex := 0
	for i := 0; i < len(inOrder); i++ {
		if inOrder[i] == rootValue {
			inOrderRootIndex = i
			break
		}
	}

	//get left tree, right tree traversals inOrder
	var leftInOrderTraversal, leftPreOrderTraversal []string
	if inOrderRootIndex > 0 {
		leftInOrderTraversal = inOrder[:inOrderRootIndex]
		leftPreOrderTraversal = preOrder[1 : inOrderRootIndex+1]
	} else {
		leftInOrderTraversal = make([]string, 0)
		leftPreOrderTraversal = make([]string, 0)
	}

	var rightInOrderTraversal, rightPreOrderTraversal []string
	if inOrderRootIndex+1 < len(inOrder) {
		rightInOrderTraversal = inOrder[inOrderRootIndex+1:]
		rightPreOrderTraversal = preOrder[inOrderRootIndex+1:]
	} else {
		rightInOrderTraversal = make([]string, 0)
		rightPreOrderTraversal = make([]string, 0)
	}

	//construct
	leftTree := reconstructTreeFromTraversals(leftPreOrderTraversal, leftInOrderTraversal)
	rightTree := reconstructTreeFromTraversals(rightPreOrderTraversal, rightInOrderTraversal)

	return &BinaryTree{rootValue, leftTree, rightTree}
}
