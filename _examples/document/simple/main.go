// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"baliance.com/gooxml/color"
	"baliance.com/gooxml/document"
	"baliance.com/gooxml/schema/soo/wordprocessingml"
)

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

	para = doc.AddParagraph()
	para.SetStyle("Heading3")
	run = para.AddRun()
	run.AddText("Some Heading Text")

	para = doc.AddParagraph()
	run = para.AddRun()
	run.AddText("A run is a string of characters with the same formatting. ")

	run = para.AddRun()
	run.SetBold(true)
	run.SetFontFamily("Courier")
	run.SetFontSize(15)
	run.SetColor(color.Red)
	run.AddText("Multiple runs with different formatting can exist in the same paragraph. ")

	run = para.AddRun()
	run.AddText("Adding breaks to a run will insert line breaks after the run. ")
	run.AddBreak()
	run.AddBreak()

	run = createParaRun(doc, "Runs support styling options:")

	run = createParaRun(doc, "small caps")
	run.SetSmallCaps(true)

	run = createParaRun(doc, "strike")
	run.SetStrikeThrough(true)

	run = createParaRun(doc, "double strike")
	run.SetDoubleStrikeThrough(true)

	run = createParaRun(doc, "outline")
	run.SetOutline(true)

	run = createParaRun(doc, "emboss")
	run.SetEmboss(true)

	run = createParaRun(doc, "shadow")
	run.SetShadow(true)

	run = createParaRun(doc, "imprint")
	run.SetImprint(true)

	run = createParaRun(doc, "highlighting")
	run.SetHighlight(wordprocessingml.ST_HighlightColorYellow)

	run = createParaRun(doc, "underline")
	run.SetUnderline(wordprocessingml.ST_UnderlineWavyDouble, color.Red)

	run = createParaRun(doc, "text effects")
	run.SetEffect(wordprocessingml.ST_TextEffectAntsRed)

	doc.SaveToFile("simple.docx")
}

func createParaRun(doc *document.Document, s string) document.Run {
	para := doc.AddParagraph()
	run := para.AddRun()
	run.AddText(s)
	return run
}
