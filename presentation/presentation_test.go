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

	want := "TestSubtestHeader 1Test ErrewrTests sf s. dsf Sdfsd xvxv Header 2Qweqre wefs dfSd f sdf"

	got, err := ExtractText(testFilePath)
	if err != nil {
		t.Errorf("Error opening pres: %v", err)
		return
	}

	if want != got {
		t.Errorf("extracted text mismatch. got\n\"%s\"\nwant\n\"%s\"", got, want)
	}
}
