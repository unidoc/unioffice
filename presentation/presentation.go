// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentation

import (
	"archive/zip"
	"fmt"
	"io"
	"os"

	"baliance.com/gooxml"
	"baliance.com/gooxml/common"
	"baliance.com/gooxml/measurement"
	"baliance.com/gooxml/schema/soo/dml"
	"baliance.com/gooxml/schema/soo/ofc/sharedTypes"
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
}

// New initializes and reurns a new presentation
func New() *Presentation {
	p := &Presentation{x: pml.NewPresentation()}

	p.x.SldIdLst = pml.NewCT_SlideIdList()
	p.x.ConformanceAttr = sharedTypes.ST_ConformanceClassTransitional
	p.AppProperties = common.NewAppProperties()
	p.CoreProperties = common.NewCoreProperties()
	p.ContentTypes = common.NewContentTypes()

	p.ContentTypes.AddOverride("/ppt/presentation.xml", "application/vnd.openxmlformats-officedocument.presentationml.presentation.main+xml")

	p.Rels = common.NewRelationships()

	p.Rels.AddRelationship("docProps/core.xml", "http://schemas.openxmlformats.org/package/2006/relationships/metadata/core-properties")
	p.Rels.AddRelationship("docProps/app.xml", "http://schemas.openxmlformats.org/officeDocument/2006/relationships/extended-properties")
	p.Rels.AddRelationship("ppt/presentation.xml", "http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument")

	p.prels = common.NewRelationships()

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

	p.ContentTypes.AddOverride("/ppt/slideMasters/slideMaster1.xml", gooxml.SlideMasterContentType)
	mrelID := p.prels.AddRelationship("slideMasters/slideMaster1.xml", gooxml.SlideMasterRelationshipType)
	smid := pml.NewCT_SlideMasterIdListEntry()
	smid.IdAttr = gooxml.Uint32(2147483648)
	smid.RIdAttr = mrelID.ID()
	p.x.SldMasterIdLst.SldMasterId = append(p.x.SldMasterIdLst.SldMasterId, smid)
	mrel := common.NewRelationships()
	p.masterRels = append(p.masterRels, mrel)

	ls := pml.NewSldLayout()
	lrid := mrel.AddRelationship("../slideLayouts/slideLayout1.xml", gooxml.SlideLayoutType)
	p.ContentTypes.AddOverride("/ppt/slideLayouts/slideLayout1.xml", gooxml.SlideLayoutContentType)
	mrel.AddRelationship("../theme/theme1.xml", gooxml.ThemeType)
	p.layouts = append(p.layouts, ls)

	m.SldLayoutIdLst = pml.NewCT_SlideLayoutIdList()
	lid := pml.NewCT_SlideLayoutIdListEntry()
	lid.IdAttr = gooxml.Uint32(2147483649)
	lid.RIdAttr = lrid.ID()
	m.SldLayoutIdLst.SldLayoutId = append(m.SldLayoutIdLst.SldLayoutId, lid)

	lrel := common.NewRelationships()
	p.layoutRels = append(p.layoutRels, lrel)
	lrel.AddRelationship("../slideMasters/slideMaster1.xml", gooxml.SlideMasterRelationshipType)
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
	p.ContentTypes.AddOverride("/ppt/theme/theme1.xml", gooxml.ThemeContentType)
	p.prels.AddRelationship("theme/theme1.xml", gooxml.ThemeType)
	return p
}

// X returns the inner wrapped XML type.
func (p *Presentation) X() *pml.Presentation {
	return p.x
}

func (p *Presentation) AddSlide() Slide {
	sd := pml.NewCT_SlideIdListEntry()
	sd.IdAttr = 256
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

	c := pml.NewCT_GroupShapeChoice()
	slide.CSld.SpTree.Choice = append(slide.CSld.SpTree.Choice, c)
	sp := pml.NewCT_Shape()
	c.Sp = append(c.Sp, sp)

	sp.NvSpPr.NvPr.Ph = pml.NewCT_Placeholder()
	sp.NvSpPr.NvPr.Ph.TypeAttr = pml.ST_PlaceholderTypeCtrTitle

	sp.TxBody = dml.NewCT_TextBody()
	para := dml.NewCT_TextParagraph()
	sp.TxBody.P = append(sp.TxBody.P, para)

	run := dml.NewEG_TextRun()
	para.EG_TextRun = append(para.EG_TextRun, run)
	run.R = dml.NewCT_RegularTextRun()
	run.R.T = "testing 123"

	p.slides = append(p.slides, slide)
	fn := fmt.Sprintf("slides/slide%d.xml", len(p.slides))
	srelID := p.prels.AddRelationship(fn, gooxml.SlideType)
	sd.RIdAttr = srelID.ID()

	p.ContentTypes.AddOverride(fmt.Sprintf("/ppt/slides/slide%d.xml", len(p.slides)), gooxml.SlideContentType)

	srel := common.NewRelationships()
	p.slideRels = append(p.slideRels, srel)
	srel.AddRelationship("../slideLayouts/slideLayout1.xml", gooxml.SlideLayoutType)

	return Slide{sd, slide}
}

// Save writes the presentation out to a writer in the Zip package format
func (p *Presentation) Save(w io.Writer) error {
	z := zip.NewWriter(w)
	defer z.Close()
	if err := zippkg.MarshalXML(z, gooxml.ContentTypesFilename, p.ContentTypes.X()); err != nil {
		return err
	}
	if err := zippkg.MarshalXML(z, "_rels/.rels", p.Rels.X()); err != nil {
		return err
	}
	if err := zippkg.MarshalXML(z, "docProps/app.xml", p.AppProperties.X()); err != nil {
		return err
	}
	if err := zippkg.MarshalXML(z, "docProps/core.xml", p.CoreProperties.X()); err != nil {
		return err
	}
	if err := zippkg.MarshalXML(z, "ppt/presentation.xml", p.x); err != nil {
		return err
	}
	if err := zippkg.MarshalXML(z, "ppt/_rels/presentation.xml.rels", p.prels.X()); err != nil {
		return err
	}
	for i, slide := range p.slides {
		spath := fmt.Sprintf("ppt/slides/slide%d.xml", i+1)
		zippkg.MarshalXML(z, spath, slide)
		rpath := zippkg.RelationsPathFor(spath)
		zippkg.MarshalXML(z, rpath, p.slideRels[i].X())
	}
	for i, m := range p.masters {
		mpath := fmt.Sprintf("ppt/slideMasters/slideMaster%d.xml", i+1)
		zippkg.MarshalXML(z, mpath, m)
		rpath := zippkg.RelationsPathFor(mpath)
		zippkg.MarshalXML(z, rpath, p.masterRels[i].X())
	}
	for i, l := range p.layouts {
		mpath := fmt.Sprintf("ppt/slideLayouts/slideLayout%d.xml", i+1)
		zippkg.MarshalXML(z, mpath, l)
		rpath := zippkg.RelationsPathFor(mpath)
		zippkg.MarshalXML(z, rpath, p.layoutRels[i].X())
	}
	for i, l := range p.themes {
		tpath := fmt.Sprintf("ppt/theme/theme%d.xml", i+1)
		zippkg.MarshalXML(z, tpath, l)
	}
	p.WriteExtraFiles(z)
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
	for i, s := range p.slides {
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
