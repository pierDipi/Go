package math

import "math"

func MaxInt(a, b int) int  {
	return int(math.Max(float64(a), float64(b)))
}

func MinInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
