// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

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
	"baliance.com/gooxml/schema/schemas.openxmlformats.org/spreadsheetml"
	"baliance.com/gooxml/spreadsheet/styles"
	"baliance.com/gooxml/zippkg"
)

// Workbook is the top level container item for a set of spreadsheets.
type Workbook struct {
	common.DocBase
	x *spreadsheetml.Workbook

	StyleSheet    styles.StyleSheet
	Theme         common.Theme
	SharedStrings SharedStrings
	xws           []*spreadsheetml.Worksheet
	xwsRels       []common.Relationships
	wbRels        common.Relationships
}

// New constructs a new workbook.
func New() *Workbook {
	wb := &Workbook{}
	wb.x = spreadsheetml.NewWorkbook()

	wb.AppProperties = common.NewAppProperties()
	wb.CoreProperties = common.NewCoreProperties()
	wb.StyleSheet = styles.NewStyleSheet()

	wb.Rels = common.NewRelationships()
	wb.wbRels = common.NewRelationships()
	wb.Rels.AddRelationship(zippkg.AppPropsFilename, common.ExtendedPropertiesType)
	wb.Rels.AddRelationship(zippkg.CorePropsFilename, common.CorePropertiesType)
	wb.Rels.AddRelationship("xl/workbook.xml", common.OfficeDocumentType)
	wb.wbRels.AddRelationship("styles.xml", common.StylesType)

	wb.ContentTypes = common.NewContentTypes()
	wb.ContentTypes.AddOverride("/xl/workbook.xml", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml")
	wb.ContentTypes.AddOverride("/xl/styles.xml", "application/vnd.openxmlformats-officedocument.spreadsheetml.styles+xml")

	wb.SharedStrings = NewSharedStrings()
	wb.ContentTypes.AddOverride("/xl/sharedStrings.xml", common.SharedStringsContentType)
	wb.wbRels.AddRelationship("sharedStrings.xml", common.SharedStingsType)

	return wb
}

// Open opens and reads a workbook from a file (.xlsx).
func Open(filename string) (*Workbook, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening %s: %s", filename, err)
	}
	defer f.Close()
	fi, err := os.Stat(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening %s: %s", filename, err)
	}
	return Read(f, fi.Size())
}

// Read reads a workbook from an io.Reader(.xlsx).
func Read(r io.ReaderAt, size int64) (*Workbook, error) {
	wb := New()
	td, err := ioutil.TempDir("", "gooxml-xlsx")
	if err != nil {
		return nil, err
	}
	wb.TmpPath = td

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
			if err := zippkg.Decode(f, wb.ContentTypes.X()); err != nil {
				return nil, err
			}
			files[i] = nil
		case zippkg.BaseRelsFilename:
			if err := zippkg.Decode(f, wb.Rels.X()); err != nil {
				return nil, err
			}
			files[i] = nil
		}
	}

	basePaths := map[interface{}]string{}
	decMap := make(map[string]interface{})
	for _, r := range wb.Rels.Relationships() {
		switch r.Type() {
		case common.OfficeDocumentType:
			wb.x = spreadsheetml.NewWorkbook()
			decMap[r.Target()] = wb.x
			// look for the workbook relationships file as well
			basePath, _ := filepath.Split(r.Target())
			wb.wbRels = common.NewRelationships()
			decMap[zippkg.RelationsPathFor(r.Target())] = wb.wbRels.X()
			basePaths[wb.wbRels] = basePath
		case common.CorePropertiesType:
			decMap[r.Target()] = wb.CoreProperties.X()
		case common.ExtendedPropertiesType:
			decMap[r.Target()] = wb.AppProperties.X()
		default:
			log.Printf("unsupported type: %s", r.Type())
		}
	}

	if err := zippkg.DecodeFromMap(files, decMap); err != nil {
		return nil, err
	}

	for _, r := range wb.wbRels.Relationships() {
		switch r.Type() {
		case common.WorksheetType:
			ws := spreadsheetml.NewWorksheet()
			wb.xws = append(wb.xws, ws)
			decMap[basePaths[wb.wbRels]+r.Target()] = ws

			// look for worksheet rels
			basePath, _ := filepath.Split(r.Target())
			wksRel := common.NewRelationships()
			decMap[basePaths[wb.wbRels]+zippkg.RelationsPathFor(r.Target())] = wksRel.X()
			basePaths[wksRel] = basePath
			wb.xwsRels = append(wb.xwsRels, wksRel)
		case common.StylesType:
			wb.StyleSheet = styles.NewStyleSheet()
			decMap[basePaths[wb.wbRels]+r.Target()] = wb.StyleSheet.X()
		case common.ThemeType:
			wb.Theme = common.NewTheme()
			decMap[basePaths[wb.wbRels]+r.Target()] = wb.Theme.X()
		case common.SharedStingsType:
			wb.SharedStrings = NewSharedStrings()
			decMap[basePaths[wb.wbRels]+r.Target()] = wb.SharedStrings.X()
		default:
			fmt.Println("unsupported worksheet rel", r)
		}
	}

	if err := zippkg.DecodeFromMap(files, decMap); err != nil {
		return nil, err
	}

	for _, f := range files {
		if f == nil {
			continue
		}
		if err := wb.AddExtraFileFromZip(f); err != nil {
			return nil, err
		}
	}
	return wb, nil
}

