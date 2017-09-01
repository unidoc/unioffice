package gooxml_test

import "baliance.com/gooxml/document"

func Example_document() {
	// see the baliance.com/gooxml/document documentation or _examples/document
	// for more examples
	doc := document.New()
	doc.AddParagraph().AddRun().AddText("Hello World!")
	doc.SaveToFile("document.docx")
}
