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
	"time"

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
	wsRel := common.NewRelationships()
	wb.xwsRels = append(wb.xwsRels, wsRel)
	ws.SheetData = sml.NewCT_SheetData()

	dt := gooxml.DocTypeSpreadsheet
	// update the references
	rid := wb.wbRels.AddAutoRelationship(dt, len(wb.x.Sheets.Sheet), gooxml.WorksheetType)
	rs.IdAttr = rid.ID()

	// add the content type
	wb.ContentTypes.AddOverride(gooxml.AbsoluteFilename(dt, gooxml.WorksheetContentType, len(wb.x.Sheets.Sheet)),
		gooxml.WorksheetContentType)

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

// Uses1904Dates returns true if the the workbook uses dates relative to
// 1 Jan 1904. This is uncommon.
func (wb *Workbook) Uses1904Dates() bool {
	if wb.x.WorkbookPr == nil || wb.x.WorkbookPr.Date1904Attr == nil {
		return false
	}
	return *wb.x.WorkbookPr.Date1904Attr
}

// Epoch returns the point at which the dates/times in the workbook are relative to.
func (wb *Workbook) Epoch() time.Time {
	if wb.Uses1904Dates() {
		time.Date(1904, 1, 1, 0, 0, 0, 0, time.UTC)
	}
	return time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)
}

