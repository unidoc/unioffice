// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentation

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"os"
	"path"
	"strings"

	"baliance.com/gooxml"
	"baliance.com/gooxml/common"
	"baliance.com/gooxml/measurement"
	"baliance.com/gooxml/schema/soo/dml"
	"baliance.com/gooxml/schema/soo/ofc/sharedTypes"
	"baliance.com/gooxml/schema/soo/pkg/relationships"
	"baliance.com/gooxml/schema/soo/pml"
	"baliance.com/gooxml/zippkg"
)

// Presentation is the a presentation base document.
type Presentation struct {
	common.DocBase
	x          *pml.Presentation
	prels      common.Relationships
	slides     []*pml.Sld
	slideRels  []common.Relationships
	masters    []*pml.SldMaster
	masterRels []common.Relationships
	layouts    []*pml.SldLayout
	layoutRels []common.Relationships
	themes     []*dml.Theme
	themeRels  []common.Relationships
}

func newEmpty() *Presentation {
	p := &Presentation{x: pml.NewPresentation()}
	p.x.SldIdLst = pml.NewCT_SlideIdList()
	p.x.ConformanceAttr = sharedTypes.ST_ConformanceClassTransitional
	p.AppProperties = common.NewAppProperties()
	p.CoreProperties = common.NewCoreProperties()
	p.ContentTypes = common.NewContentTypes()
	p.Rels = common.NewRelationships()
	p.prels = common.NewRelationships()
	return p
}

