package zipping

import (
	"archive/zip"
	"strings"

	"github.com/mec07/unioffice"
	"github.com/mec07/unioffice/common"
	"github.com/mec07/unioffice/zippkg"
)

//AddImage adds an image (either from bytes or from disk) and adds it to the zip file.
func AddImage(z *zip.Writer, img common.ImageRef, imageNum int, dt unioffice.DocType) error {
	filename := unioffice.AbsoluteImageFilename(dt, imageNum, strings.ToLower(img.Format()))
	if img.Data() != nil && len(*img.Data()) > 0 {
		if err := zippkg.AddFileFromBytes(z, filename, *img.Data()); err != nil {
			return err
		}
	} else if img.Path() != "" {
		if err := zippkg.AddFileFromDisk(z, filename, img.Path()); err != nil {
			return err
		}
	} else {
		unioffice.Log("unsupported image source: %+v", img)
	}
	return nil
}
