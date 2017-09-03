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
	"log"
	"os"

	"baliance.com/gooxml/common"
	"baliance.com/gooxml/zippkg"

	dml "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
	st "baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
	"baliance.com/gooxml/schema/schemas.openxmlformats.org/package/2006/relationships"
	wml "baliance.com/gooxml/schema/schemas.openxmlformats.org/wordprocessingml"
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

	headers     []*wml.Hdr
	footers     []*wml.Ftr
	docRels     common.Relationships
	images      []*iref
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
	d.docRels.AddRelationship("settings.xml", common.SettingsType)
	d.ContentTypes.AddOverride("/word/settings.xml", "application/vnd.openxmlformats-officedocument.wordprocessingml.settings+xml")

	d.Rels = common.NewRelationships()
	d.Rels.AddRelationship("docProps/core.xml", common.CorePropertiesType)
	d.Rels.AddRelationship("docProps/app.xml", common.ExtendedPropertiesType)
	d.Rels.AddRelationship("word/document.xml", common.OfficeDocumentType)

	d.Numbering = NewNumbering()
	d.Numbering.InitializeDefault()
	d.ContentTypes.AddOverride("/word/numbering.xml", "application/vnd.openxmlformats-officedocument.wordprocessingml.numbering+xml")
	d.docRels.AddRelationship("numbering.xml", common.NumberingType)

	d.Styles = NewStyles()
	d.Styles.InitializeDefault()
	d.ContentTypes.AddOverride("/word/styles.xml", "application/vnd.openxmlformats-officedocument.wordprocessingml.styles+xml")
	d.docRels.AddRelationship("styles.xml", common.StylesType)

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
	d.docRels.AddRelationship(path, common.HeaderType)

	d.ContentTypes.AddOverride("/word/"+path, "application/vnd.openxmlformats-officedocument.wordprocessingml.header+xml")
	return Header{d, hdr}
}

// AddFooter creates a Footer associated with the document, but doesn't add it
// to the document for display.
func (d *Document) AddFooter() Footer {
	ftr := wml.NewFtr()
	d.footers = append(d.footers, ftr)
	path := fmt.Sprintf("footer%d.xml", len(d.footers))
	d.docRels.AddRelationship(path, common.FooterType)
	d.ContentTypes.AddOverride("/word/"+path, "application/vnd.openxmlformats-officedocument.wordprocessingml.footer+xml")
	return Footer{d, ftr}
}

