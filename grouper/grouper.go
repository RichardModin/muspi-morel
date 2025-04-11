package grouper

import (
	"strings"
	"unicode"
)

// GroupWordsByLength takes input text and returns a map[int][]string
// where each key is a word length, and values are lowercase, unique word slices.
func GroupWordsByLength(input string) map[int][]string {
	words := strings.FieldsFunc(input, func(r rune) bool {
		return !unicode.IsLetter(r)
	})

	wordMap := make(map[int][]string)
	wordSet := make(map[int]map[string]bool) // Tracks seen words per length

	for _, word := range words {
		word = strings.ToLower(word)
		length := len(word)

		// Initialize set if not already created
		if _, exists := wordSet[length]; !exists {
			wordSet[length] = make(map[string]bool)
		}

		// Add only if not already in the set
		if !wordSet[length][word] {
			wordMap[length] = append(wordMap[length], word)
			wordSet[length][word] = true
		}
	}

	return wordMap
}
