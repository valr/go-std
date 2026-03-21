// Package slicesx provides utility functions for slice manipulation.
package slicesx

import (
	"log"
)

// Count the occurrences of an element in a slice.
func Count[S ~[]T, T comparable](s S, x T) int {
	count := 0
	for i := range s {
		if s[i] == x {
			count++
		}
	}
	return count
}

// Create a slice of size count with all elements initialized with value of given element x.
func Create[T any](count int, x T) []T {
	if count < 0 {
		log.Panicf("error when creating slice in slicesx.Create: %v", count)
	}
	slice := make([]T, count)
	if count > 0 {
		slice[0] = x
		for i := 1; i < count; i *= 2 {
			copy(slice[i:], slice[:i])
		}
	}
	return slice
}

// Filter returns a slice of all elements in the slice s that satisfy the predicate function f.
func Filter[S ~[]T, T any](s S, f func(x T) bool) S {
	filtered := make(S, len(s))
	count := 0
	for i := range s {
		if f(s[i]) {
			filtered[count] = s[i]
			count++
		}
	}
	return filtered[:count]
}

// Filter2 returns a slice of all elements in the slice s that satisfy the predicate function f.
func Filter2[S ~[]T, T any](s S, f func(x T, idx int) bool) S {
	filtered := make(S, len(s))
	count := 0
	for i := range s {
		if f(s[i], i) {
			filtered[count] = s[i]
			count++
		}
	}
	return filtered[:count]
}

// Flatten a nested slice into a single slice.
func Flatten[T any](s [][]T) []T {
	length := 0
	for i := range s {
		length += len(s[i])
	}
	flattened := make([]T, 0, length)
	for i := range s {
		flattened = append(flattened, s[i]...)
	}
	return flattened
}

// Index returns the index of the first occurrence of sub in s, or -1 if not present.
func Index[S ~[]T, T comparable](s S, sub S) int {
next:
	for i := range len(s) - len(sub) + 1 {
		for j := range sub {
			if s[i+j] != sub[j] {
				continue next
			}
		}
		return i
	}
	return -1
}

// Intersect returns a slice of all elements that are present in all the given slices.
// The order of the elements in the returned slice is not guaranteed to be the same as in any of the input slices.
func Intersect[T comparable](slices ...[]T) []T {
	intersection := make([]T, 0)
	if len(slices) == 0 {
		return intersection
	}
	frequency := make(map[T]int)
	for i := range slices[0] {
		frequency[slices[0][i]]++
	}
	for j := 1; j < len(slices); j++ {
		tempFrequency := make(map[T]int)
		for i := range slices[j] {
			tempFrequency[slices[j][i]]++
		}
		for e, freq1 := range frequency {
			if freq2, ok := tempFrequency[e]; ok {
				if freq2 < freq1 {
					frequency[e] = freq2
				}
			} else {
				delete(frequency, e)
			}
		}
	}
	for e, freq := range frequency {
		for range freq {
			intersection = append(intersection, e)
		}
	}
	return intersection
}

// Map returns a slice by applying function f to each element in the slice s.
func Map[S1 ~[]T1, S2 []T2, T1, T2 any](s S1, f func(x T1) T2) S2 {
	mapped := make(S2, len(s))
	for i := range s {
		mapped[i] = f(s[i])
	}
	return mapped
}

// Map2 returns a slice by applying function f to each element in the slice s.
func Map2[S1 ~[]T1, S2 []T2, T1, T2 any](s S1, f func(x T1, idx int) T2) S2 {
	mapped := make(S2, len(s))
	for i := range s {
		mapped[i] = f(s[i], i)
	}
	return mapped
}

// Reduce returns the result of applying a binary function f cumulatively to the elements of the slice s.
func Reduce[S ~[]T1, T1, T2 any](s S, init T2, f func(acc T2, x T1) T2) T2 {
	acc := init
	for i := range s {
		acc = f(acc, s[i])
	}
	return acc
}

// Reduce2 returns the result of applying a binary function f cumulatively to the elements of the slice s.
func Reduce2[S ~[]T1, T1, T2 any](s S, init T2, f func(acc T2, x T1, idx int) T2) T2 {
	acc := init
	for i := range s {
		acc = f(acc, s[i], i)
	}
	return acc
}
