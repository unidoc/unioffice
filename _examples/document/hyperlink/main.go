package main

import (
	"baliance.com/gooxml/color"
	"baliance.com/gooxml/document"
	"baliance.com/gooxml/schema/soo/wml"
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

	hl := para.AddHyperLink()
	hl.SetTarget("http://www.google.com")

	run = hl.AddRun()
	run.SetStyle("Hyperlink")
	run.AddText("Click Here to open google.com")
	hl.SetToolTip("hover to see this")

	doc.SaveToFile("hyperlink.docx")
}
