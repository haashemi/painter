package painter

import "image/color"

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

// MergeNRGBA merges two [image/color.NRGBA] colors.
func MergeColors(background, foreground color.NRGBA) color.NRGBA {
	if foreground.A == 255 || background.A == 0 {
		return foreground
	}

	oA := float64(foreground.A)
	return color.NRGBA{
		R: uint8((float64(background.R)*(255-oA) + float64(foreground.R)*oA) / 255),
		G: uint8((float64(background.G)*(255-oA) + float64(foreground.G)*oA) / 255),
		B: uint8((float64(background.B)*(255-oA) + float64(foreground.B)*oA) / 255),
		A: uint8(min(float64(background.A)+float64(foreground.A), 255)),
	}
}
