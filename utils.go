package painter

import (
	"math"
)

// clamp returns the value of x constrained to the range lowerLimit to upperLimit.
func clamp(x, lowerLimit, upperLimit float64) float64 {
	if x < lowerLimit {
		return lowerLimit
	}
	if x > upperLimit {
		return upperLimit
	}
	return x
}

// smoothStep performs smooth Hermite interpolation between 0 and 1 when edge0 < x < edge1.
// This is useful in cases where a threshold function with a smooth transition is desired.
//
// Learn more: https://en.wikipedia.org/wiki/Smoothstep
func smoothStep(edge0, edge1, x float64) float64 {
	x = clamp((x-edge0)/(edge1-edge0), 0, 1)
	return x * x * (3.0 - 2.0*x)
}

// getDistance returns the distance between two coordinates using the distance formula.
func getDistance(x1, y1, x2, y2 int) float64 {
	x, y := x2-x1, y2-y1
	return math.Sqrt(float64(x*x + y*y))
}
