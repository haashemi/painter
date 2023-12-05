package painter

import (
	"image"
	"image/color"
	"math"
)

// LinearGradient is a Renderer which draws a linear gradient on the image
// from startColor to endColor with in the passed angle in degrees.
func LinearGradient(angle float64, startColor, endColor color.RGBA) Renderer {
	angle = Radians(angle)

	return func(img image.Image, point image.Point) color.RGBA {
		u := float64(point.X)/float64(img.Bounds().Dx()) - 0.5
		v := float64(point.Y)/float64(img.Bounds().Dy()) - 0.5

		atan := math.Atan2(v, u)
		length := math.Sqrt(u*u + v*v)
		rotation := math.Cos(angle+atan)*length + 0.5

		fragColor := MixRGBA(SmoothStep(0, 1, rotation), startColor, endColor)

		return fragColor
	}
}

// RadialGradient is a Renderer which draws a radial gradient on the image
// from startColor to endColor. startColor starts from provided center point.
func RadialGradient(startColor, endColor color.RGBA, center image.Point) Renderer {
	getDistance := func(c, p image.Point) float64 {
		return math.Sqrt(float64((p.X-c.X)*(p.X-c.X)) + float64((p.Y-c.Y)*(p.Y-c.Y)))
	}

	return func(img image.Image, point image.Point) color.RGBA {
		longestDistance := max(
			getDistance(center, image.Point{0, 0}),
			getDistance(center, image.Point{img.Bounds().Dx(), 0}),
			getDistance(center, image.Point{0, img.Bounds().Dy()}),
			getDistance(center, image.Point{img.Bounds().Dx(), img.Bounds().Dy()}),
		)

		a := SmoothStep(0, 1, getDistance(center, point)/longestDistance)

		return MixRGBA(a, startColor, endColor)
	}
}
