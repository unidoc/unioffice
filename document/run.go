// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"bytes"
	"errors"
	"math/rand"

	"baliance.com/gooxml"
	"baliance.com/gooxml/color"
	"baliance.com/gooxml/common"
	"baliance.com/gooxml/measurement"
	dml "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
	pic "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/picture"
	wd "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"
	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
	wml "baliance.com/gooxml/schema/schemas.openxmlformats.org/wordprocessingml"
)

type Run struct {
	d *Document
	x *wml.CT_R
}

// X returns the inner wrapped XML type.
func (r Run) X() *wml.CT_R {
	return r.x
}

func (r Run) Text() string {
	if len(r.x.EG_RunInnerContent) == 0 {
		return ""
	}
	buf := bytes.Buffer{}
	for _, ic := range r.x.EG_RunInnerContent {
		if ic.T != nil {
			buf.WriteString(ic.T.Content)
		}
		if ic.Tab != nil {
			buf.WriteByte('\t')
		}
	}
	return buf.String()
}

func (r Run) AddText(s string) {
	ic := wml.NewEG_RunInnerContent()
	r.x.EG_RunInnerContent = append(r.x.EG_RunInnerContent, ic)
	ic.T = wml.NewCT_Text()
	ic.T.Content = s
}

func (r Run) newIC() *wml.EG_RunInnerContent {
	ic := wml.NewEG_RunInnerContent()
	r.x.EG_RunInnerContent = append(r.x.EG_RunInnerContent, ic)
	return ic
}
func (r Run) AddTab() {
	ic := r.newIC()
	ic.Tab = wml.NewCT_Empty()
}

// AddFieldWithFormatting adds a field (automatically computed text) to the
// document with field specifc formatting.
func (r Run) AddFieldWithFormatting(code string, fmt string) {
	ic := r.newIC()
	ic.FldChar = wml.NewCT_FldChar()
	ic.FldChar.FldCharTypeAttr = wml.ST_FldCharTypeBegin
	ic.FldChar.DirtyAttr = &sharedTypes.ST_OnOff{}
	ic.FldChar.DirtyAttr.Bool = gooxml.Bool(true)

	ic = r.newIC()
	ic.InstrText = wml.NewCT_Text()
	if fmt != "" {
		ic.InstrText.Content = code + " " + fmt
	} else {
		ic.InstrText.Content = code
	}

	ic = r.newIC()
	ic.FldChar = wml.NewCT_FldChar()
	ic.FldChar.FldCharTypeAttr = wml.ST_FldCharTypeEnd
}

// AddField adds a field (automatically computed text) to the document.
func (r Run) AddField(code string) {
	r.AddFieldWithFormatting(code, "")
}

func (r Run) ensureRPR() {
	if r.x.RPr == nil {
		r.x.RPr = wml.NewCT_RPr()
		b := wml.NewEG_RPrBase()
		r.x.RPr.EG_RPrBase = append(r.x.RPr.EG_RPrBase, b)
	}
}
func (r Run) SetFontFamily(family string) {
	r.ensureRPR()
	b := r.x.RPr.EG_RPrBase[0]
	if b.RFonts == nil {
		b.RFonts = wml.NewCT_Fonts()
	}
	b.RFonts.AsciiAttr = gooxml.String(family)
	b.RFonts.HAnsiAttr = gooxml.String(family)
}

func (r Run) SetFontSize(points uint64) {
	b := r.x.RPr.EG_RPrBase[0]
	b.Sz = wml.NewCT_HpsMeasure()
	b.Sz.ValAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(2 * points)
	b.SzCs = wml.NewCT_HpsMeasure()
	b.SzCs.ValAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(2 * points)
}

func (r Run) SetBold(b bool) {
	r.ensureRPR()
	if !b {
		r.x.RPr.EG_RPrBase[0].B = nil
		r.x.RPr.EG_RPrBase[0].BCs = nil
	} else {
		r.x.RPr.EG_RPrBase[0].B = wml.NewCT_OnOff()
		r.x.RPr.EG_RPrBase[0].BCs = wml.NewCT_OnOff()
	}
}

func (r Run) SetItalic(b bool) {
	r.ensureRPR()
	if !b {
		r.x.RPr.EG_RPrBase[0].I = nil
		r.x.RPr.EG_RPrBase[0].ICs = nil
	} else {
		r.x.RPr.EG_RPrBase[0].I = wml.NewCT_OnOff()
		r.x.RPr.EG_RPrBase[0].ICs = wml.NewCT_OnOff()
	}
}

