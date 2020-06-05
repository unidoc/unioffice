// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package document

import (
	"archive/zip"
	"errors"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/common"
	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/zippkg"

	"github.com/unidoc/unioffice/schema/soo/dml"
	st "github.com/unidoc/unioffice/schema/soo/ofc/sharedTypes"
	"github.com/unidoc/unioffice/schema/soo/pkg/relationships"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

// Document is a text document that can be written out in the OOXML .docx
// format. It can be opened from a file on disk and modified, or created from
// scratch.
type Document struct {
	common.DocBase
	x *wml.Document

	Settings  Settings  // document settings
	Numbering Numbering // numbering styles within the doucment
	Styles    Styles    // styles that are use and can be used within the document

	headers []*wml.Hdr
	hdrRels []common.Relationships

	footers []*wml.Ftr
	ftrRels []common.Relationships

	docRels     common.Relationships
	themes      []*dml.Theme
	webSettings *wml.WebSettings
	fontTable   *wml.Fonts
	endNotes    *wml.Endnotes
	footNotes   *wml.Footnotes
}

// New constructs an empty document that content can be added to.
func New() *Document {
	d := &Document{x: wml.NewDocument()}
	d.ContentTypes = common.NewContentTypes()
	d.x.Body = wml.NewCT_Body()
	d.x.ConformanceAttr = st.ST_ConformanceClassTransitional
	d.docRels = common.NewRelationships()

	d.AppProperties = common.NewAppProperties()
	d.CoreProperties = common.NewCoreProperties()

	d.ContentTypes.AddOverride("/word/document.xml", "application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml")

	d.Settings = NewSettings()
	d.docRels.AddRelationship("settings.xml", unioffice.SettingsType)
	d.ContentTypes.AddOverride("/word/settings.xml", "application/vnd.openxmlformats-officedocument.wordprocessingml.settings+xml")

	d.Rels = common.NewRelationships()
	d.Rels.AddRelationship(unioffice.RelativeFilename(unioffice.DocTypeDocument, "", unioffice.CorePropertiesType, 0), unioffice.CorePropertiesType)
	d.Rels.AddRelationship("docProps/app.xml", unioffice.ExtendedPropertiesType)
	d.Rels.AddRelationship("word/document.xml", unioffice.OfficeDocumentType)

	d.Numbering = NewNumbering()
	d.Numbering.InitializeDefault()
	d.ContentTypes.AddOverride("/word/numbering.xml", "application/vnd.openxmlformats-officedocument.wordprocessingml.numbering+xml")
	d.docRels.AddRelationship("numbering.xml", unioffice.NumberingType)

	d.Styles = NewStyles()
	d.Styles.InitializeDefault()
	d.ContentTypes.AddOverride("/word/styles.xml", "application/vnd.openxmlformats-officedocument.wordprocessingml.styles+xml")
	d.docRels.AddRelationship("styles.xml", unioffice.StylesType)

	d.x.Body = wml.NewCT_Body()
	return d
}

// GetOrCreateCustomProperties returns the custom properties of the document (and if they not exist yet, creating them first)
func (d *Document) GetOrCreateCustomProperties() common.CustomProperties {
	if d.CustomProperties.X() == nil {
		d.createCustomProperties()
	}
	return d.CustomProperties
}

func (d *Document) createCustomProperties() {
	d.CustomProperties = common.NewCustomProperties()
	d.addCustomRelationships()
}

func (d *Document) addCustomRelationships() {
	d.ContentTypes.AddOverride("/docProps/custom.xml", "application/vnd.openxmlformats-officedocument.custom-properties+xml")
	d.Rels.AddRelationship("docProps/custom.xml", unioffice.CustomPropertiesType)
}

// X returns the inner wrapped XML type.
func (d *Document) X() *wml.Document {
	return d.x
}

// AddHeader creates a header associated with the document, but doesn't add it
// to the document for display.
func (d *Document) AddHeader() Header {
	hdr := wml.NewHdr()
	d.headers = append(d.headers, hdr)
	path := fmt.Sprintf("header%d.xml", len(d.headers))
	d.docRels.AddRelationship(path, unioffice.HeaderType)

	d.ContentTypes.AddOverride("/word/"+path, "application/vnd.openxmlformats-officedocument.wordprocessingml.header+xml")
	d.hdrRels = append(d.hdrRels, common.NewRelationships())

	return Header{d, hdr}
}

// Headers returns the headers defined in the document.
func (d *Document) Headers() []Header {
	ret := []Header{}
	for _, h := range d.headers {
		ret = append(ret, Header{d, h})
	}
	return ret
}

// Footers returns the footers defined in the document.
func (d *Document) Footers() []Footer {
	ret := []Footer{}
	for _, f := range d.footers {
		ret = append(ret, Footer{d, f})
	}
	return ret
}

// AddFooter creates a Footer associated with the document, but doesn't add it
// to the document for display.
func (d *Document) AddFooter() Footer {
	ftr := wml.NewFtr()
	d.footers = append(d.footers, ftr)
	path := fmt.Sprintf("footer%d.xml", len(d.footers))
	d.docRels.AddRelationship(path, unioffice.FooterType)

	d.ContentTypes.AddOverride("/word/"+path, "application/vnd.openxmlformats-officedocument.wordprocessingml.footer+xml")
	d.ftrRels = append(d.ftrRels, common.NewRelationships())

	return Footer{d, ftr}
}

// BodySection returns the default body section used for all preceding
// paragraphs until the previous Section. If there is no previous sections, the
// body section applies to the entire document.
func (d *Document) BodySection() Section {
	if d.x.Body.SectPr == nil {
		d.x.Body.SectPr = wml.NewCT_SectPr()
	}
	return Section{d, d.x.Body.SectPr}
}

// Save writes the document to an io.Writer in the Zip package format.
func (d *Document) Save(w io.Writer) error {
	if err := d.x.Validate(); err != nil {
		unioffice.Log("validation error in document: %s", err)
	}
	dt := unioffice.DocTypeDocument

	if !license.GetLicenseKey().IsLicensed() && flag.Lookup("test.v") == nil {
		fmt.Println("Unlicensed version of UniOffice")
		fmt.Println("- Get a license on https://unidoc.io")
		hdr := d.AddHeader()
		para := hdr.AddParagraph()
		para.Properties().AddTabStop(2.5*measurement.Inch, wml.ST_TabJcCenter, wml.ST_TabTlcNone)
		run := para.AddRun()
		run.AddTab()
		run.AddText("Unlicensed version of UniOffice - Get a license on https://unidoc.io")
		run.Properties().SetBold(true)
		run.Properties().SetSize(14)
		run.Properties().SetColor(color.Red)
		d.BodySection().SetHeader(hdr, wml.ST_HdrFtrDefault)
	}

	z := zip.NewWriter(w)
	defer z.Close()
	if err := zippkg.MarshalXML(z, unioffice.BaseRelsFilename, d.Rels.X()); err != nil {
		return err
	}
	if err := zippkg.MarshalXMLByType(z, dt, unioffice.ExtendedPropertiesType, d.AppProperties.X()); err != nil {
		return err
	}
	if err := zippkg.MarshalXMLByType(z, dt, unioffice.CorePropertiesType, d.CoreProperties.X()); err != nil {
		return err
	}
	if d.CustomProperties.X() != nil {
		if err := zippkg.MarshalXMLByType(z, dt, unioffice.CustomPropertiesType, d.CustomProperties.X()); err != nil {
			return err
		}
	}
	if d.Thumbnail != nil {
		tn, err := z.Create("docProps/thumbnail.jpeg")
		if err != nil {
			return err
		}
		if err := jpeg.Encode(tn, d.Thumbnail, nil); err != nil {
			return err
		}
	}
	if err := zippkg.MarshalXMLByType(z, dt, unioffice.SettingsType, d.Settings.X()); err != nil {
		return err
	}
	documentFn := unioffice.AbsoluteFilename(dt, unioffice.OfficeDocumentType, 0)
	if err := zippkg.MarshalXML(z, documentFn, d.x); err != nil {
		return err
	}
	if err := zippkg.MarshalXML(z, zippkg.RelationsPathFor(documentFn), d.docRels.X()); err != nil {
		return err
	}

	if d.Numbering.X() != nil {
		if err := zippkg.MarshalXMLByType(z, dt, unioffice.NumberingType, d.Numbering.X()); err != nil {
			return err
		}
	}
	if err := zippkg.MarshalXMLByType(z, dt, unioffice.StylesType, d.Styles.X()); err != nil {
		return err
	}

	if d.webSettings != nil {
		if err := zippkg.MarshalXMLByType(z, dt, unioffice.WebSettingsType, d.webSettings); err != nil {
			return err
		}
	}
	if d.fontTable != nil {
		if err := zippkg.MarshalXMLByType(z, dt, unioffice.FontTableType, d.fontTable); err != nil {
			return err
		}
	}
	if d.endNotes != nil {
		if err := zippkg.MarshalXMLByType(z, dt, unioffice.EndNotesType, d.endNotes); err != nil {
			return err
		}
	}
	if d.footNotes != nil {
		if err := zippkg.MarshalXMLByType(z, dt, unioffice.FootNotesType, d.footNotes); err != nil {
			return err
		}
	}
	for i, thm := range d.themes {
		if err := zippkg.MarshalXMLByTypeIndex(z, dt, unioffice.ThemeType, i+1, thm); err != nil {
			return err
		}
	}
	for i, hdr := range d.headers {
		fn := unioffice.AbsoluteFilename(dt, unioffice.HeaderType, i+1)
		if err := zippkg.MarshalXML(z, fn, hdr); err != nil {
			return err
		}
		if !d.hdrRels[i].IsEmpty() {
			zippkg.MarshalXML(z, zippkg.RelationsPathFor(fn), d.hdrRels[i].X())
		}
	}
	for i, ftr := range d.footers {
		fn := unioffice.AbsoluteFilename(dt, unioffice.FooterType, i+1)
		if err := zippkg.MarshalXMLByTypeIndex(z, dt, unioffice.FooterType, i+1, ftr); err != nil {
			return err
		}
		if !d.ftrRels[i].IsEmpty() {
			zippkg.MarshalXML(z, zippkg.RelationsPathFor(fn), d.ftrRels[i].X())
		}
	}

	for i, img := range d.Images {
		if err := common.AddImageToZip(z, img, i+1, unioffice.DocTypeDocument); err != nil {
			return err
		}
	}

	if err := zippkg.MarshalXML(z, unioffice.ContentTypesFilename, d.ContentTypes.X()); err != nil {
		return err
	}
	if err := d.WriteExtraFiles(z); err != nil {
		return err
	}
	return z.Close()
}

// AddTable adds a new table to the document body.
func (d *Document) AddTable() Table {
	elts := wml.NewEG_BlockLevelElts()
	d.x.Body.EG_BlockLevelElts = append(d.x.Body.EG_BlockLevelElts, elts)
	c := wml.NewEG_ContentBlockContent()
	elts.EG_ContentBlockContent = append(elts.EG_ContentBlockContent, c)
	tbl := wml.NewCT_Tbl()
	c.Tbl = append(c.Tbl, tbl)
	return Table{d, tbl}
}

func (d *Document) InsertTableAfter(relativeTo Paragraph) Table {
	return d.insertTable(relativeTo, false)
}

func (d *Document) InsertTableBefore(relativeTo Paragraph) Table {
	return d.insertTable(relativeTo, true)
}

func (d *Document) insertTable(relativeTo Paragraph, before bool) Table {
	body := d.x.Body
	if body == nil {
		return d.AddTable()
	}
	relX := relativeTo.X()
	for i, ble := range body.EG_BlockLevelElts {
		for _, c := range ble.EG_ContentBlockContent {
			for j, p := range c.P {
				// found the paragraph
				if p == relX {
					tbl := wml.NewCT_Tbl()
					elts := wml.NewEG_BlockLevelElts()
					cbc := wml.NewEG_ContentBlockContent()
					elts.EG_ContentBlockContent = append(elts.EG_ContentBlockContent, cbc)
					cbc.Tbl = append(cbc.Tbl, tbl)
					body.EG_BlockLevelElts = append(body.EG_BlockLevelElts, nil)
					if before {
						copy(body.EG_BlockLevelElts[i+1:], body.EG_BlockLevelElts[i:])
						body.EG_BlockLevelElts[i] = elts
						if j != 0 {
							elts := wml.NewEG_BlockLevelElts()
							cbc := wml.NewEG_ContentBlockContent()
							elts.EG_ContentBlockContent = append(elts.EG_ContentBlockContent, cbc)
							cbc.P = c.P[:j]
							body.EG_BlockLevelElts = append(body.EG_BlockLevelElts, nil)
							copy(body.EG_BlockLevelElts[i+1:], body.EG_BlockLevelElts[i:])
							body.EG_BlockLevelElts[i] = elts
						}
						c.P = c.P[j:]
					} else {
						copy(body.EG_BlockLevelElts[i+2:], body.EG_BlockLevelElts[i+1:])
						body.EG_BlockLevelElts[i+1] = elts
						if j != len(c.P)-1 {
							elts := wml.NewEG_BlockLevelElts()
							cbc := wml.NewEG_ContentBlockContent()
							elts.EG_ContentBlockContent = append(elts.EG_ContentBlockContent, cbc)
							cbc.P = c.P[j+1:]
							body.EG_BlockLevelElts = append(body.EG_BlockLevelElts, nil)
							copy(body.EG_BlockLevelElts[i+3:], body.EG_BlockLevelElts[i+2:])
							body.EG_BlockLevelElts[i+2] = elts
						}
						c.P = c.P[:j+1]
					}
					return Table{d, tbl}
				}
			}

			for _, tbl := range c.Tbl {
				for _, crc := range tbl.EG_ContentRowContent {
					for _, tr := range crc.Tr {
						for _, ccc := range tr.EG_ContentCellContent {
							for _, tc := range ccc.Tc {
								for i, ble := range tc.EG_BlockLevelElts {
									for _, cbcOuter := range ble.EG_ContentBlockContent {
										for j, p := range cbcOuter.P {
											if p == relX {
												elts := wml.NewEG_BlockLevelElts()
												cbcInner := wml.NewEG_ContentBlockContent()
												elts.EG_ContentBlockContent = append(elts.EG_ContentBlockContent, cbcInner)
												tbl := wml.NewCT_Tbl()
												cbcInner.Tbl = append(cbcInner.Tbl, tbl)
												tc.EG_BlockLevelElts = append(tc.EG_BlockLevelElts, nil)
												if before {
													copy(tc.EG_BlockLevelElts[i+1:], tc.EG_BlockLevelElts[i:])
													tc.EG_BlockLevelElts[i] = elts
													if j != 0 {
														elts := wml.NewEG_BlockLevelElts()
														cbc := wml.NewEG_ContentBlockContent()
														elts.EG_ContentBlockContent = append(elts.EG_ContentBlockContent, cbc)
														cbc.P = cbcOuter.P[:j]
														tc.EG_BlockLevelElts = append(tc.EG_BlockLevelElts, nil)
														copy(tc.EG_BlockLevelElts[i+1:], tc.EG_BlockLevelElts[i:])
														tc.EG_BlockLevelElts[i] = elts
													}
													cbcOuter.P = cbcOuter.P[j:]
												} else {
													copy(tc.EG_BlockLevelElts[i+2:], tc.EG_BlockLevelElts[i+1:])
													tc.EG_BlockLevelElts[i+1] = elts
													if j != len(c.P)-1 {
														elts := wml.NewEG_BlockLevelElts()
														cbc := wml.NewEG_ContentBlockContent()
														elts.EG_ContentBlockContent = append(elts.EG_ContentBlockContent, cbc)
														cbc.P = cbcOuter.P[j+1:]
														tc.EG_BlockLevelElts = append(tc.EG_BlockLevelElts, nil)
														copy(tc.EG_BlockLevelElts[i+3:], tc.EG_BlockLevelElts[i+2:])
														tc.EG_BlockLevelElts[i+2] = elts
													}
													cbcOuter.P = cbcOuter.P[:j+1]
												}
												return Table{d, tbl}
											}
										}
									}
								}
							}
						}
					}
				}
			}

		}
	}
	return d.AddTable()
}

func (d *Document) tables(bc *wml.EG_ContentBlockContent) []Table {
	ret := []Table{}
	for _, t := range bc.Tbl {
		ret = append(ret, Table{d, t})
		for _, crc := range t.EG_ContentRowContent {
			for _, tr := range crc.Tr {
				for _, ccc := range tr.EG_ContentCellContent {
					for _, tc := range ccc.Tc {
						for _, ble := range tc.EG_BlockLevelElts {
							for _, cbc := range ble.EG_ContentBlockContent {
								for _, tbl := range d.tables(cbc) {
									ret = append(ret, tbl)
								}
							}
						}
					}
				}
			}
		}
	}

	return ret
}

// Tables returns the tables defined in the document.
func (d *Document) Tables() []Table {
	ret := []Table{}
	if d.x.Body == nil {
		return nil
	}
	for _, ble := range d.x.Body.EG_BlockLevelElts {
		for _, c := range ble.EG_ContentBlockContent {
			for _, t := range d.tables(c) {
				ret = append(ret, t)
			}
		}
	}
	return ret
}

// AddParagraph adds a new paragraph to the document body.
func (d *Document) AddParagraph() Paragraph {
	elts := wml.NewEG_BlockLevelElts()
	d.x.Body.EG_BlockLevelElts = append(d.x.Body.EG_BlockLevelElts, elts)
	c := wml.NewEG_ContentBlockContent()
	elts.EG_ContentBlockContent = append(elts.EG_ContentBlockContent, c)
	p := wml.NewCT_P()
	c.P = append(c.P, p)
	return Paragraph{d, p}
}

// RemoveParagraph removes a paragraph from a document.
func (d *Document) RemoveParagraph(p Paragraph) {
	if d.x.Body == nil {
		return
	}

	for _, ble := range d.x.Body.EG_BlockLevelElts {
		for _, c := range ble.EG_ContentBlockContent {
			for i, pa := range c.P {
				// do we need to remove this paragraph
				if pa == p.x {
					copy(c.P[i:], c.P[i+1:])
					c.P = c.P[0 : len(c.P)-1]
					return
				}
			}

			if c.Sdt != nil && c.Sdt.SdtContent != nil && c.Sdt.SdtContent.P != nil {
				for i, pa := range c.Sdt.SdtContent.P {
					if pa == p.x {
						copy(c.P[i:], c.P[i+1:])
						c.P = c.P[0 : len(c.P)-1]
						return
					}
				}
			}
		}
	}
}

// StructuredDocumentTags returns the structured document tags in the document
// which are commonly used in document templates.
func (d *Document) StructuredDocumentTags() []StructuredDocumentTag {
	ret := []StructuredDocumentTag{}
	for _, ble := range d.x.Body.EG_BlockLevelElts {
		for _, c := range ble.EG_ContentBlockContent {
			if c.Sdt != nil {
				ret = append(ret, StructuredDocumentTag{d, c.Sdt})
			}
		}
	}
	return ret
}

// Paragraphs returns all of the paragraphs in the document body including tables.
func (d *Document) Paragraphs() []Paragraph {
	ret := []Paragraph{}
	if d.x.Body == nil {
		return nil
	}
	for _, ble := range d.x.Body.EG_BlockLevelElts {
		for _, c := range ble.EG_ContentBlockContent {
			for _, p := range c.P {
				ret = append(ret, Paragraph{d, p})
			}
		}
	}

	for _, t := range d.Tables() {
		for _, r := range t.Rows() {
			for _, c := range r.Cells() {
				ret = append(ret, c.Paragraphs()...)
			}
		}
	}
	return ret
}

// HasFootnotes returns a bool based on the presence or abscence of footnotes within
// the document.
func (d *Document) HasFootnotes() bool {
	return d.footNotes != nil
}

// Footnotes returns the footnotes defined in the document.
func (d *Document) Footnotes() []Footnote {
	ret := []Footnote{}
	for _, f := range d.footNotes.CT_Footnotes.Footnote {
		ret = append(ret, Footnote{d, f})
	}
	return ret
}

// Footnote returns the footnote based on the ID; this can be used nicely with
// the run.IsFootnote() functionality.
func (d *Document) Footnote(id int64) Footnote {
	for _, f := range d.Footnotes() {
		if f.id() == id {
			return f
		}
	}
	return Footnote{}
}

// HasEndnotes returns a bool based on the presence or abscence of endnotes within
// the document.
func (d *Document) HasEndnotes() bool {
	return d.endNotes != nil
}

// Endnotes returns the endnotes defined in the document.
func (d *Document) Endnotes() []Endnote {
	ret := []Endnote{}
	for _, f := range d.endNotes.CT_Endnotes.Endnote {
		ret = append(ret, Endnote{d, f})
	}
	return ret
}

// Endnote returns the endnote based on the ID; this can be used nicely with
// the run.IsEndnote() functionality.
func (d *Document) Endnote(id int64) Endnote {
	for _, f := range d.Endnotes() {
		if f.id() == id {
			return f
		}
	}
	return Endnote{}
}

// SaveToFile writes the document out to a file.
func (d *Document) SaveToFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return d.Save(f)
}

