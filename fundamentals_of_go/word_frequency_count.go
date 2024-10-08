package fundamentalsofgo

import (
	"strings"
	"unicode"
)

func wordFrequency(s string) map[string]int {
	frequency := make(map[string]int)

	// Normalize the string to lowercase and replace punctuation with spaces
	var sb strings.Builder
	for _, r := range s {
		if !unicode.IsPunct(r) {
			sb.WriteRune(' ')
		} else {
			sb.WriteRune(unicode.ToLower(r))
		}
	}

	// Split the string into words and count the frequency of each word
	words := strings.Fields(sb.String())
	for _, word := range words {
		frequency[word]++
	}

	return frequency
}
