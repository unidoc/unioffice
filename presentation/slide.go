// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package presentation

import (
	"errors"
	"fmt"

	"github.com/unidoc/unioffice"

	"github.com/unidoc/unioffice/common"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/dml"

	"github.com/unidoc/unioffice/schema/soo/pml"
)

// Slide represents a slide of a presentation.
type Slide struct {
	sid *pml.CT_SlideIdListEntry
	x   *pml.Sld
	p   *Presentation
}

// X returns the inner wrapped XML type.
func (s Slide) X() *pml.Sld {
	return s.x
}

// PlaceHolders returns all of the content place holders within a given slide.
func (s Slide) PlaceHolders() []PlaceHolder {
	ret := []PlaceHolder{}
	for _, spChc := range s.x.CSld.SpTree.Choice {
		for _, sp := range spChc.Sp {
			if sp.NvSpPr != nil && sp.NvSpPr.NvPr != nil && sp.NvSpPr.NvPr.Ph != nil {
				ret = append(ret, PlaceHolder{sp, s.x})
			}
		}
	}
	return ret
}

// GetPlaceholder returns a placeholder given its type.  If there are multiplace
// placeholders of the same type, this method returns the first one.  You must use the
// PlaceHolders() method to access the others.
func (s Slide) GetPlaceholder(t pml.ST_PlaceholderType) (PlaceHolder, error) {
	for _, spChc := range s.x.CSld.SpTree.Choice {
		for _, sp := range spChc.Sp {
			if sp.NvSpPr != nil && sp.NvSpPr.NvPr != nil && sp.NvSpPr.NvPr.Ph != nil {
				if sp.NvSpPr.NvPr.Ph.TypeAttr == t {
					return PlaceHolder{sp, s.x}, nil
				}
			}
		}
	}
	return PlaceHolder{}, errors.New("unable to find placeholder")
}

// GetPlaceholderByIndex returns a placeholder given its index.  If there are multiplace
// placeholders of the same index, this method returns the first one.  You must use the
// PlaceHolders() method to access the others.
func (s Slide) GetPlaceholderByIndex(idx uint32) (PlaceHolder, error) {
	for _, spChc := range s.x.CSld.SpTree.Choice {
		for _, sp := range spChc.Sp {
			if sp.NvSpPr != nil && sp.NvSpPr.NvPr != nil && sp.NvSpPr.NvPr.Ph != nil {
				if (idx == 0 && sp.NvSpPr.NvPr.Ph.IdxAttr == nil) ||
					(sp.NvSpPr.NvPr.Ph.IdxAttr != nil && *sp.NvSpPr.NvPr.Ph.IdxAttr == idx) {
					return PlaceHolder{sp, s.x}, nil
				}
			}
		}
	}
	return PlaceHolder{}, errors.New("unable to find placeholder")
}

// ValidateWithPath validates the slide passing path informaton for a better
// error message.
func (s Slide) ValidateWithPath(path string) error {
	// schema checks
	if err := s.x.ValidateWithPath(path); err != nil {
		return err
	}

	// stuff we've figured out
	for _, c := range s.x.CSld.SpTree.Choice {
		for _, sp := range c.Sp {
			if sp.TxBody != nil {
				if len(sp.TxBody.P) == 0 {
					return errors.New(path + " : slide shape with a txbody must contain paragraphs")
				}
			}
		}
	}
	return nil
}

// AddTextBox adds an empty textbox to a slide.
func (s Slide) AddTextBox() TextBox {
	c := pml.NewCT_GroupShapeChoice()
	s.x.CSld.SpTree.Choice = append(s.x.CSld.SpTree.Choice, c)

	sp := pml.NewCT_Shape()
	c.Sp = append(c.Sp, sp)
	sp.SpPr = dml.NewCT_ShapeProperties()
	sp.SpPr.Xfrm = dml.NewCT_Transform2D()
	sp.SpPr.PrstGeom = dml.NewCT_PresetGeometry2D()
	sp.SpPr.PrstGeom.PrstAttr = dml.ST_ShapeTypeRect
	sp.TxBody = dml.NewCT_TextBody()
	sp.TxBody.BodyPr = dml.NewCT_TextBodyProperties()
	sp.TxBody.BodyPr.WrapAttr = dml.ST_TextWrappingTypeSquare
	sp.TxBody.BodyPr.SpAutoFit = dml.NewCT_TextShapeAutofit()

	tb := TextBox{sp}
	tb.Properties().SetWidth(3 * measurement.Inch)
	tb.Properties().SetHeight(1 * measurement.Inch)
	tb.Properties().SetPosition(0, 0)
	return tb
}

// AddImage adds an image textbox to a slide.
func (s Slide) AddImage(img common.ImageRef) Image {
	c := pml.NewCT_GroupShapeChoice()
	s.x.CSld.SpTree.Choice = append(s.x.CSld.SpTree.Choice, c)

	pic := pml.NewCT_Picture()
	c.Pic = append(c.Pic, pic)

	pic.NvPicPr.CNvPicPr = dml.NewCT_NonVisualPictureProperties()
	pic.NvPicPr.CNvPicPr.PicLocks = dml.NewCT_PictureLocking()
	pic.NvPicPr.CNvPicPr.PicLocks.NoChangeAspectAttr = unioffice.Bool(true)

	pic.BlipFill = dml.NewCT_BlipFillProperties()

	pic.BlipFill.Blip = dml.NewCT_Blip()

	imgID := s.AddImageToRels(img)
	pic.BlipFill.Blip.EmbedAttr = unioffice.String(imgID)

	pic.BlipFill.Stretch = dml.NewCT_StretchInfoProperties()
	pic.BlipFill.Stretch.FillRect = dml.NewCT_RelativeRect()

	pic.SpPr = dml.NewCT_ShapeProperties()
	pic.SpPr.PrstGeom = dml.NewCT_PresetGeometry2D()
	pic.SpPr.PrstGeom.PrstAttr = dml.ST_ShapeTypeRect

	ir := Image{pic}
	sz := img.Size()
	ir.Properties().SetWidth(measurement.Distance(sz.X) * measurement.Pixel72)
	ir.Properties().SetHeight(measurement.Distance(sz.Y) * measurement.Pixel72)
	ir.Properties().SetPosition(0, 0)
	return ir
}

// AddImageToRels adds an image relationship to a slide without putting image on the slide.
func (s Slide) AddImageToRels(img common.ImageRef) string {
	imgIdx := 0
	for i, ig := range s.p.Images {
		if ig == img {
			imgIdx = i + 1
			break
		}
	}

	var imgID string
	for i, os := range s.p.Slides() {
		if os.x == s.x {
			fn := fmt.Sprintf("../media/image%d.%s", imgIdx, img.Format())
			rel := s.p.slideRels[i].AddRelationship(fn, unioffice.ImageType)
			imgID = rel.ID()
		}
	}

	return imgID
}
