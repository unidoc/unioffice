package image

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/unidoc/unioffice/common"
	"github.com/unidoc/unioffice/internal/storage"
)

// ImageFromStorage reads an image from an internal Unioffice storage.
func ImageFromStorage(path string) (common.Image, error) {
	r := common.Image{}
	f, err := storage.Open(path)
	if err != nil {
		return r, fmt.Errorf("error reading image: %s", err)
	}
	imgDec, ifmt, err := image.Decode(f)
	if err != nil {
		return r, fmt.Errorf("unable to parse image: %s", err)
	}

	r.Path = path
	r.Format = ifmt
	r.Size = imgDec.Bounds().Size()
	return r, nil
}
