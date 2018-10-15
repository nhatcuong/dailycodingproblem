// You have an N by N board. Write a function that, given N, returns the number
// of possible arrangements of the board where N queens can be placed on the
// board without threatening each other, i.e. no two queens share the same row,
// column, or diagonal.

package main

import "fmt"

func main() {
	fmt.Printf("%v\n", totalArrangementsCount(2))
	// count for 2
	fmt.Printf("%v\n", totalArrangementsCount(4))
	// count for 4
	fmt.Printf("%v\n", totalArrangementsCount(8))
	// count for 8
}

func totalArrangementsCount(N int) int {
	emptyBoard := make([][]bool, N)
	for i := 0; i < len(emptyBoard); i++ {
		emptyBoard[i] = make([]bool, N)
	}
	return arrangementsCount(0, emptyBoard)
}

func arrangementsCount(placedQueensCount int, markedBoard [][]bool) int {
	// all placed queens are on column from 0 to n-1
	// next queen will be in column n
	// if n is the size of the board so return 1

	if placedQueensCount == len(markedBoard) {
		return 1
	}

	// find all possible squares for queen n+1
	// for all possible square, do the following:
	// mark all places on board which related to this new placement
	// compute the next arrangementsCount for n+1 and the new board
	// sum up all the result count.
	result := 0
	for i := 0; i < len(markedBoard); i++ {
		if !markedBoard[placedQueensCount][i] {
			cloneMarkedBoard := make([][]bool, len(markedBoard))
			for j := 0; j < len(markedBoard); j++ {
				cloneMarkedBoard[j] = make([]bool, len(markedBoard))
				copy(cloneMarkedBoard[j], markedBoard[j])
			}
			newMarkedBoard := markBoardFromNewPosition(placedQueensCount, i, cloneMarkedBoard)
			arrangementCountWithNewPlacement := arrangementsCount(placedQueensCount+1, newMarkedBoard)
			result += arrangementCountWithNewPlacement
		}
	}
	return result
}

func markBoardFromNewPosition(row int, column int, markedBoard [][]bool) [][]bool {
	for i := 0; i < len(markedBoard); i++ {
		markedBoard[row][i] = true
	}
	for j := 0; j < len(markedBoard); j++ {
		markedBoard[j][column] = true
	}
	for k := 0; k < len(markedBoard); k++ {
		if row+k < len(markedBoard) {
			if column+k < len(markedBoard) {
				markedBoard[row+k][column+k] = true
			}
			if column-k >= 0 {
				markedBoard[row+k][column-k] = true
			}
		}
		if row-k > len(markedBoard) {
			if column+k < len(markedBoard) {
				markedBoard[row-k][column+k] = true
			}
			if column-k >= 0 {
				markedBoard[row-k][column-k] = true
			}
		}
	}
	return markedBoard
}
