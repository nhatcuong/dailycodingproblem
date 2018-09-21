// This problem was asked by Facebook.
// Given the mapping a = 1, b = 2, ... z = 26, and an encoded message, count the number of ways it can be decoded.
// For example, the message '111' would give 3, since it could be decoded as 'aaa', 'ka', and 'ak'.
// You can assume that the messages are decodable. For example, '001' is not allowed.

package main

import (
	"fmt"
	"strconv"
)

func CountWords(input string, fromIndex int, cachedResult *map[int]int, allCharacters []string) int {
	if result, ok := (*cachedResult)[fromIndex]; ok {
		return result
	}

	if len(input)-fromIndex == 1 {
		result := countWordsLength1(input[fromIndex:], allCharacters)
		(*cachedResult)[fromIndex] = result
		return result
	}
	if len(input)-fromIndex == 2 {
		result := countWordsLength2(input[fromIndex:], allCharacters)
		(*cachedResult)[fromIndex] = result
		return result
	}

	wordCountStarted2DigitChar := 0
	wordCountStarted1DigitChar := 0
	if isStringEncodedCharacter(input[fromIndex:fromIndex+2], allCharacters) {
		wordCountStarted2DigitChar = CountWords(input, fromIndex+2, cachedResult, allCharacters)
	}
	if isStringEncodedCharacter(input[fromIndex:fromIndex+1], allCharacters) {
		wordCountStarted1DigitChar = CountWords(input, fromIndex+1, cachedResult, allCharacters)
	}
	result := wordCountStarted1DigitChar + wordCountStarted2DigitChar
	(*cachedResult)[fromIndex] = result
	return result
}

func countWordsLength1(input string, allCharacters []string) int {
	if isStringEncodedCharacter(input, allCharacters) {
		return 1
	}
	return 0
}

func countWordsLength2(input string, allCharacters []string) int {
	result := 0
	if isStringEncodedCharacter(input, allCharacters) {
		result += 1
	}
	if isStringEncodedCharacter(input[:1], allCharacters) {
		result += 1
	}
	return result
}

func isStringEncodedCharacter(input string, allCharacters []string) bool {
	for _, c := range allCharacters {
		if c == input {
			return true
		}
	}
	return false
}

func main() {
	var allCharacters []string = make([]string, 26)
	for i := 0; i < 26; i++ {
		allCharacters[i] = strconv.Itoa(i + 1)
	}
	cachedResult := make(map[int]int, 0)
	fmt.Println(allCharacters)
	fmt.Println("Result:")
	fmt.Println(CountWords("11011", 0, &cachedResult, allCharacters))
}