// Open opens and reads a document from a file (.docx).
func Open(filename string) (*Document, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening %s: %s", filename, err)
	}
	defer f.Close()
	fi, err := os.Stat(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening %s: %s", filename, err)
	}
	_ = fi
	return Read(f, fi.Size())
}

// OpenTemplate opens a document, removing all content so it can be used as a
// template.  Since Word removes unused styles from a document upon save, to
// create a template in Word add a paragraph with every style of interest.  When
// opened with OpenTemplate the document's styles will be available but the
// content will be gone.
func OpenTemplate(filename string) (*Document, error) {
	d, err := Open(filename)
	if err != nil {
		return nil, err
	}
	d.x.Body = wml.NewCT_Body()
	return d, nil
}

// Read reads a document from an io.Reader.
func Read(r io.ReaderAt, size int64) (*Document, error) {
	doc := New()
	// numbering is not required
	doc.Numbering.x = nil

	td, err := ioutil.TempDir("", "gooxml-docx")
	if err != nil {
		return nil, err
	}
	doc.TmpPath = td

	zr, err := zip.NewReader(r, size)
	if err != nil {
		return nil, fmt.Errorf("parsing zip: %s", err)
	}

	files := []*zip.File{}
	files = append(files, zr.File...)

	addCustom := false
	for _, f := range files {
		if f.FileHeader.Name == "docProps/custom.xml" {
			addCustom = true
			break
		}
	}
	if addCustom {
		doc.createCustomProperties()
	}

	ca := doc.x.ConformanceAttr
	decMap := zippkg.DecodeMap{}
	decMap.SetOnNewRelationshipFunc(doc.onNewRelationship)
	// we should discover all contents by starting with these two files
	decMap.AddTarget(unioffice.ContentTypesFilename, doc.ContentTypes.X(), "", 0)
	decMap.AddTarget(unioffice.BaseRelsFilename, doc.Rels.X(), "", 0)
	if err := decMap.Decode(files); err != nil {
		return nil, err
	}
	doc.x.ConformanceAttr = ca

	for _, f := range files {
		if f == nil {
			continue
		}
		if err := doc.AddExtraFileFromZip(f); err != nil {
			return nil, err
		}
	}

	if addCustom {
		customPropertiesExist := false
		for _, rel := range doc.Rels.X().Relationship {
			if rel.TargetAttr == "docProps/custom.xml" {
				customPropertiesExist = true
				break
			}
		}
		if !customPropertiesExist {
			doc.addCustomRelationships()
		}
	}

	return doc, nil
}

