package strconvx

import (
	"bytes"
	"log"
	"testing"
)

func TestStrToIntSuccess(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"positive", "42", 42},
		{"zero", "0", 0},
		{"negative", "-15", -15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StrToInt(tt.input)
			if result != tt.expected {
				t.Errorf("StrToInt(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestStrToIntPanicOnInvalid(t *testing.T) {
	originalOutput := log.Writer()
	var redirectedOutput bytes.Buffer
	log.SetOutput(&redirectedOutput)
	defer log.SetOutput(originalOutput)

	input := "NaN"
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic for invalid string %q, but did not panic", input)
		}
	}()
	StrToInt(input)
}

func TestIntToStr(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{"positive", 42, "42"},
		{"zero", 0, "0"},
		{"negative", -15, "-15"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IntToStr(tt.input)
			if result != tt.expected {
				t.Errorf("IntToStr(%v) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}
