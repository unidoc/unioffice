// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"log"

	"baliance.com/gooxml/spreadsheet"
)

func main() {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	dwng := ss.AddDrawing()
	chart, anc := dwng.AddChart(spreadsheet.AnchorTypeTwoCell)
	anc.MoveTo(0, 0)
	anc.SetWidth(10)

	// No cell data needed, we can supply data directly to the chart
	lc := chart.AddLineChart()
	priceSeries := lc.AddSeries()
	priceSeries.SetText("Price")
	priceSeries.CategoryAxis().SetValues([]string{"Prod 1", "Prod 2", "Prod 3", "Prod 4", "Prod 5"})
	priceSeries.Values().SetValues([]float64{5, 4, 3, 9, 2})

	soldSeries := lc.AddSeries()
	soldSeries.SetText("Sold")
	soldSeries.Values().SetValues([]float64{1, 2, 3, 4, 5})

	totalSeries := lc.AddSeries()
	totalSeries.SetText("Total")
	totalSeries.Values().SetValues([]float64{9, 2, 1, 8, 1})

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
	ss.SaveToFile("line-chart-no-data.xlsx")
}
