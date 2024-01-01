package painter

import (
	"image/color"
	"math"
)

// Clamp returns the value of x constrained to the range minVal to maxVal.
//
// more: https://thebookofshaders.com/glossary/?search=clamp
func Clamp(x, minVal, maxVal float64) float64 {
	return min(max(x, minVal), maxVal)
}

// SmoothStep performs smooth Hermite interpolation between 0 and 1 when edge0 < x < edge1.
// This is useful in cases where a threshold function with a smooth transition is desired.
//
// more: https://thebookofshaders.com/glossary/?search=smoothstep
func SmoothStep(edge0, edge1, x float64) float64 {
	t := Clamp((x-edge0)/(edge1-edge0), 0.0, 1.0)
	return t * t * (3.0 - 2.0*t)
}

// Radians converts the degrees to radians.
func Radians(angle float64) float64 {
	return angle * math.Pi / 180
}

// MixNRGBA merges two colors with the provided alpha.
//
// "a" should be between 0 and 1.
//
// TODO: Support multiple colors
func MixNRGBA(a float64, x, y color.NRGBA) color.NRGBA {
	if a == 0 {
		return x
	} else if a == 1 {
		return y
	}

	return color.NRGBA{
		R: uint8(float64(x.R)*(1-a) + float64(y.R)*a),
		G: uint8(float64(x.G)*(1-a) + float64(y.G)*a),
		B: uint8(float64(x.B)*(1-a) + float64(y.B)*a),
		A: uint8(float64(x.A)*(1-a) + float64(y.A)*a),
	}
}

// MergeNRGBA merges two [image/color.NRGBA] with taking care of their alpha value.
func MergeNRGBA(background, foreground color.NRGBA) color.NRGBA {
	oA := float64(foreground.A)
	if oA == 255 || background.A == 0 {
		return foreground
	}

	return color.NRGBA{
		R: uint8((float64(background.R)*(255-oA) + float64(foreground.R)*oA) / 255),
		G: uint8((float64(background.G)*(255-oA) + float64(foreground.G)*oA) / 255),
		B: uint8((float64(background.B)*(255-oA) + float64(foreground.B)*oA) / 255),
		A: uint8(min(float64(background.A)+float64(foreground.A), 255)),
	}
}
