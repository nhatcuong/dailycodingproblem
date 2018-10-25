// Given a string, find the longest palindromic contiguous substring. If there are more than one with the maximum length, return any one.
// For example, the longest palindromic substring of "aabcdcb" is "bcdcb". The longest palindromic substring of "bananas" is "anana".

package main

import "fmt"

func main() {
	fmt.Println(longest_palindrom_substring("abcd"))
	fmt.Println(longest_palindrom_substring("a"))
	fmt.Println(longest_palindrom_substring("bb"))
	fmt.Println(longest_palindrom_substring("bananas"))
	fmt.Println(longest_palindrom_substring("aabcdcb"))
}

func longest_palindrom_substring(input string) string {
	l := len(input)
	tempResult := make([][]bool, l)
	for i := 0; i < l; i++ {
		tempResult[i] = make([]bool, l)
	}

	iMax, jMax := 0, 0

	for d := 0; d < l; d++ {
		for i := 0; i < l-d; i++ {
			j := i + d

			if d == 0 {
				tempResult[i][j] = true
			} else if d == 1 {
				if input[i] == input[j] {
					tempResult[i][j] = true
				}
			} else {
				if input[i] == input[j] && tempResult[i+1][j-1] {
					tempResult[i][j] = true
				}
			}

			if j-i > jMax-iMax && tempResult[i][j] {
				iMax, jMax = i, j
			}

		}
	}

	return input[iMax : jMax+1]
}
