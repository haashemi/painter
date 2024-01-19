package main

import (
	"image"
	"image/color"
	"math"

	"github.com/haashemi/painter"
)

var (
	Width, Height  = 1080, 1920
	PrimaryColor   = color.NRGBA{28, 181, 224, 255}
	SecondaryColor = color.NRGBA{0, 0, 70, 255}
	TertiaryColor  = color.NRGBA{0, 0, 0, 235}
)

func main() {
	img := painter.New(Width, Height)

	painter.DrawRadialGradient(img, img.Rect, PrimaryColor, SecondaryColor, image.Point{1000, 1500})
	painter.DrawLinearGradient(img, image.Rect(0, 1200, Width, Height), math.Pi/2, TertiaryColor, color.NRGBA{})

	err := painter.SavePNG(img, "examples/gradient-rect/gradient-rect.png")
	if err != nil {
		panic(err)
	}
}
