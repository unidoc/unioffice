package common_test

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"testing"
	"time"

	"baliance.com/gooxml/common"
	"baliance.com/gooxml/document"
	"baliance.com/gooxml/testhelper"
	"baliance.com/gooxml/zippkg"
)

func TestMarshalCoreProperties(t *testing.T) {
	f, err := os.Open("testdata/core.xml")
	if err != nil {
		t.Fatalf("error reading file")
	}
	dec := xml.NewDecoder(f)
	cp := common.NewCoreProperties()
	if err := dec.Decode(cp.X()); err != nil {
		t.Errorf("error decoding: %s", err)
	}

	got := &bytes.Buffer{}
	enc := xml.NewEncoder(zippkg.SelfClosingWriter{W: got})
	if err := enc.Encode(cp.X()); err != nil {
		t.Errorf("error encoding: %s", err)
	}

	testhelper.CompareGoldenXML(t, "core.xml", got.Bytes())
}

func ExampleCoreProperties() {
	doc, _ := document.Open("document.docx")
	cp := doc.CoreProperties
	// Reading Properties
	fmt.Println("Title:", cp.Title())
	fmt.Println("Author:", cp.Author())
	fmt.Println("Description:", cp.Description())
	fmt.Println("Last Modified By:", cp.LastModifiedBy())
	fmt.Println("Category:", cp.Category())
	fmt.Println("Content Status:", cp.ContentStatus())
	fmt.Println("Created:", cp.Created())
	fmt.Println("Modified:", cp.Modified())

	// Setting Properties
	cp.SetTitle("CP Invoices")
	cp.SetAuthor("John Doe")
	cp.SetCategory("Invoices")
	cp.SetContentStatus("Draft")
	cp.SetLastModifiedBy("Jane Smith")
	cp.SetCreated(time.Now())
	cp.SetModified(time.Now())
	doc.SaveToFile("document.docx")
}
