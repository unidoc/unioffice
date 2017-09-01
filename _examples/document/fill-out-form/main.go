// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"fmt"
	"log"

	"baliance.com/gooxml/document"
)

func main() {
	doc, err := document.Open("form.docx")
	if err != nil {
		log.Fatalf("error opening form: %s", err)
	}

	// FindAllFields is a helper function that traverses the document
	// identifying fields
	fields := doc.FormFields()
	fmt.Println("found", len(fields), "fields")

	for _, fld := range fields {
		fmt.Println("- Name:", fld.Name(), "Type:", fld.Type(), "Value:", fld.Value())

		switch fld.Type() {
		case document.FormFieldTypeText:
			// you can directly set values on text fields
			fld.SetValue("testing 123")
		case document.FormFieldTypeCheckBox:
			// you can check check boxes
			fld.SetChecked(true)
		case document.FormFieldTypeDropDown:
			// and select items in a dropdown, here the value must be one of the
			// fields possible values
			lpv := len(fld.PossibleValues())
			if lpv > 0 {
				fld.SetValue(fld.PossibleValues()[lpv-1])
			}
		}
	}

	doc.SaveToFile("filled-form.docx")
}
