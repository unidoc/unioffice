package document_test

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"testing"

	"baliance.com/gooxml/document"
	"baliance.com/gooxml/testhelper"
	"baliance.com/gooxml/zippkg"
)

func TestStylesUnmarshal(t *testing.T) {
	f, err := os.Open("testdata/styles.xml")
	if err != nil {
		t.Fatalf("error reading content types file")
	}
	defer f.Close()
	dec := xml.NewDecoder(f)
	r := document.NewStyles()
	if err := dec.Decode(r.X()); err != nil {
		t.Errorf("error decoding content types: %s", err)
	}
	got := &bytes.Buffer{}
	fmt.Fprintf(got, zippkg.XMLHeader)
	enc := xml.NewEncoder(zippkg.SelfClosingWriter{W: got})
	if err := enc.Encode(r.X()); err != nil {
		t.Errorf("error encoding content types: %s", err)
	}

	testhelper.CompareGoldenXML(t, "styles.xml", got.Bytes())
}

func TestStylesList(t *testing.T) {
	f, err := os.Open("testdata/styles.xml")
	if err != nil {
		t.Fatalf("error reading content types file")
	}
	defer f.Close()
	dec := xml.NewDecoder(f)
	r := document.NewStyles()
	if err := dec.Decode(r.X()); err != nil {
		t.Errorf("error decoding content types: %s", err)
	}
	expStyleCnt := 26
	if got := len(r.Styles()); got != expStyleCnt {
		t.Errorf("expected %d total styles, got %d", expStyleCnt, got)
	}

	expParaStyleCnt := 9
	if got := len(r.ParagraphStyles()); got != expParaStyleCnt {
		t.Errorf("expected %d total paragraph styles, got %d", expStyleCnt, got)
	}
	for _, ps := range r.ParagraphStyles() {
		switch ps.StyleID() {
		case "Normal", "Heading1", "Heading2", "Heading3",
			"Title", "Subtitle", "Quote", "IntenseQuote", "ListParagraph":
		default:
			t.Errorf("unexpected paragraph style: %s", ps.StyleID())
		}
	}
}
