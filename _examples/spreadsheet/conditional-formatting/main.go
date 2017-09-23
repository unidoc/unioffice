// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"log"
	"math/rand"

	"baliance.com/gooxml/color"
	"baliance.com/gooxml/schema/soo/sml"
	"baliance.com/gooxml/spreadsheet"
)

func main() {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	for r := 0; r < 20; r++ {
		if r != 0 && r%5 == 0 {
			sheet.AddRow()
		}
		row := sheet.AddRow()
		for c := 0; c < 5; c++ {
			cell := row.AddCell()
			cell.SetNumber(float64(rand.Intn(1000)) / 100.0)
		}
	}
	{
		cfmt := sheet.AddConditionalFormatting([]string{"A1:E5"})
		r := cfmt.AddRule()
		// cell is
		r.SetType(sml.ST_CfTypeCellIs)
		// greater than
		r.SetOperator(sml.ST_ConditionalFormattingOperatorLessThan)
		// four
		r.SetConditionValue("4")
		// should be formatted with this style
		green := ss.StyleSheet.AddDifferentialStyle()
		green.Fill().SetPatternFill().SetBgColor(color.SuccessGreen)
		r.SetStyle(green)

		r = cfmt.AddRule()
		// cell is
		r.SetType(sml.ST_CfTypeCellIs)
		// greater than
		r.SetOperator(sml.ST_ConditionalFormattingOperatorGreaterThan)
		// four
		r.SetConditionValue("7")
		// should be formatted with this style
		red := ss.StyleSheet.AddDifferentialStyle()
		red.Fill().SetPatternFill().SetBgColor(color.Red)
		r.SetStyle(red)
	}
	{
		// Color gradient by value
		cfmt := sheet.AddConditionalFormatting([]string{"A7:E11"})
		r := cfmt.AddRule()
		cs := r.SetColorScale()
		cs.AddFormatValue(sml.ST_CfvoTypeMin, "0")
		cs.AddGradientStop(color.Red)
		cs.AddFormatValue(sml.ST_CfvoTypePercentile, "50")
		cs.AddGradientStop(color.Yellow)
		cs.AddFormatValue(sml.ST_CfvoTypeMax, "0")
		cs.AddGradientStop(color.SuccessGreen)
	}
	{
		// Icons
		cfmt := sheet.AddConditionalFormatting([]string{"A13:E17"})
		r := cfmt.AddRule()
		icons := r.SetIcons()
		icons.SetIcons(sml.ST_IconSetType3TrafficLights1)
		icons.AddFormatValue(sml.ST_CfvoTypePercent, "0")
		icons.AddFormatValue(sml.ST_CfvoTypePercent, "040")
		icons.AddFormatValue(sml.ST_CfvoTypePercent, "90")
	}
	{
		cfmt := sheet.AddConditionalFormatting([]string{"A19:E23"})
		r := cfmt.AddRule()
		db := r.SetDataBar()
		db.AddFormatValue(sml.ST_CfvoTypeMin, "0")
		db.AddFormatValue(sml.ST_CfvoTypeMax, "0")
		db.SetColor(color.Blue)
	}

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}

	ss.SaveToFile("conditional-formatting.xlsx")
}
