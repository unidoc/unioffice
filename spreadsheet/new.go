// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"runtime"

	"baliance.com/gooxml"
	"baliance.com/gooxml/common"
	"baliance.com/gooxml/schema/soo/sml"
)

// New constructs a new workbook.
func New() *Workbook {
	wb := &Workbook{}
	wb.x = sml.NewWorkbook()

	runtime.SetFinalizer(wb, workbookFinalizer)

	wb.AppProperties = common.NewAppProperties()
	wb.CoreProperties = common.NewCoreProperties()
	wb.StyleSheet = NewStyleSheet(wb)

	wb.Rels = common.NewRelationships()
	wb.wbRels = common.NewRelationships()

	wb.Rels.AddRelationship(gooxml.RelativeFilename(gooxml.DocTypeSpreadsheet, "", gooxml.ExtendedPropertiesType, 0), gooxml.ExtendedPropertiesType)
	wb.Rels.AddRelationship(gooxml.RelativeFilename(gooxml.DocTypeSpreadsheet, "", gooxml.CorePropertiesType, 0), gooxml.CorePropertiesType)
	wb.Rels.AddRelationship(gooxml.RelativeFilename(gooxml.DocTypeSpreadsheet, "", gooxml.OfficeDocumentType, 0), gooxml.OfficeDocumentType)
	wb.wbRels.AddRelationship(gooxml.RelativeFilename(gooxml.DocTypeSpreadsheet, gooxml.OfficeDocumentType, gooxml.StylesType, 0), gooxml.StylesType)

	wb.ContentTypes = common.NewContentTypes()
	wb.ContentTypes.AddDefault("vml", gooxml.VMLDrawingContentType)
	wb.ContentTypes.AddOverride(gooxml.AbsoluteFilename(gooxml.DocTypeSpreadsheet, gooxml.OfficeDocumentType, 0), "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml")
	wb.ContentTypes.AddOverride(gooxml.AbsoluteFilename(gooxml.DocTypeSpreadsheet, gooxml.StylesType, 0), gooxml.SMLStyleSheetContentType)

	wb.SharedStrings = NewSharedStrings()
	wb.ContentTypes.AddOverride(gooxml.AbsoluteFilename(gooxml.DocTypeSpreadsheet, gooxml.SharedStingsType, 0), gooxml.SharedStringsContentType)
	wb.wbRels.AddRelationship(gooxml.RelativeFilename(gooxml.DocTypeSpreadsheet, gooxml.OfficeDocumentType, gooxml.SharedStingsType, 0), gooxml.SharedStingsType)

	return wb
}
