// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"log"
	"math"

	"baliance.com/gooxml/common"
	"baliance.com/gooxml/measurement"
	"baliance.com/gooxml/spreadsheet"
)

func main() {
	ss := spreadsheet.New()
	// add a single sheet
	sheet := ss.AddSheet()

	img, err := common.ImageFromFile("gophercolor.png")
	if err != nil {
		log.Fatalf("unable to create image: %s", err)
	}

	iref, err := ss.AddImage(img)
	if err != nil {
		log.Fatalf("unable to add image to workbook: %s", err)
	}

	dwng := ss.AddDrawing()
	sheet.SetDrawing(dwng)
	for i := float64(0); i < 360; i += 30 {
		anc := dwng.AddImage(iref, spreadsheet.AnchorTypeAbsolute)

		ang := i * math.Pi / 180
		x := 2 + 2*math.Cos(ang)
		y := 2 + +2*math.Sin(ang)
		anc.SetColOffset(measurement.Distance(x) * measurement.Inch)
		anc.SetRowOffset(measurement.Distance(y) * measurement.Inch)

		// set the image to 1x1 inches
		var w measurement.Distance = 1 * measurement.Inch
		anc.SetWidth(w)
		anc.SetHeight(iref.RelativeHeight(w))
	}

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}

	ss.SaveToFile("image.xlsx")
}
