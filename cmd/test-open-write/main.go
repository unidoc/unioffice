// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/spreadsheet"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("usage: test-open-write filename.docx/.xlsx")
		os.Exit(1)
	}
	fn := flag.Arg(0)
	if strings.HasSuffix(strings.ToLower(fn), ".xlsx") {
		xlsx, err := spreadsheet.Open(fn)
		if err != nil {
			log.Fatalf("error opening %s: %s", fn, err)
		}
		if err := xlsx.Validate(); err != nil {
			log.Printf("validation error: %s", err)
		}
		if err := xlsx.SaveToFile("converted.xlsx"); err != nil {
			log.Fatalf("error saving: %s", err)
		}
	} else if strings.HasSuffix(strings.ToLower(fn), ".docx") {
		docx, err := document.Open(fn)
		if err != nil {
			log.Fatalf("error opening %s: %s", fn, err)
		}
		if err := docx.Validate(); err != nil {
			log.Printf("validation error: %s", err)
		}
		if err := docx.SaveToFile("converted.docx"); err != nil {
			log.Fatalf("error saving: %s", err)
		}
	}
	fmt.Println("reading", fn)
}
