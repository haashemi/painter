// Painter is a simple [image.NRGBA] utility packages which helps you with drawing on images.
package painter

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"runtime"

	_ "golang.org/x/image/webp"
)

// Workers defines the goroutines count used for some methods like Draw.
var Workers int = runtime.NumCPU()

// New returns a new [image.NRGBA] in the specified width and height.
func New(width, height int) *image.NRGBA {
	return image.NewNRGBA(image.Rect(0, 0, width, height))
}

// Decode calls the [image.Decode] and returns the image.
//
// It's useful when you're not sure about the image type and also want to be safe
// from it. supports jpeg, png, and webp.
func Decode(r io.Reader) (image.Image, error) {
	img, _, err := image.Decode(r)
	return img, err
}

// SavePNG creates a file in the specified path and writes the encoded png to it.
func SavePNG(img image.Image, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	return png.Encode(f, img)
}

// SaveJPEG creates a file in the specified path and writes the encoded jpeg to it.
func SaveJPEG(img image.Image, path string, o *jpeg.Options) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	return jpeg.Encode(f, img, o)
}