// New initializes and reurns a new presentation
func New() *Presentation {
	p := newEmpty()

	p.ContentTypes.AddOverride("/ppt/presentation.xml", "application/vnd.openxmlformats-officedocument.presentationml.presentation.main+xml")

	p.Rels.AddRelationship("docProps/core.xml", "http://schemas.openxmlformats.org/package/2006/relationships/metadata/core-properties")
	p.Rels.AddRelationship("docProps/app.xml", "http://schemas.openxmlformats.org/officeDocument/2006/relationships/extended-properties")
	p.Rels.AddRelationship("ppt/presentation.xml", "http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument")

	p.x.SldMasterIdLst = pml.NewCT_SlideMasterIdList()
	m := pml.NewSldMaster()
	m.ClrMap.Bg1Attr = dml.ST_ColorSchemeIndexLt1
	m.ClrMap.Bg2Attr = dml.ST_ColorSchemeIndexLt2
	m.ClrMap.Tx1Attr = dml.ST_ColorSchemeIndexDk1
	m.ClrMap.Tx2Attr = dml.ST_ColorSchemeIndexDk2
	m.ClrMap.Accent1Attr = dml.ST_ColorSchemeIndexAccent1
	m.ClrMap.Accent2Attr = dml.ST_ColorSchemeIndexAccent2
	m.ClrMap.Accent3Attr = dml.ST_ColorSchemeIndexAccent3
	m.ClrMap.Accent4Attr = dml.ST_ColorSchemeIndexAccent4
	m.ClrMap.Accent5Attr = dml.ST_ColorSchemeIndexAccent5
	m.ClrMap.Accent6Attr = dml.ST_ColorSchemeIndexAccent6
	m.ClrMap.HlinkAttr = dml.ST_ColorSchemeIndexHlink
	m.ClrMap.FolHlinkAttr = dml.ST_ColorSchemeIndexFolHlink

	p.masters = append(p.masters, m)

	smFn := gooxml.AbsoluteFilename(gooxml.DocTypePresentation, gooxml.SlideMasterType, 1)
	p.ContentTypes.AddOverride(smFn, gooxml.SlideMasterContentType)

	mrelID := p.prels.AddAutoRelationship(gooxml.DocTypePresentation, gooxml.OfficeDocumentType,
		1, gooxml.SlideMasterType)
	smid := pml.NewCT_SlideMasterIdListEntry()
	smid.IdAttr = gooxml.Uint32(2147483648)
	smid.RIdAttr = mrelID.ID()
	p.x.SldMasterIdLst.SldMasterId = append(p.x.SldMasterIdLst.SldMasterId, smid)
	mrel := common.NewRelationships()
	p.masterRels = append(p.masterRels, mrel)

	ls := pml.NewSldLayout()
	lrid := mrel.AddAutoRelationship(gooxml.DocTypePresentation, gooxml.SlideMasterType, 1, gooxml.SlideLayoutType)
	slfn := gooxml.AbsoluteFilename(gooxml.DocTypePresentation, gooxml.SlideLayoutType, 1)
	p.ContentTypes.AddOverride(slfn, gooxml.SlideLayoutContentType)
	mrel.AddAutoRelationship(gooxml.DocTypePresentation, gooxml.SlideMasterType, 1, gooxml.ThemeType)
	p.layouts = append(p.layouts, ls)

	m.SldLayoutIdLst = pml.NewCT_SlideLayoutIdList()
	lid := pml.NewCT_SlideLayoutIdListEntry()
	lid.IdAttr = gooxml.Uint32(2147483649)
	lid.RIdAttr = lrid.ID()
	m.SldLayoutIdLst.SldLayoutId = append(m.SldLayoutIdLst.SldLayoutId, lid)

	lrel := common.NewRelationships()
	p.layoutRels = append(p.layoutRels, lrel)
	lrel.AddAutoRelationship(gooxml.DocTypePresentation, gooxml.SlideType, 1, gooxml.SlideMasterType)
	p.x.NotesSz.CxAttr = 6858000
	p.x.NotesSz.CyAttr = 9144000

	thm := dml.NewTheme()

	thm.NameAttr = gooxml.String("gooxml Theme")
	thm.ThemeElements.ClrScheme.NameAttr = "Office"
	thm.ThemeElements.ClrScheme.Dk1.SysClr = dml.NewCT_SystemColor()
	thm.ThemeElements.ClrScheme.Dk1.SysClr.LastClrAttr = gooxml.String("000000")
	thm.ThemeElements.ClrScheme.Dk1.SysClr.ValAttr = dml.ST_SystemColorValWindowText

	thm.ThemeElements.ClrScheme.Lt1.SysClr = dml.NewCT_SystemColor()
	thm.ThemeElements.ClrScheme.Lt1.SysClr.LastClrAttr = gooxml.String("ffffff")
	thm.ThemeElements.ClrScheme.Lt1.SysClr.ValAttr = dml.ST_SystemColorValWindow

	thm.ThemeElements.ClrScheme.Dk2.SrgbClr = dml.NewCT_SRgbColor()
	thm.ThemeElements.ClrScheme.Dk2.SrgbClr.ValAttr = "44546a"

	thm.ThemeElements.ClrScheme.Lt2.SrgbClr = dml.NewCT_SRgbColor()
	thm.ThemeElements.ClrScheme.Lt2.SrgbClr.ValAttr = "e7e7e6"

	thm.ThemeElements.ClrScheme.Accent1.SrgbClr = dml.NewCT_SRgbColor()
	thm.ThemeElements.ClrScheme.Accent1.SrgbClr.ValAttr = "4472c4"

	thm.ThemeElements.ClrScheme.Accent2.SrgbClr = dml.NewCT_SRgbColor()
	thm.ThemeElements.ClrScheme.Accent2.SrgbClr.ValAttr = "ed7d31"

	thm.ThemeElements.ClrScheme.Accent3.SrgbClr = dml.NewCT_SRgbColor()
	thm.ThemeElements.ClrScheme.Accent3.SrgbClr.ValAttr = "a5a5a5"

	thm.ThemeElements.ClrScheme.Accent4.SrgbClr = dml.NewCT_SRgbColor()
	thm.ThemeElements.ClrScheme.Accent4.SrgbClr.ValAttr = "ffc000"

	thm.ThemeElements.ClrScheme.Accent5.SrgbClr = dml.NewCT_SRgbColor()
	thm.ThemeElements.ClrScheme.Accent5.SrgbClr.ValAttr = "5b9bd5"

	thm.ThemeElements.ClrScheme.Accent6.SrgbClr = dml.NewCT_SRgbColor()
	thm.ThemeElements.ClrScheme.Accent6.SrgbClr.ValAttr = "70ad47"

	thm.ThemeElements.ClrScheme.Hlink.SrgbClr = dml.NewCT_SRgbColor()
	thm.ThemeElements.ClrScheme.Hlink.SrgbClr.ValAttr = "0563c1"

	thm.ThemeElements.ClrScheme.FolHlink.SrgbClr = dml.NewCT_SRgbColor()
	thm.ThemeElements.ClrScheme.FolHlink.SrgbClr.ValAttr = "954f72"

	thm.ThemeElements.FontScheme.NameAttr = "Office"
	thm.ThemeElements.FontScheme.MajorFont.Latin.TypefaceAttr = "Calibri Light"
	thm.ThemeElements.FontScheme.MinorFont.Latin.TypefaceAttr = "Calibri"

	thm.ThemeElements.FmtScheme.NameAttr = gooxml.String("Office")
	// fills
	fp := dml.NewEG_FillProperties()
	thm.ThemeElements.FmtScheme.FillStyleLst.EG_FillProperties = append(thm.ThemeElements.FmtScheme.FillStyleLst.EG_FillProperties, fp)
	fp.SolidFill = &dml.CT_SolidColorFillProperties{
		SchemeClr: &dml.CT_SchemeColor{ValAttr: dml.ST_SchemeColorValPhClr},
	}

	// rot fill 0
	fp = dml.NewEG_FillProperties()
	thm.ThemeElements.FmtScheme.FillStyleLst.EG_FillProperties = append(thm.ThemeElements.FmtScheme.FillStyleLst.EG_FillProperties, fp)
	// add it twice so OSX word doesn't choke
	thm.ThemeElements.FmtScheme.FillStyleLst.EG_FillProperties = append(thm.ThemeElements.FmtScheme.FillStyleLst.EG_FillProperties, fp)
	fp.GradFill = &dml.CT_GradientFillProperties{RotWithShapeAttr: gooxml.Bool(true),
		GsLst: &dml.CT_GradientStopList{},
		Lin:   &dml.CT_LinearShadeProperties{}}
	fp.GradFill.Lin.AngAttr = gooxml.Int32(5400000)
	fp.GradFill.Lin.ScaledAttr = gooxml.Bool(false)

	gs := dml.NewCT_GradientStop()
	gs.PosAttr.ST_PositiveFixedPercentageDecimal = gooxml.Int32(0)
	gs.SchemeClr = &dml.CT_SchemeColor{ValAttr: dml.ST_SchemeColorValPhClr}
	fp.GradFill.GsLst.Gs = append(fp.GradFill.GsLst.Gs, gs)

	gs = dml.NewCT_GradientStop()
	gs.PosAttr.ST_PositiveFixedPercentageDecimal = gooxml.Int32(50000)
	gs.SchemeClr = &dml.CT_SchemeColor{ValAttr: dml.ST_SchemeColorValPhClr}
	fp.GradFill.GsLst.Gs = append(fp.GradFill.GsLst.Gs, gs)

	thm.ThemeElements.FmtScheme.LnStyleLst = dml.NewCT_LineStyleList()
	for i := 0; i < 3; i++ {
		lp := dml.NewCT_LineProperties()
		lp.WAttr = gooxml.Int32(int32(6350 * (i + 1)))
		lp.CapAttr = dml.ST_LineCapFlat
		lp.CmpdAttr = dml.ST_CompoundLineSng
		lp.AlgnAttr = dml.ST_PenAlignmentCtr
		thm.ThemeElements.FmtScheme.LnStyleLst.Ln = append(thm.ThemeElements.FmtScheme.LnStyleLst.Ln, lp)
	}

	thm.ThemeElements.FmtScheme.EffectStyleLst = dml.NewCT_EffectStyleList()
	for i := 0; i < 3; i++ {
		ef := dml.NewCT_EffectStyleItem()
		ef.EffectLst = dml.NewCT_EffectList()
		thm.ThemeElements.FmtScheme.EffectStyleLst.EffectStyle = append(thm.ThemeElements.FmtScheme.EffectStyleLst.EffectStyle, ef)
	}

	sf := dml.NewEG_FillProperties()
	sf.SolidFill = &dml.CT_SolidColorFillProperties{
		SchemeClr: &dml.CT_SchemeColor{ValAttr: dml.ST_SchemeColorValPhClr},
	}
	thm.ThemeElements.FmtScheme.BgFillStyleLst.EG_FillProperties = append(thm.ThemeElements.FmtScheme.BgFillStyleLst.EG_FillProperties,
		sf)
	thm.ThemeElements.FmtScheme.BgFillStyleLst.EG_FillProperties = append(thm.ThemeElements.FmtScheme.BgFillStyleLst.EG_FillProperties,
		sf)
	thm.ThemeElements.FmtScheme.BgFillStyleLst.EG_FillProperties = append(thm.ThemeElements.FmtScheme.BgFillStyleLst.EG_FillProperties,
		fp)

	p.themes = append(p.themes, thm)
	themeFn := gooxml.AbsoluteFilename(gooxml.DocTypePresentation, gooxml.ThemeType, 1)
	p.ContentTypes.AddOverride(themeFn, gooxml.ThemeContentType)
	p.prels.AddAutoRelationship(gooxml.DocTypePresentation, gooxml.OfficeDocumentType, 1, gooxml.ThemeType)

	thmRel := common.NewRelationships()
	p.themeRels = append(p.themeRels, thmRel)

	return p
}

