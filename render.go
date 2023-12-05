package painter

import (
	"image"
	"image/color"
	"runtime"
	"sync"
)

// Renderer is a function which gets the [image.Image] and the current [image.Point],
// and returns a [image/color.RGBA] for that point.
type Renderer func(img image.Image, point image.Point) color.RGBA

// Render calls the renderer in each pixel of the image
// in NumCPU*2 goroutines for better performance.
func Render(img *image.RGBA, renderer Renderer) {
	var wg sync.WaitGroup
	var queue = make(chan int, runtime.NumCPU()*2)

	rowGenerator := func() {
		for y := range queue {
			for x := 0; x < img.Rect.Dx(); x++ {
				img.SetRGBA(x, y, renderer(img, image.Point{X: x, Y: y}))
			}

			wg.Done()
		}
	}

	for i := 0; i < runtime.NumCPU()*2; i++ {
		go rowGenerator()
	}

	for y := 0; y < img.Bounds().Dy(); y++ {
		wg.Add(1)
		queue <- y
	}

	wg.Wait()
	close(queue)
}
