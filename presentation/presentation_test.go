package presentation

import (
	"io/ioutil"
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

func TestTmpFiles(t *testing.T) {
	ppt, err := Open("testdata/image.pptx")
	if err != nil {
		t.Errorf("error opening document: %s", err)
	}
	files, err := ioutil.ReadDir(ppt.TmpPath)
	if err != nil {
		t.Errorf("cannot open a workbook: %s", err)
	}
	expected := 5
	got := len(files)
	if got != expected {
		t.Errorf("should be %d files in the temp dir, found %d", expected, got)
	}
	ppt.Close()
	files, err = ioutil.ReadDir(ppt.TmpPath)
	expected = 0
	got = len(files)
	if got != expected {
		t.Errorf("should be %d files in the temp dir, found %d", expected, got)
	}
}
