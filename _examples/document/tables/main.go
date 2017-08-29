// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"baliance.com/gooxml/color"
	"baliance.com/gooxml/document"
	"baliance.com/gooxml/measurement"

	wml "baliance.com/gooxml/schema/schemas.openxmlformats.org/wordprocessingml"
)

func main() {
	doc := document.New()

	table := doc.AddTable()
	table.SetWidthPercent(100)
	row := table.AddRow()
	row.AddCell().AddParagraph().AddRun().AddText("Name")
	row.AddCell().AddParagraph().AddRun().AddText("John Smith")
	row = table.AddRow()
	row.AddCell().AddParagraph().AddRun().AddText("Street Address")
	row.AddCell().AddParagraph().AddRun().AddText("111 Country Road")

	table = doc.AddTable()
	table.SetWidth(4 * measurement.Inch)

	borders := table.Borders()
	borders.SetAll(wml.ST_BorderSingle, color.Auto, measurement.Zero)

	row = table.AddRow()
	row.AddCell().AddParagraph().AddRun().AddText("Name")
	row.AddCell().AddParagraph().AddRun().AddText("John Smith")
	row = table.AddRow()
	row.AddCell().AddParagraph().AddRun().AddText("Street Address")
	row.AddCell().AddParagraph().AddRun().AddText("111 Country Road")

	doc.SaveToFile("tables.docx")

}
