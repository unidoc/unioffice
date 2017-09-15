// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"

	"baliance.com/gooxml/spreadsheet"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	start := time.Now()
	ss := spreadsheet.New()
	nRows := 30000
	nCols := 100
	sheet := ss.AddSheet()

	// rows
	for r := 0; r < nRows; r++ {
		row := sheet.AddRow()
		// and cells
		for c := 0; c < nCols; c++ {
			cell := row.AddCell()
			cell.SetNumber(float64(r + c))
		}
	}

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}

	fmt.Printf("creating %d rows * %d cells took %s\n", nRows, nCols, time.Now().Sub(start))
	ss.SaveToFile("lots-of-rows.xlsx")

	start = time.Now()
	fmt.Printf("saving took %s\n", time.Now().Sub(start))

	start = time.Now()
	_, err := spreadsheet.Open("lots-of-rows.xlsx")
	if err != nil {
		log.Fatalf("error opening sheet: %s", err)
	}
	fmt.Printf("reading took %s\n", time.Now().Sub(start))
}
