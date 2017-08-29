// Copyright 2017 Baliance. All rights reserved.

package main

import (
	"flag"
	"fmt"
	"log"

	"baliance.com/gooxml/document"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatalf("pass a single document as a parameter")
	}
	doc, err := document.Open(flag.Arg(0))
	if err != nil {
		log.Fatalf("error opening: %s", err)
	}
	for _, para := range doc.Paragraphs() {
		for _, run := range para.Runs() {
			if s := run.Text(); s != "" {
				fmt.Print(run.Text())
			}
		}
		fmt.Println()
	}
}
