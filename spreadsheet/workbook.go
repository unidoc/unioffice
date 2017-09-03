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
	"image"
	"image/jpeg"
	"io"
	"os"

	"baliance.com/gooxml"
	"baliance.com/gooxml/common"
	"baliance.com/gooxml/zippkg"

	dml "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
	crt "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/chart"
	sd "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"
	"baliance.com/gooxml/schema/schemas.openxmlformats.org/package/2006/relationships"
	sml "baliance.com/gooxml/schema/schemas.openxmlformats.org/spreadsheetml"
)

// Workbook is the top level container item for a set of spreadsheets.
type Workbook struct {
	common.DocBase
	x *sml.Workbook

	StyleSheet    StyleSheet
	SharedStrings SharedStrings

	xws         []*sml.Worksheet
	xwsRels     []common.Relationships
	wbRels      common.Relationships
	themes      []*dml.Theme
	drawings    []*sd.WsDr
	drawingRels []common.Relationships
	charts      []*crt.ChartSpace
}

// X returns the inner wrapped XML type.
func (wb *Workbook) X() *sml.Workbook {
	return wb.x
}

// AddSheet adds a new sheet with a given name to a workbook.
func (wb *Workbook) AddSheet() Sheet {
	rs := sml.NewCT_Sheet()

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
	ws := sml.NewWorksheet()
	ws.Dimension = sml.NewCT_SheetDimension()
	ws.Dimension.RefAttr = "A1"
	wb.xws = append(wb.xws, ws)
	wb.xwsRels = append(wb.xwsRels, common.NewRelationships())
	ws.SheetData = sml.NewCT_SheetData()

	// update the references
	rid := wb.wbRels.AddRelationship(fmt.Sprintf("worksheets/sheet%d.xml", len(wb.x.Sheets.Sheet)),
		"http://schemas.openxmlformats.org/officeDocument/2006/relationships/worksheet")
	rs.IdAttr = rid.ID()

	// add the content type
	wb.ContentTypes.AddOverride(fmt.Sprintf("/xl/worksheets/sheet%d.xml", len(wb.x.Sheets.Sheet)),
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
	for i, thm := range wb.themes {
		if err := zippkg.MarshalXML(z, fmt.Sprintf("xl/theme/theme%d.xml", i+1), thm); err != nil {
			return err
		}
	}
	for i, sheet := range wb.xws {
		fn := fmt.Sprintf("xl/worksheets/sheet%d.xml", i+1)
		zippkg.MarshalXML(z, fn, sheet)
		zippkg.MarshalXML(z, zippkg.RelationsPathFor(fn), wb.xwsRels[i].X())
	}
	if wb.Thumbnail != nil {
		tn, err := z.Create("docProps/thumbnail.jpeg")
		if err != nil {
			return err
		}
		if err := jpeg.Encode(tn, wb.Thumbnail, nil); err != nil {
			return err
		}
	}
	for i, chart := range wb.charts {
		fn := fmt.Sprintf("xl/charts/chart%d.xml", i+1)
		zippkg.MarshalXML(z, fn, chart)
	}
	for i, drawing := range wb.drawings {
		fn := fmt.Sprintf("xl/drawings/drawing%d.xml", i+1)
		zippkg.MarshalXML(z, fn, drawing)
		zippkg.MarshalXML(z, zippkg.RelationsPathFor(fn), wb.drawingRels[i].X())
	}
	if err := wb.WriteExtraFiles(z); err != nil {
		return err
	}
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

	// Excel doesn't like reused sheet names
	usedNames := map[string]struct{}{}
	for i, s := range wb.x.Sheets.Sheet {
		sw := Sheet{wb, s, wb.xws[i]}
		if _, ok := usedNames[sw.Name()]; ok {
			return fmt.Errorf("workbook/Sheet[%d] has duplicate name '%s'", i, sw.Name())
		}
		usedNames[sw.Name()] = struct{}{}
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

func (wb *Workbook) onNewRelationship(decMap *zippkg.DecodeMap, target, typ string, files []*zip.File, rel *relationships.Relationship) error {
	switch typ {
	case gooxml.OfficeDocumentType:
		wb.x = sml.NewWorkbook()
		decMap.AddTarget(target, wb.x)
		// look for the workbook relationships file as well
		wb.wbRels = common.NewRelationships()
		decMap.AddTarget(zippkg.RelationsPathFor(target), wb.wbRels.X())

	case gooxml.CorePropertiesType:
		decMap.AddTarget(target, wb.CoreProperties.X())

	case gooxml.ExtendedPropertiesType:
		decMap.AddTarget(target, wb.AppProperties.X())

	case gooxml.WorksheetType:
		ws := sml.NewWorksheet()
		wb.xws = append(wb.xws, ws)
		decMap.AddTarget(target, ws)
		// look for worksheet rels
		wksRel := common.NewRelationships()
		decMap.AddTarget(zippkg.RelationsPathFor(target), wksRel.X())
		wb.xwsRels = append(wb.xwsRels, wksRel)
		// fix the relationship target so it points to where we'll save
		// the worksheet
		rel.TargetAttr = fmt.Sprintf("worksheets/sheet%d.xml", len(wb.xws))

	case gooxml.StylesType:
		wb.StyleSheet = NewStyleSheet()
		decMap.AddTarget(target, wb.StyleSheet.X())

	case gooxml.ThemeType:
		thm := dml.NewTheme()
		wb.themes = append(wb.themes, thm)
		decMap.AddTarget(target, thm)

	case gooxml.SharedStingsType:
		wb.SharedStrings = NewSharedStrings()
		decMap.AddTarget(target, wb.SharedStrings.X())

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
				wb.Thumbnail, _, err = image.Decode(rc)
				rc.Close()
				if err != nil {
					return fmt.Errorf("error decoding thumbnail: %s", err)
				}
				files[i] = nil
			}
		}

	case gooxml.DrawingType:
		drawing := sd.NewWsDr()
		decMap.AddTarget(target, drawing)
		wb.drawings = append(wb.drawings, drawing)

		drel := common.NewRelationships()
		decMap.AddTarget(zippkg.RelationsPathFor(target), drel.X())
		wb.drawingRels = append(wb.drawingRels, drel)

	case gooxml.ChartType:
		chart := crt.NewChartSpace()
		decMap.AddTarget(target, chart)
		wb.charts = append(wb.charts, chart)

	default:
		fmt.Println("unsupported relationship", target, typ)
	}
	return nil
}