func (r Run) SetAllCaps(b bool) {
	r.ensureRPR()
	if !b {
		r.x.RPr.EG_RPrBase[0].Caps = nil
	} else {
		r.x.RPr.EG_RPrBase[0].Caps = wml.NewCT_OnOff()
	}
}

func (r Run) SetSmallCaps(b bool) {
	r.ensureRPR()
	if !b {
		r.x.RPr.EG_RPrBase[0].SmallCaps = nil
	} else {
		r.x.RPr.EG_RPrBase[0].SmallCaps = wml.NewCT_OnOff()
	}
}

func (r Run) SetUnderline(style wml.ST_Underline, c color.Color) {
	r.ensureRPR()
	if style == wml.ST_UnderlineUnset {
		r.x.RPr.EG_RPrBase[0].U = nil
	} else {
		r.x.RPr.EG_RPrBase[0].U = wml.NewCT_Underline()
		r.x.RPr.EG_RPrBase[0].U.ColorAttr = &wml.ST_HexColor{}
		r.x.RPr.EG_RPrBase[0].U.ColorAttr.ST_HexColorRGB = c.AsRGBString()
		r.x.RPr.EG_RPrBase[0].U.ValAttr = style
	}
}

func (r Run) SetStrikeThrough(b bool) {
	r.ensureRPR()
	if !b {
		r.x.RPr.EG_RPrBase[0].Strike = nil
	} else {
		r.x.RPr.EG_RPrBase[0].Strike = wml.NewCT_OnOff()
	}
}

func (r Run) SetDoubleStrikeThrough(b bool) {
	r.ensureRPR()
	if !b {
		r.x.RPr.EG_RPrBase[0].Dstrike = nil
	} else {
		r.x.RPr.EG_RPrBase[0].Dstrike = wml.NewCT_OnOff()
	}
}

func (r Run) SetOutline(b bool) {
	r.ensureRPR()
	if !b {
		r.x.RPr.EG_RPrBase[0].Outline = nil
	} else {
		r.x.RPr.EG_RPrBase[0].Outline = wml.NewCT_OnOff()
	}
}

func (r Run) SetShadow(b bool) {
	r.ensureRPR()
	if !b {
		r.x.RPr.EG_RPrBase[0].Shadow = nil
	} else {
		r.x.RPr.EG_RPrBase[0].Shadow = wml.NewCT_OnOff()
	}
}

func (r Run) SetEmboss(b bool) {
	r.ensureRPR()
	if !b {
		r.x.RPr.EG_RPrBase[0].Emboss = nil
	} else {
		r.x.RPr.EG_RPrBase[0].Emboss = wml.NewCT_OnOff()
	}
}

func (r Run) SetImprint(b bool) {
	r.ensureRPR()
	if !b {
		r.x.RPr.EG_RPrBase[0].Imprint = nil
	} else {
		r.x.RPr.EG_RPrBase[0].Imprint = wml.NewCT_OnOff()
	}
}

func (r Run) ClearColor() {
	r.ensureRPR()
	r.x.RPr.EG_RPrBase[0].Color = nil
}

func (r Run) SetColor(c color.Color) {
	r.ensureRPR()
	r.x.RPr.EG_RPrBase[0].Color = wml.NewCT_Color()
	r.x.RPr.EG_RPrBase[0].Color.ValAttr.ST_HexColorRGB = c.AsRGBString()
}

func (r Run) SetHighlight(c wml.ST_HighlightColor) {
	r.ensureRPR()
	r.x.RPr.EG_RPrBase[0].Highlight = wml.NewCT_Highlight()
	r.x.RPr.EG_RPrBase[0].Highlight.ValAttr = c
}

// TODO: test this on Windows
func (r Run) SetEffect(e wml.ST_TextEffect) {
	r.ensureRPR()
	if e == wml.ST_TextEffectUnset {
		r.x.RPr.EG_RPrBase[0].Effect = nil
	} else {
		r.x.RPr.EG_RPrBase[0].Effect = wml.NewCT_TextEffect()
		r.x.RPr.EG_RPrBase[0].Effect.ValAttr = wml.ST_TextEffectShimmer
	}
}

// AddBreak adds a line break to a run.
func (r Run) AddBreak() {
	ic := r.newIC()
	ic.Br = wml.NewCT_Br()
}