// X returns the inner wrapped XML type.
func (p *Presentation) X() *pml.Presentation {
	return p.x
}

func (p *Presentation) nextSlideID() uint32 {
	id := uint32(256)
	for _, s := range p.x.SldIdLst.SldId {
		if s.IdAttr >= id {
			id = s.IdAttr + 1
		}
	}
	return id
}

// AddSlide adds a new slide to the presentation.
func (p *Presentation) AddSlide() Slide {
	sd := pml.NewCT_SlideIdListEntry()
	sd.IdAttr = p.nextSlideID()
	p.x.SldIdLst.SldId = append(p.x.SldIdLst.SldId, sd)

	slide := pml.NewSld()
	slide.CSld.SpTree.NvGrpSpPr.CNvPr.IdAttr = 1
	slide.CSld.SpTree.GrpSpPr.Xfrm = dml.NewCT_GroupTransform2D()
	slide.CSld.SpTree.GrpSpPr.Xfrm.Off = dml.NewCT_Point2D()
	slide.CSld.SpTree.GrpSpPr.Xfrm.Off.XAttr.ST_CoordinateUnqualified = gooxml.Int64(0)
	slide.CSld.SpTree.GrpSpPr.Xfrm.Off.YAttr.ST_CoordinateUnqualified = gooxml.Int64(0)
	slide.CSld.SpTree.GrpSpPr.Xfrm.Ext = dml.NewCT_PositiveSize2D()
	slide.CSld.SpTree.GrpSpPr.Xfrm.Ext.CxAttr = int64(0 * measurement.Point)
	slide.CSld.SpTree.GrpSpPr.Xfrm.Ext.CyAttr = int64(0 * measurement.Point)
	slide.CSld.SpTree.GrpSpPr.Xfrm.ChOff = slide.CSld.SpTree.GrpSpPr.Xfrm.Off
	slide.CSld.SpTree.GrpSpPr.Xfrm.ChExt = slide.CSld.SpTree.GrpSpPr.Xfrm.Ext

	p.slides = append(p.slides, slide)
	srelID := p.prels.AddAutoRelationship(gooxml.DocTypePresentation, gooxml.OfficeDocumentType,
		len(p.slides), gooxml.SlideType)
	sd.RIdAttr = srelID.ID()

	slidefn := gooxml.AbsoluteFilename(gooxml.DocTypePresentation, gooxml.SlideType, len(p.slides))
	p.ContentTypes.AddOverride(slidefn, gooxml.SlideContentType)

	srel := common.NewRelationships()
	p.slideRels = append(p.slideRels, srel)
	// TODO: make the slide layout configurable
	srel.AddAutoRelationship(gooxml.DocTypePresentation, gooxml.SlideType,
		len(p.layouts), gooxml.SlideLayoutType)

	return Slide{sd, slide, p}
}

