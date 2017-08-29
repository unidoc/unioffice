// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"log"

	"baliance.com/gooxml/document"
	"baliance.com/gooxml/measurement"
	wd "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"
)

var lorem = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin lobortis, lectus dictum feugiat tempus, sem neque finibus enim, sed eleifend sem nunc ac diam. Vestibulum tempus sagittis elementum`

func main() {
	doc := document.New()

	img, err := document.ImageFromFile("gophercolor.png")
	if err != nil {
		log.Fatalf("unable to create image: %s", err)
	}

	iref, err := doc.AddImage(img)
	if err != nil {
		log.Fatalf("unable to add image to document: %s", err)
	}

	para := doc.AddParagraph()
	anchored, err := para.AddRun().AddDrawingAnchored(iref)
	if err != nil {
		log.Fatalf("unable to add anchored image: %s", err)
	}
	anchored.SetName("Gopher")
	anchored.SetSize(2*measurement.Inch, 2*measurement.Inch)
	anchored.SetOrigin(wd.ST_RelFromHPage, wd.ST_RelFromVTopMargin)
	anchored.SetHAlignment(wd.ST_AlignHCenter)
	anchored.SetYOffset(3 * measurement.Inch)
	anchored.SetTextWrapSquare(wd.ST_WrapTextBothSides)

	run := para.AddRun()
	for i := 0; i < 16; i++ {
		run.AddText(lorem)
	}
	doc.SaveToFile("image.docx")
}
