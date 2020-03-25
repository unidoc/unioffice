// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package document

import (
	"bytes"
	"errors"
	"math/rand"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/common"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/dml"
	pic "github.com/unidoc/unioffice/schema/soo/dml/picture"
	"github.com/unidoc/unioffice/schema/soo/ofc/sharedTypes"
	"github.com/unidoc/unioffice/schema/soo/wml"
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

// ClearContent clears any content in the run (text, tabs, breaks, etc.)
func (r Run) ClearContent() {
	r.x.EG_RunInnerContent = nil
}

// AddText adds tet to a run.
func (r Run) AddText(s string) {
	ic := wml.NewEG_RunInnerContent()
	r.x.EG_RunInnerContent = append(r.x.EG_RunInnerContent, ic)
	ic.T = wml.NewCT_Text()
	if unioffice.NeedsSpacePreserve(s) {
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

// Clear removes all of the content from within a run.
func (r Run) Clear() {
	r.x.EG_RunInnerContent = nil
}

// AddTab adds tab to a run and can be used with the the Paragraph's tab stops.
func (r Run) AddTab() {
	ic := r.newIC()
	ic.Tab = wml.NewCT_Empty()
}

// AddFieldWithFormatting adds a field (automatically computed text) to the
// document with field specifc formatting.
func (r Run) AddFieldWithFormatting(code string, fmt string, isDirty bool) {
	ic := r.newIC()
	ic.FldChar = wml.NewCT_FldChar()
	ic.FldChar.FldCharTypeAttr = wml.ST_FldCharTypeBegin
	if isDirty {
		ic.FldChar.DirtyAttr = &sharedTypes.ST_OnOff{}
		ic.FldChar.DirtyAttr.Bool = unioffice.Bool(true)
	}

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
	r.AddFieldWithFormatting(code, "", true)
}

// Properties returns the run properties.
func (r Run) Properties() RunProperties {
	if r.x.RPr == nil {
		r.x.RPr = wml.NewCT_RPr()
	}
	return RunProperties{r.x.RPr}
}

// AddBreak adds a line break to a run.
func (r Run) AddBreak() {
	ic := r.newIC()
	ic.Br = wml.NewCT_Br()
}

// AddPageBreak adds a page break to a run.
func (r Run) AddPageBreak() {
	ic := r.newIC()
	ic.Br = wml.NewCT_Br()
	ic.Br.TypeAttr = wml.ST_BrTypePage
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
	anchor.SimplePosAttr = unioffice.Bool(false)

	anchor.AllowOverlapAttr = true
	anchor.CNvGraphicFramePr = dml.NewCT_NonVisualGraphicFrameProperties()

	ic.Drawing.Anchor = append(ic.Drawing.Anchor, anchor)
	anchor.Graphic = dml.NewGraphic()
	anchor.Graphic.GraphicData = dml.NewCT_GraphicalObjectData()
	anchor.Graphic.GraphicData.UriAttr = "http://schemas.openxmlformats.org/drawingml/2006/picture"
	anchor.SimplePos.XAttr.ST_CoordinateUnqualified = unioffice.Int64(0)
	anchor.SimplePos.YAttr.ST_CoordinateUnqualified = unioffice.Int64(0)
	anchor.PositionH.RelativeFromAttr = wml.WdST_RelFromHPage
	anchor.PositionH.Choice = &wml.WdCT_PosHChoice{}
	anchor.PositionH.Choice.PosOffset = unioffice.Int32(0)

	anchor.PositionV.RelativeFromAttr = wml.WdST_RelFromVPage
	anchor.PositionV.Choice = &wml.WdCT_PosVChoice{}
	anchor.PositionV.Choice.PosOffset = unioffice.Int32(0)

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
	p.SpPr.Xfrm.Off.XAttr.ST_CoordinateUnqualified = unioffice.Int64(0)
	p.SpPr.Xfrm.Off.YAttr.ST_CoordinateUnqualified = unioffice.Int64(0)
	p.SpPr.Xfrm.Ext = dml.NewCT_PositiveSize2D()
	p.SpPr.Xfrm.Ext.CxAttr = int64(img.Size().X * measurement.Point)
	p.SpPr.Xfrm.Ext.CyAttr = int64(img.Size().Y * measurement.Point)
	// required by Word on OSX for the image to display
	p.SpPr.PrstGeom = dml.NewCT_PresetGeometry2D()
	p.SpPr.PrstGeom.PrstAttr = dml.ST_ShapeTypeRect

	return ad, nil
}

// AddDrawingInline adds an inline drawing from an ImageRef.
func (r Run) AddDrawingInline(img common.ImageRef) (InlineDrawing, error) {
	ic := r.newIC()
	ic.Drawing = wml.NewCT_Drawing()

	inl := wml.NewWdInline()
	inline := InlineDrawing{r.d, inl}

	// required by Word on OSX for the file to open
	//anchor.SimplePosAttr = gooxml.Bool(false)

	//anchor.AllowOverlapAttr = true
	inl.CNvGraphicFramePr = dml.NewCT_NonVisualGraphicFrameProperties()

	ic.Drawing.Inline = append(ic.Drawing.Inline, inl)
	inl.Graphic = dml.NewGraphic()
	inl.Graphic.GraphicData = dml.NewCT_GraphicalObjectData()
	inl.Graphic.GraphicData.UriAttr = "http://schemas.openxmlformats.org/drawingml/2006/picture"

	inl.DistTAttr = unioffice.Uint32(0)
	inl.DistLAttr = unioffice.Uint32(0)
	inl.DistBAttr = unioffice.Uint32(0)
	inl.DistRAttr = unioffice.Uint32(0)

	inl.Extent.CxAttr = int64(float64(img.Size().X*measurement.Pixel72) / measurement.EMU)
	inl.Extent.CyAttr = int64(float64(img.Size().Y*measurement.Pixel72) / measurement.EMU)

	// Mac Word chokes if the ID is greater than an int32, even though the field is a
	// uint32 in the XSD
	randID := 0x7FFFFFFF & rand.Uint32()
	inl.DocPr.IdAttr = randID
	p := pic.NewPic()
	p.NvPicPr.CNvPr.IdAttr = randID

	// find the reference to the actual image file in the document relationships
	// so we can embed via the relationship ID
	imgID := img.RelID()
	if imgID == "" {
		return inline, errors.New("couldn't find reference to image within document relations")
	}

	inl.Graphic.GraphicData.Any = append(inl.Graphic.GraphicData.Any, p)
	p.BlipFill = dml.NewCT_BlipFillProperties()
	p.BlipFill.Blip = dml.NewCT_Blip()
	p.BlipFill.Blip.EmbedAttr = &imgID
	p.BlipFill.Stretch = dml.NewCT_StretchInfoProperties()
	p.BlipFill.Stretch.FillRect = dml.NewCT_RelativeRect()

	p.SpPr = dml.NewCT_ShapeProperties()
	// Required to allow resizing
	p.SpPr.Xfrm = dml.NewCT_Transform2D()
	p.SpPr.Xfrm.Off = dml.NewCT_Point2D()
	p.SpPr.Xfrm.Off.XAttr.ST_CoordinateUnqualified = unioffice.Int64(0)
	p.SpPr.Xfrm.Off.YAttr.ST_CoordinateUnqualified = unioffice.Int64(0)
	p.SpPr.Xfrm.Ext = dml.NewCT_PositiveSize2D()
	p.SpPr.Xfrm.Ext.CxAttr = int64(img.Size().X * measurement.Point)
	p.SpPr.Xfrm.Ext.CyAttr = int64(img.Size().Y * measurement.Point)
	// required by Word on OSX for the image to display
	p.SpPr.PrstGeom = dml.NewCT_PresetGeometry2D()
	p.SpPr.PrstGeom.PrstAttr = dml.ST_ShapeTypeRect

	return inline, nil
}

// IsFootnote returns a bool based on whether the run has a
// footnote or not. Returns both a bool as to whether it has
// a footnote as well as the ID of the footnote.
func (r Run) IsFootnote() (bool, int64) {
	if r.x.EG_RunInnerContent != nil {
		if r.x.EG_RunInnerContent[0].FootnoteReference != nil {
			return true, r.x.EG_RunInnerContent[0].FootnoteReference.IdAttr
		}
	}
	return false, 0
}

// IsEndnote returns a bool based on whether the run has a
// footnote or not. Returns both a bool as to whether it has
// a footnote as well as the ID of the footnote.
func (r Run) IsEndnote() (bool, int64) {
	if r.x.EG_RunInnerContent != nil {
		if r.x.EG_RunInnerContent[0].EndnoteReference != nil {
			return true, r.x.EG_RunInnerContent[0].EndnoteReference.IdAttr
		}
	}
	return false, 0
}
