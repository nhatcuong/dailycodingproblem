// This problem was asked by Facebook.
// Given the mapping a = 1, b = 2, ... z = 26, and an encoded message, count the number of ways it can be decoded.
// For example, the message '111' would give 3, since it could be decoded as 'aaa', 'ka', and 'ak'.
// You can assume that the messages are decodable. For example, '001' is not allowed.

package main

import (
	"fmt"
	"strconv"
)

func CountWords(input string, allCharacters []string) int {
	if len(input) == 1 {
		return countWordsLength1(input, allCharacters)
	}
	if len(input) == 2 {
		return countWordsLength2(input, allCharacters)
	}

	wordCountStarted2DigitChar := 0
	wordCountStarted1DigitChar := 0
	if isStringEncodedCharacter(input[:2], allCharacters) {
		wordCountStarted2DigitChar = CountWords(input[2:], allCharacters)
	}
	if isStringEncodedCharacter(input[:1], allCharacters) {
		wordCountStarted1DigitChar = CountWords(input[1:], allCharacters)
	}
	return wordCountStarted1DigitChar + wordCountStarted2DigitChar
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
	if isStringEncodedCharacter(input[0:0], allCharacters) {
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
	fmt.Println(allCharacters)
	fmt.Println("Result:")
	fmt.Println(CountWords("100", allCharacters))
}