// Save writes the workbook out to a writer in the zipped xlsx format.
func (wb *Workbook) Save(w io.Writer) error {
	z := zip.NewWriter(w)
	defer z.Close()
	dt := gooxml.DocTypeSpreadsheet

	if err := zippkg.MarshalXML(z, zippkg.BaseRelsFilename, wb.Rels.X()); err != nil {
		return err
	}
	if err := zippkg.MarshalXMLByType(z, dt, gooxml.ExtendedPropertiesType, wb.AppProperties.X()); err != nil {
		return err
	}
	if err := zippkg.MarshalXMLByType(z, dt, gooxml.CorePropertiesType, wb.CoreProperties.X()); err != nil {
		return err
	}

	workbookFn := gooxml.AbsoluteFilename(dt, gooxml.OfficeDocumentType, 0)
	if err := zippkg.MarshalXML(z, workbookFn, wb.x); err != nil {
		return err
	}
	if err := zippkg.MarshalXML(z, zippkg.RelationsPathFor(workbookFn), wb.wbRels.X()); err != nil {
		return err
	}

	if err := zippkg.MarshalXMLByType(z, dt, gooxml.StylesType, wb.StyleSheet.X()); err != nil {
		return err
	}

	for i, thm := range wb.themes {
		if err := zippkg.MarshalXMLByTypeIndex(z, dt, gooxml.ThemeType, i+1, thm); err != nil {
			return err
		}
	}
	for i, sheet := range wb.xws {
		fn := gooxml.AbsoluteFilename(dt, gooxml.WorksheetType, i+1)
		zippkg.MarshalXML(z, fn, sheet)
		zippkg.MarshalXML(z, zippkg.RelationsPathFor(fn), wb.xwsRels[i].X())
	}
	if err := zippkg.MarshalXMLByType(z, dt, gooxml.SharedStingsType, wb.SharedStrings.X()); err != nil {
		return err
	}

	if wb.Thumbnail != nil {
		fn := gooxml.AbsoluteFilename(dt, gooxml.ThumbnailType, 0)
		tn, err := z.Create(fn)
		if err != nil {
			return err
		}
		if err := jpeg.Encode(tn, wb.Thumbnail, nil); err != nil {
			return err
		}
	}
	for i, chart := range wb.charts {
		fn := gooxml.AbsoluteFilename(dt, gooxml.ChartType, i+1)
		zippkg.MarshalXML(z, fn, chart)
	}
	for i, drawing := range wb.drawings {
		fn := gooxml.AbsoluteFilename(dt, gooxml.DrawingType, i+1)
		zippkg.MarshalXML(z, fn, drawing)
		if !wb.drawingRels[i].IsEmpty() {
			zippkg.MarshalXML(z, zippkg.RelationsPathFor(fn), wb.drawingRels[i].X())
		}
	}
	if err := zippkg.MarshalXML(z, zippkg.ContentTypesFilename, wb.ContentTypes.X()); err != nil {
		return err
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

		if err := sw.Validate(); err != nil {
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
	dt := gooxml.DocTypeSpreadsheet

	switch typ {
	case gooxml.OfficeDocumentType:
		wb.x = sml.NewWorkbook()
		decMap.AddTarget(target, wb.x)
		// look for the workbook relationships file as well
		wb.wbRels = common.NewRelationships()
		decMap.AddTarget(zippkg.RelationsPathFor(target), wb.wbRels.X())
		rel.TargetAttr = gooxml.RelativeFilename(dt, typ, 0)

	case gooxml.CorePropertiesType:
		decMap.AddTarget(target, wb.CoreProperties.X())
		rel.TargetAttr = gooxml.RelativeFilename(dt, typ, 0)

	case gooxml.ExtendedPropertiesType:
		decMap.AddTarget(target, wb.AppProperties.X())
		rel.TargetAttr = gooxml.RelativeFilename(dt, typ, 0)

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
		rel.TargetAttr = gooxml.RelativeFilename(dt, typ, len(wb.xws))

	case gooxml.StylesType:
		wb.StyleSheet = NewStyleSheet(wb)
		decMap.AddTarget(target, wb.StyleSheet.X())
		rel.TargetAttr = gooxml.RelativeFilename(dt, typ, 0)

	case gooxml.ThemeType:
		thm := dml.NewTheme()
		wb.themes = append(wb.themes, thm)
		decMap.AddTarget(target, thm)
		rel.TargetAttr = gooxml.RelativeFilename(dt, typ, len(wb.themes))

	case gooxml.SharedStingsType:
		wb.SharedStrings = NewSharedStrings()
		decMap.AddTarget(target, wb.SharedStrings.X())
		rel.TargetAttr = gooxml.RelativeFilename(dt, typ, 0)

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
		rel.TargetAttr = gooxml.RelativeFilename(dt, typ, len(wb.drawings))

	case gooxml.ChartType:
		chart := crt.NewChartSpace()
		decMap.AddTarget(target, chart)
		wb.charts = append(wb.charts, chart)
		rel.TargetAttr = gooxml.RelativeFilename(dt, typ, len(wb.charts))

	default:
		fmt.Println("unsupported relationship", target, typ)
	}
	return nil
}

func (wb *Workbook) AddDrawing() Drawing {
	drawing := sd.NewWsDr()
	wb.drawings = append(wb.drawings, drawing)
	fn := gooxml.AbsoluteFilename(gooxml.DocTypeSpreadsheet, gooxml.DrawingType, len(wb.drawings))
	wb.ContentTypes.AddOverride(fn, gooxml.DrawingContentType)
	wb.drawingRels = append(wb.drawingRels, common.NewRelationships())
	d := Drawing{wb, drawing}
	d.InitializeDefaults()
	return d
}

// AddDefinedName adds a name for a cell or range reference that can be used in
// formulas and charts.
func (wb *Workbook) AddDefinedName(name, ref string) DefinedName {
	if wb.x.DefinedNames == nil {
		wb.x.DefinedNames = sml.NewCT_DefinedNames()
	}
	dn := sml.NewCT_DefinedName()
	dn.Content = ref
	dn.NameAttr = name
	wb.x.DefinedNames.DefinedName = append(wb.x.DefinedNames.DefinedName, dn)
	return DefinedName{dn}
}

// DefinedNames returns a slice of all defined names in the workbook.
func (wb *Workbook) DefinedNames() []DefinedName {
	if wb.x.DefinedNames == nil {
		return nil
	}
	ret := []DefinedName{}
	for _, dn := range wb.x.DefinedNames.DefinedName {
		ret = append(ret, DefinedName{dn})
	}
	return ret
}
