// Copyright 2017 Baliance. All rights reserved.

package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"

	"baliance.com/gooxml/spreadsheet"
	"baliance.com/gooxml/spreadsheet/format"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatalf("pass a single document as a parameter")
	}
	f, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatalf("error opening: %s", err)
	}

	cv := csv.NewReader(f)
	wb := spreadsheet.New()
	sheet := wb.AddSheet()
	for {
		rec, err := cv.Read()
		if err != nil {
			break
		}
		// one row per CSV row
		row := sheet.AddRow()
		for _, c := range rec {
			cell := row.AddCell()
			// go ahead and use SetNumber/SetString so the cell types get set
			// correctly, in practice this doesn't matter too much as Excel will
			// automatically treat number-like string as numbers.
			if format.IsNumber(c) {
				v, _ := strconv.ParseFloat(c, 64)
				cell.SetNumber(v)
			} else {
				cell.SetString(c)
			}
		}

	}

	if err := wb.Validate(); err != nil {
		log.Fatalf("error validating spreadsheet: %s", err)
	}

	outFile := strings.Replace(flag.Arg(0), ".csv", ".xlsx", -1)
	if err := wb.SaveToFile(outFile); err != nil {
		log.Fatalf("error saving spreadsheet: %s", err)
	}
}
