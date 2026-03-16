package stringsx

import (
	"slices"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty", "", ""},
		{"string", "coucou", "uocuoc"},
		{"palindrome", "kayak", "kayak"},
		{"unicode", "カッコウ", "ウコッカ"},
		{"emoji", "😻👍", "👍😻"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reverse(tt.input)
			if result != tt.expected {
				t.Errorf("Reverse(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSplitAny(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		separators string
		expected   []string
	}{
		{"none", "one,two,three", "", []string{"one,two,three"}},
		{"comma", "one,two,three", ",", []string{"one", "two", "three"}},
		{"multiple", "one;two,three four", ",; ", []string{"one", "two", "three", "four"}},
		{"leading separators", ",one,two", ",", []string{"one", "two"}},
		{"leading multiple separators", ";,one,two", ",;", []string{"one", "two"}},
		{"trailing separators", "one,two,", ",", []string{"one", "two"}},
		{"trailing multiple separators", "one,two,;", ",;", []string{"one", "two"}},
		{"consecutive separators", "one,,two", ",", []string{"one", "two"}},
		{"consecutive multiple separators", "one,;,two", ",;", []string{"one", "two"}},
		{"unicode separators", "one🐈two🐈three", "🐈", []string{"one", "two", "three"}},
		{"unicode multiple separators", "one🐈two🐱three", "🐱🐈", []string{"one", "two", "three"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SplitAny(tt.input, tt.separators)
			if !slices.Equal(result, tt.expected) {
				t.Errorf("SplitAny(%q, %q) returned %v; want %v", tt.input, tt.separators, result, tt.expected)
			}
		})
	}
}