// Validate validates the structure and in cases where it't possible, the ranges
// of elements within a document. A validation error dones't mean that the
// document won't work in MS Word or LibreOffice, but it's worth checking into.
func (d *Document) Validate() error {
	if d == nil || d.x == nil {
		return errors.New("document not initialized correctly, nil base")
	}

	for _, v := range []func() error{d.validateTableCells, d.validateBookmarks} {
		if err := v(); err != nil {
			return err
		}
	}
	if err := d.x.Validate(); err != nil {
		return err
	}
	return nil
}

func (d *Document) validateBookmarks() error {
	bmnames := make(map[string]struct{})
	for _, bm := range d.Bookmarks() {
		if _, ok := bmnames[bm.Name()]; ok {
			return fmt.Errorf("duplicate bookmark %s found", bm.Name())
		}
		bmnames[bm.Name()] = struct{}{}
	}
	return nil
}
func (d *Document) validateTableCells() error {
	for _, elt := range d.x.Body.EG_BlockLevelElts {
		for _, c := range elt.EG_ContentBlockContent {
			for _, t := range c.Tbl {
				for _, rc := range t.EG_ContentRowContent {
					for _, row := range rc.Tr {
						hasCell := false
						for _, ecc := range row.EG_ContentCellContent {
							cellHasPara := false
							for _, cell := range ecc.Tc {
								hasCell = true
								for _, cellElt := range cell.EG_BlockLevelElts {
									for _, cellCont := range cellElt.EG_ContentBlockContent {
										if len(cellCont.P) > 0 {
											cellHasPara = true
											break
										}
									}
								}
							}
							if !cellHasPara {
								return errors.New("table cell must contain a paragraph")
							}
						}
						// OSX Word requires this and won't open the file otherwise
						if !hasCell {
							return errors.New("table row must contain a cell")
						}
					}
				}
			}
		}
	}
	return nil
}

