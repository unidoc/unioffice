// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"log"
	"os"
	"path/filepath"

	"baliance.com/gooxml/document"
	"github.com/go-ole/go-ole/oleutil"
)

// NOTE: This example can only run on Windows and requires that Word be installed.

func main() {
	doc := document.New()

	para := doc.AddParagraph()
	run := para.AddRun()
	para.SetStyle("Title")
	run.AddText("Simple Document Formatting")

	para = doc.AddParagraph()
	para.SetStyle("Heading1")
	run = para.AddRun()
	run.AddText("Some Heading Text")

	para = doc.AddParagraph()
	para.SetStyle("Heading2")
	run = para.AddRun()
	run.AddText("Some Heading Text")
	doc.SaveToFile("simple.docx")

	cwd, _ := os.Getwd()
	ConvertToPDF(filepath.Join(cwd, "simple.docx"), filepath.Join(cwd, "simple.pdf"))
}

// ConvertToPDF uses go-ole to convert a docx to a PDF using the Word application
func ConvertToPDF(source, destination string) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	iunk, err := oleutil.CreateObject("Word.Application")
	if err != nil {
		log.Fatalf("error creating Word object: %s", err)
	}

	word := iunk.MustQueryInterface(ole.IID_IDispatch)
	defer word.Release()

	// opening then saving works due to the call to doc.Settings.SetUpdateFieldsOnOpen(true) above

	docs := oleutil.MustGetProperty(word, "Documents").ToIDispatch()
	wordDoc := oleutil.MustCallMethod(docs, "Open", source).ToIDispatch()

	// file format constant comes from https://msdn.microsoft.com/en-us/vba/word-vba/articles/wdsaveformat-enumeration-word
	const wdFormatPDF = 17
	oleutil.MustCallMethod(wordDoc, "SaveAs2", destination, wdFormatPDF)
	oleutil.MustCallMethod(wordDoc, "Close")
	oleutil.MustCallMethod(word, "Quit")
}
