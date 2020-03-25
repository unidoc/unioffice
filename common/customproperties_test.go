package common_test

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/unidoc/unioffice/common"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/schema/soo/ofc/docPropsVTypes"
	"github.com/unidoc/unioffice/testhelper"
	"github.com/unidoc/unioffice/zippkg"
)

func TestMarshalCustomProperties(t *testing.T) {
	f, err := os.Open("testdata/custom.xml")
	if err != nil {
		t.Fatalf("error reading file")
	}
	dec := xml.NewDecoder(f)
	cp := common.NewCustomProperties()
	if err := dec.Decode(cp.X()); err != nil {
		t.Errorf("error decoding: %s", err)
	}

	got := &bytes.Buffer{}
	enc := xml.NewEncoder(zippkg.SelfClosingWriter{W: got})
	if err := enc.Encode(cp.X()); err != nil {
		t.Errorf("error encoding: %s", err)
	}

	testhelper.CompareGoldenXML(t, "custom.xml", got.Bytes())
}

func TestCustomPropertiesNew(t *testing.T) {
	cp := common.NewCustomProperties()
	expName := "Foo"

	if got := cp.PropertiesList(); len(got) != 0 {
		t.Errorf("expected empty properties list, got %d elements", len(got))
	}

	if got := cp.GetPropertyByName(expName); got != nil {
		t.Errorf("expected nil Foo property, got %p", got)
	}
}

func TestCustomPropertiesSettersStrings(t *testing.T) {
	cp := common.NewCustomProperties()
	expName := "Foo"
	expValue := "Bar"

	cp.SetPropertyAsLpstr(expName, expValue)
	if got := cp.GetPropertyByName(expName); *got.Lpstr != expValue {
		t.Errorf("expected value of %s=%s, got %s", expName, expValue, *got.Lpstr)
	}

	cp.SetPropertyAsLpwstr(expName, expValue)
	if got := cp.GetPropertyByName(expName); *got.Lpwstr != expValue {
		t.Errorf("expected value of %s=%s, got %s", expName, expValue, *got.Lpwstr)
	}

	cp.SetPropertyAsBlob(expName, expValue)
	if got := cp.GetPropertyByName(expName); *got.Blob != expValue {
		t.Errorf("expected value of %s=%s, got %s", expName, expValue, *got.Blob)
	}

	cp.SetPropertyAsOblob(expName, expValue)
	if got := cp.GetPropertyByName(expName); *got.Oblob != expValue {
		t.Errorf("expected value of %s=%s, got %s", expName, expValue, *got.Oblob)
	}

	cp.SetPropertyAsStream(expName, expValue)
	if got := cp.GetPropertyByName(expName); *got.Stream != expValue {
		t.Errorf("expected value of %s=%s, got %s", expName, expValue, *got.Stream)
	}

	cp.SetPropertyAsOstream(expName, expValue)
	if got := cp.GetPropertyByName(expName); *got.Ostream != expValue {
		t.Errorf("expected value of %s=%s, got %s", expName, expValue, *got.Ostream)
	}

	cp.SetPropertyAsStorage(expName, expValue)
	if got := cp.GetPropertyByName(expName); *got.Storage != expValue {
		t.Errorf("expected value of %s=%s, got %s", expName, expValue, *got.Storage)
	}

	cp.SetPropertyAsOstorage(expName, expValue)
	if got := cp.GetPropertyByName(expName); *got.Ostorage != expValue {
		t.Errorf("expected value of %s=%s, got %s", expName, expValue, *got.Ostorage)
	}

	cp.SetPropertyAsClsid(expName, expValue)
	if got := cp.GetPropertyByName(expName); *got.Clsid != expValue {
		t.Errorf("expected value of %s=%s, got %s", expName, expValue, *got.Clsid)
	}

	cp.SetPropertyAsCy(expName, expValue)
	if got := cp.GetPropertyByName(expName); *got.Cy != expValue {
		t.Errorf("expected value of %s=%s, got %s", expName, expValue, *got.Cy)
	}

	cp.SetPropertyAsError(expName, expValue)
	if got := cp.GetPropertyByName(expName); *got.Error != expValue {
		t.Errorf("expected value of %s=%s, got %s", expName, expValue, *got.Error)
	}
}

func TestCustomPropertiesSettersNumbers(t *testing.T) {
	cp := common.NewCustomProperties()
	expName := "Foo"
	expInt := 42
	expInt64 := int64(42)
	expFloat64 := 3.14
	expFloat32 := float32(3.14)

	cp.SetPropertyAsInt(expName, expInt)
	if got := cp.GetPropertyByName(expName); *got.Int != int32(expInt) {
		t.Errorf("expected value of %s=%v, got %v", expName, expInt, *got.Int)
	}

	cp.SetPropertyAsI8(expName, expInt64)
	if got := cp.GetPropertyByName(expName); *got.I8 != expInt64 {
		t.Errorf("expected value of %s=%v, got %v", expName, expInt64, *got.I8)
	}

	cp.SetPropertyAsR4(expName, expFloat32)
	if got := cp.GetPropertyByName(expName); *got.R4 != expFloat32 {
		t.Errorf("expected value of %s=%v, got %v", expName, expFloat32, *got.R4)
	}

	cp.SetPropertyAsDecimal(expName, expFloat64)
	if got := cp.GetPropertyByName(expName); *got.Decimal != expFloat64 {
		t.Errorf("expected value of %s=%v, got %v", expName, expFloat64, *got.Decimal)
	}
}

func TestCustomPropertiesSettersDates(t *testing.T) {
	cp := common.NewCustomProperties()
	expName := "dateExample"
	expValue := time.Date(2017, 1, 2, 3, 4, 5, 0, time.UTC)

	cp.SetPropertyAsDate(expName, expValue)
	if got := cp.GetPropertyByName(expName); *got.Date != expValue {
		t.Errorf("expected created =%v, got %v", expValue, *got.Date)
	}

	cp.SetPropertyAsFiletime(expName, expValue)
	if got := cp.GetPropertyByName(expName); *got.Filetime != expValue {
		t.Errorf("expected created =%v, got %v", expValue, *got.Filetime)
	}
}

func TestCustomPropertiesSettersVstream(t *testing.T) {
	cp := common.NewCustomProperties()
	expName := "dateExample"

	newVstream := docPropsVTypes.NewVstream()
	newVstream.VersionAttr = "v0"
	newVstream.Content = "c1"
	cp.SetPropertyAsVstream(expName, newVstream)

	if got := cp.GetPropertyByName(expName); *got.Vstream != *newVstream {
		t.Errorf("expected created =%v, got %v", *newVstream, *got.Vstream)
	}
}

func ExampleCustomProperties() {
	doc, _ := document.Open("document.docx")

	cp := doc.CustomProperties
	// You can read properties from the document
	fmt.Println("AppVersion", *cp.GetPropertyByName("AppVersion").Lpwstr)
	fmt.Println("category", *cp.GetPropertyByName("category").Lpwstr)
	fmt.Println("contentStatus", *cp.GetPropertyByName("contentStatus").Lpwstr)
	fmt.Println("HyperlinksChanged", *cp.GetPropertyByName("HyperlinksChanged").Bool)
	fmt.Println("Non-existent", cp.GetPropertyByName("nonexistentproperty"))

	// And change them as well
	cp.SetPropertyAsLpwstr("Another text property", "My text value") // text
	cp.SetPropertyAsI4("Another integer number property", 42)        // int23
	cp.SetPropertyAsR8("Another float number property", 3.14)        // float64
	cp.SetPropertyAsDate("Another date property", time.Now())        // date

	doc.SaveToFile("document.docx")
}
