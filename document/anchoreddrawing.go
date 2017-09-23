// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"baliance.com/gooxml"
	"baliance.com/gooxml/common"
	"baliance.com/gooxml/measurement"
	pic "baliance.com/gooxml/schema/soo/dml/picture"
	"baliance.com/gooxml/schema/soo/wml"
)

// AnchoredDrawing is an absolutely positioned image within a document page.
type AnchoredDrawing struct {
	d *Document
	x *wml.WdAnchor
}

// X returns the inner wrapped XML type.
func (a AnchoredDrawing) X() *wml.WdAnchor {
	return a.x
}

// GetImage returns the ImageRef associated with an AnchoredDrawing.
func (a AnchoredDrawing) GetImage() (common.ImageRef, bool) {
	any := a.x.Graphic.GraphicData.Any
	if len(any) > 0 {
		p, ok := any[0].(*pic.Pic)
		if ok {
			if p.BlipFill != nil && p.BlipFill.Blip != nil && p.BlipFill.Blip.EmbedAttr != nil {
				return a.d.GetImageByRelID(*p.BlipFill.Blip.EmbedAttr)
			}
		}
	}
	return common.ImageRef{}, false
}

// SetName sets the name of the image, visible in the properties of the image
// within Word.
func (a AnchoredDrawing) SetName(name string) {
	a.x.DocPr.NameAttr = name
	for _, a := range a.x.Graphic.GraphicData.Any {
		if p, ok := a.(*pic.Pic); ok {
			p.NvPicPr.CNvPr.DescrAttr = gooxml.String(name)
		}
	}
}

// SetOrigin sets the origin of the image.  It defaults to ST_RelFromHPage and
// ST_RelFromVPage
func (a AnchoredDrawing) SetOrigin(h wml.WdST_RelFromH, v wml.WdST_RelFromV) {
	a.x.PositionH.RelativeFromAttr = h
	a.x.PositionV.RelativeFromAttr = v
}

// SetOffset sets the offset of the image relative to the origin, which by
// default this is the top-left corner of the page. Offset is incompatible with
// SetAlignment, whichever is called last is applied.
func (a AnchoredDrawing) SetOffset(x, y measurement.Distance) {
	a.SetXOffset(x)
	a.SetYOffset(y)
}

// SetXOffset sets the X offset for an image relative to the origin.
func (a AnchoredDrawing) SetXOffset(x measurement.Distance) {
	a.x.PositionH.Choice = &wml.WdCT_PosHChoice{}
	a.x.PositionH.Choice.PosOffset = gooxml.Int32(int32(x / measurement.EMU))
}

// SetYOffset sets the Y offset for an image relative to the origin.
func (a AnchoredDrawing) SetYOffset(y measurement.Distance) {
	a.x.PositionV.Choice = &wml.WdCT_PosVChoice{}
	a.x.PositionV.Choice.PosOffset = gooxml.Int32(int32(y / measurement.EMU))
}

// SetAlignment positions an anchored image via alignment.  Offset is
// incompatible with SetOffset, whichever is called last is applied.
func (a AnchoredDrawing) SetAlignment(h wml.WdST_AlignH, v wml.WdST_AlignV) {
	a.SetHAlignment(h)
	a.SetVAlignment(v)
}

// SetHAlignment sets the horizontal alignment for an anchored image.
func (a AnchoredDrawing) SetHAlignment(h wml.WdST_AlignH) {
	a.x.PositionH.Choice = &wml.WdCT_PosHChoice{}
	a.x.PositionH.Choice.Align = h
}

// SetVAlignment sets the vertical alignment for an anchored image.
func (a AnchoredDrawing) SetVAlignment(v wml.WdST_AlignV) {
	a.x.PositionV.Choice = &wml.WdCT_PosVChoice{}
	a.x.PositionV.Choice.Align = v
}

// SetSize sets the size of the displayed image on the page.
func (a AnchoredDrawing) SetSize(w, h measurement.Distance) {
	a.x.Extent.CxAttr = int64(float64(w*measurement.Pixel72) / measurement.EMU)
	a.x.Extent.CyAttr = int64(float64(h*measurement.Pixel72) / measurement.EMU)
}

// SetTextWrapNone unsets text wrapping so the image can float on top of the
// text. When used in conjunction with X/Y Offset relative to the page it can be
// used to place a logo at the top of a page at an absolute position that
// doesn't interfere with text.
func (a AnchoredDrawing) SetTextWrapNone() {
	a.x.Choice = &wml.WdEG_WrapTypeChoice{}
	a.x.Choice.WrapNone = wml.NewWdCT_WrapNone()
}

// SetTextWrapSquare sets the text wrap to square with a given wrap type.
func (a AnchoredDrawing) SetTextWrapSquare(t wml.WdST_WrapText) {
	a.x.Choice = &wml.WdEG_WrapTypeChoice{}
	a.x.Choice.WrapSquare = wml.NewWdCT_WrapSquare()
	a.x.Choice.WrapSquare.WrapTextAttr = t
}
