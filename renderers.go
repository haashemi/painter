package painter

import (
	"image"
	"image/color"
	"math"
)

// LinearGradient is a Renderer which draws a linear gradient on the image from
// startColor to endColor with in the passed angle in degrees.
func LinearGradient(angle float64, startColor, endColor color.RGBA) Renderer {
	angle = Radians(angle)

	return func(rect image.Rectangle, point image.Point) color.RGBA {
		u := float64(point.X-rect.Min.X)/float64(rect.Max.X-rect.Min.X) - 0.5
		v := float64(point.Y-rect.Min.Y)/float64(rect.Max.Y-rect.Min.Y) - 0.5

		atan := math.Atan2(v, u)
		length := math.Sqrt(u*u + v*v)
		rotation := math.Cos(angle+atan)*length + 0.5

		fragColor := MixRGBA(SmoothStep(0, 1, rotation), startColor, endColor)

		return fragColor
	}
}

// RadialGradient is a Renderer which draws a radial gradient on the image from
// startColor to endColor. startColor starts from provided center point.
func RadialGradient(startColor, endColor color.RGBA, center image.Point) Renderer {
	getDistance := func(c, p image.Point) float64 {
		return math.Sqrt(float64((p.X-c.X)*(p.X-c.X)) + float64((p.Y-c.Y)*(p.Y-c.Y)))
	}

	var longestDistance float64

	return func(rect image.Rectangle, point image.Point) color.RGBA {
		// calculate only if it's not calculated, yet. (i.e. calculate once)
		if longestDistance == 0 {
			longestDistance = max(
				getDistance(center, image.Point{rect.Min.X, rect.Min.Y}),
				getDistance(center, image.Point{rect.Min.X, rect.Max.Y}),
				getDistance(center, image.Point{rect.Max.X, rect.Min.Y}),
				getDistance(center, image.Point{rect.Max.X, rect.Max.Y}),
			)
		}

		a := SmoothStep(0, 1, getDistance(center, point)/longestDistance)
		return MixRGBA(a, startColor, endColor)
	}
}

// FillColor is a Renderer which draws the specified color on every pixels.
func FillColor(c color.RGBA) Renderer {
	return func(rect image.Rectangle, point image.Point) color.RGBA {
		return c
	}
}