// AddSlideWithLayout adds a new slide with content copied from a layout.  Normally you should
// use AddDefaultSlideWithLayout as it will do some post processing similar to PowerPoint to
// clear place holder text, etc.
func (p *Presentation) AddSlideWithLayout(l SlideLayout) (Slide, error) {
	sd := pml.NewCT_SlideIdListEntry()
	sd.IdAttr = 256
	for _, id := range p.x.SldIdLst.SldId {
		if id.IdAttr >= sd.IdAttr {
			sd.IdAttr = id.IdAttr + 1
		}
	}
	p.x.SldIdLst.SldId = append(p.x.SldIdLst.SldId, sd)

	slide := pml.NewSld()

	buf := bytes.Buffer{}
	enc := xml.NewEncoder(&buf)
	start := xml.StartElement{Name: xml.Name{Local: "slide"}}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/presentationml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:p"}, Value: "http://schemas.openxmlformats.org/presentationml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:sh"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})

	if err := l.x.CSld.MarshalXML(enc, start); err != nil {
		return Slide{}, err
	}
	enc.Flush()

	dec := xml.NewDecoder(&buf)
	slide.CSld = pml.NewCT_CommonSlideData()
	if err := dec.Decode(slide.CSld); err != nil {
		return Slide{}, err
	}

	// clear layout name on the slide
	slide.CSld.NameAttr = nil
	//, chc := range slide.CSld.SpTree.Choice {
	for i := 0; i < len(slide.CSld.SpTree.Choice); i++ {
		chc := slide.CSld.SpTree.Choice[i]
		for len(chc.Pic) > 0 {
			copy(slide.CSld.SpTree.Choice[i:], slide.CSld.SpTree.Choice[i+1:])
			slide.CSld.SpTree.Choice = slide.CSld.SpTree.Choice[0 : len(slide.CSld.SpTree.Choice)-1]
			chc = slide.CSld.SpTree.Choice[i]
		}
	}

	p.slides = append(p.slides, slide)

	srelID := p.prels.AddAutoRelationship(gooxml.DocTypePresentation, gooxml.OfficeDocumentType,
		len(p.slides), gooxml.SlideType)
	sd.RIdAttr = srelID.ID()

	slidefn := gooxml.AbsoluteFilename(gooxml.DocTypePresentation, gooxml.SlideType, len(p.slides))
	p.ContentTypes.AddOverride(slidefn, gooxml.SlideContentType)

	srel := common.NewRelationships()
	p.slideRels = append(p.slideRels, srel)
	for i, lout := range p.layouts {
		if lout == l.X() {
			srel.AddAutoRelationship(gooxml.DocTypePresentation, gooxml.SlideType,
				i+1, gooxml.SlideLayoutType)
		}
	}
	csld := Slide{sd, slide, p}

	return csld, nil
}