// BodySection returns the default body section used for all preceeding
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
		log.Printf("validation error in document: %s", err)
	}
	z := zip.NewWriter(w)
	defer z.Close()
	if err := zippkg.MarshalXML(z, "_rels/.rels", d.Rels.X()); err != nil {
		return err
	}
	if err := zippkg.MarshalXML(z, "docProps/app.xml", d.AppProperties.X()); err != nil {
		return err
	}
	if err := zippkg.MarshalXML(z, "docProps/core.xml", d.CoreProperties.X()); err != nil {
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
	if err := zippkg.MarshalXML(z, "word/settings.xml", d.Settings.X()); err != nil {
		return err
	}
	if err := zippkg.MarshalXML(z, "word/document.xml", d.x); err != nil {
		return err
	}
	if d.Numbering.X() != nil {
		if err := zippkg.MarshalXML(z, "word/numbering.xml", d.Numbering.X()); err != nil {
			return err
		}
	}
	if err := zippkg.MarshalXML(z, "word/styles.xml", d.Styles.X()); err != nil {
		return err
	}
	if err := zippkg.MarshalXML(z, "word/_rels/document.xml.rels", d.docRels.X()); err != nil {
		return err
	}
	if d.webSettings != nil {
		if err := zippkg.MarshalXML(z, "word/webSettings.xml", d.webSettings); err != nil {
			return err
		}
	}
	if d.fontTable != nil {
		if err := zippkg.MarshalXML(z, "word/fontTable.xml", d.fontTable); err != nil {
			return err
		}
	}
	if d.endNotes != nil {
		if err := zippkg.MarshalXML(z, "word/endnotes.xml", d.endNotes); err != nil {
			return err
		}
	}
	if d.footNotes != nil {
		if err := zippkg.MarshalXML(z, "word/footnotes.xml", d.footNotes); err != nil {
			return err
		}
	}
	for i, thm := range d.themes {
		if err := zippkg.MarshalXML(z, fmt.Sprintf("word/theme/theme%d.xml", i+1), thm); err != nil {
			return err
		}
	}
	for i, hdr := range d.headers {
		fn := fmt.Sprintf("word/header%d.xml", i+1)
		if err := zippkg.MarshalXML(z, fn, hdr); err != nil {
			return err
		}
	}
	for i, ftr := range d.footers {
		fn := fmt.Sprintf("word/footer%d.xml", i+1)
		if err := zippkg.MarshalXML(z, fn, ftr); err != nil {
			return err
		}
	}
	for i, img := range d.images {
		fn := fmt.Sprintf("word/media/image%d.png", i+1)
		if img.path != "" {
			if err := zippkg.AddFileFromDisk(z, fn, img.path); err != nil {
				return err
			}
		} else {
			log.Printf("unsupported image source: %+v", img)
		}
	}
	if err := zippkg.MarshalXML(z, "[Content_Types].xml", d.ContentTypes.X()); err != nil {
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
	decMap.AddTarget(zippkg.ContentTypesFilename, doc.ContentTypes.X())
	decMap.AddTarget(zippkg.BaseRelsFilename, doc.Rels.X())
	decMap.Decode(files)

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

	if err := d.x.Validate(); err != nil {
		return err
	}
	return nil
}

// AddImage adds an image to the document package, returning a reference that
// can be used to add the image to a run and place it in the document contents.
func (d *Document) AddImage(i Image) (ImageRef, error) {
	r := ImageRef{img: i, d: d}
	if i.Path != "" {
		r.ref = &iref{path: i.Path}
		d.images = append(d.images, r.ref)
	} else {
		return r, errors.New("image must have a path")
	}
	if i.Format == "" {
		return r, errors.New("image must have a valid format")
	}
	if i.Size.X == 0 || i.Size.Y == 0 {
		return r, errors.New("image must have a valid size")
	}

	fn := fmt.Sprintf("media/image%d.%s", len(d.images), i.Format)
	d.docRels.AddRelationship(fn, common.ImageType)
	return r, nil
}

// GetImageByRelID returns an ImageRef with the associated relation ID in the
// document.
func (d *Document) GetImageByRelID(relID string) (ImageRef, bool) {
	for _, img := range d.Images() {
		if img.RelID() == relID {
			return img, true
		}
	}
	return ImageRef{}, false
}

// Images returns all images mentioned in the document relationships.
func (d *Document) Images() []ImageRef {
	ret := []ImageRef{}
	t := Image{}
	_ = t
	imgIdx := 0
	for _, rel := range d.docRels.Relationships() {
		if rel.Type() != common.ImageType {
			continue
		}
		if imgIdx < len(d.images) {
			iref := d.images[imgIdx]
			img, err := ImageFromFile(iref.path)
			if err != nil {
				// TODO: report this error?
			} else {
				rimg := ImageRef{ref: iref, img: img, d: d}
				ret = append(ret, rimg)
			}
		}
		imgIdx++
	}
	return ret
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

func (doc *Document) onNewRelationship(decMap *zippkg.DecodeMap, target, typ string, files []*zip.File, rel *relationships.Relationship) error {
	switch typ {
	case common.OfficeDocumentType:
		doc.x = wml.NewDocument()
		decMap.AddTarget(target, doc.x)
		// look for the document relationships file as well
		decMap.AddTarget(zippkg.RelationsPathFor(target), doc.docRels.X())
	case common.CorePropertiesType:
		decMap.AddTarget(target, doc.CoreProperties.X())
	case common.ExtendedPropertiesType:
		decMap.AddTarget(target, doc.AppProperties.X())
	case common.ThumbnailType:
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
				doc.Thumbnail, _, err = image.Decode(rc)
				rc.Close()
				if err != nil {
					return fmt.Errorf("error decoding thumbnail: %s", err)
				}
				files[i] = nil
			}
		}
	case common.SettingsType:
		decMap.AddTarget(target, doc.Settings.X())
	case common.NumberingType:
		doc.Numbering = NewNumbering()
		decMap.AddTarget(target, doc.Numbering.X())
	case common.StylesType:
		doc.Styles.Clear()
		decMap.AddTarget(target, doc.Styles.X())
	case common.HeaderType:
		hdr := wml.NewHdr()
		doc.headers = append(doc.headers, hdr)
		decMap.AddTarget(target, hdr)
	case common.FooterType:
		ftr := wml.NewFtr()
		doc.footers = append(doc.footers, ftr)
		decMap.AddTarget(target, ftr)
	case common.ThemeType:
		thm := dml.NewTheme()
		doc.themes = append(doc.themes, thm)
		decMap.AddTarget(target, thm)
	case common.WebSettingsType:
		doc.webSettings = wml.NewWebSettings()
		decMap.AddTarget(target, doc.webSettings)
	case common.FontTableType:
		doc.fontTable = wml.NewFonts()
		decMap.AddTarget(target, doc.fontTable)
	case common.EndNotesType:
		doc.endNotes = wml.NewEndnotes()
		decMap.AddTarget(target, doc.endNotes)
	case common.FootNotesType:
		doc.footNotes = wml.NewFootnotes()
		decMap.AddTarget(target, doc.footNotes)
	case common.ImageType:
		for i, f := range files {
			if f == nil {
				continue
			}
			if f.Name == target {
				path, err := zippkg.ExtractToDiskTmp(f, doc.TmpPath)
				if err != nil {
					return err
				}
				img, err := ImageFromFile(path)
				if err != nil {
					return err
				}
				_ = img
				ref := &iref{path: img.Path}
				doc.images = append(doc.images, ref)
				files[i] = nil
			}
		}
	default:
		log.Printf("unsupported relationship type: %s tgt: %s", typ, target)
	}
	return nil
}