// AddImage adds an image to the document package, returning a reference that
// can be used to add the image to a run and place it in the document contents.
func (d *Document) AddImage(i common.Image) (common.ImageRef, error) {
	r := common.MakeImageRef(i, &d.DocBase, d.docRels)
	if i.Data == nil && i.Path == "" {
		return r, errors.New("image must have data or a path")
	}

	if i.Format == "" {
		return r, errors.New("image must have a valid format")
	}
	if i.Size.X == 0 || i.Size.Y == 0 {
		return r, errors.New("image must have a valid size")
	}

	d.Images = append(d.Images, r)
	fn := fmt.Sprintf("media/image%d.%s", len(d.Images), i.Format)
	rel := d.docRels.AddRelationship(fn, unioffice.ImageType)
	d.ContentTypes.EnsureDefault("png", "image/png")
	d.ContentTypes.EnsureDefault("jpeg", "image/jpeg")
	d.ContentTypes.EnsureDefault("jpg", "image/jpeg")
	d.ContentTypes.EnsureDefault("wmf", "image/x-wmf")
	d.ContentTypes.EnsureDefault(i.Format, "image/"+i.Format)
	r.SetRelID(rel.X().IdAttr)
	return r, nil
}

// GetImageByRelID returns an ImageRef with the associated relation ID in the
// document.
func (d *Document) GetImageByRelID(relID string) (common.ImageRef, bool) {
	for _, img := range d.Images {
		if img.RelID() == relID {
			return img, true
		}
	}
	return common.ImageRef{}, false
}

