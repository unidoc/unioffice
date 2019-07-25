// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"log"

	"github.com/unidoc/unioffice/measurement"

	"github.com/unidoc/unioffice/common"

	"github.com/unidoc/unioffice/presentation"
)

func main() {
	ppt := presentation.New()
	imgColor, err := common.ImageFromFile("gophercolor.png")
	if err != nil {
		log.Fatalf("unable to create image: %s", err)
	}
	imgBW, err := common.ImageFromFile("gopher.png")
	if err != nil {
		log.Fatalf("unable to create image: %s", err)
	}

	irefColor, err := ppt.AddImage(imgColor)
	if err != nil {
		log.Fatal(err)
	}

	irefBW, err := ppt.AddImage(imgBW)
	if err != nil {
		log.Fatal(err)
	}

	slide := ppt.AddSlide()

	ibColor := slide.AddImage(irefColor)
	ibColor.Properties().SetWidth(2 * measurement.Inch)
	ibColor.Properties().SetHeight(irefColor.RelativeHeight(2 * measurement.Inch))

	ibBW := slide.AddImage(irefBW)
	ibBW.Properties().SetWidth(2 * measurement.Inch)
	ibBW.Properties().SetHeight(irefBW.RelativeHeight(2 * measurement.Inch))
	ibBW.Properties().SetPosition(4*measurement.Inch, 4*measurement.Inch)

	if err := ppt.Validate(); err != nil {
		log.Fatal(err)
	}
	ppt.SaveToFile("image.pptx")
}
