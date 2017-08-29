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
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"baliance.com/gooxml/common"
	dml "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
	st "baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
	wml "baliance.com/gooxml/schema/schemas.openxmlformats.org/wordprocessingml"

	"baliance.com/gooxml/zippkg"
)

// Document is a text document that can be written out in the OOXML .docx format.
type Document struct {
	common.DocBase
	x           *wml.Document
	Settings    Settings
	Numbering   Numbering
	Styles      Styles
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

func (d *Document) ensureSectPr() {
	if d.x.Body.SectPr == nil {
		d.x.Body.SectPr = wml.NewCT_SectPr()
	}
}

// AddHeader creates a header, but doesn't add it to the document for display.
func (d *Document) AddHeader() Header {
	hdr := wml.NewHdr()
	d.headers = append(d.headers, hdr)
	path := fmt.Sprintf("header%d.xml", len(d.headers))
	d.docRels.AddRelationship(path, common.HeaderType)

	d.ContentTypes.AddOverride("/word/"+path, "application/vnd.openxmlformats-officedocument.wordprocessingml.header+xml")
	return Header{d, hdr}
}

// AddFooter creates a Footer, but doesn't add it to the document for display.
func (d *Document) AddFooter() Footer {
	ftr := wml.NewFtr()
	d.footers = append(d.footers, ftr)
	path := fmt.Sprintf("footer%d.xml", len(d.footers))
	d.docRels.AddRelationship(path, common.FooterType)
	d.ContentTypes.AddOverride("/word/"+path, "application/vnd.openxmlformats-officedocument.wordprocessingml.footer+xml")
	return Footer{d, ftr}
}

// BodySection returns the default body section used for all preceeding
// paragraphs until the previous Section. If there is no previous section, it
// applies to the entire document.
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
	d.WriteExtraFiles(z)
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
	// first pass, identify the files that should always be there
	for i, f := range files {
		switch f.Name {
		case zippkg.ContentTypesFilename:
			if err := zippkg.Decode(f, doc.ContentTypes.X()); err != nil {
				return nil, err
			}
			files[i] = nil
		case zippkg.BaseRelsFilename:
			if err := zippkg.Decode(f, doc.Rels.X()); err != nil {
				return nil, err
			}
			files[i] = nil
		}
	}

	basePaths := map[interface{}]string{}
	decMap := make(map[string]interface{})
	for _, r := range doc.Rels.Relationships() {
		switch r.Type() {
		case common.OfficeDocumentType:
			doc.x = wml.NewDocument()
			decMap[r.Target()] = doc.x

			// look for the document relationships file as well
			basePath, _ := filepath.Split(r.Target())
			decMap[zippkg.RelationsPathFor(r.Target())] = doc.docRels.X()
			basePaths[doc.docRels] = basePath
		case common.CorePropertiesType:
			decMap[r.Target()] = doc.CoreProperties.X()
		case common.ExtendedPropertiesType:
			decMap[r.Target()] = doc.AppProperties.X()
		default:
			log.Printf("unsupported type: %s", r.Type())
		}
	}

	if err := zippkg.DecodeFromMap(files, decMap); err != nil {
		return nil, err
	}

	for _, r := range doc.docRels.Relationships() {
		switch r.Type() {
		case common.SettingsType:
			decMap[basePaths[doc.docRels]+r.Target()] = doc.Settings.X()
		case common.NumberingType:
			doc.Numbering = NewNumbering()
			decMap[basePaths[doc.docRels]+r.Target()] = doc.Numbering.X()
		case common.StylesType:
			doc.Styles.Clear()
			decMap[basePaths[doc.docRels]+r.Target()] = doc.Styles.X()
		case common.HeaderType:
			hdr := wml.NewHdr()
			doc.headers = append(doc.headers, hdr)
			decMap[basePaths[doc.docRels]+r.Target()] = hdr
		case common.FooterType:
			ftr := wml.NewFtr()
			doc.footers = append(doc.footers, ftr)
			decMap[basePaths[doc.docRels]+r.Target()] = ftr
		case common.ThemeType:
			thm := dml.NewTheme()
			doc.themes = append(doc.themes, thm)
			decMap[basePaths[doc.docRels]+r.Target()] = thm
		case common.WebSettingsType:
			doc.webSettings = wml.NewWebSettings()
			decMap[basePaths[doc.docRels]+r.Target()] = doc.webSettings
		case common.FontTableType:
			doc.fontTable = wml.NewFonts()
			decMap[basePaths[doc.docRels]+r.Target()] = doc.fontTable
		case common.EndNotesType:
			doc.endNotes = wml.NewEndnotes()
			decMap[basePaths[doc.docRels]+r.Target()] = doc.endNotes
		case common.FootNotesType:
			doc.footNotes = wml.NewFootnotes()
			decMap[basePaths[doc.docRels]+r.Target()] = doc.footNotes
		case common.ImageType:
			imgPath := basePaths[doc.docRels] + r.Target()
			for i, f := range files {
				if f == nil {
					continue
				}
				if f.Name == imgPath {
					path, err := zippkg.ExtractToDiskTmp(f, doc.TmpPath)
					if err != nil {
						return nil, err
					}
					img, err := ImageFromFile(path)
					if err != nil {
						return nil, err
					}
					_ = img
					ref := &iref{path: img.Path}
					doc.images = append(doc.images, ref)
					files[i] = nil
				}
			}
		default:
			fmt.Println("unsupported document rel", r)
		}
	}
	if err := zippkg.DecodeFromMap(files, decMap); err != nil {
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

// Validate attempts to validate the structure of a document.
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
	r := ImageRef{img: i}
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
