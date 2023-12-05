// Painter is a simple [image.RGBA] utility packages which helps you with drawing on images.
package painter

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

// New returns a new [image.RGBA] in the specified width and height.
func New(width, height int) *image.RGBA {
	return image.NewRGBA(image.Rect(0, 0, width, height))
}

// SavePNG creates a file in the specified path and encodes the image as png to it.
func SavePNG(img image.Image, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	return png.Encode(f, img)
}

// SaveJPEG creates a file in the specified path and encodes the image as jpeg to it.
func SaveJPEG(img image.Image, path string, o *jpeg.Options) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	return jpeg.Encode(f, img, o)
}
