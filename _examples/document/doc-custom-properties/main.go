// Copyright 2019 FoxyUtils ehf. All rights reserved.

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/unidoc/unioffice/document"
)

func main() {
	doc, err := document.Open("document.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}

	cp := doc.GetOrCreateCustomProperties()

	// You can read properties from the document
	fmt.Println("AppVersion", *cp.GetPropertyByName("AppVersion").Lpwstr)
	fmt.Println("Company", *cp.GetPropertyByName("Company").Lpwstr)
	fmt.Println("DocSecurity", *cp.GetPropertyByName("DocSecurity").I4)
	fmt.Println("LinksUpToDate", *cp.GetPropertyByName("LinksUpToDate").Bool)
	fmt.Println("Non-existent", cp.GetPropertyByName("nonexistentproperty"))

	// And change them as well
	cp.SetPropertyAsLpwstr("Company", "Another company") // text, existing property
	fmt.Println("Company", *cp.GetPropertyByName("Company").Lpwstr)

	// Adding new properties
	cp.SetPropertyAsLpwstr("Another text property", "My text value") // text
	cp.SetPropertyAsI4("Another integer number property", 42)        // int32
	cp.SetPropertyAsR8("Another float number property", 3.14)        // float64
	cp.SetPropertyAsDate("Another date property", time.Now())        // date

	doc.SaveToFile("document_customized.docx")

	// For new documents all is the same
	docNew := document.New()

	cpNew := docNew.GetOrCreateCustomProperties()
	cpNew.SetPropertyAsLpwstr("Another text property", "My text value") // text
	cpNew.SetPropertyAsI4("Another integer number property", 42)        // int23
	cpNew.SetPropertyAsR8("Another float number property", 3.14)        // float64
	cpNew.SetPropertyAsDate("Another date property", time.Now())        // date

	docNew.SaveToFile("document_new.docx")
}
