// Copyright 2017 Baliance. All rights reserved.

package main

import (
	"baliance.com/gooxml/document"
	"baliance.com/gooxml/measurement"
	"baliance.com/gooxml/schema/soo/wml"
)

func main() {
	doc := document.New()

	hdr := doc.AddHeader()
	para := hdr.AddParagraph()
	para.Properties().AddTabStop(2.5*measurement.Inch, wml.ST_TabJcCenter, wml.ST_TabTlcNone)
	run := para.AddRun()
	run.AddTab()
	run.AddText("My Document Title")

	// Headers and footers are not immediately associated with a document as a
	// document can have multiple headers and footers for different sections.
	doc.BodySection().SetHeader(hdr, wml.ST_HdrFtrDefault)

	ftr := doc.AddFooter()
	para = ftr.AddParagraph()
	para.Properties().AddTabStop(6*measurement.Inch, wml.ST_TabJcRight, wml.ST_TabTlcNone)
	run = para.AddRun()
	run.AddText("Some subtitle goes here")
	run.AddTab()
	run.AddText("Pg ")
	run.AddField(document.FieldCurrentPage)
	run.AddText(" of ")
	run.AddField(document.FieldNumberOfPages)
	doc.BodySection().SetFooter(ftr, wml.ST_HdrFtrDefault)

	lorem := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin lobortis, lectus dictum feugiat tempus, sem neque finibus enim, sed eleifend sem nunc ac diam. Vestibulum tempus sagittis elementum`

	for i := 0; i < 5; i++ {
		para = doc.AddParagraph()
		run = para.AddRun()
		run.AddText(lorem)
	}

	doc.SaveToFile("header-footer.docx")
}
