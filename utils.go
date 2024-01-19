package painter

import (
	"math"
)

// clamp returns the value of x constrained to the range minVal to maxVal.
//
// more: https://thebookofshaders.com/glossary/?search=clamp
func clamp(x, minVal, maxVal float64) float64 {
	return min(max(x, minVal), maxVal)
}

// smoothStep performs smooth Hermite interpolation between 0 and 1 when edge0 < x < edge1.
// This is useful in cases where a threshold function with a smooth transition is desired.
//
// more: https://thebookofshaders.com/glossary/?search=smoothstep
func smoothStep(edge0, edge1, x float64) float64 {
	t := clamp((x-edge0)/(edge1-edge0), 0.0, 1.0)
	return t * t * (3.0 - 2.0*t)
}

// getDistance returns the distance between two coordinates using the distance formula.
func getDistance(x1, y1, x2, y2 int) float64 {
	x, y := x2-x1, y2-y1
	return math.Sqrt(float64(x*x + y*y))
}
