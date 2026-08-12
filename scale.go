package rtcm

import "math"

func twoPow(n int) float64 {
	return math.Ldexp(1, n)
}
