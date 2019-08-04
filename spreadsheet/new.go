// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"runtime"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/common"
	"github.com/unidoc/unioffice/schema/soo/sml"
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

	wb.Rels.AddRelationship(unioffice.RelativeFilename(unioffice.DocTypeSpreadsheet, "", unioffice.ExtendedPropertiesType, 0), unioffice.ExtendedPropertiesType)
	wb.Rels.AddRelationship(unioffice.RelativeFilename(unioffice.DocTypeSpreadsheet, "", unioffice.CorePropertiesType, 0), unioffice.CorePropertiesType)
	wb.Rels.AddRelationship(unioffice.RelativeFilename(unioffice.DocTypeSpreadsheet, "", unioffice.OfficeDocumentType, 0), unioffice.OfficeDocumentType)
	wb.wbRels.AddRelationship(unioffice.RelativeFilename(unioffice.DocTypeSpreadsheet, unioffice.OfficeDocumentType, unioffice.StylesType, 0), unioffice.StylesType)

	wb.ContentTypes = common.NewContentTypes()
	wb.ContentTypes.AddDefault("vml", unioffice.VMLDrawingContentType)
	wb.ContentTypes.AddOverride(unioffice.AbsoluteFilename(unioffice.DocTypeSpreadsheet, unioffice.OfficeDocumentType, 0), "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml")
	wb.ContentTypes.AddOverride(unioffice.AbsoluteFilename(unioffice.DocTypeSpreadsheet, unioffice.StylesType, 0), unioffice.SMLStyleSheetContentType)

	wb.SharedStrings = NewSharedStrings()
	wb.ContentTypes.AddOverride(unioffice.AbsoluteFilename(unioffice.DocTypeSpreadsheet, unioffice.SharedStingsType, 0), unioffice.SharedStringsContentType)
	wb.wbRels.AddRelationship(unioffice.RelativeFilename(unioffice.DocTypeSpreadsheet, unioffice.OfficeDocumentType, unioffice.SharedStingsType, 0), unioffice.SharedStingsType)

	return wb
}
