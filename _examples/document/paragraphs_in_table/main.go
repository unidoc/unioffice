package main

import (
	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

func main() {
	doc := document.New()

	paraBeforeTable := doc.AddParagraph()
	paraBeforeTable.AddRun().AddText("before table")

	table := doc.InsertTableAfter(paraBeforeTable)
	table.Properties().Borders().SetAll(wml.ST_BorderBasicBlackDots, color.AliceBlue, measurement.Point*2)
	tablePara1 := table.AddRow().AddCell().AddParagraph()
	tablePara1.AddRun().AddText("table paragraph 1")

	paraAfterTable := doc.AddParagraph()
	paraAfterTable.AddRun().AddText("after table")

	tablePara2 := doc.InsertParagraphAfter(tablePara1)
	tablePara2.AddRun().AddText("table paragraph after table paragraph 1")

	tablePara3 := doc.InsertParagraphBefore(tablePara1)
	tablePara3.AddRun().AddText("table paragraph before table paragraph 1")

	doc.SaveToFile("out.docx")
}
