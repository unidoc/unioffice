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
	"strings"
	"time"

	"baliance.com/gooxml"
	"baliance.com/gooxml/common"
	"baliance.com/gooxml/vmldrawing"
	"baliance.com/gooxml/zippkg"

	"baliance.com/gooxml/schema/soo/dml"
	crt "baliance.com/gooxml/schema/soo/dml/chart"
	sd "baliance.com/gooxml/schema/soo/dml/spreadsheetDrawing"
	"baliance.com/gooxml/schema/soo/pkg/relationships"
	"baliance.com/gooxml/schema/soo/sml"
)

// ErrorNotFound is returned when something is not found
var ErrorNotFound = errors.New("not found")

// Workbook is the top level container item for a set of spreadsheets.
type Workbook struct {
	common.DocBase
	x *sml.Workbook

	StyleSheet    StyleSheet
	SharedStrings SharedStrings

	comments    []*sml.Comments
	xws         []*sml.Worksheet
	xwsRels     []common.Relationships
	wbRels      common.Relationships
	themes      []*dml.Theme
	drawings    []*sd.WsDr
	drawingRels []common.Relationships
	vmlDrawings []*vmldrawing.Container
	charts      []*crt.ChartSpace
	tables      []*sml.Table
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

	wb.comments = append(wb.comments, nil)

	dt := gooxml.DocTypeSpreadsheet

	// update the references
	rid := wb.wbRels.AddAutoRelationship(dt, gooxml.OfficeDocumentType, len(wb.x.Sheets.Sheet), gooxml.WorksheetType)
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

	if err := zippkg.MarshalXML(z, gooxml.BaseRelsFilename, wb.Rels.X()); err != nil {
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
		// recalculate sheet dimensions
		sheet.Dimension.RefAttr = Sheet{wb, nil, sheet}.Extents()

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
	for i, tbl := range wb.tables {
		fn := gooxml.AbsoluteFilename(dt, gooxml.TableType, i+1)
		zippkg.MarshalXML(z, fn, tbl)
	}
	for i, drawing := range wb.drawings {
		fn := gooxml.AbsoluteFilename(dt, gooxml.DrawingType, i+1)
		zippkg.MarshalXML(z, fn, drawing)
		if !wb.drawingRels[i].IsEmpty() {
			zippkg.MarshalXML(z, zippkg.RelationsPathFor(fn), wb.drawingRels[i].X())
		}
	}
	for i, drawing := range wb.vmlDrawings {
		zippkg.MarshalXML(z, gooxml.AbsoluteFilename(dt, gooxml.VMLDrawingType, i+1), drawing)
		// never seen relationships for a VML drawing yet
	}

	for i, img := range wb.Images {
		fn := fmt.Sprintf("xl/media/image%d.%s", i+1, img.Format())
		if img.Path() != "" {
			if err := zippkg.AddFileFromDisk(z, fn, img.Path()); err != nil {
				return err
			}
		} else {
			gooxml.Log("unsupported image source: %+v", img)
		}
	}

	if err := zippkg.MarshalXML(z, gooxml.ContentTypesFilename, wb.ContentTypes.X()); err != nil {
		return err
	}
	for i, cmt := range wb.comments {
		if cmt == nil {
			continue
		}
		zippkg.MarshalXML(z, gooxml.AbsoluteFilename(dt, gooxml.CommentsType, i+1), cmt)
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

func (wb *Workbook) onNewRelationship(decMap *zippkg.DecodeMap, target, typ string, files []*zip.File, rel *relationships.Relationship, src zippkg.Target) error {
	dt := gooxml.DocTypeSpreadsheet

	switch typ {
	case gooxml.OfficeDocumentType:
		wb.x = sml.NewWorkbook()
		decMap.AddTarget(target, wb.x, typ, 0)
		// look for the workbook relationships file as well
		wb.wbRels = common.NewRelationships()
		decMap.AddTarget(zippkg.RelationsPathFor(target), wb.wbRels.X(), typ, 0)
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, 0)

	case gooxml.CorePropertiesType:
		decMap.AddTarget(target, wb.CoreProperties.X(), typ, 0)
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, 0)

	case gooxml.ExtendedPropertiesType:
		decMap.AddTarget(target, wb.AppProperties.X(), typ, 0)
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, 0)

	case gooxml.WorksheetType:
		ws := sml.NewWorksheet()
		idx := uint32(len(wb.xws))
		wb.xws = append(wb.xws, ws)
		decMap.AddTarget(target, ws, typ, idx)
		// look for worksheet rels
		wksRel := common.NewRelationships()
		decMap.AddTarget(zippkg.RelationsPathFor(target), wksRel.X(), typ, 0)
		wb.xwsRels = append(wb.xwsRels, wksRel)

		// add a comments placeholder that will be replaced if we see a comments
		// relationship for the current sheet
		wb.comments = append(wb.comments, nil)

		// fix the relationship target so it points to where we'll save
		// the worksheet
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, len(wb.xws))

	case gooxml.StylesType:
		wb.StyleSheet = NewStyleSheet(wb)
		decMap.AddTarget(target, wb.StyleSheet.X(), typ, 0)
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, 0)

	case gooxml.ThemeType:
		thm := dml.NewTheme()
		wb.themes = append(wb.themes, thm)
		decMap.AddTarget(target, thm, typ, 0)
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, len(wb.themes))

	case gooxml.SharedStingsType:
		wb.SharedStrings = NewSharedStrings()
		decMap.AddTarget(target, wb.SharedStrings.X(), typ, 0)
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
				wb.Thumbnail, _, err = image.Decode(rc)
				rc.Close()
				if err != nil {
					return fmt.Errorf("error decoding thumbnail: %s", err)
				}
				files[i] = nil
			}
		}

	case gooxml.ImageType:
		for i, f := range files {
			if f == nil {
				continue
			}
			if f.Name == target {
				path, err := zippkg.ExtractToDiskTmp(f, wb.TmpPath)
				if err != nil {
					return err
				}
				img, err := common.ImageFromFile(path)
				if err != nil {
					return err
				}
				iref := common.MakeImageRef(img, &wb.DocBase, wb.wbRels)
				wb.Images = append(wb.Images, iref)
				files[i] = nil
			}
		}
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, len(wb.Images))

	case gooxml.DrawingType:
		drawing := sd.NewWsDr()
		idx := uint32(len(wb.drawings))
		decMap.AddTarget(target, drawing, typ, idx)
		wb.drawings = append(wb.drawings, drawing)

		drel := common.NewRelationships()
		decMap.AddTarget(zippkg.RelationsPathFor(target), drel.X(), typ, idx)
		wb.drawingRels = append(wb.drawingRels, drel)
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, len(wb.drawings))

	case gooxml.VMLDrawingType:
		vd := vmldrawing.NewContainer()
		idx := uint32(len(wb.vmlDrawings))
		decMap.AddTarget(target, vd, typ, idx)
		wb.vmlDrawings = append(wb.vmlDrawings, vd)

	case gooxml.CommentsType:
		wb.comments[src.Index] = sml.NewComments()
		decMap.AddTarget(target, wb.comments[src.Index], typ, src.Index)
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, len(wb.comments))

	case gooxml.ChartType:
		chart := crt.NewChartSpace()
		idx := uint32(len(wb.charts))
		decMap.AddTarget(target, chart, typ, idx)
		wb.charts = append(wb.charts, chart)
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, len(wb.charts))

	case gooxml.TableType:
		tbl := sml.NewTable()
		idx := uint32(len(wb.tables))
		decMap.AddTarget(target, tbl, typ, idx)
		wb.tables = append(wb.tables, tbl)
		rel.TargetAttr = gooxml.RelativeFilename(dt, src.Typ, typ, len(wb.tables))
	default:
		gooxml.Log("unsupported relationship %s %s", target, typ)
	}
	return nil
}

