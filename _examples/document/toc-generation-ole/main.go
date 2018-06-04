// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"log"
	"os"
	"path/filepath"

	"baliance.com/gooxml/document"
	"baliance.com/gooxml/measurement"
	"baliance.com/gooxml/schema/soo/wml"

	"github.com/go-ole/go-ole/oleutil"
)

// NOTE: This example can only run on Windows and requires that Word be installed.

var lorem = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin lobortis, lectus dictum feugiat tempus, sem neque finibus enim, sed eleifend sem nunc ac diam. Vestibulum tempus sagittis elementum`

func main() {
	doc := document.New()

	// Force the TOC to update upon opening the document
	doc.Settings.SetUpdateFieldsOnOpen(true)

	// Add a TOC
	doc.AddParagraph().AddRun().AddField(document.FieldTOC)
	// followed by a page break
	doc.AddParagraph().Properties().AddSection(wml.ST_SectionMarkNextPage)

	nd := doc.Numbering.AddDefinition()
	for i := 0; i < 9; i++ {
		lvl := nd.AddLevel()
		lvl.SetFormat(wml.ST_NumberFormatNone)
		lvl.SetAlignment(wml.ST_JcLeft)
		if i%2 == 0 {
			lvl.SetFormat(wml.ST_NumberFormatBullet)
			lvl.RunProperties().SetFontFamily("Symbol")
			lvl.SetText("")
		}
		lvl.Properties().SetLeftIndent(0.5 * measurement.Distance(i) * measurement.Inch)
	}

	// and finally paragraphs at different heading levels
	for i := 0; i < 4; i++ {
		para := doc.AddParagraph()
		para.SetNumberingDefinition(nd)
		para.Properties().SetHeadingLevel(1)
		para.AddRun().AddText("First Level")

		doc.AddParagraph().AddRun().AddText(lorem)
		for i := 0; i < 3; i++ {
			para := doc.AddParagraph()
			para.SetNumberingDefinition(nd)
			para.Properties().SetHeadingLevel(2)
			para.AddRun().AddText("Second Level")
			doc.AddParagraph().AddRun().AddText(lorem)

			para = doc.AddParagraph()
			para.SetNumberingDefinition(nd)
			para.Properties().SetHeadingLevel(3)
			para.AddRun().AddText("Third Level")
			doc.AddParagraph().AddRun().AddText(lorem)
		}
	}
	doc.SaveToFile("toc.docx")

	cwd, _ := os.Getwd()
	UpdateFields(filepath.Join(cwd, "toc.docx"))
}

// UpdateFields uses go-ole to convert a docx to a PDF using the Word application
func UpdateFields(source string) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	iunk, err := oleutil.CreateObject("Word.Application")
	if err != nil {
		log.Fatalf("error creating Word object: %s", err)
	}
	defer iunk.Release()

	word := iunk.MustQueryInterface(ole.IID_IDispatch)
	defer word.Release()

	docs := oleutil.MustGetProperty(word, "Documents").ToIDispatch()
	defer docs.Release()
	wordDoc := oleutil.MustCallMethod(docs, "Open", source).ToIDispatch()
	defer wordDoc.Release()

	const wdFormatXMLDocument = 12
	oleutil.MustCallMethod(wordDoc, "SaveAs2", source, wdFormatXMLDocument)
	oleutil.MustCallMethod(wordDoc, "Close")
	oleutil.MustCallMethod(word, "Quit")
}