// FormFields extracts all of the fields from a document.  They can then be
// manipulated via the methods on the field and the document saved.
func (d *Document) FormFields() []FormField {
	ret := []FormField{}
	for _, p := range d.Paragraphs() {
		runs := p.Runs()
		for i, r := range runs {
			for _, ic := range r.x.EG_RunInnerContent {
				// skip non form fields
				if ic.FldChar == nil || ic.FldChar.FfData == nil {
					continue
				}

				// found a begin form field
				if ic.FldChar.FldCharTypeAttr == wml.ST_FldCharTypeBegin {
					// ensure it has a name
					if len(ic.FldChar.FfData.Name) == 0 || ic.FldChar.FfData.Name[0].ValAttr == nil {
						continue
					}

					field := FormField{x: ic.FldChar.FfData}
					// for text input boxes, we need a pointer to where to set
					// the text as well
					if ic.FldChar.FfData.TextInput != nil {

						// ensure we always have at lest two IC's
						for j := i + 1; j < len(runs)-1; j++ {
							if len(runs[j].x.EG_RunInnerContent) == 0 {
								continue
							}
							ic := runs[j].x.EG_RunInnerContent[0]
							// look for the 'separate' field
							if ic.FldChar != nil && ic.FldChar.FldCharTypeAttr == wml.ST_FldCharTypeSeparate {
								if len(runs[j+1].x.EG_RunInnerContent) == 0 {
									continue
								}
								// the value should be the text in the next inner content that is not a field char
								if runs[j+1].x.EG_RunInnerContent[0].FldChar == nil {
									field.textIC = runs[j+1].x.EG_RunInnerContent[0]
									break
								}
							}
						}
					}
					ret = append(ret, field)
				}
			}
		}
	}
	return ret
}

