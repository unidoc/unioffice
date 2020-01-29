// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package unioffice_test

import "testing"
import "github.com/unidoc/unioffice"

func TestWMLFilenames(t *testing.T) {
	td := []struct {
		Idx    int
		Type   string
		ExpAbs string
	}{
		{0, unioffice.CorePropertiesType, "docProps/core.xml"},
		{0, unioffice.ExtendedPropertiesType, "docProps/app.xml"},
		{0, unioffice.ThumbnailType, "docProps/thumbnail.jpeg"},
		{0, unioffice.StylesType, "word/styles.xml"},

		{0, unioffice.OfficeDocumentType, "word/document.xml"},
		{0, unioffice.FontTableType, "word/fontTable.xml"},
		{0, unioffice.EndNotesType, "word/endnotes.xml"},
		{0, unioffice.FootNotesType, "word/footnotes.xml"},
		{0, unioffice.NumberingType, "word/numbering.xml"},
		{0, unioffice.WebSettingsType, "word/webSettings.xml"},
		{0, unioffice.SettingsType, "word/settings.xml"},
		{23, unioffice.HeaderType, "word/header23.xml"},
		{15, unioffice.FooterType, "word/footer15.xml"},
		{1, unioffice.ThemeType, "word/theme/theme1.xml"},
	}
	for _, tc := range td {
		abs := unioffice.AbsoluteFilename(unioffice.DocTypeDocument, tc.Type, tc.Idx)
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
		{0, unioffice.CorePropertiesType, "docProps/core.xml"},
		{0, unioffice.ExtendedPropertiesType, "docProps/app.xml"},
		{0, unioffice.ThumbnailType, "docProps/thumbnail.jpeg"},
		{0, unioffice.StylesType, "xl/styles.xml"},

		{0, unioffice.OfficeDocumentType, "xl/workbook.xml"},
		{15, unioffice.ChartType, "xl/charts/chart15.xml"},
		{12, unioffice.DrawingType, "xl/drawings/drawing12.xml"},
		{13, unioffice.TableType, "xl/tables/table13.xml"},
		{2, unioffice.CommentsType, "xl/comments2.xml"},
		{15, unioffice.WorksheetType, "xl/worksheets/sheet15.xml"},
		{2, unioffice.VMLDrawingType, "xl/drawings/vmlDrawing2.vml"},
		{0, unioffice.SharedStringsType, "xl/sharedStrings.xml"},
		{1, unioffice.ThemeType, "xl/theme/theme1.xml"},
		{2, unioffice.ImageType, "xl/media/image2.png"},
	}
	for _, tc := range td {
		abs := unioffice.AbsoluteFilename(unioffice.DocTypeSpreadsheet, tc.Type, tc.Idx)
		if abs != tc.ExpAbs {
			t.Errorf("expected absolute filename of %s for document %s/%d, got %s", tc.ExpAbs, tc.Type, tc.Idx, abs)
		}
	}
}

func TestPMLFilenames(t *testing.T) {
	td := []struct {
		Idx    int
		Type   string
		ExpAbs string
	}{
		{0, unioffice.CorePropertiesType, "docProps/core.xml"},
		{0, unioffice.ExtendedPropertiesType, "docProps/app.xml"},
		{0, unioffice.ThumbnailType, "docProps/thumbnail.jpeg"},
		{0, unioffice.StylesType, "ppt/styles.xml"},

		{0, unioffice.OfficeDocumentType, "ppt/presentation.xml"},
		{4, unioffice.SlideType, "ppt/slides/slide4.xml"},
		{5, unioffice.SlideLayoutType, "ppt/slideLayouts/slideLayout5.xml"},
		{6, unioffice.SlideMasterType, "ppt/slideMasters/slideMaster6.xml"},
		{7, unioffice.ThemeType, "ppt/theme/theme7.xml"},
	}
	for _, tc := range td {
		abs := unioffice.AbsoluteFilename(unioffice.DocTypePresentation, tc.Type, tc.Idx)
		if abs != tc.ExpAbs {
			t.Errorf("expected absolute filename of %s for document %s/%d, got %s", tc.ExpAbs, tc.Type, tc.Idx, abs)
		}
	}
}
