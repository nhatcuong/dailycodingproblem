// This problem was asked by Google.
// The power set of a set is the set of all its subsets. Write a function that, given a set, generates its power set.
// For example, given the set {1, 2, 3}, it should return {{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}.

package main

import "fmt"

func main() {
	fmt.Printf("%v\n", powerSet([]int{1, 2, 3}))
	fmt.Printf("%v\n", powerSet([]int{}))
	// power set of 1 2 3
	// power set of {}
}

func powerSet(inputSet []int) [][]int {
	if len(inputSet) == 0 {
		result := make([][]int, 0)
		result = append(result, make([]int, 0))
		return result
	}
	head, queue := inputSet[0], inputSet[1:]
	powerSetQueue := powerSet(queue)
	powerSetWithHead := make([][]int, 0)
	for _, s := range powerSetQueue {
		newSubset := append([]int{head}, s...)
		powerSetWithHead = append(powerSetWithHead, newSubset)
	}
	return append(powerSetQueue, powerSetWithHead...)
}
