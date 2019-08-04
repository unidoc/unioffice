// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

var lorem = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin lobortis, lectus dictum feugiat tempus, sem neque finibus enim, sed eleifend sem nunc ac diam. Vestibulum tempus sagittis elementum`

func main() {
	doc := document.New()

	// Headers/footers apply to the preceding paragraphs in the document. There
	// is a section properties on the document body itself acessible via
	// BodySection().  To have multiple different headers (aside from the
	// supported even/odd/first), we need to add multiple sections.

	// First add some content
	for i := 0; i < 5; i++ {
		para := doc.AddParagraph()
		run := para.AddRun()
		run.AddText(lorem)
	}

	// Construct our header
	hdr := doc.AddHeader()
	para := hdr.AddParagraph()
	para.Properties().AddTabStop(2.5*measurement.Inch, wml.ST_TabJcCenter, wml.ST_TabTlcNone)
	run := para.AddRun()
	run.AddTab()
	run.AddText("My Document Title")

	// Create a new section and apply the header
	para = doc.AddParagraph()
	section := para.Properties().AddSection(wml.ST_SectionMarkNextPage)
	section.SetHeader(hdr, wml.ST_HdrFtrDefault)

	// Add some more content
	for i := 0; i < 5; i++ {
		para := doc.AddParagraph()
		run := para.AddRun()
		run.AddText(lorem)
	}

	hdr = doc.AddHeader()
	para = hdr.AddParagraph()
	para.Properties().AddTabStop(2.5*measurement.Inch, wml.ST_TabJcCenter, wml.ST_TabTlcNone)
	run = para.AddRun()
	run.AddTab()
	run.AddText("Different Title")
	doc.BodySection().SetHeader(hdr, wml.ST_HdrFtrDefault)

	doc.SaveToFile("header-footer-multiple.docx")
}