func (d *Document) onNewRelationship(decMap *zippkg.DecodeMap, target, typ string, files []*zip.File, rel *relationships.Relationship, src zippkg.Target) error {

	dt := unioffice.DocTypeDocument

	switch typ {
	case unioffice.OfficeDocumentType, unioffice.OfficeDocumentTypeStrict:
		d.x = wml.NewDocument()
		decMap.AddTarget(target, d.x, typ, 0)
		// look for the document relationships file as well
		decMap.AddTarget(zippkg.RelationsPathFor(target), d.docRels.X(), typ, 0)
		rel.TargetAttr = unioffice.RelativeFilename(dt, src.Typ, typ, 0)

	case unioffice.CorePropertiesType:
		decMap.AddTarget(target, d.CoreProperties.X(), typ, 0)
		rel.TargetAttr = unioffice.RelativeFilename(dt, src.Typ, typ, 0)

	case unioffice.CustomPropertiesType:
		decMap.AddTarget(target, d.CustomProperties.X(), typ, 0)
		rel.TargetAttr = unioffice.RelativeFilename(dt, src.Typ, typ, 0)

	case unioffice.ExtendedPropertiesType, unioffice.ExtendedPropertiesTypeStrict:
		decMap.AddTarget(target, d.AppProperties.X(), typ, 0)
		rel.TargetAttr = unioffice.RelativeFilename(dt, src.Typ, typ, 0)

	case unioffice.ThumbnailType, unioffice.ThumbnailTypeStrict:
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
				d.Thumbnail, _, err = image.Decode(rc)
				rc.Close()
				if err != nil {
					return fmt.Errorf("error decoding thumbnail: %s", err)
				}
				files[i] = nil
			}
		}

	case unioffice.SettingsType, unioffice.SettingsTypeStrict:
		decMap.AddTarget(target, d.Settings.X(), typ, 0)
		rel.TargetAttr = unioffice.RelativeFilename(dt, src.Typ, typ, 0)

	case unioffice.NumberingType, unioffice.NumberingTypeStrict:
		d.Numbering = NewNumbering()
		decMap.AddTarget(target, d.Numbering.X(), typ, 0)
		rel.TargetAttr = unioffice.RelativeFilename(dt, src.Typ, typ, 0)

	case unioffice.StylesType, unioffice.StylesTypeStrict:
		d.Styles.Clear()
		decMap.AddTarget(target, d.Styles.X(), typ, 0)
		rel.TargetAttr = unioffice.RelativeFilename(dt, src.Typ, typ, 0)

	case unioffice.HeaderType, unioffice.HeaderTypeStrict:
		hdr := wml.NewHdr()
		decMap.AddTarget(target, hdr, typ, uint32(len(d.headers)))
		d.headers = append(d.headers, hdr)
		rel.TargetAttr = unioffice.RelativeFilename(dt, src.Typ, typ, len(d.headers))

		// look for header rels
		hdrRel := common.NewRelationships()
		decMap.AddTarget(zippkg.RelationsPathFor(target), hdrRel.X(), typ, 0)
		d.hdrRels = append(d.hdrRels, hdrRel)

	case unioffice.FooterType, unioffice.FooterTypeStrict:
		ftr := wml.NewFtr()
		decMap.AddTarget(target, ftr, typ, uint32(len(d.footers)))
		d.footers = append(d.footers, ftr)
		rel.TargetAttr = unioffice.RelativeFilename(dt, src.Typ, typ, len(d.footers))

		// look for footer rels
		ftrRel := common.NewRelationships()
		decMap.AddTarget(zippkg.RelationsPathFor(target), ftrRel.X(), typ, 0)
		d.ftrRels = append(d.ftrRels, ftrRel)

	case unioffice.ThemeType, unioffice.ThemeTypeStrict:
		thm := dml.NewTheme()
		decMap.AddTarget(target, thm, typ, uint32(len(d.themes)))
		d.themes = append(d.themes, thm)
		rel.TargetAttr = unioffice.RelativeFilename(dt, src.Typ, typ, len(d.themes))

	case unioffice.WebSettingsType, unioffice.WebSettingsTypeStrict:
		d.webSettings = wml.NewWebSettings()
		decMap.AddTarget(target, d.webSettings, typ, 0)
		rel.TargetAttr = unioffice.RelativeFilename(dt, src.Typ, typ, 0)

	case unioffice.FontTableType, unioffice.FontTableTypeStrict:
		d.fontTable = wml.NewFonts()
		decMap.AddTarget(target, d.fontTable, typ, 0)
		rel.TargetAttr = unioffice.RelativeFilename(dt, src.Typ, typ, 0)

	case unioffice.EndNotesType, unioffice.EndNotesTypeStrict:
		d.endNotes = wml.NewEndnotes()
		decMap.AddTarget(target, d.endNotes, typ, 0)
		rel.TargetAttr = unioffice.RelativeFilename(dt, src.Typ, typ, 0)

	case unioffice.FootNotesType, unioffice.FootNotesTypeStrict:
		d.footNotes = wml.NewFootnotes()
		decMap.AddTarget(target, d.footNotes, typ, 0)
		rel.TargetAttr = unioffice.RelativeFilename(dt, src.Typ, typ, 0)

	case unioffice.ImageType, unioffice.ImageTypeStrict:
		var iref common.ImageRef
		for i, f := range files {
			if f == nil {
				continue
			}
			if f.Name == target {
				path, err := zippkg.ExtractToDiskTmp(f, d.TmpPath)
				if err != nil {
					return err
				}
				img, err := common.ImageFromFile(path)
				if err != nil {
					return err
				}
				iref = common.MakeImageRef(img, &d.DocBase, d.docRels)
				d.Images = append(d.Images, iref)
				files[i] = nil
			}
		}

		ext := "." + strings.ToLower(iref.Format())
		rel.TargetAttr = unioffice.RelativeFilename(dt, src.Typ, typ, len(d.Images))
		// ensure we don't change image formats
		if newExt := filepath.Ext(rel.TargetAttr); newExt != ext {
			rel.TargetAttr = rel.TargetAttr[0:len(rel.TargetAttr)-len(newExt)] + ext
		}

	default:
		unioffice.Log("unsupported relationship type: %s tgt: %s", typ, target)
	}
	return nil
}

