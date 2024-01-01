package main

import (
	"image"
	"image/color"

	"github.com/haashemi/painter"
)

var (
	Width, Height  = 1920, 1080
	PrimaryColor   = color.NRGBA{28, 181, 224, 255}
	SecondaryColor = color.NRGBA{0, 0, 70, 255}
	TertiaryColor  = color.NRGBA{48, 43, 99, 255}
)

func main() {
	img := painter.New(Width, Height)

	painter.Render(img, painter.RadialGradient(PrimaryColor, SecondaryColor, image.Point{1000, 350}))

	painter.Render(img, painter.RadialGradient(TertiaryColor, color.NRGBA{}, image.Point{1700, 950}))

	err := painter.SavePNG(img, "examples/radial-gradient/radial-gradient.png")
	if err != nil {
		panic(err)
	}
}
