package helpers

import "strings"

func IsWord(s string) bool {
	for _, r := range s {
		if !IsLetter(r) {
			return false
		}
	}
	return true
}

func IsLetter(r rune) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') || r == '\''
}

func IsCapitalized(s string) bool {
	return len(s) > 0 && strings.ToUpper(string(s[0])) == string(s[0])
}
