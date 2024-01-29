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

// DrawConicGradient fills the rect of img with a conic gradient from startColor
// to endColor, oriented at angle in radians and centered at the center point.
func DrawConicGradient(img *image.NRGBA, rect image.Rectangle, startColor, endColor color.NRGBA, angle float64, center image.Point) {
	// width and height as float64 to avoid repetitive type castings.
	width, height := float64(rect.Dx()), float64(rect.Dy())

	// X and Y of the center point in in range of 0 to 1 as float64 to avoid repetitive type castings.
	cx, cy := float64(center.X)/width, float64(center.Y)/height

	Draw(img, rect, func(xp, yp int) color.NRGBA {
		// x and y as float64 in range of 0 to 1
		x, y := float64(xp)/width, float64(yp)/height

		// subtract x and y from the center point so we could simulate a
		// graph from -1(x,y) to +1x(x,y).
		x, y = cx-x, cy-y

		// if it has to be rotated, we should do some extra calculations on x and y.
		if angle > 0 {
			// make a copy from old x,y
			ox, oy := x, y

			// rotate the x,y with the angle
			sin, cos := math.Sincos(angle)
			x = ox*cos - oy*sin
			y = oy*cos + ox*sin
		}

		// the main formula for this type of gradient, nothing special.
		a := math.Atan2(y, x)

		// Normalization step:
		// change the range from (±Pi to +2Pi), then to (0 to Pi), then to (0 to 1)
		a = (a + math.Pi) / 2 / math.Pi

		return MixNRGBA(a, startColor, endColor)
	})
}
