package replacer

import (
	"math/rand"
	"regexp"
	"reverse-lorem-ipsum/helpers"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ReplaceWordsByLength(text string, wordMap map[int][]string) string {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Match word and punctuation as separate tokens
	re := regexp.MustCompile(`[\w']+|[^\w\s]+|\s+`)

	tokens := re.FindAllString(text, -1)

	// Create a Title casing transformer
	titleCaser := cases.Title(language.Und)

	for i, token := range tokens {
		if helpers.IsWord(token) {
			length := len(token)
			if replacements, exists := wordMap[length]; exists && len(replacements) > 0 {
				replacement := replacements[rng.Intn(len(replacements))]

				// Preserve original casing
				if helpers.IsCapitalized(token) {
					replacement = titleCaser.String(replacement)
				}
				tokens[i] = replacement
			}
		}
	}

	return strings.Join(tokens, "")
}
