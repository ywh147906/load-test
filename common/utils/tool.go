package utils

import (
	"math/rand"

	"github.com/ywh147906/load-test/common/utils/generic"
)

func MaxNumber[T generic.Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func MinNumber[T generic.Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func MaxInt64(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func MinInt64(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

func MaxFloat64(a, b float64) float64 {
	if a >= b {
		return a
	}
	return b
}

func MinFloat64(a, b float64) float64 {
	if a <= b {
		return a
	}
	return b
}

func Rand100Per() float64 {
	return float64(rand.Intn(101)) / 100
}
