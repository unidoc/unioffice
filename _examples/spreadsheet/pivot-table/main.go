// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"baliance.com/gooxml/schema/soo/sml"

	"baliance.com/gooxml/spreadsheet"
)

func main() {
	wb := spreadsheet.New()
	sheet := wb.AddSheet()

	dateStyle := wb.StyleSheet.AddCellStyle()
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
				cell.SetFormulaRaw(fmt.Sprintf("%d*%d", rand.Intn(25)+1, rand.Intn(25)+1))
			case 4:
				cell.SetNumber(float64(rand.Intn(1000) + 50))
			}
		}
	}

	pivot := wb.AddPivotTable()
	pivot.SetLocation("H5:M20")
	pivot.AddPivotField()
	pivot.AddPivotField()
	pivot.AddPivotField()
	pivot.AddPivotField()
	pivot.AddPivotField()

	pivot.AddRowField().SetX(2)
	pivot.AddColumnFIeld().SetX(-2)

	d1 := pivot.AddDataField()
	d1.SetField(3)
	d1.SetSubtotal(sml.ST_DataConsolidateFunctionAverage)

	d2 := pivot.AddDataField()
	d2.SetField(4)
	d2.SetSubtotal(sml.ST_DataConsolidateFunctionAverage)

	sheet.AddPivotTable(pivot)
	pivot.SetSource(sheet, "A1:E26")
	pivot.Recalculate()
	if err := wb.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}

	wb.SaveToFile("pivot-table.xlsx")
}

func randomDate() time.Time {
	return time.Date(2017, time.Month(rand.Intn(12)+1), rand.Intn(28)+1, rand.Intn(24)+1, rand.Intn(60)+1, rand.Intn(60)+1, 0, time.Local)
}
