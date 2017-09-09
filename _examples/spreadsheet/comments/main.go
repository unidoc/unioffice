// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"log"

	"baliance.com/gooxml/spreadsheet"
)

func main() {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()

	sheet.Cell("A1").SetString("Hello World!")
	sheet.Comments().AddCommentWithStyle("A1", "Gopher", "This looks interesting.")
	sheet.Comments().AddCommentWithStyle("C10", "Gopher", "This is a different comment.")

	if err := ss.Validate(); err != nil {
		log.Fatalf("error validating sheet: %s", err)
	}

	ss.SaveToFile("comments.xlsx")
}
