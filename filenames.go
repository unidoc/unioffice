// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package unioffice

import (
	"fmt"
	"strings"

	"github.com/unidoc/unioffice/algo"
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
	case CustomPropertiesType:
		return "docProps/custom.xml"
	case ExtendedPropertiesType, ExtendedPropertiesTypeStrict:
		return "docProps/app.xml"
	case ThumbnailType, ThumbnailTypeStrict:
		return "docProps/thumbnail.jpeg"

	case OfficeDocumentType, OfficeDocumentTypeStrict:
		switch dt {
		case DocTypeSpreadsheet:
			return "xl/workbook.xml"
		case DocTypeDocument:
			return "word/document.xml"
		case DocTypePresentation:
			return "ppt/presentation.xml"
		default:
			Log("unsupported type %s pair and %v", typ, dt)
		}

	case ThemeType, ThemeTypeStrict, ThemeContentType:
		switch dt {
		case DocTypeSpreadsheet:
			return fmt.Sprintf("xl/theme/theme%d.xml", index)
		case DocTypeDocument:
			return fmt.Sprintf("word/theme/theme%d.xml", index)
		case DocTypePresentation:
			return fmt.Sprintf("ppt/theme/theme%d.xml", index)
		default:
			Log("unsupported type %s pair and %v", typ, dt)
		}

	case StylesType, StylesTypeStrict:
		switch dt {
		case DocTypeSpreadsheet:
			return "xl/styles.xml"
		case DocTypeDocument:
			return "word/styles.xml"
		case DocTypePresentation:
			return "ppt/styles.xml"
		default:
			Log("unsupported type %s pair and %v", typ, dt)
		}

	case ChartType, ChartTypeStrict, ChartContentType:
		switch dt {
		case DocTypeSpreadsheet:
			return fmt.Sprintf("xl/charts/chart%d.xml", index)
		default:
			Log("unsupported type %s pair and %v", typ, dt)
		}
	case TableType, TableTypeStrict, TableContentType:
		return fmt.Sprintf("xl/tables/table%d.xml", index)

	case DrawingType, DrawingTypeStrict, DrawingContentType:
		switch dt {
		case DocTypeSpreadsheet:
			return fmt.Sprintf("xl/drawings/drawing%d.xml", index)
		default:
			Log("unsupported type %s pair and %v", typ, dt)
		}

	case CommentsType, CommentsTypeStrict, CommentsContentType:
		switch dt {
		case DocTypeSpreadsheet:
			return fmt.Sprintf("xl/comments%d.xml", index)
		default:
			Log("unsupported type %s pair and %v", typ, dt)
		}

	case VMLDrawingType, VMLDrawingTypeStrict, VMLDrawingContentType:
		switch dt {
		case DocTypeSpreadsheet:
			return fmt.Sprintf("xl/drawings/vmlDrawing%d.vml", index)
		default:
			Log("unsupported type %s pair and %v", typ, dt)
		}

	case ImageType, ImageTypeStrict:
		switch dt {
		case DocTypeDocument:
			return fmt.Sprintf("word/media/image%d.png", index)
		case DocTypeSpreadsheet:
			return fmt.Sprintf("xl/media/image%d.png", index)
		case DocTypePresentation:
			return fmt.Sprintf("ppt/media/image%d.png", index)
		default:
			Log("unsupported type %s pair and %v", typ, dt)
		}
	// SML
	case WorksheetType, WorksheetTypeStrict, WorksheetContentType:
		return fmt.Sprintf("xl/worksheets/sheet%d.xml", index)
	case SharedStingsType, SharedStingsTypeStrict, SharedStringsContentType:
		return "xl/sharedStrings.xml"

	// WML
	case FontTableType, FontTableTypeStrict:
		return "word/fontTable.xml"
	case EndNotesType, EndNotesTypeStrict:
		return "word/endnotes.xml"
	case FootNotesType, FootNotesTypeStrict:
		return "word/footnotes.xml"
	case NumberingType, NumberingTypeStrict:
		return "word/numbering.xml"
	case WebSettingsType, WebSettingsTypeStrict:
		return "word/webSettings.xml"
	case SettingsType, SettingsTypeStrict:
		return "word/settings.xml"
	case HeaderType, HeaderTypeStrict:
		return fmt.Sprintf("word/header%d.xml", index)
	case FooterType, FooterTypeStrict:
		return fmt.Sprintf("word/footer%d.xml", index)

	// PML
	case SlideType, SlideTypeStrict:
		return fmt.Sprintf("ppt/slides/slide%d.xml", index)
	case SlideLayoutType:
		return fmt.Sprintf("ppt/slideLayouts/slideLayout%d.xml", index)
	case SlideMasterType:
		return fmt.Sprintf("ppt/slideMasters/slideMaster%d.xml", index)

	default:
		Log("unsupported type %s", typ)
	}
	return ""
}

// AbsoluteImageFilename returns the full path to an image from the root of the
// zip container.
func AbsoluteImageFilename(dt DocType, index int, fileExtension string) string {
	filename := AbsoluteFilename(dt, ImageType, index)
	// replace "png" with the actual file extension
	return filename[0:len(filename)-3] + fileExtension
}
