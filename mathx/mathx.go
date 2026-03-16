package mathx

type Number interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int8 | ~int16 | ~int32 | ~int64 |
		~int | ~float32 | ~float64
}

// Compute the absolute value of a number.
func Abs[T Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// Return the product of all numbers.
func Product[T Number](x ...T) T {
	var product T
	if len(x) > 0 {
		product = x[0]
		for i := 1; i < len(x); i++ {
			product *= x[i]
		}
	}
	return product
}

// Return the sum of all numbers.
func Sum[T Number](x ...T) T {
	var sum T
	for i := range x {
		sum += x[i]
	}
	return sum
}