// X returns the inner wrapped XML type.
func (wb *Workbook) X() *spreadsheetml.Workbook {
	return wb.x
}

// AddSheet adds a new sheet with a given name to a workbook.
func (wb *Workbook) AddSheet(name string) Sheet {
	rs := spreadsheetml.NewCT_Sheet()
	rs.NameAttr = name
	// Assign a unique sheet ID
	rs.SheetIdAttr = 1
	for _, s := range wb.x.Sheets.Sheet {
		if rs.SheetIdAttr <= s.SheetIdAttr {
			rs.SheetIdAttr = s.SheetIdAttr + 1
		}
	}
	wb.x.Sheets.Sheet = append(wb.x.Sheets.Sheet, rs)

	rs.NameAttr = fmt.Sprintf("Sheet %d", rs.SheetIdAttr)

	// create the actual worksheet
	ws := spreadsheetml.NewWorksheet()
	ws.Dimension = spreadsheetml.NewCT_SheetDimension()
	ws.Dimension.RefAttr = "A1"
	wb.xws = append(wb.xws, ws)

	ws.SheetData = spreadsheetml.NewCT_SheetData()

	// update the references
	rid := wb.wbRels.AddRelationship(fmt.Sprintf("worksheets/sheet%d.xml", rs.SheetIdAttr),
		"http://schemas.openxmlformats.org/officeDocument/2006/relationships/worksheet")
	rs.IdAttr = rid.ID()

	// add the content type
	wb.ContentTypes.AddOverride(fmt.Sprintf("/xl/worksheets/sheet%d.xml", rs.SheetIdAttr),
		"application/vnd.openxmlformats-officedocument.spreadsheetml.worksheet+xml")

	return Sheet{wb, rs, ws}
}

// SaveToFile writes the workbook out to a file.
func (wb *Workbook) SaveToFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return wb.Save(f)
}

// Save writes the workbook out to a writer in the zipped xlsx format.
func (wb *Workbook) Save(w io.Writer) error {
	z := zip.NewWriter(w)
	defer z.Close()
	if err := zippkg.MarshalXML(z, zippkg.ContentTypesFilename, wb.ContentTypes.X()); err != nil {
		return err
	}
	if err := zippkg.MarshalXML(z, zippkg.BaseRelsFilename, wb.Rels.X()); err != nil {
		return err
	}
	if err := zippkg.MarshalXML(z, zippkg.AppPropsFilename, wb.AppProperties.X()); err != nil {
		return err
	}
	if err := zippkg.MarshalXML(z, zippkg.CorePropsFilename, wb.CoreProperties.X()); err != nil {
		return err
	}
	if err := zippkg.MarshalXML(z, "xl/workbook.xml", wb.x); err != nil {
		return err
	}
	if err := zippkg.MarshalXML(z, "xl/styles.xml", wb.StyleSheet.X()); err != nil {
		return err
	}
	if err := zippkg.MarshalXML(z, "xl/sharedStrings.xml", wb.SharedStrings.X()); err != nil {
		return err
	}
	if err := zippkg.MarshalXML(z, "xl/_rels/workbook.xml.rels", wb.wbRels.X()); err != nil {
		return err
	}
	for i, sheet := range wb.xws {
		wbs := wb.x.Sheets.Sheet[i]
		zippkg.MarshalXML(z, fmt.Sprintf("xl/worksheets/sheet%d.xml", wbs.SheetIdAttr), sheet)
	}
	wb.WriteExtraFiles(z)
	return z.Close()
}

// Validate attempts to validate the structure of a workbook.
func (wb *Workbook) Validate() error {
	if wb == nil || wb.x == nil {
		return errors.New("workbook not initialized correctly, nil base")
	}

	maxID := uint32(0)
	for _, s := range wb.x.Sheets.Sheet {
		if s.SheetIdAttr > maxID {
			maxID = s.SheetIdAttr
		}
	}

	if maxID != uint32(len(wb.xws)) {
		return fmt.Errorf("found %d worksheet descriptions and %d worksheets", maxID, len(wb.xws))
	}

	for i, s := range wb.x.Sheets.Sheet {
		sw := Sheet{wb, s, wb.xws[i]}
		if err := sw.ValidateWithPath(fmt.Sprintf("workbook/Sheet[%d]", i)); err != nil {
			return err
		}
	}
	return nil
}

// Sheets returns the sheets from the workbook.
func (wb *Workbook) Sheets() []Sheet {
	ret := []Sheet{}
	for i, wks := range wb.xws {
		r := wb.x.Sheets.Sheet[i]
		ret = append(ret, Sheet{wb, r, wks})
	}
	return ret
}

// SheetCount returns the number of sheets in the workbook.
func (wb Workbook) SheetCount() int {

	return len(wb.xws)
}
