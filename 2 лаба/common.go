package main

import "math"

const eps = 1e-9

// Сравнение чисел с учётом погрешности float64
func equal(a, b float64) bool {
	return math.Abs(a-b) < eps
}
