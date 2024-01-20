package painter

import (
	"image"
	"image/draw"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

// FontFace holds a font face and provides some helper methods on top of it.
//
// Note: Its use-case may change in the future.
type FontFace struct{ Face font.Face }

// Width returns the width of the text.
func (ff *FontFace) Width(text string) int {
	return (&font.Drawer{Face: ff.Face}).MeasureString(text).Round()
}

// Height returns the font face's height from the top of a line to its baseline.
func (ff *FontFace) Height() int {
	return ff.Face.Metrics().Ascent.Ceil()
}

// Write writes the text on dst at point of x and y with the color of src.
//
// src can be an image, but if you want a static color, use [image.Uniform].
func (ff *FontFace) Write(dst draw.Image, src image.Image, text string, x, y int) {
	drawer := &font.Drawer{Dst: dst, Src: src, Face: ff.Face, Dot: fixed.P(x, y+ff.Height())}
	drawer.DrawString(text)
}
