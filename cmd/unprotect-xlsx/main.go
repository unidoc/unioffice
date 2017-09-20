// Copyright 2017 Baliance. All rights reserved.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"

	"baliance.com/gooxml/spreadsheet"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatalf("pass a single document as a parameter")
	}
	fn := flag.Arg(0)
	fmt.Println("reading", fn)
	wb, err := spreadsheet.Open(fn)
	if err != nil {
		log.Fatalf("error opening: %s", err)
	}

	prot := wb.Protection()
	// just print out some info on what is protected
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	if pwh := prot.PasswordHash(); pwh != "" {
		fmt.Fprintf(tw, "password hash\t%v\n", pwh)
	}
	fmt.Fprintf(tw, "locked structure\t%v\n", prot.IsStructureLocked())
	fmt.Fprintf(tw, "locked windows\t%v\n", prot.IsWindowLocked())
	for _, s := range wb.Sheets() {
		fmt.Fprintf(tw, "Sheet '%s'\n", s.Name())
		sp := s.Protection()
		if pwh := sp.PasswordHash(); pwh != "" {
			fmt.Fprintf(tw, " - password hash\t%v\n", pwh)
		}
		fmt.Fprintf(tw, " - sheet locked\t%v\n", sp.IsSheetLocked())
		fmt.Fprintf(tw, " - objects locked\t%v\n", sp.IsObjectLocked())

		s.ClearProtection()
	}
	tw.Flush()

	// then clear protection and resave the workbook
	wb.ClearProtection()
	op := strings.Replace(fn, filepath.Ext(fn), "-unprotected.xlsx", 1)
	fmt.Println("saving unprotected workbook to", op)
	wb.SaveToFile(op)
}
