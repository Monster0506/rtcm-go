package core

import "math"

func TwoPow(n int) float64 {
	return math.Ldexp(1, n)
}
