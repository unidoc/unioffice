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

func TestCorePropertiesSettersStrings(t *testing.T) {
	cp := common.NewCoreProperties()
	exp := "Foo"

	if got := cp.Author(); got != "" {
		t.Errorf("expected empty author, got %s", got)
	}
	if got := cp.Title(); got != "" {
		t.Errorf("expected empty title, got %s", got)
	}
	if got := cp.Description(); got != "" {
		t.Errorf("expected empty description, got %s", got)
	}
	if got := cp.Category(); got != "" {
		t.Errorf("expected empty category, got %s", got)
	}
	if got := cp.ContentStatus(); got != "" {
		t.Errorf("expected empty contentStatus, got %s", got)
	}
	if got := cp.LastModifiedBy(); got != "" {
		t.Errorf("expected empty lastModifiedBy, got %s", got)
	}
	cp.SetAuthor(exp)
	if got := cp.Author(); got != exp {
		t.Errorf("expected author=%s, got %s", exp, got)
	}
	cp.SetTitle(exp)
	if got := cp.Title(); got != exp {
		t.Errorf("expected title=%s, got %s", exp, got)
	}
	cp.SetDescription(exp)
	if got := cp.Description(); got != exp {
		t.Errorf("expected description=%s, got %s", exp, got)
	}
	cp.SetCategory(exp)
	if got := cp.Category(); got != exp {
		t.Errorf("expected category=%s, got %s", exp, got)
	}
	cp.SetContentStatus(exp)
	if got := cp.ContentStatus(); got != exp {
		t.Errorf("expected contentStatus=%s, got %s", exp, got)
	}
	cp.SetLastModifiedBy(exp)
	if got := cp.LastModifiedBy(); got != exp {
		t.Errorf("expected lastModifiedBy=%s, got %s", exp, got)
	}
}

func TestCorePropertiesSettersDates(t *testing.T) {
	cp := common.NewCoreProperties()
	exp := time.Date(2017, 1, 2, 3, 4, 5, 0, time.UTC)
	if !cp.Created().IsZero() {
		t.Errorf("expected zero created time, got %v", exp)
	}
	if !cp.Modified().IsZero() {
		t.Errorf("expected zero Modified time, got %v", exp)
	}

	cp.SetCreated(exp)
	if got := cp.Created(); got != exp {
		t.Errorf("expected created =%v, got %v", exp, got)
	}

	cp.SetModified(exp)
	if got := cp.Modified(); got != exp {
		t.Errorf("expected modified =%v, got %v", exp, got)
	}
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
