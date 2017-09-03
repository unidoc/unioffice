// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"baliance.com/gooxml/common"
	sml "baliance.com/gooxml/schema/schemas.openxmlformats.org/spreadsheetml"
	"baliance.com/gooxml/zippkg"
)

// New constructs a new workbook.
func New() *Workbook {
	wb := &Workbook{}
	wb.x = sml.NewWorkbook()

	wb.AppProperties = common.NewAppProperties()
	wb.CoreProperties = common.NewCoreProperties()
	wb.StyleSheet = NewStyleSheet()

	wb.Rels = common.NewRelationships()
	wb.wbRels = common.NewRelationships()
	wb.Rels.AddRelationship(zippkg.AppPropsFilename, common.ExtendedPropertiesType)
	wb.Rels.AddRelationship(zippkg.CorePropsFilename, common.CorePropertiesType)
	wb.Rels.AddRelationship("xl/workbook.xml", common.OfficeDocumentType)
	wb.wbRels.AddRelationship("styles.xml", common.StylesType)

	wb.ContentTypes = common.NewContentTypes()
	wb.ContentTypes.AddOverride("/xl/workbook.xml", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml")
	wb.ContentTypes.AddOverride("/xl/styles.xml", "application/vnd.openxmlformats-officedocument.spreadsheetml.styles+xml")

	wb.SharedStrings = NewSharedStrings()
	wb.ContentTypes.AddOverride("/xl/sharedStrings.xml", common.SharedStringsContentType)
	wb.wbRels.AddRelationship("sharedStrings.xml", common.SharedStingsType)

	return wb
}
