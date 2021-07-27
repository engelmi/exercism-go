package grains

import (
	"errors"
	"math"
)

func square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("invalid input")
	}
	return uint64(math.Pow(2, float64(n-1))), nil
}

func Square(n int) (uint64, error) {
	return square(n)
}

func Total() uint64 {
	max, _ := square(64)
	return (max * 2) - 1
}
