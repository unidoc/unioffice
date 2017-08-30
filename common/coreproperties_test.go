package common_test

import (
	"bytes"
	"encoding/xml"
	"os"
	"testing"

	"baliance.com/gooxml/common"
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
