// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document_test

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/unidoc/unioffice/document"
)

func TestFootnotesLoad(t *testing.T) {
	doc := loadTestData(t)
	if doc == nil {
		t.Fatal("could not read test document")
	}
	saveTestData(t, doc)
}

func TestHasFootnotes(t *testing.T) {
	doc := loadTestData(t)
	if doc == nil {
		t.Fatal("could not read test document")
	}
	if !doc.HasFootnotes() {
		t.Error("document should contain footnotes")
	}
	doc = document.New()
	if doc.HasFootnotes() {
		t.Error("document should not contain footnotes")
	}
}

func TestFootnotesListFootnotes(t *testing.T) {
	doc := loadTestData(t)
	if doc == nil {
		t.Fatal("could not read test document")
	}
	if len(doc.Footnotes()) == 0 {
		t.Error("document should contain footnotes")
	}
	if len(doc.Footnotes()) <= 2 {
		t.Error("document should contain footnotes array that includes 2 added by the system")
	}
	saveTestData(t, doc)
}

func TestFootnotesPullByID(t *testing.T) {
	doc := loadTestData(t)
	if doc == nil {
		t.Fatal("could not read test document")
	}
	fn := doc.Footnote(2)
	if fn.X().IdAttr != 2 {
		t.Error("test did not retrieve the proper footnote")
	}
	saveTestData(t, doc)
}

func TestFootnotesPullParagraphs(t *testing.T) {
	doc := loadTestData(t)
	if doc == nil {
		t.Fatal("could not read test document")
	}
	fn := doc.Footnote(2)
	if len(fn.Paragraphs()) == 0 {
		t.Error("test did not retrieve the paragraphs properly")
	}
	saveTestData(t, doc)
}

func TestFootnotesLinkedToRuns(t *testing.T) {
	doc := loadTestData(t)
	if doc == nil {
		t.Fatal("could not read test document")
	}
	var linkedRuns int = 0
	for _, p := range doc.Paragraphs() {
		for _, r := range p.Runs() {
			if ok, _ := r.IsFootnote(); ok {
				linkedRuns++
			}
		}
	}
	if linkedRuns == 0 {
		t.Error("document not properly linked to footnotes")
	}
	saveTestData(t, doc)
}

func TestFootnotesLinkedToRunsProperly(t *testing.T) {
	doc := loadTestData(t)
	if doc == nil {
		t.Fatal("could not read test document")
	}
	linkedFns := make([]document.Footnote, 0)
	for _, p := range doc.Paragraphs() {
		for _, r := range p.Runs() {
			if ok, fnID := r.IsFootnote(); ok {
				linkedFns = append(linkedFns, doc.Footnote(fnID))
			}
		}
	}
	if len(linkedFns) == 0 {
		t.Error("document not properly linked to footnotes")
	}
	saveTestData(t, doc)
}

