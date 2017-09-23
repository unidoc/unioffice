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
	"baliance.com/gooxml/schema/soo/dml"
	pic "baliance.com/gooxml/schema/soo/dml/picture"
	"baliance.com/gooxml/schema/soo/ofc/sharedTypes"
	"baliance.com/gooxml/schema/soo/wml"
)

// Run is a run of text within a paragraph that shares the same formatting.
type Run struct {
	d *Document
	x *wml.CT_R
}

// X returns the inner wrapped XML type.
func (r Run) X() *wml.CT_R {
	return r.x
}

// Text returns the underlying tet in the run.
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

// AddText adds tet to a run.
func (r Run) AddText(s string) {
	ic := wml.NewEG_RunInnerContent()
	r.x.EG_RunInnerContent = append(r.x.EG_RunInnerContent, ic)
	ic.T = wml.NewCT_Text()
	if gooxml.NeedsSpacePreserve(s) {
		p := "preserve"
		ic.T.SpaceAttr = &p
	}
	ic.T.Content = s
}

func (r Run) newIC() *wml.EG_RunInnerContent {
	ic := wml.NewEG_RunInnerContent()
	r.x.EG_RunInnerContent = append(r.x.EG_RunInnerContent, ic)
	return ic
}

// AddTab adds tab to a run and can be used with the the Paragraph's tab stops.
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
	}
}

// SetFontFamily sets the Ascii & HAnsi fonly family for a run.
func (r Run) SetFontFamily(family string) {
	r.ensureRPR()
	if r.x.RPr.RFonts == nil {
		r.x.RPr.RFonts = wml.NewCT_Fonts()
	}
	r.x.RPr.RFonts.AsciiAttr = gooxml.String(family)
	r.x.RPr.RFonts.HAnsiAttr = gooxml.String(family)
}

// SetFontSize sets the font size.
func (r Run) SetFontSize(sz measurement.Distance) {
	r.x.RPr.Sz = wml.NewCT_HpsMeasure()
	// size is measured in half points
	r.x.RPr.Sz.ValAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(uint64(sz / measurement.HalfPoint))
	r.x.RPr.SzCs = wml.NewCT_HpsMeasure()
	r.x.RPr.SzCs.ValAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(uint64(sz / measurement.HalfPoint))
}

// IsBold returns true if the run has been set to bold.
func (r Run) IsBold() bool {
	if r.x.RPr == nil {
		return false
	}
	return r.x.RPr.B != nil
}

// SetBold sets the run to bold.
func (r Run) SetBold(b bool) {
	r.ensureRPR()
	if !b {
		r.x.RPr.B = nil
		r.x.RPr.BCs = nil
	} else {
		r.x.RPr.B = wml.NewCT_OnOff()
		r.x.RPr.BCs = wml.NewCT_OnOff()
	}
}

// IsItalic returns true if the run was set to bold.
func (r Run) IsItalic() bool {
	if r.x.RPr == nil {
		return false
	}
	return r.x.RPr.I != nil
}

// SetItalic sets the run to italic.
func (r Run) SetItalic(b bool) {
	r.ensureRPR()
	if !b {
		r.x.RPr.I = nil
		r.x.RPr.ICs = nil
	} else {
		r.x.RPr.I = wml.NewCT_OnOff()
		r.x.RPr.ICs = wml.NewCT_OnOff()
	}
}

// SetAllCaps sets the run to all caps.
func (r Run) SetAllCaps(b bool) {
	r.ensureRPR()
	if !b {
		r.x.RPr.Caps = nil
	} else {
		r.x.RPr.Caps = wml.NewCT_OnOff()
	}
}

// SetSmallCaps sets the run to small caps.
func (r Run) SetSmallCaps(b bool) {
	r.ensureRPR()
	if !b {
		r.x.RPr.SmallCaps = nil
	} else {
		r.x.RPr.SmallCaps = wml.NewCT_OnOff()
	}
}

// SetUnderline sets the run to underline with a particular style and color.
func (r Run) SetUnderline(style wml.ST_Underline, c color.Color) {
	r.ensureRPR()
	if style == wml.ST_UnderlineUnset {
		r.x.RPr.U = nil
	} else {
		r.x.RPr.U = wml.NewCT_Underline()
		r.x.RPr.U.ColorAttr = &wml.ST_HexColor{}
		r.x.RPr.U.ColorAttr.ST_HexColorRGB = c.AsRGBString()
		r.x.RPr.U.ValAttr = style
	}
}

// SetStrikeThrough sets the run to strike-through.
func (r Run) SetStrikeThrough(b bool) {
	r.ensureRPR()
	if !b {
		r.x.RPr.Strike = nil
	} else {
		r.x.RPr.Strike = wml.NewCT_OnOff()
	}
}

// SetDoubleStrikeThrough sets the run to double strike-through.
func (r Run) SetDoubleStrikeThrough(b bool) {
	r.ensureRPR()
	if !b {
		r.x.RPr.Dstrike = nil
	} else {
		r.x.RPr.Dstrike = wml.NewCT_OnOff()
	}
}

