// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package gooxml

import (
	"fmt"
	"log"
	"strings"

	"baliance.com/gooxml/algo"
)

// Common filenames used in zip packages.
const (
	ContentTypesFilename = "[Content_Types].xml"
	BaseRelsFilename     = "_rels/.rels"
)

// DocType represents one of the three document types supported (docx/xlsx/pptx)
type DocType byte

// Document Type constants
const (
	Unknown DocType = iota
	DocTypeSpreadsheet
	DocTypeDocument
	DocTypePresentation
)

// RelativeFilename returns a filename relative to the source file referenced
// from a relationships file. Index is used in some cases for files which there
// may be more than one of (e.g. worksheets/drawings/charts)
func RelativeFilename(dt DocType, relToTyp, typ string, index int) string {
	orig := AbsoluteFilename(dt, typ, index)
	if relToTyp == "" {
		return orig
	}

	relTo := AbsoluteFilename(dt, relToTyp, index)
	relToSp := strings.Split(relTo, "/")
	origSp := strings.Split(orig, "/")

	// determine how many segments match
	matching := 0
	for i := 0; i < len(relToSp); i++ {
		if relToSp[i] == origSp[i] {
			matching++
		}
		if i+1 == len(origSp) {
			break
		}
	}
	relToSp = relToSp[matching:]
	origSp = origSp[matching:]
	nm := len(relToSp) - 1
	if nm > 0 {
		return algo.RepeatString("../", nm) + strings.Join(origSp, "/")
	}
	return strings.Join(origSp, "/")
}

// AbsoluteFilename returns the full path to a file from the root of the zip
// container. Index is used in some cases for files which there may be more than
// one of (e.g. worksheets/drawings/charts)
func AbsoluteFilename(dt DocType, typ string, index int) string {
	switch typ {
	case CorePropertiesType:
		return "docProps/core.xml"
	case ExtendedPropertiesType:
		return "docProps/app.xml"
	case ThumbnailType:
		return "docProps/thumbnail.jpeg"

	case OfficeDocumentType:
		switch dt {
		case DocTypeSpreadsheet:
			return "xl/workbook.xml"
		case DocTypeDocument:
			return "word/document.xml"
		default:
			log.Printf("unsupported type %s pair and %v", typ, dt)
		}

	case ThemeType, ThemeContentType:
		switch dt {
		case DocTypeSpreadsheet:
			return fmt.Sprintf("xl/theme/theme%d.xml", index)
		case DocTypeDocument:
			return fmt.Sprintf("word/theme/theme%d.xml", index)
		default:
			log.Printf("unsupported type %s pair and %v", typ, dt)
		}

	case StylesType:
		switch dt {
		case DocTypeSpreadsheet:
			return "xl/styles.xml"
		case DocTypeDocument:
			return "word/styles.xml"
		default:
			log.Printf("unsupported type %s pair and %v", typ, dt)
		}

	case ChartType, ChartContentType:
		switch dt {
		case DocTypeSpreadsheet:
			return fmt.Sprintf("xl/charts/chart%d.xml", index)
		default:
			log.Printf("unsupported type %s pair and %v", typ, dt)
		}
	case TableType, TableContentType:
		return fmt.Sprintf("xl/tables/table%d.xml", index)

	case DrawingType, DrawingContentType:
		switch dt {
		case DocTypeSpreadsheet:
			return fmt.Sprintf("xl/drawings/drawing%d.xml", index)
		default:
			log.Printf("unsupported type %s pair and %v", typ, dt)
		}

	case CommentsType, CommentsContentType:
		switch dt {
		case DocTypeSpreadsheet:
			return fmt.Sprintf("xl/comments%d.xml", index)
		default:
			log.Printf("unsupported type %s pair and %v", typ, dt)
		}

	case VMLDrawingType, VMLDrawingContentType:
		switch dt {
		case DocTypeSpreadsheet:
			return fmt.Sprintf("xl/drawings/vmlDrawing%d.vml", index)
		default:
			log.Printf("unsupported type %s pair and %v", typ, dt)
		}

	// SML
	case WorksheetType, WorksheetContentType:
		return fmt.Sprintf("xl/worksheets/sheet%d.xml", index)
	case SharedStingsType, SharedStringsContentType:
		return "xl/sharedStrings.xml"
	case PivotTableType:
		return fmt.Sprintf("xl/pivotTables/pivotTable%d.xml", index)
	case PivotCacheDefinitionType:
		return fmt.Sprintf("xl/pivotCache/pivotCacheDefinition%d.xml", index)
	case PivotCacheRecordsType:
		return fmt.Sprintf("xl/pivotCache/pivotCacheRecords%d.xml", index)

	// WML
	case FontTableType:
		return "word/fontTable.xml"
	case EndNotesType:
		return "word/endnotes.xml"
	case FootNotesType:
		return "word/footnotes.xml"
	case NumberingType:
		return "word/numbering.xml"
	case WebSettingsType:
		return "word/webSettings.xml"
	case SettingsType:
		return "word/settings.xml"
	case HeaderType:
		return fmt.Sprintf("word/header%d.xml", index)
	case FooterType:
		return fmt.Sprintf("word/footer%d.xml", index)

	default:
		log.Printf("unsupported type %s", typ)
	}
	return ""
}
