package util

import (
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestSourceInfo(t *testing.T) {
	info := SourceInfo()
	subs := strings.Split(info, ":")
	if len(subs) != 2 {
		t.Fatalf("SourceInfo() returned %q; expected format \"file:line\"", info)
	}
	if subs[0] != "util_test.go" {
		t.Errorf("SourceInfo() filename %q is not \"util_test.go\"", subs[0])
	}
	if _, err := strconv.Atoi(subs[1]); err != nil {
		t.Errorf("SourceInfo() line number %q is not a valid integer", subs[1])
	}
}

func TestWrapSlices(t *testing.T) {
	tests := []struct {
		name     string
		input    []any
		expected []any
	}{
		{"empty", []any{}, []any{}},
		{"single int", []any{42}, []any{42}},
		{"multiple ints", []any{1, 2, 3}, []any{1, 2, 3}},
		{"single string", []any{"coucou"}, []any{"coucou"}},
		{"multiple strings", []any{"hello", "world"}, []any{"hello", "world"}},
		{"mixed types", []any{1, "two", 3.0, true, nil}, []any{1, "two", 3.0, true, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Wrap(tt.input...)
			if !slices.Equal(result, tt.expected) {
				t.Errorf("Wrap(%v) returned %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestWrapVariadic(t *testing.T) {
	tests := []struct {
		name     string
		fn       func() []any
		expected []any
	}{
		{"Wrap(1, 2, 3)", func() []any { return Wrap(1, 2, 3) }, []any{1, 2, 3}},
		{"Wrap(\"a\", \"b\")", func() []any { return Wrap("a", "b") }, []any{"a", "b"}},
		{"Wrap(true, false)", func() []any { return Wrap(true, false) }, []any{true, false}},
		{"Wrap(1, \"two\", 3.0)", func() []any { return Wrap(1, "two", 3.0) }, []any{1, "two", 3.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.fn()
			if !slices.Equal(result, tt.expected) {
				t.Errorf("%v returned %v; want %v", tt.name, result, tt.expected)
			}
		})
	}
}
