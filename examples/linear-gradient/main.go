package main

import (
	"image/color"

	"github.com/haashemi/painter"
)

var (
	Width, Height  = 1920, 1080
	PrimaryColor   = color.NRGBA{28, 181, 224, 255}
	SecondaryColor = color.NRGBA{0, 0, 70, 255}
)

func main() {
	img := painter.New(Width, Height)

	painter.Render(img, painter.LinearGradient(45, PrimaryColor, SecondaryColor))

	err := painter.SavePNG(img, "examples/linear-gradient/linear-gradient.png")
	if err != nil {
		panic(err)
	}
}