// AddDefaultSlideWithLayout tries to replicate what PowerPoint does when
// inserting a slide with a new style by clearing placeholder content and removing
// some placeholders.  Use AddSlideWithLayout if you need more control.
func (p *Presentation) AddDefaultSlideWithLayout(l SlideLayout) (Slide, error) {
	sld, err := p.AddSlideWithLayout(l)

	for _, ph := range sld.PlaceHolders() {
		// clear all placeholder content
		ph.Clear()
		// and drop some of the placeholders (footer, slide date/time, slide number)
		switch ph.Type() {
		case pml.ST_PlaceholderTypeFtr, pml.ST_PlaceholderTypeDt, pml.ST_PlaceholderTypeSldNum:
			ph.Remove()
		}
	}

	return sld, err
}

// Save writes the presentation out to a writer in the Zip package format
func (p *Presentation) Save(w io.Writer) error {
	if err := p.x.Validate(); err != nil {
		log.Printf("validation error in document: %s", err)
	}

	dt := gooxml.DocTypePresentation

	z := zip.NewWriter(w)
	defer z.Close()
	if err := zippkg.MarshalXML(z, gooxml.BaseRelsFilename, p.Rels.X()); err != nil {
		return err
	}
	if err := zippkg.MarshalXMLByType(z, dt, gooxml.ExtendedPropertiesType, p.AppProperties.X()); err != nil {
		return err
	}
	if err := zippkg.MarshalXMLByType(z, dt, gooxml.CorePropertiesType, p.CoreProperties.X()); err != nil {
		return err
	}
	if p.Thumbnail != nil {
		tn, err := z.Create("docProps/thumbnail.jpeg")
		if err != nil {
			return err
		}
		if err := jpeg.Encode(tn, p.Thumbnail, nil); err != nil {
			return err
		}
	}

	documentFn := gooxml.AbsoluteFilename(dt, gooxml.OfficeDocumentType, 0)
	if err := zippkg.MarshalXML(z, documentFn, p.x); err != nil {
		return err
	}
	if err := zippkg.MarshalXML(z, zippkg.RelationsPathFor(documentFn), p.prels.X()); err != nil {
		return err
	}

	for i, slide := range p.slides {
		spath := gooxml.AbsoluteFilename(gooxml.DocTypePresentation, gooxml.SlideType, i+1)
		zippkg.MarshalXML(z, spath, slide)
		if !p.slideRels[i].IsEmpty() {
			rpath := zippkg.RelationsPathFor(spath)
			zippkg.MarshalXML(z, rpath, p.slideRels[i].X())
		}
	}
	for i, m := range p.masters {
		mpath := gooxml.AbsoluteFilename(gooxml.DocTypePresentation, gooxml.SlideMasterType, i+1)
		zippkg.MarshalXML(z, mpath, m)
		if !p.masterRels[i].IsEmpty() {
			rpath := zippkg.RelationsPathFor(mpath)
			zippkg.MarshalXML(z, rpath, p.masterRels[i].X())
		}
	}
	for i, l := range p.layouts {
		mpath := gooxml.AbsoluteFilename(gooxml.DocTypePresentation, gooxml.SlideLayoutType, i+1)
		zippkg.MarshalXML(z, mpath, l)
		if !p.layoutRels[i].IsEmpty() {
			rpath := zippkg.RelationsPathFor(mpath)
			zippkg.MarshalXML(z, rpath, p.layoutRels[i].X())
		}
	}
	for i, l := range p.themes {
		mpath := gooxml.AbsoluteFilename(gooxml.DocTypePresentation, gooxml.ThemeType, i+1)
		zippkg.MarshalXML(z, mpath, l)
		if !p.themeRels[i].IsEmpty() {
			rpath := zippkg.RelationsPathFor(mpath)
			zippkg.MarshalXML(z, rpath, p.themeRels[i].X())
		}
	}

	for i, img := range p.Images {
		fn := gooxml.AbsoluteFilename(gooxml.DocTypePresentation, gooxml.ImageType, i+1)
		fn = fn[0:len(fn)-3] + strings.ToLower(img.Format())
		if img.Path() != "" {
			if err := zippkg.AddFileFromDisk(z, fn, img.Path()); err != nil {
				return err
			}
		} else {
			gooxml.Log("unsupported image source: %+v", img)
		}
	}

	if err := zippkg.MarshalXML(z, gooxml.ContentTypesFilename, p.ContentTypes.X()); err != nil {
		return err
	}
	if err := p.WriteExtraFiles(z); err != nil {
		return err
	}
	return nil
}