// AddDrawingAnchored adds an anchored drawing that
func (r Run) AddDrawingAnchored(img ImageRef) (AnchoredDrawing, error) {
	ic := r.newIC()
	ic.Drawing = wml.NewCT_Drawing()
	anchor := wd.NewAnchor()

	ad := AnchoredDrawing{r.d, anchor}

	// required by Word on OSX for the file to open
	anchor.SimplePosAttr = gooxml.Bool(false)

	anchor.AllowOverlapAttr = true
	anchor.CNvGraphicFramePr = dml.NewCT_NonVisualGraphicFrameProperties()

	ic.Drawing.Anchor = append(ic.Drawing.Anchor, anchor)
	anchor.Graphic = dml.NewGraphic()
	anchor.Graphic.GraphicData = dml.NewCT_GraphicalObjectData()
	anchor.Graphic.GraphicData.UriAttr = "http://schemas.openxmlformats.org/drawingml/2006/picture"
	anchor.SimplePos.XAttr.ST_CoordinateUnqualified = gooxml.Int64(0)
	anchor.SimplePos.YAttr.ST_CoordinateUnqualified = gooxml.Int64(0)
	anchor.PositionH.RelativeFromAttr = wd.ST_RelFromHPage
	anchor.PositionH.Choice = &wd.CT_PosHChoice{}
	anchor.PositionH.Choice.PosOffset = gooxml.Int32(0)

	anchor.PositionV.RelativeFromAttr = wd.ST_RelFromVPage
	anchor.PositionV.Choice = &wd.CT_PosVChoice{}
	anchor.PositionV.Choice.PosOffset = gooxml.Int32(0)

	anchor.Extent.CxAttr = int64(float64(img.img.Size.X*measurement.Pixel72) / measurement.EMU)
	anchor.Extent.CyAttr = int64(float64(img.img.Size.Y*measurement.Pixel72) / measurement.EMU)
	anchor.Choice = &wd.EG_WrapTypeChoice{}
	anchor.Choice.WrapSquare = wd.NewCT_WrapSquare()
	anchor.Choice.WrapSquare.WrapTextAttr = wd.ST_WrapTextBothSides

	// Mac Word chokes if the ID is greater than an int32, even though the field is a
	// uint32 in the XSD
	randID := uint32(0x7FFFFFFF & rand.Uint32())
	anchor.DocPr.IdAttr = randID
	p := pic.NewPic()
	p.NvPicPr.CNvPr.IdAttr = randID

	// find the reference to the actual image file in the document relationships
	// so we can embed via the relationship ID
	imgIdx := -1
	for i, ir := range r.d.images {
		if img.ref == ir {
			imgIdx = i
		}
	}
	if imgIdx == -1 {
		return ad, errors.New("couldn't find reference to image within document")
	}
	imgID := r.d.docRels.FindRIDForN(imgIdx, common.ImageType)
	if imgID == "" {
		return ad, errors.New("couldn't find reference to image within document relations")
	}

	anchor.Graphic.GraphicData.Any = append(anchor.Graphic.GraphicData.Any, p)
	p.BlipFill = dml.NewCT_BlipFillProperties()
	p.BlipFill.Blip = dml.NewCT_Blip()
	p.BlipFill.Blip.EmbedAttr = &imgID
	p.BlipFill.Stretch = dml.NewCT_StretchInfoProperties()
	p.BlipFill.Stretch.FillRect = dml.NewCT_RelativeRect()

	p.SpPr = dml.NewCT_ShapeProperties()
	// Required to allow resizing
	p.SpPr.Xfrm = dml.NewCT_Transform2D()
	p.SpPr.Xfrm.Off = dml.NewCT_Point2D()
	p.SpPr.Xfrm.Off.XAttr.ST_CoordinateUnqualified = gooxml.Int64(0)
	p.SpPr.Xfrm.Off.YAttr.ST_CoordinateUnqualified = gooxml.Int64(0)
	p.SpPr.Xfrm.Ext = dml.NewCT_PositiveSize2D()
	p.SpPr.Xfrm.Ext.CxAttr = int64(img.img.Size.X * measurement.Point)
	p.SpPr.Xfrm.Ext.CyAttr = int64(img.img.Size.Y * measurement.Point)
	// required by Word on OSX for the image to display
	p.SpPr.PrstGeom = dml.NewCT_PresetGeometry2D()
	p.SpPr.PrstGeom.PrstAttr = dml.ST_ShapeTypeRect

	return ad, nil
}
