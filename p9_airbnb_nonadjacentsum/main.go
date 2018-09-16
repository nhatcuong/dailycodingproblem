package main

import "fmt"

func MaxNonAdjacentSum(inputArray []int) int {
	switch l := len(inputArray); l {
	case 0:
		return 0
	case 1:
		return inputArray[0]
	default:
		maxWithFirstElt := 0
		if inputArray[0] > 0 {
			maxWithFirstElt = inputArray[0]
		}
		maxWithFirstElt = maxWithFirstElt + MaxNonAdjacentSum(inputArray[2:])

		maxWithoutFirstElt := MaxNonAdjacentSum(inputArray[1:])

		if maxWithFirstElt > maxWithoutFirstElt {
			return maxWithFirstElt
		}
		return maxWithoutFirstElt
	}
}

func main() {
	fmt.Println(MaxNonAdjacentSum([]int{1, -9, 3, 7}))
}
