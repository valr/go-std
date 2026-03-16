package slicesx

import (
	"bytes"
	"log"
	"slices"
	"testing"
)

func TestCount(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		value    int
		expected int
	}{
		{"empty slice", []int{}, 1, 0},
		{"no occurrences", []int{1, 2, 3}, 4, 0},
		{"single occurrence", []int{1, 2, 3}, 2, 1},
		{"multiple occurrences", []int{1, 2, 1, 2, 1}, 1, 3},
		{"all same occurrences", []int{5, 5, 5}, 5, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Count(tt.slice, tt.value)
			if result != tt.expected {
				t.Errorf("Count(%v, %v) = %v; want %v", tt.slice, tt.value, result, tt.expected)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	tests := []struct {
		name     string
		count    int
		value    int
		expected []int
	}{
		{"zero count", 0, 5, []int{}},
		{"single element", 1, 5, []int{5}},
		{"multiple elements", 8, 8, []int{8, 8, 8, 8, 8, 8, 8, 8}},
		{"multiple elements", 9, 9, []int{9, 9, 9, 9, 9, 9, 9, 9, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Create(tt.count, tt.value)
			if !slices.Equal(result, tt.expected) {
				t.Errorf("Create(%v, %v) returned %v; want %v", tt.count, tt.value, result, tt.expected)
			}
		})
	}
}

func TestCreatePanicOnNegative(t *testing.T) {
	originalOutput := log.Writer()
	var redirectedOutput bytes.Buffer
	log.SetOutput(&redirectedOutput)
	defer log.SetOutput(originalOutput)

	input := -1
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic for negative count %v, but did not panic", input)
		}
	}()
	Create(input, 5)
}

func TestFilter(t *testing.T) {
	isEven := func(x int) bool { return x%2 == 0 }

	tests := []struct {
		name     string
		slice    []int
		expected []int
	}{
		{"empty slice", []int{}, []int{}},
		{"single element", []int{1}, []int{}},
		{"no matches", []int{1, 3, 5}, []int{}},
		{"all matches", []int{2, 4, 6}, []int{2, 4, 6}},
		{"mixed matches", []int{1, 3, 2, 4}, []int{2, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Filter(tt.slice, isEven)
			if !slices.Equal(result, tt.expected) {
				t.Errorf("Filter(%v, isEven) returned %v; want %v", tt.slice, result, tt.expected)
			}
		})
	}
}

func TestFilter2(t *testing.T) {
	isValueEqualsIndexPlusOne := func(x, idx int) bool { return x == idx+1 }

	tests := []struct {
		name     string
		slice    []int
		expected []int
	}{
		{"empty slice", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"no matches", []int{2, 3, 4}, []int{}},
		{"all matches", []int{1, 2, 3}, []int{1, 2, 3}},
		{"mixed matches", []int{1, 3, 2, 4}, []int{1, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Filter2(tt.slice, isValueEqualsIndexPlusOne)
			if !slices.Equal(result, tt.expected) {
				t.Errorf("Filter2(%v, isValueEqualsIndexPlusOne) returned %v; want %v", tt.slice, result, tt.expected)
			}
		})
	}
}

func TestFlatten(t *testing.T) {
	tests := []struct {
		name     string
		slices   [][]int
		expected []int
	}{
		{"empty", [][]int{}, []int{}},
		{"single empty", [][]int{{}}, []int{}},
		{"single slice", [][]int{{1, 2}}, []int{1, 2}},
		{"multiple slices", [][]int{{1, 2}, {3, 4}, {5}}, []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Flatten(tt.slices)
			if !slices.Equal(result, tt.expected) {
				t.Errorf("Flatten(%v) returned %v; want %v", tt.slices, result, tt.expected)
			}
		})
	}
}

