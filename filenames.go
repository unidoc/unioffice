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

type DocType byte

const (
	Unknown DocType = iota
	DocTypeSpreadsheet
	DocTypeDocument
	DocTypePresentation
)

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
	case ChartType:
		return fmt.Sprintf("../charts/chart%d.xml", index)
	case DrawingType:
		return fmt.Sprintf("drawings/drawing%d.xml", index)

	case ThemeType, ThemeContentType:
		return fmt.Sprintf("theme/theme%d.xml", index)
	case OfficeDocumentType:
		switch dt {
		case DocTypeSpreadsheet:
			return "xl/workbook.xml"
		default:
			log.Printf("unsupported type %s pair and %v", typ, dt)
		}

		// SML
	case WorksheetType, WorksheetContentType:
		return fmt.Sprintf("worksheets/sheet%d.xml", index)

	case SharedStingsType, SharedStringsContentType:
		return "sharedStrings.xml"
	default:
		log.Printf("unsupported type %s", typ)
	}
	return ""
}

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
		default:
			log.Printf("unsupported type %s pair and %v", typ, dt)
		}

	case ThemeType, ThemeContentType:
		switch dt {
		case DocTypeSpreadsheet:
			return fmt.Sprintf("xl/theme/theme%d.xml", index)
		default:
			log.Printf("unsupported type %s pair and %v", typ, dt)
		}

	case StylesType:
		switch dt {
		case DocTypeSpreadsheet:
			return "xl/styles.xml"
		}

	case ChartType:
		switch dt {
		case DocTypeSpreadsheet:
			return fmt.Sprintf("xl/charts/chart%d.xml", index)
		default:
			log.Printf("unsupported type %s pair and %v", typ, dt)
		}

	case DrawingType:
		switch dt {
		case DocTypeSpreadsheet:
			return fmt.Sprintf("xl/drawings/drawing%d.xml", index)
		default:
			log.Fatalf("unsupported type %s pair and %v", typ, dt)
		}
	// SML
	case WorksheetType, WorksheetContentType:
		return fmt.Sprintf("xl/worksheets/sheet%d.xml", index)
	case SharedStingsType, SharedStringsContentType:
		return "xl/sharedStrings.xml"

	default:
		log.Fatalf("unsupported type %s", typ)
	}
	return ""
}
