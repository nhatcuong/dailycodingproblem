// This problem was asked by Google.
// Implement locking in a binary tree. A binary tree node can be locked or
// unlocked only if all of its descendants or ancestors are not locked.
// Design a binary tree node class with the following methods:

// is_locked, which returns whether the node is locked

// lock, which attempts to lock the node. If it cannot be locked, then it should
// return false. Otherwise, it should lock it and return true.

// unlock, which unlocks the node. If it cannot be unlocked, then it should
// return false. Otherwise, it should unlock it and return true.

// You may augment the node to add parent pointers or any other property you
// would like. You may assume the class is used in a single-threaded program, so
// there is no need for actual locks or mutexes. Each method should run in O(h),
// where h is the height of the tree.

package main

import "fmt"

func main() {
	root := initLockingTree("root")
	l := root.appendLeftChild("l")
	root.appendRightChild("r")
	ll := l.appendLeftChild("ll")
	ll.appendLeftChild("lll")
	l.appendRightChild("lr")
	llr := ll.appendRightChild("llr")
	fmt.Printf("l is lock available %v\n", l.isLockAvailable())
	fmt.Printf("Lock l done: %v\n", l.lock())
	fmt.Printf("Lock llr done: %v\n", llr.lock())
	fmt.Printf("Unlock l done: %v\n", l.unlock())
	fmt.Printf("Lock llr done: %v\n", llr.lock())
	fmt.Printf("Lock l done: %v\n", l.lock())
}

type LockingTree struct {
	value               string
	isLocked            bool
	lockedChildrenCount int
	leftChild           *LockingTree
	rightChild          *LockingTree
	parent              *LockingTree
}

func initLockingTree(value string) *LockingTree {
	return &LockingTree{
		value:               value,
		isLocked:            false,
		lockedChildrenCount: 0,
		leftChild:           nil,
		rightChild:          nil,
		parent:              nil,
	}
}

func (t *LockingTree) appendLeftChild(value string) *LockingTree {
	newTree := initLockingTree(value)
	t.leftChild = newTree
	newTree.parent = t
	return newTree
}

func (t *LockingTree) appendRightChild(value string) *LockingTree {
	newTree := initLockingTree(value)
	t.rightChild = newTree
	newTree.parent = t
	return newTree
}

func (t *LockingTree) augmentAscendantsLockedChildrenCount() {
	currentTree := t
	for currentTree.parent != nil {
		currentTree.parent.lockedChildrenCount++
		currentTree = currentTree.parent
	}
}

func (t *LockingTree) isLockAvailable() bool {
	hasLockedChildren := t.lockedChildrenCount > 0
	if hasLockedChildren {
		return false
	}
	currentTree := t
	for currentTree.parent != nil {
		if currentTree.parent.isLocked {
			return false
		}
		currentTree = currentTree.parent
	}
	return true
}

func (t *LockingTree) lock() bool {
	if !t.isLocked && t.isLockAvailable() {
		t.isLocked = true
		t.augmentAscendantsLockedChildrenCount()
		return true
	}
	return false
}

func (t *LockingTree) decreeaseAscendantsLockedChildrenCount() {
	currentTree := t
	for currentTree.parent != nil {
		currentTree.parent.lockedChildrenCount--
		currentTree = currentTree.parent
	}
}

func (t *LockingTree) unlock() bool {
	if t.isLocked && t.isLockAvailable() {
		t.isLocked = false
		t.decreeaseAscendantsLockedChildrenCount()
		return true
	}
	return false
}
