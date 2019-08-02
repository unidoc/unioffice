// Copyright 2017 FoxyUtils ehf. All rights reserved.

package main

import (
	"log"

	"github.com/unidoc/unioffice/common"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

func main() {
	doc := document.New()

	img, err := common.ImageFromFile("gophercolor.png")
	if err != nil {
		log.Fatalf("unable to create image: %s", err)
	}

	hdr := doc.AddHeader()
	// We need to add a reference of the image to the header instead of to the
	// document
	iref, err := hdr.AddImage(img)
	if err != nil {
		log.Fatalf("unable to to add image to document: %s", err)
	}

	para := hdr.AddParagraph()
	para.Properties().AddTabStop(2.5*measurement.Inch, wml.ST_TabJcCenter, wml.ST_TabTlcNone)
	run := para.AddRun()
	run.AddTab()
	run.AddText("My Document Title")

	imgInl, _ := para.AddRun().AddDrawingInline(iref)
	imgInl.SetSize(1*measurement.Inch, 1*measurement.Inch)

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
