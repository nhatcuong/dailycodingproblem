// This problem was asked by Airbnb.
// Given a list of integers, write a function that returns the largest sum of non-adjacent numbers. Numbers can be 0 or negative.
// For example, [2, 4, 6, 2, 5] should return 13, since we pick 2, 6, and 5. [5, 1, 1, 5] should return 10, since we pick 5 and 5.

package main

import "fmt"

func MaxNonAdjacentSum(inputArray []int, fromIndex int, mapTempResult map[int]int) int {
	if val, ok := mapTempResult[fromIndex]; ok {
		return val
	}
	switch l := len(inputArray) - fromIndex; l {
	case 0:
		mapTempResult[fromIndex] = 0
		return 0
	case 1:
		mapTempResult[fromIndex] = inputArray[fromIndex]
		return inputArray[fromIndex]
	default:
		maxWithFirstElt := 0
		if firstElt := inputArray[fromIndex]; firstElt > 0 {
			maxWithFirstElt = firstElt
		}
		maxWithFirstElt = maxWithFirstElt + MaxNonAdjacentSum(inputArray, fromIndex+2, mapTempResult)

		maxWithoutFirstElt := MaxNonAdjacentSum(inputArray, fromIndex+1, mapTempResult)

		if maxWithFirstElt > maxWithoutFirstElt {
			mapTempResult[fromIndex] = maxWithFirstElt
			return maxWithFirstElt
		}
		mapTempResult[fromIndex] = maxWithoutFirstElt
		return maxWithoutFirstElt
	}
}

func main() {
	fmt.Println(MaxNonAdjacentSum([]int{5, 1, 1, 5}, 0, make(map[int]int)))
}
