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

	choices = removeChoicesWithPics(choices)
	if len(choices) != 1 {
		t.Fatalf("expected choices to be length 1, found it to be: %d", len(choices))
	}
}