// SaveToFile writes the Presentation out to a file.
func (p *Presentation) SaveToFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return p.Save(f)
}

func (p *Presentation) Validate() error {
	if err := p.x.Validate(); err != nil {
		return err
	}
	for i, s := range p.Slides() {
		if err := s.ValidateWithPath(fmt.Sprintf("Slide[%d]", i)); err != nil {
			return err
		}
	}

	for i, sm := range p.masters {
		if err := sm.ValidateWithPath(fmt.Sprintf("SlideMaster[%d]", i)); err != nil {
			return err
		}
	}
	for i, sl := range p.layouts {
		if err := sl.ValidateWithPath(fmt.Sprintf("SlideLayout[%d]", i)); err != nil {
			return err
		}
	}
	return nil
}

// SlideMasters returns the slide masters defined in the presentation.
func (p *Presentation) SlideMasters() []SlideMaster {
	ret := []SlideMaster{}
	for i, m := range p.masters {

		ret = append(ret, SlideMaster{p, p.masterRels[i], m})
	}
	return ret
}

// SlideLayouts returns the slide layouts defined in the presentation.
func (p *Presentation) SlideLayouts() []SlideLayout {
	ret := []SlideLayout{}
	for _, l := range p.layouts {
		ret = append(ret, SlideLayout{l})
	}
	return ret
}

