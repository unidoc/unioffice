package main

import (
	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

func main() {

	doc := document.New()

	hlStyle := doc.Styles.AddStyle("Hyperlink", wml.ST_StyleTypeCharacter, false)
	hlStyle.SetName("Hyperlink")
	hlStyle.SetBasedOn("DefaultParagraphFont")
	hlStyle.RunProperties().Color().SetThemeColor(wml.ST_ThemeColorHyperlink)
	clr := color.FromHex("#0563C1")
	hlStyle.RunProperties().Color().SetColor(clr)
	hlStyle.RunProperties().SetUnderline(wml.ST_UnderlineSingle, clr)

	para := doc.AddParagraph()
	run := para.AddRun()
	run.AddText("Hello World! ")
	bm := para.AddBookmark("_bookmark1")
	addBlankLines(para)

	// first link to a URL
	hl := para.AddHyperLink()
	hl.SetTarget("http://www.google.com")
	run = hl.AddRun()
	run.Properties().SetStyle("Hyperlink")
	run.AddText("Click Here to open google.com")
	hl.SetToolTip("hover to see this")

	addBlankLines(para)
	// second link to a bookmark
	hl = para.AddHyperLink()
	hl.SetTargetBookmark(bm)

	run = hl.AddRun()
	run.AddText("Click Here to jump to the bookmark")

	doc.SaveToFile("hyperlink.docx")
}

func addBlankLines(p document.Paragraph) {
	run := p.AddRun()
	for i := 0; i < 4; i++ {
		run.AddBreak()
	}
}