// SetOutline sets the run to outlined text.
func (r Run) SetOutline(b bool) {
	r.ensureRPR()
	if !b {
		r.x.RPr.Outline = nil
	} else {
		r.x.RPr.Outline = wml.NewCT_OnOff()
	}
}

// SetShadow sets the run to shadowed text.
func (r Run) SetShadow(b bool) {
	r.ensureRPR()
	if !b {
		r.x.RPr.Shadow = nil
	} else {
		r.x.RPr.Shadow = wml.NewCT_OnOff()
	}
}

// SetEmboss sets the run to embossed text.
func (r Run) SetEmboss(b bool) {
	r.ensureRPR()
	if !b {
		r.x.RPr.Emboss = nil
	} else {
		r.x.RPr.Emboss = wml.NewCT_OnOff()
	}
}

// SetImprint sets the run to imprinted text.
func (r Run) SetImprint(b bool) {
	r.ensureRPR()
	if !b {
		r.x.RPr.Imprint = nil
	} else {
		r.x.RPr.Imprint = wml.NewCT_OnOff()
	}
}

// ClearColor clears the text color.
func (r Run) ClearColor() {
	r.ensureRPR()
	r.x.RPr.Color = nil
}

// SetColor sets the text color.
func (r Run) SetColor(c color.Color) {
	r.ensureRPR()
	r.x.RPr.Color = wml.NewCT_Color()
	r.x.RPr.Color.ValAttr.ST_HexColorRGB = c.AsRGBString()
}

// SetHighlight highlights text in a specified color.
func (r Run) SetHighlight(c wml.ST_HighlightColor) {
	r.ensureRPR()
	r.x.RPr.Highlight = wml.NewCT_Highlight()
	r.x.RPr.Highlight.ValAttr = c
}

// SetEffect sets a text effect on the run.
func (r Run) SetEffect(e wml.ST_TextEffect) {
	r.ensureRPR()
	if e == wml.ST_TextEffectUnset {
		r.x.RPr.Effect = nil
	} else {
		r.x.RPr.Effect = wml.NewCT_TextEffect()
		r.x.RPr.Effect.ValAttr = wml.ST_TextEffectShimmer
	}
}

// AddBreak adds a line break to a run.
func (r Run) AddBreak() {
	ic := r.newIC()
	ic.Br = wml.NewCT_Br()
}

// DrawingAnchored returns a slice of AnchoredDrawings.
func (r Run) DrawingAnchored() []AnchoredDrawing {
	ret := []AnchoredDrawing{}
	for _, ic := range r.x.EG_RunInnerContent {
		if ic.Drawing == nil {
			continue
		}
		for _, anc := range ic.Drawing.Anchor {
			ret = append(ret, AnchoredDrawing{r.d, anc})
		}
	}
	return ret
}

// AddDrawingAnchored adds an anchored (floating) drawing from an ImageRef.
func (r Run) AddDrawingAnchored(img common.ImageRef) (AnchoredDrawing, error) {
	ic := r.newIC()
	ic.Drawing = wml.NewCT_Drawing()
	anchor := wml.NewWdAnchor()

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
	anchor.PositionH.RelativeFromAttr = wml.WdST_RelFromHPage
	anchor.PositionH.Choice = &wml.WdCT_PosHChoice{}
	anchor.PositionH.Choice.PosOffset = gooxml.Int32(0)

	anchor.PositionV.RelativeFromAttr = wml.WdST_RelFromVPage
	anchor.PositionV.Choice = &wml.WdCT_PosVChoice{}
	anchor.PositionV.Choice.PosOffset = gooxml.Int32(0)

	anchor.Extent.CxAttr = int64(float64(img.Size().X*measurement.Pixel72) / measurement.EMU)
	anchor.Extent.CyAttr = int64(float64(img.Size().Y*measurement.Pixel72) / measurement.EMU)
	anchor.Choice = &wml.WdEG_WrapTypeChoice{}
	anchor.Choice.WrapSquare = wml.NewWdCT_WrapSquare()
	anchor.Choice.WrapSquare.WrapTextAttr = wml.WdST_WrapTextBothSides

	// Mac Word chokes if the ID is greater than an int32, even though the field is a
	// uint32 in the XSD
	randID := 0x7FFFFFFF & rand.Uint32()
	anchor.DocPr.IdAttr = randID
	p := pic.NewPic()
	p.NvPicPr.CNvPr.IdAttr = randID

	// find the reference to the actual image file in the document relationships
	// so we can embed via the relationship ID
	imgID := img.RelID()
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
	p.SpPr.Xfrm.Ext.CxAttr = int64(img.Size().X * measurement.Point)
	p.SpPr.Xfrm.Ext.CyAttr = int64(img.Size().Y * measurement.Point)
	// required by Word on OSX for the image to display
	p.SpPr.PrstGeom = dml.NewCT_PresetGeometry2D()
	p.SpPr.PrstGeom.PrstAttr = dml.ST_ShapeTypeRect

	return ad, nil
}
