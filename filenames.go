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

// RelativeFilename returns a filename relative to where it is normally
// referenced from a relationships file. Index is used in some cases for files
// which there may be more than one of (e.g. worksheets/drawings/charts)
func RelativeFilename(dt DocType, typ string, index int) string {
	switch typ {
	case CorePropertiesType:
		return "docProps/core.xml"
	case ExtendedPropertiesType:
		return "docProps/app.xml"
	case ThumbnailType:
		return "docProps/thumbnail.jpeg"

	case StylesType:
		return "styles.xml"
	case ChartType, ChartContentType:
		return fmt.Sprintf("../charts/chart%d.xml", index)
	case DrawingType, DrawingContentType:
		return fmt.Sprintf("../drawings/drawing%d.xml", index)
	case CommentsType, CommentsContentType:
		return fmt.Sprintf("../comments%d.xml", index)

	case VMLDrawingType, VMLDrawingContentType:
		return fmt.Sprintf("../drawings/vmlDrawing%d.vml", index)

	case TableType, TableContentType:
		return fmt.Sprintf("../tables/table%d.xml", index)

	case ThemeType, ThemeContentType:
		return fmt.Sprintf("theme/theme%d.xml", index)
	case OfficeDocumentType:
		switch dt {
		case DocTypeSpreadsheet:
			return "xl/workbook.xml"
		case DocTypeDocument:
			return "word/document.xml"
		default:
			log.Printf("unsupported type %s pair and %v", typ, dt)
		}

		// SML
	case WorksheetType, WorksheetContentType:
		return fmt.Sprintf("worksheets/sheet%d.xml", index)

	case SharedStingsType, SharedStringsContentType:
		return "sharedStrings.xml"

		// WML
	case FontTableType:
		return "fontTable.xml"
	case EndNotesType:
		return "endnotes.xml"
	case FootNotesType:
		return "footnotes.xml"
	case NumberingType:
		return "numbering.xml"
	case WebSettingsType:
		return "webSettings.xml"
	case SettingsType:
		return "settings.xml"
	case HeaderType:
		return fmt.Sprintf("header%d.xml", index)
	case FooterType:
		return fmt.Sprintf("footer%d.xml", index)
	default:
		log.Printf("unsupported type %s", typ)
	}
	return ""
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