// InsertParagraphAfter adds a new empty paragraph after the relativeTo
// paragraph.
func (d *Document) InsertParagraphAfter(relativeTo Paragraph) Paragraph {
	return d.insertParagraph(relativeTo, false)
}

// InsertParagraphBefore adds a new empty paragraph before the relativeTo
// paragraph.
func (d *Document) InsertParagraphBefore(relativeTo Paragraph) Paragraph {
	return d.insertParagraph(relativeTo, true)
}

func (d *Document) insertParagraph(relativeTo Paragraph, before bool) Paragraph {
	if d.x.Body == nil {
		return d.AddParagraph()
	}

	relX := relativeTo.X()

	for _, ble := range d.x.Body.EG_BlockLevelElts {
		for _, c := range ble.EG_ContentBlockContent {
			for i, p := range c.P {
				// found the paragraph
				if p == relX {
					p := wml.NewCT_P()
					c.P = append(c.P, nil)
					if before {
						copy(c.P[i+1:], c.P[i:])
						c.P[i] = p
					} else {
						copy(c.P[i+2:], c.P[i+1:])
						c.P[i+1] = p
					}
					return Paragraph{d, p}
				}
			}

			for _, tbl := range c.Tbl {
				for _, crc := range tbl.EG_ContentRowContent {
					for _, tr := range crc.Tr {
						for _, ccc := range tr.EG_ContentCellContent {
							for _, tc := range ccc.Tc {
								for _, ble := range tc.EG_BlockLevelElts {
									for _, cbc := range ble.EG_ContentBlockContent {
										for i, p := range cbc.P {
											if p == relX {
												p := wml.NewCT_P()
												cbc.P = append(cbc.P, nil)
												if before {
													copy(cbc.P[i+1:], cbc.P[i:])
													cbc.P[i] = p
												} else {
													copy(cbc.P[i+2:], cbc.P[i+1:])
													cbc.P[i+1] = p
												}
												return Paragraph{d, p}
											}
										}
									}
								}
							}
						}
					}
				}
			}

			if c.Sdt != nil && c.Sdt.SdtContent != nil && c.Sdt.SdtContent.P != nil {
				for i, p := range c.Sdt.SdtContent.P {
					if p == relX {
						p := wml.NewCT_P()
						c.Sdt.SdtContent.P = append(c.Sdt.SdtContent.P, nil)
						if before {
							copy(c.Sdt.SdtContent.P[i+1:], c.Sdt.SdtContent.P[i:])
							c.Sdt.SdtContent.P[i] = p
						} else {
							copy(c.Sdt.SdtContent.P[i+2:], c.Sdt.SdtContent.P[i+1:])
							c.Sdt.SdtContent.P[i+1] = p
						}
						return Paragraph{d, p}
					}
				}
			}
		}
	}
	return d.AddParagraph()
}

