// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"baliance.com/gooxml/document"
	"baliance.com/gooxml/spreadsheet"
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
