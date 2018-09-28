// This problem was asked by Twitter.
// Implement an autocomplete system. That is, given a query string s and a set of all possible query strings, return all strings in the set that have s as a prefix.
// For example, given the query string de and the set of strings [dog, deer, deal], return [deer, deal].
// Hint: Try preprocessing the dictionary into a more efficient data structure to speed up queries.

// I use here a structure that resembles to both Tree and Automate, since
// a word make transition from one node to another, but never backward.
// In this solution, adding a word to the index costs linearly to the length
// of the word; querying suggestions costs linearly to the sum of the size of
// all results.

package main

import "fmt"

func main() {
	root := initAutocompleteDictNode()
	root.digestWord("deer")
	root.digestWord("dear")
	root.digestWord("doe")
	root.digestWord("art")
	fmt.Println(root.getAllWords())
	fmt.Printf("Suggestions are: %v\n", root.retrieveAutocompleteWords("a"))
}

type AutocompleteDictNode struct {
	nextWordMap     map[byte]*AutocompleteDictNode
	isSingleWordEnd bool
	// can find autocomplete for words
}

func initAutocompleteDictNode() AutocompleteDictNode {
	result := AutocompleteDictNode{
		nextWordMap:     make(map[byte]*AutocompleteDictNode),
		isSingleWordEnd: false,
	}
	return result
}

func (node *AutocompleteDictNode) digestWord(word string) {
	currentNode := node
	for i := 0; i < len(word); i++ {
		if nextNode, ok := currentNode.nextWordMap[word[i]]; ok {
			currentNode = nextNode
		} else {
			currentNode = currentNode.addCharacter(word[i])
		}
	}
	currentNode.markEndNode()
}

func (node *AutocompleteDictNode) addCharacter(character byte) *AutocompleteDictNode {
	newNode := initAutocompleteDictNode()
	node.nextWordMap[character] = &newNode
	return &newNode
}

func (node *AutocompleteDictNode) markEndNode() {
	node.isSingleWordEnd = true
}

func (node *AutocompleteDictNode) getAllWords() []string {
	if node.isSingleWordEnd {
		return []string{""}
	}
	results := make([]string, 0)
	for k, v := range node.nextWordMap {
		subWords := v.getAllWords()
		for i := 0; i < len(subWords); i++ {
			subWords[i] = string(k) + subWords[i]
		}
		results = append(results, subWords...)
	}
	return results
}

func (dict *AutocompleteDictNode) retrieveAutocompleteWords(inputWord string) []string {
	// getAutocompleteNodeForWord
	autocompleteNode := dict
	for i := 0; i < len(inputWord); i++ {
		if nextNode, ok := autocompleteNode.nextWordMap[inputWord[i]]; ok {
			autocompleteNode = nextNode
		} else {
			break
		}
	}
	possibleSuffixes := autocompleteNode.getAllWords()
	suggestions := make([]string, len(possibleSuffixes))
	for i := 0; i < len(possibleSuffixes); i++ {
		suggestions[i] = inputWord + possibleSuffixes[i]
	}
	return suggestions
}
