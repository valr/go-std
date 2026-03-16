package mathx

import (
	"testing"
)

func TestAbsInt(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"positive", 5, 5},
		{"negative", -3, 3},
		{"zero", 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Abs(tt.input)
			if result != tt.expected {
				t.Errorf("Abs(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestAbsFloat64(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"positive", 5.5, 5.5},
		{"negative", -3.5, 3.5},
		{"zero", 0.0, 0.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Abs(tt.input)
			if result != tt.expected {
				t.Errorf("Abs(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestProductInt(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"empty", []int{}, 0},
		{"single", []int{5}, 5},
		{"multiple", []int{2, 3, 4}, 24},
		{"with zero", []int{2, 0, 4}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Product(tt.input...)
			if result != tt.expected {
				t.Errorf("Product(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestProductFloat64(t *testing.T) {
	tests := []struct {
		name     string
		input    []float64
		expected float64
	}{
		{"empty", []float64{}, 0.0},
		{"single", []float64{5.5}, 5.5},
		{"multiple", []float64{2.5, 3.0, 4.7}, 35.25},
		{"with zero", []float64{2.5, 0.0, 4.7}, 0.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Product(tt.input...)
			if result != tt.expected {
				t.Errorf("Product(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSumInt(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"empty", []int{}, 0},
		{"single", []int{5}, 5},
		{"multiple", []int{1, 2, 3}, 6},
		{"negative", []int{-1, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sum(tt.input...)
			if result != tt.expected {
				t.Errorf("Sum(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSumFloat64(t *testing.T) {
	tests := []struct {
		name     string
		input    []float64
		expected float64
	}{
		{"empty", []float64{}, 0.0},
		{"single", []float64{5.5}, 5.5},
		{"multiple", []float64{1.5, 2.2, 3.7}, 7.4},
		{"negative", []float64{-1.5, 1.5}, 0.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sum(tt.input...)
			if result != tt.expected {
				t.Errorf("Sum(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
