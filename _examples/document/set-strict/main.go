// Copyright 2017 FoxyUtils ehf. All rights reserved.

package main

import (
	"github.com/unidoc/unioffice/document"
	st "github.com/unidoc/unioffice/schema/soo/ofc/sharedTypes"
)

func main() {
	doc, err := document.Open("document.docx")
	if err != nil {
		panic(err)
	}
	doc.SetStrict(false) // document will be saved as Word document (this is a default option for new files)
	doc.SaveToFile("conformance_transitional.docx")
	doc.SetStrict(true) // document will be saved in the Strict mode
	doc.SaveToFile("conformance_strict.docx")
	doc.SetConformance(st.ST_ConformanceClassUnset) // Conformance attribute will be unset, which also leads to saving as Word document
	doc.SaveToFile("conformance_unset.docx")
}
