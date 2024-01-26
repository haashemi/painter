package painter

import (
	"fmt"
	"image"
	"image/color"
	"math"
)

var (
	color1 = color.NRGBA{67, 206, 162, 255}
	color2 = color.NRGBA{24, 91, 157, 255}
	color3 = color.NRGBA{19, 48, 145, 255}
)

func ExampleDrawColor() {
	img := New(1920, 1080)

	DrawColor(img, img.Rect, color1)

	err := SavePNG(img, "tests/Example-DrawColor.png")
	fmt.Println(err)

	// Output: <nil>
}

func ExampleDrawColor_second() {
	img := New(1920, 1080)

	DrawColor(img, image.Rect(0, 0, 1000, 500), color1)
	DrawColor(img, image.Rect(0, 500, 1000, 1080), color2)
	DrawColor(img, image.Rect(1000, 0, 1920, 1080), color3)

	err := SavePNG(img, "tests/Example-DrawColor_second.png")
	fmt.Println(err)

	// Output: <nil>
}

func ExampleDrawRadialGradient() {
	img := New(1920, 1080)

	DrawRadialGradient(img, img.Rect, color1, color2, image.Point{1000, 350})

	err := SavePNG(img, "tests/Example-DrawRadialGradient.png")
	fmt.Println(err)

	// Output: <nil>
}

func ExampleDrawRadialGradient_second() {
	img := New(1920, 1080)

	DrawRadialGradient(img, img.Rect, color1, color2, image.Point{1000, 350})
	DrawRadialGradient(img, img.Rect, color3, color.NRGBA{}, image.Point{1700, 950})

	err := SavePNG(img, "tests/Example-DrawRadialGradient_second.png")
	fmt.Println(err)

	// Output: <nil>
}

func ExampleDrawLinearGradient() {
	img := New(1920, 1080)

	DrawLinearGradient(img, img.Rect, math.Pi/4, color1, color2)

	err := SavePNG(img, "tests/Example-DrawLinearGradient.png")
	fmt.Println(err)

	// Output: <nil>
}

func ExampleDrawLinearGradient_second() {
	img := New(1920, 1080)

	DrawLinearGradient(img, img.Rect, math.Pi/4, color1, color2)
	DrawLinearGradient(img, img.Rect, math.Pi*3/4, color3, color.NRGBA{})

	err := SavePNG(img, "tests/Example-DrawLinearGradient_second.png")
	fmt.Println(err)

	// Output: <nil>
}
