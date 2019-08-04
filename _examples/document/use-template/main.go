// Copyright 2017 FoxyUtils ehf. All rights reserved.

package main

import (
	"fmt"
	"log"

	"github.com/unidoc/unioffice/document"
)

var lorem = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin lobortis, lectus dictum feugiat tempus, sem neque finibus enim, sed eleifend sem nunc ac diam. Vestibulum tempus sagittis elementum`

func main() {
	// When Word saves a document, it removes all unused styles.  This means to
	// copy the styles from an existing document, you must first create a
	// document that contains text in each style of interest.  As an example,
	// see the template.docx in this directory.  It contains a paragraph set in
	// each style that Word supports by default.
	doc, err := document.OpenTemplate("template.docx")
	if err != nil {
		log.Fatalf("error opening Windows Word 2016 document: %s", err)
	}

	// We can now print out all styles in the document, verifying that they
	// exist.
	for _, s := range doc.Styles.Styles() {
		fmt.Println("style", s.Name(), "has ID of", s.StyleID(), "type is", s.Type())
	}

	// And create documents setting their style to the style ID (not style name).
	para := doc.AddParagraph()
	para.SetStyle("Title")
	para.AddRun().AddText("My Document Title")

	para = doc.AddParagraph()
	para.SetStyle("Subtitle")
	para.AddRun().AddText("Document Subtitle")

	para = doc.AddParagraph()
	para.SetStyle("Heading1")
	para.AddRun().AddText("Major Section")
	para = doc.AddParagraph()
	para = doc.AddParagraph()
	for i := 0; i < 4; i++ {
		para.AddRun().AddText(lorem)
	}

	para = doc.AddParagraph()
	para.SetStyle("Heading2")
	para.AddRun().AddText("Minor Section")
	para = doc.AddParagraph()
	for i := 0; i < 4; i++ {
		para.AddRun().AddText(lorem)
	}

	// using a pre-defined table style
	table := doc.AddTable()
	table.Properties().SetWidthPercent(90)
	table.Properties().SetStyle("GridTable4-Accent1")
	look := table.Properties().TableLook()
	// these have default values in the style, so we manually turn some of them off
	look.SetFirstColumn(false)
	look.SetFirstRow(true)
	look.SetLastColumn(false)
	look.SetLastRow(true)
	look.SetHorizontalBanding(true)

	for r := 0; r < 5; r++ {
		row := table.AddRow()
		for c := 0; c < 5; c++ {
			cell := row.AddCell()
			cell.AddParagraph().AddRun().AddText(fmt.Sprintf("row %d col %d", r+1, c+1))
		}
	}
	doc.SaveToFile("use-template.docx")
}
