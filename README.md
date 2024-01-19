# 🎨 Painter (WIP)

A simple utility Golang package to help you with drawing on images

> This package is currently work-in-progress and subject to change at any time.

## Installation

```bash
$ go get github.com/haashemi/painter@main
```

## Example

There are some examples in the [examples folder](/examples/), but here is one of them.

```go
package main

import (
	"image/color"
	"math"

	"github.com/haashemi/painter"
)

func main() {
	img := painter.New(width, height)

	black := color.NRGBA{0, 0, 0, 255}
	white := color.NRGBA{255, 255, 255, 255}

	painter.DrawLinearGradient(img, img.Rect, math.Pi/4, black, white)

	painter.SavePNG(img, "painter.png")
}
```

## Contributions

All type of contributions are highly appreciated. <3
