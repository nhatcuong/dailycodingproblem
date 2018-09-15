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
	wordCountStarted1DigitChar = CountWords(input[1:], allCharacters)
	return wordCountStarted1DigitChar + wordCountStarted2DigitChar
}

func countWordsLength1(input string, allCharacters []string) int {
	return 1
}

func countWordsLength2(input string, allCharacters []string) int {
	if isStringEncodedCharacter(input, allCharacters) {
		return 2
	}
	return 1
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
	fmt.Println(CountWords("111", allCharacters))
}
