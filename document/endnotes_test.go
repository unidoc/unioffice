package document_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/unidoc/unioffice/document"
)

func TestEndnotesLoad(t *testing.T) {
	doc := loadTestData(t)
	if doc == nil {
		t.Fatal("could not read test document")
	}
	saveTestData(t, doc)
}

func TestHasEndnotes(t *testing.T) {
	doc := loadTestData(t)
	if doc == nil {
		t.Fatal("could not read test document")
	}
	if doc.HasEndnotes() == false {
		t.Error("document should contain endnotes")
	}
	doc = document.New()
	if doc.HasEndnotes() == true {
		t.Error("document should not contain endnotes")
	}
}

func TestEndnotesListEndnotes(t *testing.T) {
	doc := loadTestData(t)
	if doc == nil {
		t.Fatal("could not read test document")
	}
	if len(doc.Endnotes()) == 0 {
		t.Error("document should contain footnotes")
	}
	if len(doc.Endnotes()) <= 2 {
		t.Error("document should contain footnotes array that includes 2 added by the system")
	}
	saveTestData(t, doc)
}

func TestEndnotesPullByID(t *testing.T) {
	doc := loadTestData(t)
	if doc == nil {
		t.Fatal("could not read test document")
	}
	fn := doc.Endnote(2)
	if fn.X().IdAttr != 2 {
		t.Error("test did not retrieve the proper footnote")
	}
	saveTestData(t, doc)
}

func TestEndnotesPullParagraphs(t *testing.T) {
	doc := loadTestData(t)
	if doc == nil {
		t.Fatal("could not read test document")
	}
	fn := doc.Endnote(2)
	if len(fn.Paragraphs()) == 0 {
		t.Error("test did not retrieve the paragraphs properly")
	}
	saveTestData(t, doc)
}

func TestEndnotesLinkedToRuns(t *testing.T) {
	doc := loadTestData(t)
	if doc == nil {
		t.Fatal("could not read test document")
	}
	var linkedRuns int
	for _, p := range doc.Paragraphs() {
		for _, r := range p.Runs() {
			if ok, _ := r.IsEndnote(); ok {
				linkedRuns++
			}
		}
	}
	if linkedRuns == 0 {
		t.Error("document not properly linked to footnotes")
	}
	saveTestData(t, doc)
}

func TestEndnotesLinkedToRunsProperly(t *testing.T) {
	doc := loadTestData(t)
	if doc == nil {
		t.Fatal("could not read test document")
	}
	linkedEnds := make([]document.Endnote, 0)
	for _, p := range doc.Paragraphs() {
		for _, r := range p.Runs() {
			if ok, enID := r.IsEndnote(); ok {
				linkedEnds = append(linkedEnds, doc.Endnote(enID))
			}
		}
	}
	if len(linkedEnds) == 0 {
		t.Error("document not properly linked to footnotes")
	}
	saveTestData(t, doc)
}

func TestEndnotesHaveText(t *testing.T) {
	doc := loadTestData(t)
	if doc == nil {
		t.Fatal("could not read test document")
	}
	fn := doc.Endnote(2)
	var buff bytes.Buffer
	for _, p1 := range fn.Paragraphs() {
		for _, r2 := range p1.Runs() {
			buff.WriteString(r2.Text())
		}
	}
	if buff.String() == "" {
		t.Error("content not accessible from footnote")
	}
	saveTestData(t, doc)
}

func TestEndnotesClearContent(t *testing.T) {
	doc := loadTestData(t)
	if doc == nil {
		t.Fatal("could not read test document")
	}
	fn := doc.Endnote(2)
	var buff bytes.Buffer
	for _, p1 := range fn.Paragraphs() {
		for _, r2 := range p1.Runs() {
			r2.Clear()
			buff.WriteString(r2.Text())
		}
	}
	output := buff.String()
	if output != "" {
		t.Errorf("content not cleared from footnote: %s", output)
	}
	saveTestData(t, doc)
}

