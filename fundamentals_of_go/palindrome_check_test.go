package fundamentalsofgo

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"aba", true},
		{"abc", false},
		{"abA", true},
	}

	for _, test := range tests {
		if output := IsPalindrome(test.input); output != test.expected {
			t.Errorf("IsPalindrome(%q) = %v, want %v", test.input, output, test.expected)
		}
	}
}
