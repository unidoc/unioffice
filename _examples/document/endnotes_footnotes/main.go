// Copyright 2017 FoxyUtils ehf. All rights reserved.

package main

import (
	"fmt"
	"log"

	"github.com/unidoc/unioffice/document"
)

func main() {
	doc, err := document.Open("footnotes_endnotes.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}

	if doc.HasFootnotes() {
		fmt.Printf("Document has %02d footnotes.\n", len(doc.Footnotes()))
	} else {
		fmt.Println("Document has no footnotes")
	}

	if doc.HasEndnotes() {
		fmt.Printf("Document has %02d endnotes.\n", len(doc.Endnotes()))
	} else {
		fmt.Println("Document has no endnotes")
	}
}
