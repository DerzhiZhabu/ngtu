package trian

import (
	"math"
)

func verification(a, b, c float64) bool {
	if a <= 0 || b <= 0 || c <= 0 || a+b <= c || a+c <= b || b+c <= a {
		return false
	}
	return true
}

func perimeter(a, b, c float64) float64 {
	p := a + b + c
	return p
}

func area(a, b, c, p float64) float64 {
	p = p / 2
	s := math.Sqrt(p * (p - a) * (p - b) * (p - c))
	return s
}

func equilateral(a, b, c float64) bool {
	if a == b || b == c || c == a {
		return true
	}
	return false
}
