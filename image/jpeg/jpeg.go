package jpeg

import (
	"image"
	"image/jpeg"
	"io"
)

type JPEGImageProcess struct {
}

func New() {
	return JPEGImageProcess{}
}

func (p JPEGImageProcess) Decode(r io.Reader) (image.Image, error) {
	return jpeg.Decode(r)
}

func (p JPEGImageProcess) Encode(w io.Writer, i image.Image) error {
	return jpeg.Encode(w, i, nil)
}
