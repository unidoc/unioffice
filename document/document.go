// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"archive/zip"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"baliance.com/gooxml"
	"baliance.com/gooxml/common"
	"baliance.com/gooxml/zippkg"

	"baliance.com/gooxml/schema/soo/dml"
	st "baliance.com/gooxml/schema/soo/ofc/sharedTypes"
	"baliance.com/gooxml/schema/soo/pkg/relationships"
	"baliance.com/gooxml/schema/soo/wml"
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
	d.docRels.AddRelationship("settings.xml", gooxml.SettingsType)
	d.ContentTypes.AddOverride("/word/settings.xml", "application/vnd.openxmlformats-officedocument.wordprocessingml.settings+xml")

	d.Rels = common.NewRelationships()
	d.Rels.AddRelationship(gooxml.RelativeFilename(gooxml.DocTypeDocument, "", gooxml.CorePropertiesType, 0), gooxml.CorePropertiesType)
	d.Rels.AddRelationship("docProps/app.xml", gooxml.ExtendedPropertiesType)
	d.Rels.AddRelationship("word/document.xml", gooxml.OfficeDocumentType)

	d.Numbering = NewNumbering()
	d.Numbering.InitializeDefault()
	d.ContentTypes.AddOverride("/word/numbering.xml", "application/vnd.openxmlformats-officedocument.wordprocessingml.numbering+xml")
	d.docRels.AddRelationship("numbering.xml", gooxml.NumberingType)

	d.Styles = NewStyles()
	d.Styles.InitializeDefault()
	d.ContentTypes.AddOverride("/word/styles.xml", "application/vnd.openxmlformats-officedocument.wordprocessingml.styles+xml")
	d.docRels.AddRelationship("styles.xml", gooxml.StylesType)

	d.x.Body = wml.NewCT_Body()
	return d
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
	d.docRels.AddRelationship(path, gooxml.HeaderType)

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
	d.docRels.AddRelationship(path, gooxml.FooterType)

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
		gooxml.Log("validation error in document: %s", err)
	}
	dt := gooxml.DocTypeDocument

	z := zip.NewWriter(w)
	defer z.Close()
	if err := zippkg.MarshalXML(z, gooxml.BaseRelsFilename, d.Rels.X()); err != nil {
		return err
	}
	if err := zippkg.MarshalXMLByType(z, dt, gooxml.ExtendedPropertiesType, d.AppProperties.X()); err != nil {
		return err
	}
	if err := zippkg.MarshalXMLByType(z, dt, gooxml.CorePropertiesType, d.CoreProperties.X()); err != nil {
		return err
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
	if err := zippkg.MarshalXMLByType(z, dt, gooxml.SettingsType, d.Settings.X()); err != nil {
		return err
	}
	documentFn := gooxml.AbsoluteFilename(dt, gooxml.OfficeDocumentType, 0)
	if err := zippkg.MarshalXML(z, documentFn, d.x); err != nil {
		return err
	}
	if err := zippkg.MarshalXML(z, zippkg.RelationsPathFor(documentFn), d.docRels.X()); err != nil {
		return err
	}

	if d.Numbering.X() != nil {
		if err := zippkg.MarshalXMLByType(z, dt, gooxml.NumberingType, d.Numbering.X()); err != nil {
			return err
		}
	}
	if err := zippkg.MarshalXMLByType(z, dt, gooxml.StylesType, d.Styles.X()); err != nil {
		return err
	}

	if d.webSettings != nil {
		if err := zippkg.MarshalXMLByType(z, dt, gooxml.WebSettingsType, d.webSettings); err != nil {
			return err
		}
	}
	if d.fontTable != nil {
		if err := zippkg.MarshalXMLByType(z, dt, gooxml.FontTableType, d.fontTable); err != nil {
			return err
		}
	}
	if d.endNotes != nil {
		if err := zippkg.MarshalXMLByType(z, dt, gooxml.EndNotesType, d.endNotes); err != nil {
			return err
		}
	}
	if d.footNotes != nil {
		if err := zippkg.MarshalXMLByType(z, dt, gooxml.FootNotesType, d.footNotes); err != nil {
			return err
		}
	}
	for i, thm := range d.themes {
		if err := zippkg.MarshalXMLByTypeIndex(z, dt, gooxml.ThemeType, i+1, thm); err != nil {
			return err
		}
	}
	for i, hdr := range d.headers {
		fn := gooxml.AbsoluteFilename(dt, gooxml.HeaderType, i+1)
		if err := zippkg.MarshalXML(z, fn, hdr); err != nil {
			return err
		}
		if !d.hdrRels[i].IsEmpty() {
			zippkg.MarshalXML(z, zippkg.RelationsPathFor(fn), d.hdrRels[i].X())
		}
	}
	for i, ftr := range d.footers {
		fn := gooxml.AbsoluteFilename(dt, gooxml.FooterType, i+1)
		if err := zippkg.MarshalXMLByTypeIndex(z, dt, gooxml.FooterType, i+1, ftr); err != nil {
			return err
		}
		if !d.ftrRels[i].IsEmpty() {
			zippkg.MarshalXML(z, zippkg.RelationsPathFor(fn), d.ftrRels[i].X())
		}
	}

	for i, img := range d.Images {
		fn := fmt.Sprintf("word/media/image%d.%s", i+1, strings.ToLower(img.Format()))
		if img.Path() != "" {
			if err := zippkg.AddFileFromDisk(z, fn, img.Path()); err != nil {
				return err
			}
		} else {
			gooxml.Log("unsupported image source: %+v", img)
		}
	}

	if err := zippkg.MarshalXML(z, gooxml.ContentTypesFilename, d.ContentTypes.X()); err != nil {
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
	if d.x.Body == nil {
		return d.AddTable()
	}
	for i, ble := range d.x.Body.EG_BlockLevelElts {
		for _, c := range ble.EG_ContentBlockContent {
			for j, p := range c.P {
				// found the paragraph
				if p == relativeTo.X() {
					tbl := wml.NewCT_Tbl()
					elts := wml.NewEG_BlockLevelElts()
					cbc := wml.NewEG_ContentBlockContent()
					elts.EG_ContentBlockContent = append(elts.EG_ContentBlockContent, cbc)
					cbc.Tbl = append(cbc.Tbl, tbl)
					d.x.Body.EG_BlockLevelElts = append(d.x.Body.EG_BlockLevelElts, nil)
					if before {
						copy(d.x.Body.EG_BlockLevelElts[i+1:], d.x.Body.EG_BlockLevelElts[i:])
						d.x.Body.EG_BlockLevelElts[i] = elts
						if j != 0 {
							elts := wml.NewEG_BlockLevelElts()
							cbc := wml.NewEG_ContentBlockContent()
							elts.EG_ContentBlockContent = append(elts.EG_ContentBlockContent, cbc)
							cbc.P = c.P[:j]
							d.x.Body.EG_BlockLevelElts = append(d.x.Body.EG_BlockLevelElts, nil)
							copy(d.x.Body.EG_BlockLevelElts[i+1:], d.x.Body.EG_BlockLevelElts[i:])
							d.x.Body.EG_BlockLevelElts[i] = elts
						}
						c.P = c.P[j:]
					} else {
						copy(d.x.Body.EG_BlockLevelElts[i+2:], d.x.Body.EG_BlockLevelElts[i+1:])
						d.x.Body.EG_BlockLevelElts[i+1] = elts
						if j != len(c.P)-1 {
							elts := wml.NewEG_BlockLevelElts()
							cbc := wml.NewEG_ContentBlockContent()
							elts.EG_ContentBlockContent = append(elts.EG_ContentBlockContent, cbc)
							cbc.P = c.P[j+1:]
							d.x.Body.EG_BlockLevelElts = append(d.x.Body.EG_BlockLevelElts, nil)
							copy(d.x.Body.EG_BlockLevelElts[i+3:], d.x.Body.EG_BlockLevelElts[i+2:])
							d.x.Body.EG_BlockLevelElts[i+2] = elts
						}
						c.P = c.P[:j+1]
					}
					return Table{d, tbl}
				}
			}
		}
	}
	return d.AddTable()
}

// Tables returns the tables defined in the document.
func (d *Document) Tables() []Table {
	ret := []Table{}
	if d.x.Body == nil {
		return nil
	}
	for _, ble := range d.x.Body.EG_BlockLevelElts {
		for _, c := range ble.EG_ContentBlockContent {
			for _, t := range c.Tbl {
				ret = append(ret, Table{d, t})
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

// Paragraphs returns all of the paragraphs in the document body.
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
	return ret
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

	decMap := zippkg.DecodeMap{}
	decMap.SetOnNewRelationshipFunc(doc.onNewRelationship)
	// we should discover all contents by starting with these two files
	decMap.AddTarget(gooxml.ContentTypesFilename, doc.ContentTypes.X(), "", 0)
	decMap.AddTarget(gooxml.BaseRelsFilename, doc.Rels.X(), "", 0)
	if err := decMap.Decode(files); err != nil {
		return nil, err
	}

	for _, f := range files {
		if f == nil {
			continue
		}
		if err := doc.AddExtraFileFromZip(f); err != nil {
			return nil, err
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
	if i.Path == "" {
		return r, errors.New("image must have a path")
	}

	if i.Format == "" {
		return r, errors.New("image must have a valid format")
	}
	if i.Size.X == 0 || i.Size.Y == 0 {
		return r, errors.New("image must have a valid size")
	}

	d.Images = append(d.Images, r)
	fn := fmt.Sprintf("media/image%d.%s", len(d.Images), i.Format)
	rel := d.docRels.AddRelationship(fn, gooxml.ImageType)
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
	dt := gooxml.DocTypeDocument

	switch typ {
	case gooxml.OfficeDocumentType:
		d.x = wml.NewDocument()
		decMap.AddTarget(target, d.x, typ, 0)
		// look for the document relationships file as well
		decMap.AddTarget(zippkg.RelationsPathFor(target), d.docRels.X(), typ, 0)
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, 0)

	case gooxml.CorePropertiesType:
		decMap.AddTarget(target, d.CoreProperties.X(), typ, 0)
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, 0)

	case gooxml.ExtendedPropertiesType:
		decMap.AddTarget(target, d.AppProperties.X(), typ, 0)
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, 0)

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
				d.Thumbnail, _, err = image.Decode(rc)
				rc.Close()
				if err != nil {
					return fmt.Errorf("error decoding thumbnail: %s", err)
				}
				files[i] = nil
			}
		}

	case gooxml.SettingsType:
		decMap.AddTarget(target, d.Settings.X(), typ, 0)
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, 0)

	case gooxml.NumberingType:
		d.Numbering = NewNumbering()
		decMap.AddTarget(target, d.Numbering.X(), typ, 0)
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, 0)

	case gooxml.StylesType:
		d.Styles.Clear()
		decMap.AddTarget(target, d.Styles.X(), typ, 0)
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, 0)

	case gooxml.HeaderType:
		hdr := wml.NewHdr()
		decMap.AddTarget(target, hdr, typ, uint32(len(d.headers)))
		d.headers = append(d.headers, hdr)
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, len(d.headers))

		// look for header rels
		hdrRel := common.NewRelationships()
		decMap.AddTarget(zippkg.RelationsPathFor(target), hdrRel.X(), typ, 0)
		d.hdrRels = append(d.hdrRels, hdrRel)

	case gooxml.FooterType:
		ftr := wml.NewFtr()
		decMap.AddTarget(target, ftr, typ, uint32(len(d.footers)))
		d.footers = append(d.footers, ftr)
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, len(d.footers))

		// look for footer rels
		ftrRel := common.NewRelationships()
		decMap.AddTarget(zippkg.RelationsPathFor(target), ftrRel.X(), typ, 0)
		d.ftrRels = append(d.ftrRels, ftrRel)

	case gooxml.ThemeType:
		thm := dml.NewTheme()
		decMap.AddTarget(target, thm, typ, uint32(len(d.themes)))
		d.themes = append(d.themes, thm)
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, len(d.themes))

	case gooxml.WebSettingsType:
		d.webSettings = wml.NewWebSettings()
		decMap.AddTarget(target, d.webSettings, typ, 0)
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, 0)

	case gooxml.FontTableType:
		d.fontTable = wml.NewFonts()
		decMap.AddTarget(target, d.fontTable, typ, 0)
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, 0)

	case gooxml.EndNotesType:
		d.endNotes = wml.NewEndnotes()
		decMap.AddTarget(target, d.endNotes, typ, 0)
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, 0)

	case gooxml.FootNotesType:
		d.footNotes = wml.NewFootnotes()
		decMap.AddTarget(target, d.footNotes, typ, 0)
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, 0)

	case gooxml.ImageType:
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
				iref := common.MakeImageRef(img, &d.DocBase, d.docRels)
				d.Images = append(d.Images, iref)
				files[i] = nil
			}
		}
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, len(d.Images))
	default:
		gooxml.Log("unsupported relationship type: %s tgt: %s", typ, target)
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

	for _, ble := range d.x.Body.EG_BlockLevelElts {
		for _, c := range ble.EG_ContentBlockContent {
			for i, p := range c.P {
				// foudn the paragraph
				if p == relativeTo.X() {
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

			if c.Sdt != nil && c.Sdt.SdtContent != nil && c.Sdt.SdtContent.P != nil {
				for i, p := range c.Sdt.SdtContent.P {
					if p == relativeTo.X() {
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

// Bookmarks returns all of the bookmarks defined in the document.
func (d Document) Bookmarks() []Bookmark {
	if d.x.Body == nil {
		return nil
	}
	ret := []Bookmark{}
	for _, ble := range d.x.Body.EG_BlockLevelElts {
		for _, bc := range ble.EG_ContentBlockContent {
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
			for _, re := range bc.EG_RunLevelElts {
				for _, rm := range re.EG_RangeMarkupElements {
					if rm.BookmarkStart != nil {
						ret = append(ret, Bookmark{rm.BookmarkStart})
					}
				}
			}
		}
	}
	return ret
}
