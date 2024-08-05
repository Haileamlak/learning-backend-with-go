package fundamentalsofgo

import "testing"

func TestWordFrequency(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{"", map[string]int{}},
		{"a", map[string]int{"a": 1}},
		{"aa ,", map[string]int{"aa": 1}},
		{"ab -", map[string]int{"ab": 1}},
		{"aba ==", map[string]int{"aba": 1}},
		{"abc", map[string]int{"abc": 1}},
		{"a  a", map[string]int{"a": 2}},
		{"a a !- a", map[string]int{"a": 3}},
		{"a b a", map[string]int{"a": 2, "b": 1}},
		{"a b c", map[string]int{"a": 1, "b": 1, "c": 1}},
		{"a b c a", map[string]int{"a": 2, "b": 1, "c": 1}},
		{"a b c a b", map[string]int{"a": 2, "b": 2, "c": 1}},
		{"a b c a b c", map[string]int{"a": 2, "b": 2, "c": 2}},
		{"a b c a b c a", map[string]int{"a": 3, "b": 2, "c": 2}},
		{"a b c a b c a b", map[string]int{"a": 3, "b": 3, "c": 2}},
		{"a b c a b c a b c", map[string]int{"a": 3, "b": 3, "c": 3}},
		{"a b c a b c a b c a", map[string]int{"a": 4, "b": 3, "c": 3}},
		{"a b c a b c a b c a b", map[string]int{"a": 4, "b": 4, "c": 3}},
		{"a b c a b c a b c a b c", map[string]int{"a": 4, "b": 4, "c": 4}},
		{"a b c a b c a b c a b c a", map[string]int{"a": 5, "b": 4, "c": 4}},
		{"a b c a b c a b c a b c a b", map[string]int{"a": 5, "b": 5, "c": 4}},
		{"a b c a b c a b c a b c a b c", map[string]int{"a": 5, "b": 5, "c": 5}},
	}

	for _, test := range tests {
		if output := wordFrequency(test.input); !equal(output, test.expected) {
			t.Errorf("wordFrequency(%q) = %v, want %v", test.input, output, test.expected)
		}
	}

}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