// AddHyperlink adds a hyperlink to a document. Adding the hyperlink to a document
// and setting it on a cell is more efficient than setting hyperlinks directly
// on a cell.
func (d Document) AddHyperlink(url string) common.Hyperlink {
	return d.docRels.AddHyperlink(url)
}

func bookmarks(bc *wml.EG_ContentBlockContent) []Bookmark {
	ret := []Bookmark{}

	// bookmarks within paragraphs
	for _, p := range bc.P {
		for _, ec := range p.EG_PContent {
			for _, ecr := range ec.EG_ContentRunContent {
				for _, re := range ecr.EG_RunLevelElts {
					for _, rm := range re.EG_RangeMarkupElements {
						if rm.BookmarkStart != nil {
							ret = append(ret, Bookmark{rm.BookmarkStart})
						}
					}
				}
			}
		}
	}
	// bookmarks within block runs
	for _, re := range bc.EG_RunLevelElts {
		for _, rm := range re.EG_RangeMarkupElements {
			if rm.BookmarkStart != nil {
				ret = append(ret, Bookmark{rm.BookmarkStart})
			}
		}
	}
	// bookmarks within tables, potentially nested
	for _, tbl := range bc.Tbl {
		for _, crc := range tbl.EG_ContentRowContent {
			for _, tr := range crc.Tr {
				for _, ccc := range tr.EG_ContentCellContent {
					for _, tc := range ccc.Tc {
						for _, ble := range tc.EG_BlockLevelElts {
							for _, bc := range ble.EG_ContentBlockContent {
								for _, b := range bookmarks(bc) {
									ret = append(ret, b)
								}
							}
						}
					}
				}
			}
		}
	}
	return ret
}

// Bookmarks returns all of the bookmarks defined in the document.
func (d Document) Bookmarks() []Bookmark {
	if d.x.Body == nil {
		return nil
	}
	ret := []Bookmark{}
	for _, ble := range d.x.Body.EG_BlockLevelElts {
		for _, bc := range ble.EG_ContentBlockContent {
			for _, b := range bookmarks(bc) {
				ret = append(ret, b)
			}
		}
	}
	return ret
}

// SetConformance sets conformance attribute of the document
// as one of these values from github.com/unidoc/unioffice/schema/soo/ofc/sharedTypes:
// ST_ConformanceClassUnset, ST_ConformanceClassStrict or ST_ConformanceClassTransitional.
func (d Document) SetConformance(conformanceAttr st.ST_ConformanceClass) {
	d.x.ConformanceAttr = conformanceAttr
}

// SetStrict is a shortcut for document.SetConformance,
// as one of these values from github.com/unidoc/unioffice/schema/soo/ofc/sharedTypes:
// ST_ConformanceClassUnset, ST_ConformanceClassStrict or ST_ConformanceClassTransitional.
func (d Document) SetStrict(strict bool) {
	if strict {
		d.x.ConformanceAttr = st.ST_ConformanceClassStrict
	} else {
		d.x.ConformanceAttr = st.ST_ConformanceClassTransitional
	}
}

func getBool(onOff *wml.CT_OnOff) bool {
	return onOff != nil
}
