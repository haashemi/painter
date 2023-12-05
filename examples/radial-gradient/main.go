package main

import (
	"image"
	"image/color"

	"github.com/haashemi/painter"
)

var (
	Width, Height  = 1920, 1080
	PrimaryColor   = color.RGBA{28, 181, 224, 255}
	SecondaryColor = color.RGBA{0, 0, 70, 255}
)

func main() {
	img := painter.New(Width, Height)

	painter.Render(img, painter.RadialGradient(PrimaryColor, SecondaryColor, image.Point{1000, 350}))

	err := painter.SavePNG(img, "examples/radial-gradient/radial-gradient.png")
	if err != nil {
		panic(err)
	}
}
