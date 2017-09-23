// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"log"
	"math/rand"
	"time"

	"baliance.com/gooxml/spreadsheet"
)

func main() {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	dateStyle := ss.StyleSheet.AddCellStyle()
	dateStyle.SetNumberFormatStandard(spreadsheet.StandardFormatDate)
	hdr := sheet.AddRow()
	hdr.AddCell().SetString("Date")
	hdr.AddCell().SetString("Product")
	hdr.AddCell().SetString("City")
	hdr.AddCell().SetString("Sold")
	hdr.AddCell().SetString("Remaining")
	products := []string{"Apples", "Oranges", "Grapes", "Strawberries", "Pears"}
	cities := []string{"New York", "Chicago", "Louisville", "New Orleans"}

	for r := 0; r < 25; r++ {
		row := sheet.AddRow()
		// and cells
		for c := 0; c < 5; c++ {
			cell := row.AddCell()
			switch c {
			case 0:
				cell.SetDate(randomDate())
				cell.SetStyle(dateStyle)
			case 1:
				cell.SetString(products[rand.Intn(len(products))])
			case 2:
				cell.SetString(cities[rand.Intn(len(cities))])
			case 3:
				cell.SetNumber(float64(rand.Intn(1000) + 50))
			case 4:
				cell.SetNumber(float64(rand.Intn(1000) + 50))
			}
		}
	}

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}

	ss.SaveToFile("pivot-table.xlsx")
}

func randomDate() time.Time {
	return time.Date(2017, time.Month(rand.Intn(12)+1), rand.Intn(28)+1, rand.Intn(24)+1, rand.Intn(60)+1, rand.Intn(60)+1, 0, time.Local)
}
