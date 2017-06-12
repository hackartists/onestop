package png

import (
	"image"
	"image/png"
	"io"
)

type PNGImageProcess struct {
}

func New() {
	return PNGImageProcess{}
}

func (p PNGImageProcess) Decode(r io.Reader) (image.Image, error) {
	return png.Decode(r)
}

func (p PNGImageProcess) Encode(w io.Writer, i image.Image) error {
	return png.Encode(w, i)
}
