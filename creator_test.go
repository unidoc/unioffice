// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package gooxml_test

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"baliance.com/gooxml/document"

	"baliance.com/gooxml"

	"baliance.com/gooxml/schema/soo/ofc/sharedTypes"
	"baliance.com/gooxml/schema/soo/wml"
	"baliance.com/gooxml/zippkg"
)

func TestCreatorUnknownType(t *testing.T) {
	el, err := gooxml.CreateElement(xml.StartElement{Name: xml.Name{Local: "foo", Space: "bar"}})
	if el == nil || err != nil {
		t.Errorf("CreateElement should never return nil: %s", err)
	}
	if _, ok := el.(*gooxml.XSDAny); !ok {
		t.Errorf("CreateElement should return XSDAny for unknown types")
	}
}

func TestCreatorKnownType(t *testing.T) {
	el, err := gooxml.CreateElement(xml.StartElement{Name: xml.Name{Local: "CT_Settings", Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"}})
	if el == nil || err != nil {
		t.Errorf("CreateElement should never return nil: %s", err)
	}
	if _, ok := el.(*wml.CT_Settings); !ok {
		t.Errorf("CreateElement should return the element requested, got %T", el)
	}
}
func TestRawEncode(t *testing.T) {
	f, err := os.Open("testdata/settings.xml")
	if err != nil {
		t.Fatalf("error reading settings file")
	}
	dec := xml.NewDecoder(f)
	var got *bytes.Buffer

	// should round trip multiple times with no changes after
	// the first encoding
	for i := 0; i < 5; i++ {
		stng := wml.NewSettings()
		if err := dec.Decode(stng); err != nil {
			t.Errorf("error decoding settings: %s", err)
		}
		got = &bytes.Buffer{}
		fmt.Fprintf(got, zippkg.XMLHeader)
		enc := xml.NewEncoder(zippkg.SelfClosingWriter{W: got})
		if err := enc.Encode(stng); err != nil {
			t.Errorf("error encoding settings: %s", err)
		}

		dec = xml.NewDecoder(bytes.NewReader(got.Bytes()))
	}
	xmlStr := got.String()
	beg := strings.LastIndex(xmlStr, "<w:hdrShapeDefaults>")
	end := strings.LastIndex(xmlStr, "</w:hdrShapeDefaults>")

	gotRaw := xmlStr[beg+20 : end]
	exp := "<o:shapedefaults xmlns=\"urn:schemas-microsoft-com:office:office\" xmlns:o=\"urn:schemas-microsoft-com:office:office\" xmlns:r=\"http://schemas.openxmlformats.org/officeDocument/2006/relationships\" xmlns:s=\"http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes\" xmlns:v=\"urn:schemas-microsoft-com:vml\" xmlns:xml=\"http://www.w3.org/XML/1998/namespace\" spidmax=\"2049\" ext=\"edit\"/>"
	if gotRaw != exp {
		t.Errorf("expected\n%q\ngot\n%q\n", exp, gotRaw)
	}

}

func TestDocProcessingIssue241(t *testing.T) {
	doc, err := document.Open("./testdata/before.docx")
	if err != nil {
		t.Fatalf("error opening document: %v\n", err)
	}

	var (
		tableIdxReq1  = 40
		tableIdxReq12 = 51

		testProcNoRegexp = regexp.MustCompile(`^\s*(\d+(\.\d+)+(\.[a-z]){0,1})\s`)

		reportTxtHolder = "<Report Findings Here>"
		reportTxtPrefix = "ref-"
	)

	// Add report text refs
	for i := tableIdxReq1; i <= tableIdxReq12; i++ {
		addReportTextRefs(doc, i, reportTxtHolder, reportTxtPrefix, testProcNoRegexp)
	}
	doc.SaveToFile("./testdata/after.docx")

	afterdoc, err := document.Open("./testdata/after.docx")

	if afterdoc.X().Body == nil {
		t.Errorf("Doc body is nil after processing")
	}
}

// addReportTextRefs iterates the table at index tableIdx and adds the
// correct report text refs. The doc should contain text matching ReportTxtHolder
// in all cells requiring report text, otherwise the numbering will be off. Caller
// is responsible for setting SetUpdateFieldsOnOpen(true) on doc.Settings and
// saving the result.
func addReportTextRefs(doc *document.Document, tableIdx int, reportTxtHolder, reportTxtPrefix string,
	testProcNoRegexp *regexp.Regexp) {
	tpNo := ""
	rtIdx := 1
	for _, row := range doc.Tables()[tableIdx].Rows() {
		if tpn := testProcNo(row, testProcNoRegexp); tpn != "" {
			tpNo = tpn
			rtIdx = 1
		}
		for _, cell := range row.Cells() {
			for _, para := range cell.Paragraphs() {
				for _, run := range para.Runs() {
					fieldDefault := findFldCharText(run, reportTxtHolder)
					if fieldDefault == nil {
						// possibly check if run has a FldChar, but wasn't matched
						continue
					}
					rtVal := fmt.Sprintf("%s%s-%d", reportTxtPrefix, tpNo, rtIdx)
					fieldDefault.ValAttr = rtVal
					fmt.Println(rtVal)
					rtIdx++
				}
			}
		}
	}
}

func testProcNo(run document.Row, testProcNoRegexp *regexp.Regexp) string {
	if len(run.Cells()) < 2 {
		return ""
	}
	paraTxt := mergeParas(run.Cells()[0].Paragraphs())
	matchGroups := testProcNoRegexp.FindStringSubmatch(paraTxt)
	if len(matchGroups) >= 2 {
		return matchGroups[1]
	}
	return ""
}

// findFldCharText searches all content elements in a run for a FldChar
// with default text matching s. The string pointer is returned so that the
// text can be manipulated by the caller.
func findFldCharText(run document.Run, s string) *wml.CT_String {
	for _, ic := range run.X().EG_RunInnerContent {
		if ic.FldChar != nil && ic.FldChar.FfData != nil && ic.FldChar.FfData.TextInput != nil && ic.FldChar.FfData.TextInput.Default != nil {
			// mark them dirty here for now, but really should be done elsewhere:
			ic.FldChar.DirtyAttr = &sharedTypes.ST_OnOff{}
			ic.FldChar.DirtyAttr.Bool = gooxml.Bool(true)

			// there is some caching of field text being displayed. try setting SetUpdateFieldsOnOpen
			// ic.FldChar.FfData.CalcOnExit[0].ValAttr = &sharedTypes.ST_OnOff{Bool: &[]bool{true}[0]}
			return ic.FldChar.FfData.TextInput.Default
		}
	}
	return nil
}

func mergeRuns(runs []document.Run) string {
	s := ""
	for _, run := range runs {
		s += run.Text()
	}
	return s
}

func mergeParas(paras []document.Paragraph) string {
	s := ""
	for _, para := range paras {
		s += mergeRuns(para.Runs())
	}
	return s
}
