// Copyright 2017 Baliance. All rights reserved.

package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"

	"baliance.com/gooxml/spreadsheet"
)

func main() {
	raw := flag.Bool("raw", false, "print raw values instead of formatted")
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatalf("pass a single document as a parameter")
	}
	wb, err := spreadsheet.Open(flag.Arg(0))
	if err != nil {
		log.Fatalf("error opening: %s", err)
	}

	for _, sheet := range wb.Sheets() {
		f, err := os.Create(sheet.Name() + ".csv")
		if err != nil {
			log.Fatalf("error creating sheet: %s", err)
		}
		cw := csv.NewWriter(f)
		sc, _, ec, _ := sheet.ExtentsIndex()
		scIdx := spreadsheet.ColumnToIndex(sc)
		ecIdx := spreadsheet.ColumnToIndex(ec)
		for _, r := range sheet.Rows() {
			record := []string{}
			for c := scIdx; c <= ecIdx; c++ {
				cell := r.Cell(spreadsheet.IndexToColumn(c))
				if !*raw {
					record = append(record, cell.GetFormattedValue())
				} else {
					v, _ := cell.GetRawValue()
					record = append(record, v)
				}
			}
			cw.Write(record)
		}
		cw.Flush()
		f.Close()
	}
}
