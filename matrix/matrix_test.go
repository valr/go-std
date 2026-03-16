package matrix

import (
	"slices"
	"testing"
)

func TestRotateStrMatrix(t *testing.T) {
	tests := []struct {
		name     string
		matrix   []string
		expected []string
	}{
		{
			"2x2",
			[]string{"ab", "cd"},
			[]string{"ac", "bd"},
		},
		{
			"3x3",
			[]string{"abc", "def", "ghi"},
			[]string{"adg", "beh", "cfi"},
		},
		{
			"2x3",
			[]string{"abc", "def"},
			[]string{"ad", "be", "cf"},
		},
		{
			"3x2",
			[]string{"ab", "cd", "ef"},
			[]string{"ace", "bdf"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RotateStrMatrix(tt.matrix)
			if !slices.Equal(result, tt.expected) {
				t.Errorf("RotateStrMatrix(%v) returned %v; want %v", tt.matrix, result, tt.expected)
			}
		})
	}
}

func TestDiagonalStrMatrix(t *testing.T) {
	tests := []struct {
		name     string
		matrix   []string
		expected []string
	}{
		{
			"2x2",
			[]string{"ab", "cd"},
			[]string{"a", "cb", "d"},
		},
		{
			"3x3",
			[]string{"abc", "def", "ghi"},
			[]string{"a", "db", "gec", "hf", "i"},
		},
		{
			"2x3",
			[]string{"abc", "def"},
			[]string{"a", "db", "ec", "f"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DiagonalStrMatrix(tt.matrix)
			if !slices.Equal(result, tt.expected) {
				t.Errorf("DiagonalStrMatrix(%v) returned %v; want %v", tt.matrix, result, tt.expected)
			}
		})
	}
}

func TestCounterDiagonalStrMatrix(t *testing.T) {
	tests := []struct {
		name     string
		matrix   []string
		expected []string
	}{
		{
			"2x2",
			[]string{"ab", "cd"},
			[]string{"b", "da", "c"},
		},
		{
			"3x3",
			[]string{"abc", "def", "ghi"},
			[]string{"c", "fb", "iea", "hd", "g"},
		},
		{
			"2x3",
			[]string{"abc", "def"},
			[]string{"c", "fb", "ea", "d"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CounterDiagonalStrMatrix(tt.matrix)
			if !slices.Equal(result, tt.expected) {
				t.Errorf("CounterDiagonalStrMatrix(%v) returned %v; want %v", tt.matrix, result, tt.expected)
			}
		})
	}
}