func (p *Presentation) onNewRelationship(decMap *zippkg.DecodeMap, target, typ string, files []*zip.File, rel *relationships.Relationship, src zippkg.Target) error {
	dt := gooxml.DocTypePresentation

	switch typ {
	case gooxml.OfficeDocumentType:
		p.x = pml.NewPresentation()
		decMap.AddTarget(target, p.x, typ, 0)
		decMap.AddTarget(zippkg.RelationsPathFor(target), p.prels.X(), typ, 0)
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, 0)

	case gooxml.CorePropertiesType:
		decMap.AddTarget(target, p.CoreProperties.X(), typ, 0)
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, 0)

	case gooxml.ExtendedPropertiesType:
		decMap.AddTarget(target, p.AppProperties.X(), typ, 0)
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, 0)

	case gooxml.SlideType:
		sld := pml.NewSld()
		p.slides = append(p.slides, sld)
		decMap.AddTarget(target, sld, typ, uint32(len(p.slides)))
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, len(p.slides))

		slRel := common.NewRelationships()
		decMap.AddTarget(zippkg.RelationsPathFor(target), slRel.X(), typ, 0)
		p.slideRels = append(p.slideRels, slRel)

	case gooxml.SlideMasterType:
		sm := pml.NewSldMaster()
		if !decMap.AddTarget(target, sm, typ, uint32(len(p.masters)+1)) {
			return nil
		}
		p.masters = append(p.masters, sm)
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, len(p.masters))

		// look for master rels
		smRel := common.NewRelationships()
		decMap.AddTarget(zippkg.RelationsPathFor(target), smRel.X(), typ, 0)
		p.masterRels = append(p.masterRels, smRel)

	case gooxml.SlideLayoutType:
		sl := pml.NewSldLayout()
		if !decMap.AddTarget(target, sl, typ, uint32(len(p.layouts)+1)) {
			return nil
		}
		p.layouts = append(p.layouts, sl)
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, len(p.layouts))

		// look for layout rels
		slRel := common.NewRelationships()
		decMap.AddTarget(zippkg.RelationsPathFor(target), slRel.X(), typ, 0)
		p.layoutRels = append(p.layoutRels, slRel)

	case gooxml.ThumbnailType:
		// read our thumbnail
		for i, f := range files {
			if f == nil {
				continue
			}
			if f.Name == target {
				rc, err := f.Open()
				if err != nil {
					return fmt.Errorf("error reading thumbnail: %s", err)
				}
				p.Thumbnail, _, err = image.Decode(rc)
				rc.Close()
				if err != nil {
					return fmt.Errorf("error decoding thumbnail: %s", err)
				}
				files[i] = nil
			}
		}

	case gooxml.ThemeType:
		thm := dml.NewTheme()
		if !decMap.AddTarget(target, thm, typ, uint32(len(p.themes)+1)) {
			return nil
		}
		p.themes = append(p.themes, thm)
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, len(p.themes))

		// look for theme rels
		thmRel := common.NewRelationships()
		decMap.AddTarget(zippkg.RelationsPathFor(target), thmRel.X(), typ, 0)
		p.themeRels = append(p.themeRels, thmRel)

	case gooxml.ImageType:
		// we use path.Clean instead of filepath.Clean to ensure we
		// end up with forward separators
		target = path.Clean(target)
		for i, f := range files {
			if f == nil {
				continue
			}
			if f.Name == target {
				path, err := zippkg.ExtractToDiskTmp(f, p.TmpPath)
				if err != nil {
					return err
				}
				img, err := common.ImageFromFile(path)
				if err != nil {
					return err
				}
				iref := common.MakeImageRef(img, &p.DocBase, p.prels)
				p.Images = append(p.Images, iref)
				files[i] = nil
				decMap.RecordIndex(target, len(p.Images))
				break
			}
		}
		idx := decMap.IndexFor(target)
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, idx)

	default:
		gooxml.Log("unsupported relationship type: %s tgt: %s", typ, target)
	}
	return nil
}

