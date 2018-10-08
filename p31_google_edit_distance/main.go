// This problem was asked by Google.
//
// The edit distance between two strings refers to the minimum number of
// character insertions, deletions, and substitutions required to change
// one string to the other. For example, the edit distance between “kitten” and
// “sitting” is three: substitute the “k” for “s”, substitute the “e” for “i”, and append a “g”.
//
// Given two strings, compute the edit distance between them.
//
// The solution here has to cost at len(a) * len(b) in worst case

package main

import "fmt"

func main() {
	// 2 strings with no CommonSubSequence
	// 2 strings with 1 char common sub sequence
	// 2 strings with the first one englobing the other
	// 2 strings with the common subsequence starts at 0 and end at last
	// 2 strings with the common subsequence spreads along
	fmt.Printf("abcdefg 123456 %v %v\n", editDistance("abcdefg", "123456"), longestCommonSubSequence("abcdefg", "123456"))
	fmt.Printf("abcdefg 123c456 %v %v\n", editDistance("abcdefg", "123c456"), longestCommonSubSequence("abcdefg", "123c456"))
	fmt.Printf("abcdefg cdef %v %v\n", editDistance("abcdefg", "cdef"), longestCommonSubSequence("abcdefg", "cdef"))
	fmt.Printf("abcdefg a12c34ed56e56g %v %v\n", editDistance("abcdefg", "a12c34ed56e56g"), longestCommonSubSequence("abcdefg", "a12c34ed56e56g"))
	fmt.Printf("abcdefg 12b3fg4ed56e56g %v %v\n", editDistance("abcdefg", "12b3fg4ed56e56g"), longestCommonSubSequence("abcdefg", "12b3fg4ed56e56g"))
}

func editDistance(a, b string) int {
	c := longestCommonSubSequence(a, b)
	distanceAC := len(a) - len(c)
	distanceBC := len(b) - len(c)
	if distanceAC > distanceBC {
		return distanceAC
	}
	return distanceBC
}

func longestCommonSubSequence(a, b string) string {
	return longestCommonSubSequenceRecurrent(a, b, 0, 0, make(map[[2]int]string))
}

func longestCommonSubSequenceRecurrent(a, b string, indexA, indexB int, tempResult map[[2]int]string) string {
	switch {
	case indexA >= len(a) || indexB >= len(b):
		return ""
	case a[indexA] == b[indexB]:
		result := string(a[indexA]) + longestCommonSubSequenceRecurrent(a, b, indexA+1, indexB+1, tempResult)
		tempResult[[2]int{indexA, indexB}] = result
		return result
	default:
		longestWithIndexA := longestCommonSubSequenceRecurrent(a, b, indexA, indexB+1, tempResult)
		longestWithIndexB := longestCommonSubSequenceRecurrent(a, b, indexA+1, indexB, tempResult)
		if len(longestWithIndexA) > len(longestWithIndexB) {
			return longestWithIndexA
		}
		return longestWithIndexB
	}
}
