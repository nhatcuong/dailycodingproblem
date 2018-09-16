// This problem was asked by Google.
// A unival tree (which stands for "universal value") is a tree where all nodes under it have the same value.
// Given the root to a binary tree, count the number of unival subtrees.
// For example, the following tree has 5 unival subtrees:
//    0
//   / \
//  1   0
//     / \
//    1   0
//   / \
//  1   1

package main

import "fmt"

type BinaryTree struct {
	value       int
	left, right *BinaryTree
}

type BinaryTreeUnivalInspectionResult struct {
	univalSubTreeCount int
	isUnival           bool
	uniValue           int
}

func inspectTree(tree *BinaryTree) BinaryTreeUnivalInspectionResult {
	univalSubTreeCount := 0
	isUnival := false
	uniValue := 0
	if tree.left == nil && tree.right == nil {
		return BinaryTreeUnivalInspectionResult{1, true, tree.value}
	}
	if tree.left != nil && tree.right != nil {
		leftInspectionResult := inspectTree(tree.left)
		rightInspectionResult := inspectTree(tree.right)
		isUnival = (leftInspectionResult.isUnival &&
			rightInspectionResult.isUnival &&
			tree.value == leftInspectionResult.uniValue &&
			tree.value == rightInspectionResult.uniValue)
		if isUnival {
			uniValue = tree.value
			univalSubTreeCount = 1
		}
		univalSubTreeCount += (leftInspectionResult.univalSubTreeCount + rightInspectionResult.univalSubTreeCount)
		return BinaryTreeUnivalInspectionResult{
			univalSubTreeCount,
			isUnival,
			uniValue}
	}

	var onlyTreeChild *BinaryTree
	if tree.left != nil {
		onlyTreeChild = tree.left
	} else {
		onlyTreeChild = tree.right
	}
	childInspectionResult := inspectTree(onlyTreeChild)
	isUnival = childInspectionResult.isUnival && childInspectionResult.uniValue == tree.value
	if isUnival {
		uniValue = tree.value
		univalSubTreeCount = 1
	}
	univalSubTreeCount += childInspectionResult.univalSubTreeCount
	return BinaryTreeUnivalInspectionResult{univalSubTreeCount, isUnival, uniValue}
}

func main() {
	left0 := BinaryTree{1, nil, nil}
	right0 := BinaryTree{1, nil, nil}
	left1 := BinaryTree{1, &left0, &right0}
	right1 := BinaryTree{0, nil, nil}
	right2 := BinaryTree{0, &left1, &right1}
	left2 := BinaryTree{1, nil, nil}
	tree := BinaryTree{1, &left2, &right2}
	fmt.Println(inspectTree(&tree).univalSubTreeCount)
}