// Slides returns the slides in the presentation.
func (p *Presentation) Slides() []Slide {
	ret := []Slide{}
	for i, v := range p.slides {
		ret = append(ret, Slide{p.x.SldIdLst.SldId[i], v, p})
	}
	return ret
}

// RemoveSlide removes a slide from a presentation.
func (p *Presentation) RemoveSlide(s Slide) error {
	removed := false
	slideIdx := 0
	for i, v := range p.slides {
		if v == s.x {
			if p.x.SldIdLst.SldId[i] != s.sid {
				return errors.New("inconsistency in slides and ID list")
			}
			copy(p.slides[i:], p.slides[i+1:])
			p.slides = p.slides[0 : len(p.slides)-1]

			copy(p.slideRels[i:], p.slideRels[i+1:])
			p.slideRels = p.slideRels[0 : len(p.slideRels)-1]

			copy(p.x.SldIdLst.SldId[i:], p.x.SldIdLst.SldId[i+1:])
			p.x.SldIdLst.SldId = p.x.SldIdLst.SldId[0 : len(p.x.SldIdLst.SldId)-1]

			removed = true
			slideIdx = i
		}
	}

	if !removed {
		return errors.New("unable to find slide")
	}

	// remove it from content types
	fn := gooxml.AbsoluteFilename(gooxml.DocTypePresentation, gooxml.SlideType, slideIdx+1)
	p.ContentTypes.RemoveOverride(fn)
	return nil
}

// GetLayoutByName retrieves a slide layout given a layout name.
func (p *Presentation) GetLayoutByName(name string) (SlideLayout, error) {
	for _, l := range p.layouts {
		if l.CSld.NameAttr != nil && name == *l.CSld.NameAttr {
			return SlideLayout{l}, nil
		}
	}
	return SlideLayout{}, errors.New("unable to find layout with that name")
}

// AddImage adds an image to the document package, returning a reference that
// can be used to add the image to a run and place it in the document contents.
func (p *Presentation) AddImage(i common.Image) (common.ImageRef, error) {
	r := common.MakeImageRef(i, &p.DocBase, p.prels)
	if i.Path == "" {
		return r, errors.New("image must have a path")
	}

	if i.Format == "" {
		return r, errors.New("image must have a valid format")
	}
	if i.Size.X == 0 || i.Size.Y == 0 {
		return r, errors.New("image must have a valid size")
	}

	p.Images = append(p.Images, r)
	fn := fmt.Sprintf("media/image%d.%s", len(p.Images), i.Format)
	p.prels.AddRelationship(fn, gooxml.ImageType)
	return r, nil
}

// GetImageByRelID returns an ImageRef with the associated relation ID in the
// document.
func (p *Presentation) GetImageByRelID(relID string) (common.ImageRef, bool) {
	for _, img := range p.Images {
		if img.RelID() == relID {
			return img, true
		}
	}
	return common.ImageRef{}, false
}
