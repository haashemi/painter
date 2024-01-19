package painter

import (
	"image"
	"image/color"
	"math"
)

// DrawColor replaces each pixel's color with the specified color c.
// It fully takes care of transparent colors.
func DrawColor(img *image.NRGBA, rect image.Rectangle, c color.NRGBA) {
	Draw(img, rect, func(_, _ int) color.NRGBA { return c })
}

// DrawRadialGradient draws a radial gradient on the image from startColor to endColor.
// startColor starts from provided center point. It fully takes care of transparent colors.
func DrawRadialGradient(img *image.NRGBA, rect image.Rectangle, startColor, endColor color.NRGBA, center image.Point) {
	longestDistance := max(
		getDistance(center.X, center.Y, rect.Min.X, rect.Min.Y),
		getDistance(center.X, center.Y, rect.Min.X, rect.Max.Y),
		getDistance(center.X, center.Y, rect.Max.X, rect.Min.Y),
		getDistance(center.X, center.Y, rect.Max.X, rect.Max.Y),
	)

	Draw(img, rect, func(x, y int) color.NRGBA {
		a := smoothStep(0, 1, getDistance(center.X, center.Y, x, y)/longestDistance)
		return MixNRGBA(a, startColor, endColor)
	})
}

// LinearGradient draws a linear gradient on the image from startColor to endColor
// with in the passed angle in radians. It fully takes care of transparent colors.
func DrawLinearGradient(img *image.NRGBA, rect image.Rectangle, angle float64, startColor, endColor color.NRGBA) {
	width := float64(rect.Max.X - rect.Min.X)
	height := float64(rect.Max.Y - rect.Min.Y)

	Draw(img, rect, func(x, y int) color.NRGBA {
		u := float64(x-rect.Min.X)/width - 0.5
		v := float64(y-rect.Min.Y)/height - 0.5

		atan := math.Atan2(v, u)
		length := math.Sqrt(u*u + v*v)
		rotation := math.Cos(angle+atan)*length + 0.5

		fragColor := MixNRGBA(smoothStep(0, 1, rotation), startColor, endColor)
		return fragColor
	})
}