func TestEndnotesSimplifyContent(t *testing.T) {
	newTxt := "This be shortened."
	doc := loadTestData(t)
	if doc == nil {
		t.Fatal("could not read test document")
	}

	en := doc.Endnote(2)
	for i, p := range en.Paragraphs() {
		if i > 0 {
			en.RemoveParagraph(p)
		}
		for j, r := range p.Runs() {
			if j > 0 {
				p.RemoveRun(r)
			}
			r.Clear()
			r.AddText(newTxt)
		}
	}

	en2 := doc.Endnote(2)
	if len(en2.Paragraphs()) != 1 {
		t.Error("wrong number of paragraphs")
	}
	if len(en2.Paragraphs()[0].Runs()) != 1 {
		t.Errorf("wrong number of runs: expect %d, got %d", 1, len(en2.Paragraphs()[0].Runs()))
	}
	if en2.Paragraphs()[0].Runs()[0].Text() != newTxt {
		t.Error("incorrect modification to footnote text")
	}
	saveTestData(t, doc)
}

func TestEndnotesExpandContent(t *testing.T) {
	newTxt := "This be expanded."
	doc := loadTestData(t)
	if doc == nil {
		t.Fatal("could not read test document")
	}
	en := doc.Endnote(2)

	para := en.AddParagraph()
	run := para.AddRun()
	run.AddText(newTxt)

	en2 := doc.Endnote(2)
	if len(en2.Paragraphs()) != 2 {
		t.Error("wrong number of paragraphs")
	}
	if len(en2.Paragraphs()[1].Runs()) != 1 {
		t.Errorf("wrong number of runs: expect %d, got %d", 1, len(en2.Paragraphs()[0].Runs()))
	}
	if en2.Paragraphs()[1].Runs()[0].Text() != newTxt {
		t.Error("incorrect modification to footnote text")
	}
	saveTestData(t, doc)
}

func TestEndnotesRemoveEndnote(t *testing.T) {
	doc := loadTestData(t)
	if doc == nil {
		t.Fatal("could not read test document")
	}
	enB4 := len(doc.Endnotes())
	var runsInP int
	var p document.Paragraph
	for _, p1 := range doc.Paragraphs() {
		for _, r := range p1.Runs() {
			if ok, enID := r.IsEndnote(); ok {
				if enID == 2 {
					p = p1
					runsInP = len(p.Runs())
				}
			}
		}
	}
	p.RemoveEndnote(2)

	if len(doc.Endnotes()) == enB4 {
		fmt.Println(doc.Endnotes())
		t.Fatal("did not properly remove footnote")
	}

	if len(p.Runs()) >= runsInP {
		t.Fatal("did not properly remove footnote")
	}
}

func TestEndnotesAddEndnote(t *testing.T) {
	doc := loadTestData(t)
	if doc == nil {
		t.Fatal("could not read test document")
	}
	enB4 := len(doc.Endnotes())
	var enB4InP int
	var p document.Paragraph
	for _, p1 := range doc.Paragraphs() {
		for _, r := range p1.Runs() {
			if ok, enID := r.IsEndnote(); ok {
				if enID == 2 {
					p = p1
					enB4InP++
				}
			}
		}
	}
	en := p.AddEndnote("testing")

	if len(doc.Endnotes()) == enB4 {
		t.Fatal("did not properly add footnote")
	}

	var enInP int
	for _, r := range p.Runs() {
		if ok, _ := r.IsEndnote(); ok {
			enInP++
		}
	}

	if enInP <= enB4InP {
		t.Error("footnotes not properly added to run")
	}

	pinner := en.AddParagraph()
	rinner := pinner.AddRun()
	rinner.AddText("more test goes in this paragraph")

	if len(en.Paragraphs()) != 2 {
		t.Error("wrong number of paragraphs in footnote")
	}
	for _, p := range en.Paragraphs() {
		if p.Properties().Style() != "Endnote" {
			t.Error("wrong style set for the paragraph")
		}
	}

	saveTestData(t, doc)
}
