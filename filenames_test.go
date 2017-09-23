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
		ExpRel string
		ExpAbs string
	}{
		{0, gooxml.CorePropertiesType, "docProps/core.xml", "docProps/core.xml"},
		{0, gooxml.ExtendedPropertiesType, "docProps/app.xml", "docProps/app.xml"},
		{0, gooxml.ThumbnailType, "docProps/thumbnail.jpeg", "docProps/thumbnail.jpeg"},
		{0, gooxml.StylesType, "styles.xml", "word/styles.xml"},

		{0, gooxml.OfficeDocumentType, "word/document.xml", "word/document.xml"},
		{0, gooxml.FontTableType, "fontTable.xml", "word/fontTable.xml"},
		{0, gooxml.EndNotesType, "endnotes.xml", "word/endnotes.xml"},
		{0, gooxml.FootNotesType, "footnotes.xml", "word/footnotes.xml"},
		{0, gooxml.NumberingType, "numbering.xml", "word/numbering.xml"},
		{0, gooxml.WebSettingsType, "webSettings.xml", "word/webSettings.xml"},
		{0, gooxml.SettingsType, "settings.xml", "word/settings.xml"},
		{23, gooxml.HeaderType, "header23.xml", "word/header23.xml"},
		{15, gooxml.FooterType, "footer15.xml", "word/footer15.xml"},
		{1, gooxml.ThemeType, "theme/theme1.xml", "word/theme/theme1.xml"},
	}
	for _, tc := range td {
		rel := gooxml.RelativeFilename(gooxml.DocTypeDocument, tc.Type, tc.Idx)
		abs := gooxml.AbsoluteFilename(gooxml.DocTypeDocument, tc.Type, tc.Idx)
		if rel != tc.ExpRel {
			t.Errorf("expected relative filename of %s for document %s/%d, got %s", tc.ExpRel, tc.Type, tc.Idx, rel)
		}
		if abs != tc.ExpAbs {
			t.Errorf("expected absolute filename of %s for document %s/%d, got %s", tc.ExpAbs, tc.Type, tc.Idx, abs)
		}
	}
}

func TestSMLFilenames(t *testing.T) {
	td := []struct {
		Idx    int
		Type   string
		ExpRel string
		ExpAbs string
	}{
		{0, gooxml.CorePropertiesType, "docProps/core.xml", "docProps/core.xml"},
		{0, gooxml.ExtendedPropertiesType, "docProps/app.xml", "docProps/app.xml"},
		{0, gooxml.ThumbnailType, "docProps/thumbnail.jpeg", "docProps/thumbnail.jpeg"},
		{0, gooxml.StylesType, "styles.xml", "xl/styles.xml"},

		{0, gooxml.OfficeDocumentType, "xl/workbook.xml", "xl/workbook.xml"},
		{15, gooxml.ChartType, "../charts/chart15.xml", "xl/charts/chart15.xml"},
		{12, gooxml.DrawingType, "../drawings/drawing12.xml", "xl/drawings/drawing12.xml"},
		{13, gooxml.TableType, "../tables/table13.xml", "xl/tables/table13.xml"},
		{2, gooxml.CommentsType, "../comments2.xml", "xl/comments2.xml"},
		{15, gooxml.WorksheetType, "worksheets/sheet15.xml", "xl/worksheets/sheet15.xml"},
		{2, gooxml.VMLDrawingType, "../drawings/vmlDrawing2.vml", "xl/drawings/vmlDrawing2.vml"},
		{0, gooxml.SharedStingsType, "sharedStrings.xml", "xl/sharedStrings.xml"},
		{1, gooxml.ThemeType, "theme/theme1.xml", "xl/theme/theme1.xml"},
	}
	for _, tc := range td {
		rel := gooxml.RelativeFilename(gooxml.DocTypeSpreadsheet, tc.Type, tc.Idx)
		abs := gooxml.AbsoluteFilename(gooxml.DocTypeSpreadsheet, tc.Type, tc.Idx)
		if rel != tc.ExpRel {
			t.Errorf("expected relative filename of %s for document %s/%d, got %s", tc.ExpRel, tc.Type, tc.Idx, rel)
		}
		if abs != tc.ExpAbs {
			t.Errorf("expected absolute filename of %s for document %s/%d, got %s", tc.ExpAbs, tc.Type, tc.Idx, abs)
		}
	}
}
