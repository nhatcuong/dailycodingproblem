// You run an e-commerce website and want to record the last N order ids in a log. Implement a data structure to accomplish this, with the following API:
//
// record(order_id): adds the order_id to the log
// get_last(i): gets the ith last element from the log. i is guaranteed to be smaller than or equal to N.
// You should be as efficient with time and space as possible.

// This solution here is linear in space all the time, constant in time when record and linear in time when query

package main

import "fmt"

type LogNode struct {
	previousNode *LogNode // closest sooner node
	nextNode     *LogNode // closest later node
	value        int
}

type LogStack struct {
	capacity       int
	size           int
	mostRecentNode *LogNode
	oldestNode     *LogNode
}

func NewLogStack(maxSize int) LogStack {
	return LogStack{
		capacity:       maxSize,
		size:           0,
		mostRecentNode: nil,
		oldestNode:     nil,
	}
}

func (ls *LogStack) record(orderId int) {
	newNode := LogNode{
		previousNode: nil,
		nextNode:     nil,
		value:        orderId,
	}
	if ls.oldestNode == nil { // empty LogStack
		ls.oldestNode = &newNode
		ls.mostRecentNode = &newNode
		ls.size = 1
		return
	}

	newNode.previousNode = ls.mostRecentNode
	ls.mostRecentNode.nextNode = &newNode
	ls.mostRecentNode = &newNode

	if ls.size == ls.capacity {
		secondOldestNode := ls.oldestNode.nextNode
		secondOldestNode.previousNode = nil
		ls.oldestNode = secondOldestNode
	} else {
		ls.size += 1
	}
}

func (ls *LogStack) getLast(index int) int {
	if index >= ls.size {
		return 0
	}
	result := ls.mostRecentNode
	for i := 0; i < index; i++ {
		result = result.previousNode
	}
	return result.value
}

func main() {
	logStack := NewLogStack(4)
	logStack.record(0)
	logStack.record(1)
	fmt.Printf("getLast 2: %v\n", logStack.getLast(2)) //must be 0
	fmt.Printf("getLast 0: %v\n", logStack.getLast(0)) //must be 1
	logStack.record(2)
	logStack.record(3)
	logStack.record(4)
	fmt.Printf("getLast 2: %v\n", logStack.getLast(2)) //must be 2
}
