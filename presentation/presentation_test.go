package presentation

import (
	"testing"

	"github.com/unidoc/unioffice/schema/soo/pml"
)

func TestRemoveChoicesWithPics(t *testing.T) {
	var choices []*pml.CT_GroupShapeChoice
	var pics []*pml.CT_Picture
	pics = append(pics, &pml.CT_Picture{})
	choices = append(choices, pml.NewCT_GroupShapeChoice())
	choices = append(choices, &pml.CT_GroupShapeChoice{
		Pic: pics,
	})
	choices = append(choices, &pml.CT_GroupShapeChoice{
		Pic: pics,
	})
	choices = append(choices, pml.NewCT_GroupShapeChoice())

	choices = removeChoicesWithPics(choices)
	for _, choice := range choices {
		if len(choice.Pic) > 0 {
			t.Fatal("expected to have removed all choices with pics")
		}
	}
}

func TestExtractText(t *testing.T) {
	testFilePath := "./testdata/extraction.pptx"

	// test the whole text extraction
	want := "Test\nSubtest\nHeader 1\nTest \nErrewr\nTests sf s. dsf \nSdfsd xvxv \nHeader 2\n" +
		"Qweqre wefs df\nSd f\n sdf"

	got, err := ExtractText(testFilePath, nil)
	if err != nil {
		t.Fatalf("error opening pres: %v", err)
	}

	if want != got {
		t.Fatalf("extracted text mismatch. got\n\"%s\"\nwant\n\"%s\"", got, want)
	}

	// test extraction from specific slide
	want = "Header 1\nTest \nErrewr\nTests sf s. dsf \nSdfsd xvxv "

	got, err = ExtractText(testFilePath, []int{1})
	if err != nil {
		t.Fatalf("error opening pres: %v", err)
	}

	if want != got {
		t.Fatalf("extracted text mismatch. got\n\"%s\"\nwant\n\"%s\"", got, want)
	}
}
