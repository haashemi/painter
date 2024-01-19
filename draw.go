package painter

import (
	"image"
	"image/color"
	"sync"

	"image/draw"
)

// PixelProcessor returns a [image/color.NRGBA] for the specified image coordinates.
// It will be called concurrently.
type PixelProcessor func(x, y int) color.NRGBA

// Draw is a low-level method which loops through each pixel in rect's range
// concurrently and sets the merged & processed color on them.
func Draw(img *image.NRGBA, rect image.Rectangle, process PixelProcessor) {
	workers := Workers

	var wg sync.WaitGroup
	var queue = make(chan int, workers)

	rowGenerator := func() {
		for y := range queue {
			for x := rect.Min.X; x < rect.Max.X; x++ {
				c := MergeColors(img.NRGBAAt(x, y), process(x, y))
				img.SetNRGBA(x, y, c)
			}

			wg.Done()
		}
	}

	for i := 0; i < workers; i++ {
		go rowGenerator()
	}

	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		wg.Add(1)
		queue <- y
	}

	wg.Wait()
	close(queue)
}

func Paste(dst *image.NRGBA, src image.Image, x, y int) {
	draw.Draw(dst, src.Bounds().Add(image.Point{x, y}), src, image.Point{}, draw.Over)
}
