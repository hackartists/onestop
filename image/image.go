package image

import (
	"bytes"
	"errors"
	"image"
	"io"

	"github.com/nfnt/resize"
	"github.com/pwnartist/onestop/image/gif"
	"github.com/pwnartist/onestop/image/jpeg"
	"github.com/pwnartist/onestop/image/png"
)

type OnestopImageProcessor interface {
	Decode(io.Reader) (image.Image, error)
	Encode(io.Writer, image.Image) error
}

var ImageProcs = map[string]OnestopImageProcessor{
	"jpeg": jpeg.New(),
	"png":  png.New(),
	"gif":  gif.New(),
}

func ResizeImage(r io.Reader, ext string, x uint, y uint) (*bytes.Buffer, error) {
	if imgProc, ok := ImageProcs[ext]; ok {
		img, err := imgProc.Decode(r)

		if err != nil {
			return nil, err
		}
		buf := new(bytes.Buffer)
		m := resize.Resize(x, y, img, resize.Lanczos3)

		imgProc.Encode(buf, m)

		return buf, nil
	}

	return nil, errors.New("Unknown image type: use jpeg, png and gif")
}
