// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package gooxml_test

import "testing"
import "baliance.com/gooxml"

func TestWMLFilenames(t *testing.T) {
	td := []struct {
		Idx    int
		Type   string
		ExpAbs string
	}{
		{0, gooxml.CorePropertiesType, "docProps/core.xml"},
		{0, gooxml.ExtendedPropertiesType, "docProps/app.xml"},
		{0, gooxml.ThumbnailType, "docProps/thumbnail.jpeg"},
		{0, gooxml.StylesType, "word/styles.xml"},

		{0, gooxml.OfficeDocumentType, "word/document.xml"},
		{0, gooxml.FontTableType, "word/fontTable.xml"},
		{0, gooxml.EndNotesType, "word/endnotes.xml"},
		{0, gooxml.FootNotesType, "word/footnotes.xml"},
		{0, gooxml.NumberingType, "word/numbering.xml"},
		{0, gooxml.WebSettingsType, "word/webSettings.xml"},
		{0, gooxml.SettingsType, "word/settings.xml"},
		{23, gooxml.HeaderType, "word/header23.xml"},
		{15, gooxml.FooterType, "word/footer15.xml"},
		{1, gooxml.ThemeType, "word/theme/theme1.xml"},
	}
	for _, tc := range td {
		abs := gooxml.AbsoluteFilename(gooxml.DocTypeDocument, tc.Type, tc.Idx)
		if abs != tc.ExpAbs {
			t.Errorf("expected absolute filename of %s for document %s/%d, got %s", tc.ExpAbs, tc.Type, tc.Idx, abs)
		}
	}
}

func TestSMLFilenames(t *testing.T) {
	td := []struct {
		Idx    int
		Type   string
		ExpAbs string
	}{
		{0, gooxml.CorePropertiesType, "docProps/core.xml"},
		{0, gooxml.ExtendedPropertiesType, "docProps/app.xml"},
		{0, gooxml.ThumbnailType, "docProps/thumbnail.jpeg"},
		{0, gooxml.StylesType, "xl/styles.xml"},

		{0, gooxml.OfficeDocumentType, "xl/workbook.xml"},
		{15, gooxml.ChartType, "xl/charts/chart15.xml"},
		{12, gooxml.DrawingType, "xl/drawings/drawing12.xml"},
		{13, gooxml.TableType, "xl/tables/table13.xml"},
		{2, gooxml.CommentsType, "xl/comments2.xml"},
		{15, gooxml.WorksheetType, "xl/worksheets/sheet15.xml"},
		{2, gooxml.VMLDrawingType, "xl/drawings/vmlDrawing2.vml"},
		{0, gooxml.SharedStingsType, "xl/sharedStrings.xml"},
		{1, gooxml.ThemeType, "xl/theme/theme1.xml"},
		{2, gooxml.PivotTableType, "xl/pivotTables/pivotTable2.xml"},
		{3, gooxml.PivotCacheDefinitionType, "xl/pivotCache/pivotCacheDefinition3.xml"},
		{4, gooxml.PivotCacheRecordsType, "xl/pivotCache/pivotCacheRecords4.xml"},
	}
	for _, tc := range td {
		abs := gooxml.AbsoluteFilename(gooxml.DocTypeSpreadsheet, tc.Type, tc.Idx)
		if abs != tc.ExpAbs {
			t.Errorf("expected absolute filename of %s for document %s/%d, got %s", tc.ExpAbs, tc.Type, tc.Idx, abs)
		}
	}
}
