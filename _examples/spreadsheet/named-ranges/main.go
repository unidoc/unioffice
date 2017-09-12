// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"fmt"
	"log"

	"baliance.com/gooxml/spreadsheet"
)

func main() {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	// Create all of our data
	row := sheet.AddRow()
	row.AddCell().SetString("Item")
	row.AddCell().SetString("Price")
	row.AddCell().SetString("# Sold")
	row.AddCell().SetString("Total")
	for r := 0; r < 5; r++ {
		row := sheet.AddRow()
		row.AddCell().SetString(fmt.Sprintf("Product %d", r+1))
		row.AddCell().SetNumber(1.23 * float64(r+1))
		row.AddCell().SetNumber(float64(r%3 + 1))
		row.AddCell().SetFormulaRaw(fmt.Sprintf("C%d*B%d", r+2, r+2))
	}

	// create some defined names for various ranges that we can use
	// instead of the sheet/cell references
	productNames := ss.AddDefinedName("ProductNames", sheet.RangeReference("A2:A6"))
	prices := ss.AddDefinedName("Prices", sheet.RangeReference("B2:B6"))
	sold := ss.AddDefinedName("Sold", sheet.RangeReference("C2:C6"))
	total := ss.AddDefinedName("Total", sheet.RangeReference("D2:D6"))

	for _, dn := range ss.DefinedNames() {
		fmt.Println("- defined name", dn.Name(), "=", dn.Content())
	}
	// Charts need to reside in a drawing
	dwng := ss.AddDrawing()
	chart, anc := dwng.AddChart(spreadsheet.AnchorTypeTwoCell)
	anc.SetWidthCells(10)

	lc := chart.AddLineChart()
	priceSeries := lc.AddSeries()
	priceSeries.SetText("Price")

	// Set a category axis reference on the first series to pull the product names
	priceSeries.CategoryAxis().SetLabelReference(`'Sheet 1'!` + productNames.Name())
	priceSeries.Values().SetReference(`'Sheet 1'!` + prices.Name())

	soldSeries := lc.AddSeries()
	soldSeries.SetText("Sold")
	soldSeries.Values().SetReference(`'Sheet 1'!` + sold.Name())

	totalSeries := lc.AddSeries()
	totalSeries.SetText("Total")
	totalSeries.Values().SetReference(`'Sheet 1'!` + total.Name())

	// the line chart accepts up to two axes
	ca := chart.AddCategoryAxis()
	va := chart.AddValueAxis()
	lc.AddAxis(ca)
	lc.AddAxis(va)

	ca.SetCrosses(va)
	va.SetCrosses(ca)

	// add a title and legend
	title := chart.AddTitle()
	title.SetText("Items Sold")
	chart.AddLegend()

	// and finally add the chart to the sheet
	sheet.SetDrawing(dwng)

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}
	ss.SaveToFile("named-ranges.xlsx")
}
