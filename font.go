package painter

import (
	"io"
	"sync"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
	"golang.org/x/image/math/fixed"
)

// Font holds an embedded sfnt.Font and a buffer for it.
//
// Font is parallel safe.
type Font struct {
	*sfnt.Font

	mut sync.Mutex
	buf *sfnt.Buffer
}

// ParseFont parses an opentype font and returns a Font of the parsed font.
func ParseFont(src []byte) (*Font, error) {
	f, err := sfnt.Parse(src)
	if err != nil {
		return nil, err
	}

	return &Font{Font: f, buf: &sfnt.Buffer{}}, nil
}

// NewFace creates a new opentype font face at Size of size, DPI of 72,
// and Hinting of HintingFull, and returns it as a FontFace
func (f *Font) NewFace(size float64) (face *FontFace) {
	// it actually never returns any errors. (checked in: x/image v0.15.0)
	ff, _ := opentype.NewFace(f.Font, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	return &FontFace{Face: ff}
}

// Bounds calls sfnt.Font's Bounds with a font buffer.
func (f *Font) Bounds(ppem fixed.Int26_6, h font.Hinting) (fixed.Rectangle26_6, error) {
	f.mut.Lock()
	defer f.mut.Unlock()

	return f.Font.Bounds(f.buf, ppem, h)
}

// GlyphAdvance calls sfnt.Font's GlyphAdvance with a font buffer.
func (f *Font) GlyphAdvance(x sfnt.GlyphIndex, ppem fixed.Int26_6, h font.Hinting) (fixed.Int26_6, error) {
	f.mut.Lock()
	defer f.mut.Unlock()

	return f.Font.GlyphAdvance(f.buf, x, ppem, h)
}

// GlyphBounds calls sfnt.Font's GlyphBounds with a font buffer.
func (f *Font) GlyphBounds(x sfnt.GlyphIndex, ppem fixed.Int26_6, h font.Hinting) (bounds fixed.Rectangle26_6, advance fixed.Int26_6, err error) {
	f.mut.Lock()
	defer f.mut.Unlock()

	return f.Font.GlyphBounds(f.buf, x, ppem, h)
}

// GlyphIndex calls sfnt.Font's GlyphIndex with a font buffer.
func (f *Font) GlyphIndex(r rune) (sfnt.GlyphIndex, error) {
	f.mut.Lock()
	defer f.mut.Unlock()

	return f.Font.GlyphIndex(f.buf, r)
}

// GlyphName calls sfnt.Font's GlyphName with a font buffer.
func (f *Font) GlyphName(x sfnt.GlyphIndex) (string, error) {
	f.mut.Lock()
	defer f.mut.Unlock()

	return f.Font.GlyphName(f.buf, x)
}

// Kern calls sfnt.Font's Kern with a font buffer.
func (f *Font) Kern(x0 sfnt.GlyphIndex, x1 sfnt.GlyphIndex, ppem fixed.Int26_6, h font.Hinting) (fixed.Int26_6, error) {
	f.mut.Lock()
	defer f.mut.Unlock()

	return f.Font.Kern(f.buf, x0, x1, ppem, h)
}

// LoadGlyph calls sfnt.Font's LoadGlyph with a font buffer.
func (f *Font) LoadGlyph(x sfnt.GlyphIndex, ppem fixed.Int26_6, opts *sfnt.LoadGlyphOptions) (sfnt.Segments, error) {
	f.mut.Lock()
	defer f.mut.Unlock()

	return f.Font.LoadGlyph(f.buf, x, ppem, opts)
}

// Metrics calls sfnt.Font's Metrics with a font buffer.
func (f *Font) Metrics(ppem fixed.Int26_6, h font.Hinting) (font.Metrics, error) {
	f.mut.Lock()
	defer f.mut.Unlock()

	return f.Font.Metrics(f.buf, ppem, h)
}

// Name calls sfnt.Font's Name with a font buffer.
func (f *Font) Name(id sfnt.NameID) (string, error) {
	f.mut.Lock()
	defer f.mut.Unlock()

	return f.Font.Name(f.buf, id)
}

// WriteSourceTo calls sfnt.Font's WriteSourceTo with a font buffer.
func (f *Font) WriteSourceTo(w io.Writer) (int64, error) {
	f.mut.Lock()
	defer f.mut.Unlock()

	return f.Font.WriteSourceTo(f.buf, w)
}
