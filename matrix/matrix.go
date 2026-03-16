package matrix

import (
	"github.com/valr/go-std/stringsx"
)

// Return the rotated string matrix.
func RotateStrMatrix(matrix []string) []string {
	var rotated []string
	for i := range len(matrix[0]) {
		var row []rune
		for j := range matrix {
			row = append(row, []rune(matrix[j])[i])
		}
		rotated = append(rotated, string(row))
	}
	return rotated
}

// Return the string matrix read diagonally.
func DiagonalStrMatrix(matrix []string) []string {
	var diagonals []string
	rows, cols := len(matrix), len(matrix[0])
	for i := range rows + cols - 1 {
		var diag []rune
		for j := max(0, i-cols+1); j < min(rows, i+1); j++ {
			diag = append([]rune{[]rune(matrix[j])[i-j]}, diag...)
		}
		diagonals = append(diagonals, string(diag))
	}
	return diagonals
}

// Return the string matrix read counter diagonally.
func CounterDiagonalStrMatrix(matrix []string) []string {
	var counterDiagonals []string
	for i := range matrix {
		counterDiagonals = append(counterDiagonals, stringsx.Reverse(matrix[i]))
	}
	return DiagonalStrMatrix(counterDiagonals)
}
