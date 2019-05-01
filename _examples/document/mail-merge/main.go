package main

import (
	"log"

	"github.com/unidoc/unioffice/document"
)

func main() {
	d, err := document.Open("mm.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	for _, v := range d.MergeFields() {
		log.Println("replacing", v)
	}
	rep := map[string]string{}
	rep["Title"] = "mr."      // has a \* Upper attribute on the field
	rep["FirstName"] = "JOHN" // has a \* Lower attribute on the field
	d.MailMerge(rep)
	d.SaveToFile("merged.docx")
}