func TestFootnotesHaveText(t *testing.T) {
	doc := loadTestData(t)
	if doc == nil {
		t.Fatal("could not read test document")
	}
	fn := doc.Footnote(2)
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

func TestFootnotesClearContent(t *testing.T) {
	doc := loadTestData(t)
	if doc == nil {
		t.Fatal("could not read test document")
	}
	fn := doc.Footnote(2)
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

func TestFootnotesSimplifyContent(t *testing.T) {
	newTxt := "This will be shortened."
	doc := loadTestData(t)
	if doc == nil {
		t.Fatal("could not read test document")
	}

	fn := doc.Footnote(2)
	for i, p := range fn.Paragraphs() {
		if i > 0 {
			fn.RemoveParagraph(p)
		}
		for j, r := range p.Runs() {
			if j > 0 {
				p.RemoveRun(r)
			}
			r.Clear()
			r.AddText(newTxt)
		}
	}

	fn2 := doc.Footnote(2)
	if len(fn2.Paragraphs()) != 1 {
		t.Error("wrong number of paragraphs")
	}
	if len(fn2.Paragraphs()[0].Runs()) != 1 {
		t.Errorf("wrong number of runs: expect %d, got %d", 1, len(fn2.Paragraphs()[0].Runs()))
	}
	if fn2.Paragraphs()[0].Runs()[0].Text() != newTxt {
		t.Error("incorrect modification to footnote text")
	}
	saveTestData(t, doc)
}

func TestFootnotesExpandContent(t *testing.T) {
	newTxt := "This will be expanded."
	doc := loadTestData(t)
	if doc == nil {
		t.Fatal("could not read test document")
	}
	fn := doc.Footnote(3)

	para := fn.AddParagraph()
	run := para.AddRun()
	run.AddText(newTxt)

	fn2 := doc.Footnote(3)
	if len(fn2.Paragraphs()) != 2 {
		t.Error("wrong number of paragraphs")
	}
	if len(fn2.Paragraphs()[1].Runs()) != 1 {
		t.Errorf("wrong number of runs: expect %d, got %d", 1, len(fn2.Paragraphs()[0].Runs()))
	}
	if fn2.Paragraphs()[1].Runs()[0].Text() != newTxt {
		t.Error("incorrect modification to footnote text")
	}
	saveTestData(t, doc)
}

func TestFootnotesAddFootnote(t *testing.T) {
	doc := loadTestData(t)
	if doc == nil {
		t.Fatal("could not read test document")
	}
	fnB4 := len(doc.Footnotes())
	var fnB4InP int
	var p document.Paragraph
	for _, p1 := range doc.Paragraphs() {
		for _, r := range p1.Runs() {
			if ok, fnID := r.IsFootnote(); ok {
				if fnID == 2 {
					p = p1
					fnB4InP++
				}
			}
		}
	}
	fn := p.AddFootnote("testing")

	if len(doc.Footnotes()) == fnB4 {
		t.Fatal("did not properly add footnote")
	}

	var fnInP int
	for _, r := range p.Runs() {
		if ok, _ := r.IsFootnote(); ok {
			fnInP++
		}
	}

	if fnInP <= fnB4InP {
		t.Error("footnotes not properly added to run")
	}

	pinner := fn.AddParagraph()
	rinner := pinner.AddRun()
	rinner.AddText("more test goes in this paragraph")

	if len(fn.Paragraphs()) != 2 {
		t.Error("wrong number of paragraphs in footnote")
	}
	for _, p := range fn.Paragraphs() {
		if p.Properties().Style() != "Footnote" {
			t.Error("wrong style set for the paragraph")
		}
	}

	saveTestData(t, doc)
}

func TestFootnotesRemoveFootnote(t *testing.T) {
	doc := loadTestData(t)
	if doc == nil {
		t.Fatal("could not read test document")
	}
	fnB4 := len(doc.Footnotes())
	var runsInP int
	var p document.Paragraph
	for _, p1 := range doc.Paragraphs() {
		for _, r := range p1.Runs() {
			if ok, fnID := r.IsFootnote(); ok {
				if fnID == 2 {
					p = p1
					runsInP = len(p.Runs())
				}
			}
		}
	}
	p.RemoveFootnote(2)

	if len(doc.Footnotes()) == fnB4 {
		fmt.Println(doc.Footnotes())
		t.Fatal("did not properly remove footnote")
	}

	if len(p.Runs()) >= runsInP {
		t.Fatal("did not properly remove footnote")
	}
}

func loadTestData(t *testing.T) *document.Document {
	docBytes, err := ioutil.ReadFile(filepath.Join("testdata", "footnotes_endnotes.docx"))
	if err != nil {
		t.Error("cannot read text fixture from disk")
		return nil
	}
	doc, err := document.Read(bytes.NewReader(docBytes), int64(len(docBytes)))
	if err != nil {
		t.Error("cannot unmarshal test fixture into document struct")
		return nil
	}
	return doc
}

func saveTestData(t *testing.T, doc *document.Document) {
	// Ensure the document can render
	var b bytes.Buffer
	buff := bufio.NewWriter(&b)
	err := doc.Save(buff)
	if err != nil {
		t.Error("failure rendering document to docx:", err)
	}
}
