package presentation

import (
	"io/ioutil"
	"testing"

	"github.com/unidoc/unioffice/common"
	"github.com/unidoc/unioffice/schema/soo/dml"
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

func TestFillPlaceholderText(t *testing.T) {
	ppt, err := loadTemplateHelper()
	if err != nil {
		t.Fatal(err)
	}

	// Add new slide from template
	layout, err := ppt.GetLayoutByName("Picture with Caption")
	if err != nil {
		t.Fatal(err)
	}

	slide, err := ppt.AddDefaultSlideWithLayout(layout)
	if err != nil {
		t.Fatal(err)
	}

	// Inject content into placeholders
	title, _ := slide.GetPlaceholder(pml.ST_PlaceholderTypeTitle)
	title.SetText("New title")

	body, _ := slide.GetPlaceholder(pml.ST_PlaceholderTypeBody)
	body.SetText("New body text")

	titleExpected := "New title"
	titleGot := slide.X().CSld.SpTree.Choice[0].Sp[0].TxBody.P[0].EG_TextRun[0].R.T
	if titleExpected != titleGot {
		t.Errorf("Expected title to be %s, got %s", titleExpected, titleGot)
	}

	bodyTextExpected := "New body text"
	bodyTextGot := slide.X().CSld.SpTree.Choice[2].Sp[0].TxBody.P[0].EG_TextRun[0].R.T
	if bodyTextExpected != bodyTextGot {
		t.Errorf("Expected body text to be %s, got %s", bodyTextExpected, bodyTextGot)
	}
}

func TestFillPlaceholderImage(t *testing.T) {
	ppt, err := loadTemplateHelper()
	if err != nil {
		t.Fatal(err)
	}

	// Add local image to pptx
	image, err := common.ImageFromFile("testdata/gophercolor.png")
	if err != nil {
		t.Fatal(err)
	}

	iRef, err := ppt.AddImage(image)
	if err != nil {
		t.Fatal(err)
	}

	// Add new slide from template
	layout, err := ppt.GetLayoutByName("Picture with Caption")
	if err != nil {
		t.Fatal(err)
	}

	slide, err := ppt.AddDefaultSlideWithLayout(layout)
	if err != nil {
		t.Fatal(err)
	}

	imageRelID := slide.AddImageToRels(iRef)

	pic, err := slide.GetPlaceholder(pml.ST_PlaceholderTypePic)
	if err != nil {
		t.Fatal(err)
	}

	spPr := dml.NewCT_ShapeProperties()
	spPr.BlipFill = dml.NewCT_BlipFillProperties()
	spPr.BlipFill.Blip = dml.NewCT_Blip()
	spPr.BlipFill.Blip.EmbedAttr = &imageRelID
	spPr.BlipFill.Stretch = dml.NewCT_StretchInfoProperties() // stretch to parent block with default values

	pic.X().SpPr = spPr

	imageRelIDGot := *slide.X().CSld.SpTree.Choice[1].Sp[0].SpPr.BlipFill.Blip.EmbedAttr
	if imageRelIDGot != imageRelID {
		t.Errorf("Expected image relationship id to be %s, got %s", imageRelID, imageRelIDGot)
	}
	if len(ppt.Images) != 1 {
		t.Fatalf("Expected number of images id to be %d, got %d", 1, len(ppt.Images))
	}

	img := ppt.Images[0]

	formatExpected := "png"
	formatGot := img.Format()
	if formatExpected != formatGot {
		t.Errorf("Expected image format to be %s, got %s", formatExpected, formatGot)
	}

	pathExpected := "testdata/gophercolor.png"
	pathGot := img.Path()
	if pathExpected != pathGot {
		t.Errorf("Expected image path to be %s, got %s", pathExpected, pathGot)
	}

	sizeXExpected := 400
	sizeXGot := img.Size().X
	if sizeXExpected != sizeXGot {
		t.Errorf("Expected image sizeX to be %d, got %d", sizeXExpected, sizeXGot)
	}

	sizeYExpected := 400
	sizeYGot := img.Size().Y
	if sizeYExpected != sizeYGot {
		t.Errorf("Expected image sizeY to be %d, got %d", sizeYExpected, sizeYGot)
	}
}

func loadTemplateHelper() (*Presentation, error) {
	ppt, err := OpenTemplate("testdata/template.potx")
	if err != nil {
		return nil, err
	}

	for _, s := range ppt.Slides() {
		if err = ppt.RemoveSlide(s); err != nil {
			return nil, err
		}
	}

	return ppt, nil
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
