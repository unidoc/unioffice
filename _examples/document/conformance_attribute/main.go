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
	doc.SetConformance(st.ST_ConformanceClassTransitional)
	doc.SaveToFile("conformance_transitional.docx")
	doc.SetConformance(st.ST_ConformanceClassStrict)
	doc.SaveToFile("conformance_strict.docx")
	doc.SetConformance(st.ST_ConformanceClassUnset)
	doc.SaveToFile("conformance_unset.docx")
}