func TestIndex(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		sub      []int
		expected int
	}{
		{"empty slice and sub", []int{}, []int{}, 0},
		{"empty slice", []int{}, []int{4, 5}, -1},
		{"empty sub", []int{1, 2, 3}, []int{}, 0},
		{"sub not found", []int{1, 2, 3}, []int{4, 5}, -1},
		{"sub at start", []int{1, 2, 3, 4}, []int{1, 2}, 0},
		{"sub in middle", []int{1, 2, 3, 4}, []int{2, 3}, 1},
		{"sub at end", []int{1, 2, 3, 4}, []int{3, 4}, 2},
		{"single element", []int{1, 2, 3}, []int{2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Index(tt.slice, tt.sub)
			if result != tt.expected {
				t.Errorf("Index(%v, %v) = %v; want %v", tt.slice, tt.sub, result, tt.expected)
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	tests := []struct {
		name     string
		slices   [][]int
		expected []int
	}{
		{"empty slice", [][]int{}, []int{}},
		{"single slice", [][]int{{1, 2, 3}}, []int{1, 2, 3}},
		{"two slices", [][]int{{1, 2, 3}, {2, 3, 4}}, []int{2, 3}},
		{"three slices", [][]int{{1, 2, 3}, {2, 3, 4}, {2, 3, 5}}, []int{2, 3}},
		{"no intersection", [][]int{{1, 2}, {3, 4}}, []int{}},
		{"duplicates no intersection", [][]int{{1, 1, 2}, {1, 2, 2}}, []int{1, 2}},
		{"duplicates with intersection", [][]int{{1, 2, 2, 3, 4}, {6, 5, 2, 2, 3, 7, 8}}, []int{2, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Intersect(tt.slices...)
			slices.Sort(result)
			slices.Sort(tt.expected)
			if !slices.Equal(result, tt.expected) {
				t.Errorf("Intersect(%v) returned %v; want %v", tt.slices, result, tt.expected)
			}
		})
	}
}

func TestMap(t *testing.T) {
	double := func(x int) int { return x * 2 }

	tests := []struct {
		name     string
		slice    []int
		expected []int
	}{
		{"empty", []int{}, []int{}},
		{"single", []int{5}, []int{10}},
		{"multiple", []int{1, 2, 3}, []int{2, 4, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Map(tt.slice, double)
			if !slices.Equal(result, tt.expected) {
				t.Errorf("Map(%v, double) returned %v; want %v", tt.slice, result, tt.expected)
			}
		})
	}
}

func TestMap2(t *testing.T) {
	multiplyByIndex := func(x, idx int) int { return x * idx }

	tests := []struct {
		name     string
		slice    []int
		expected []int
	}{
		{"empty", []int{}, []int{}},
		{"single", []int{5}, []int{0}},
		{"multiple", []int{10, 20, 30}, []int{0, 20, 60}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Map2(tt.slice, multiplyByIndex)
			if !slices.Equal(result, tt.expected) {
				t.Errorf("Map2(%v, multiplyByIndex) returned %v; want %v", tt.slice, result, tt.expected)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	sum := func(acc, x int) int { return acc + x }

	tests := []struct {
		name     string
		slice    []int
		init     int
		expected int
	}{
		{"empty", []int{}, 0, 0},
		{"single", []int{5}, 0, 5},
		{"multiple", []int{1, 2, 3, 4}, 0, 10},
		{"empty with init", []int{}, 5, 5},
		{"single with init", []int{5}, 5, 10},
		{"multiple with init", []int{1, 2, 3, 4}, 5, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reduce(tt.slice, tt.init, sum)
			if result != tt.expected {
				t.Errorf("Reduce(%v, %v, sum) = %v; want %v", tt.slice, tt.init, result, tt.expected)
			}
		})
	}
}

func TestReduce2(t *testing.T) {
	sumWithIndex := func(acc, x, idx int) int { return acc + x + idx }

	tests := []struct {
		name     string
		slice    []int
		init     int
		expected int
	}{
		{"empty", []int{}, 0, 0},
		{"single", []int{5}, 0, 5},
		{"multiple", []int{1, 2, 3, 4}, 0, 16},
		{"empty with init", []int{}, 5, 5},
		{"single with init", []int{5}, 5, 10},
		{"multiple with init", []int{1, 2, 3, 4}, 5, 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reduce2(tt.slice, tt.init, sumWithIndex)
			if result != tt.expected {
				t.Errorf("Reduce2(%v, %v, sumWithIndex) = %v; want %v", tt.slice, tt.init, result, tt.expected)
			}
		})
	}
}
