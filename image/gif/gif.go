package gif

import (
	"image"
	"image/gif"
	"io"
)

type GIFImageProcess struct {
}

func New() {
	return GIFImageProcess{}
}

func (p GIFImageProcess) Decode(r io.Reader) (image.Image, error) {
	return gif.Decode(r)
}

func (p GIFImageProcess) Encode(w io.Writer, i image.Image) error {
	return gif.Encode(w, i, nil)
}
