package document_test

import (
	"fmt"
	"log"

	"baliance.com/gooxml/document"
)

func ExampleNew() {
	doc := document.New()
	doc.AddParagraph().AddRun().AddText("Hello World!")
	doc.SaveToFile("document.docx")
}

func ExampleOpen() {
	doc, err := document.Open("existing.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	for _, para := range doc.Paragraphs() {
		for _, run := range para.Runs() {
			fmt.Print(run.Text())
		}
		fmt.Println()
	}
}

func ExampleOpenTemplate() {
	doc, err := document.OpenTemplate("existing.docx")
	if err != nil {
		log.Fatalf("error opening document template: %s", err)
	}
	para := doc.AddParagraph()
	para.SetStyle("Title")
	para.AddRun().AddText("My Document Title")

	para = doc.AddParagraph()
	para.SetStyle("Subtitle")
	para.AddRun().AddText("Document Subtitle")

	para = doc.AddParagraph()
	para.SetStyle("Heading1")
	para.AddRun().AddText("Major Section")
	doc.SaveToFile("ouput.docx")
}

func ExampleDocument_FormFields() {
	doc, err := document.Open("invitation.docx")
	if err != nil {
		log.Fatalf("error opening document form: %s", err)
	}
	for _, field := range doc.FormFields() {
		switch field.Name() {
		case "attendingEvent":
			if field.Type() == document.FormFieldTypeCheckBox {
				field.SetChecked(true)
			}
		case "name":
			if field.Type() == document.FormFieldTypeText {
				field.SetValue("John Smith")
			}
		}
	}
	doc.SaveToFile("invitation-respoonse.docx")
}