// AddDrawing adds a drawing to a workbook.  However the drawing is not actually
// displayed or used until it's set on a sheet.
func (wb *Workbook) AddDrawing() Drawing {
	drawing := sd.NewWsDr()
	wb.drawings = append(wb.drawings, drawing)
	fn := gooxml.AbsoluteFilename(gooxml.DocTypeSpreadsheet, gooxml.DrawingType, len(wb.drawings))
	wb.ContentTypes.AddOverride(fn, gooxml.DrawingContentType)
	wb.drawingRels = append(wb.drawingRels, common.NewRelationships())
	return Drawing{wb, drawing}
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

// RemoveDefinedName removes an existing defined name.
func (wb *Workbook) RemoveDefinedName(dn DefinedName) error {
	if dn.X() == nil {
		return errors.New("attempt to remove nil DefinedName")
	}
	for i, sdn := range wb.x.DefinedNames.DefinedName {
		if sdn == dn.X() {
			copy(wb.x.DefinedNames.DefinedName[i:], wb.x.DefinedNames.DefinedName[i+1:])
			wb.x.DefinedNames.DefinedName[len(wb.x.DefinedNames.DefinedName)-1] = nil
			wb.x.DefinedNames.DefinedName = wb.x.DefinedNames.DefinedName[:len(wb.x.DefinedNames.DefinedName)-1]
			return nil
		}
	}
	return errors.New("defined name not found")
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

// ClearCachedFormulaResults clears any computed formula values that are stored
// in the sheet. This may be required if you modify cells that are used as a
// formula input to force the formulas to be recomputed the next time the sheet
// is opened in Excel.
func (wb *Workbook) ClearCachedFormulaResults() {
	for _, s := range wb.Sheets() {
		s.ClearCachedFormulaResults()
	}
}

// RecalculateFormulas re-computes any computed formula values that are stored
// in the sheet. As gooxml formula support is still new and not all functins are
// supported,  if formula execution fails either due to a parse error or missing
// function, or erorr in the result (even if expected) the cached value will be
// left empty allowing Excel to recompute it on load.
func (wb *Workbook) RecalculateFormulas() {
	for _, s := range wb.Sheets() {
		s.RecalculateFormulas()
	}
}

// AddImage adds an image to the workbook package, returning a reference that
// can be used to add the image to a drawing.
func (wb *Workbook) AddImage(i common.Image) (common.ImageRef, error) {
	r := common.MakeImageRef(i, &wb.DocBase, wb.wbRels)
	if i.Path == "" {
		return r, errors.New("image must have a path")
	}

	if i.Format == "" {
		return r, errors.New("image must have a valid format")
	}
	if i.Size.X == 0 || i.Size.Y == 0 {
		return r, errors.New("image must have a valid size")
	}

	wb.Images = append(wb.Images, r)
	fn := fmt.Sprintf("media/image%d.%s", len(wb.Images), i.Format)
	wb.wbRels.AddRelationship(fn, gooxml.ImageType)
	return r, nil
}

// SetActiveSheet sets the active sheet which will be the tab displayed when the
// spreadsheet is initially opened.
func (wb *Workbook) SetActiveSheet(s Sheet) {
	for i, st := range wb.xws {
		if s.x == st {
			wb.SetActiveSheetIndex(uint32(i))
		}
	}
}

// SetActiveSheetIndex sets the index of the active sheet (0-n) which will be
// the tab displayed when the spreadsheet is initially opened.
func (wb *Workbook) SetActiveSheetIndex(idx uint32) {
	if wb.x.BookViews == nil {
		wb.x.BookViews = sml.NewCT_BookViews()
	}
	if len(wb.x.BookViews.WorkbookView) == 0 {
		wb.x.BookViews.WorkbookView = append(wb.x.BookViews.WorkbookView, sml.NewCT_BookView())
	}

	wb.x.BookViews.WorkbookView[0].ActiveTabAttr = gooxml.Uint32(idx)
}

// Tables returns a slice of all defined tables in the workbook.
func (wb *Workbook) Tables() []Table {
	if wb.tables == nil {
		return nil
	}
	ret := []Table{}
	for _, t := range wb.tables {
		ret = append(ret, Table{t})
	}
	return ret
}

// ClearProtection clears all workbook protections.
func (wb *Workbook) ClearProtection() {
	wb.x.WorkbookProtection = nil
}

// Protection allows control over the workbook protections.
func (wb *Workbook) Protection() WorkbookProtection {
	if wb.x.WorkbookProtection == nil {
		wb.x.WorkbookProtection = sml.NewCT_WorkbookProtection()
	}
	return WorkbookProtection{wb.x.WorkbookProtection}
}

// GetSheet returns a sheet by name, or an error if a sheet by the given name
// was not found.
func (wb *Workbook) GetSheet(name string) (Sheet, error) {
	for _, s := range wb.Sheets() {
		if s.Name() == name {
			return s, nil
		}
	}
	return Sheet{}, ErrorNotFound
}

func workbookFinalizer(wb *Workbook) {
	wb.Close()
}

// Close closes the workbook, removing any temporary files that might have been
// created when opening a document.
func (wb *Workbook) Close() error {
	if wb.TmpPath != "" && strings.HasPrefix(wb.TmpPath, os.TempDir()) {
		return os.RemoveAll(wb.TmpPath)
	}
	return nil
}
