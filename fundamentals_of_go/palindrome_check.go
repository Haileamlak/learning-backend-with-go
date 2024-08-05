package fundamentalsofgo

import "unicode"

func IsPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !unicode.IsLetter(rune(s[left])) {
			left++
		}
		for left < right && !unicode.IsLetter(rune(s[right])) {
			right--
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}