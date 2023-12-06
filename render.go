package painter

import (
	"image"
	"image/color"
	"runtime"
	"sync"
)

// Renderer is a function which gets the [image.Image] and the current [image.Point],
// and returns a [image/color.RGBA] for that point.
type Renderer func(rect image.Rectangle, point image.Point) color.RGBA

// Render calls the renderer in each pixel of the image
// in NumCPU*2 goroutines for better performance.
func Render(img *image.RGBA, renderer Renderer) {
	RenderRect(img, img.Rect, renderer)
}

func RenderRect(img *image.RGBA, rect image.Rectangle, renderer Renderer) {
	var wg sync.WaitGroup
	var queue = make(chan int, runtime.NumCPU()*2)

	rowGenerator := func() {
		for y := range queue {

			for x := rect.Min.X; x < rect.Max.X; x++ {
				rc := renderer(rect, image.Point{X: x, Y: y})

				img.SetRGBA(x, y, MergeRGBA(img.RGBAAt(x, y), rc))
			}

			wg.Done()
		}
	}

	for i := 0; i < runtime.NumCPU()*2; i++ {
		go rowGenerator()
	}

	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		wg.Add(1)
		queue <- y
	}

	wg.Wait()
	close(queue)
}
