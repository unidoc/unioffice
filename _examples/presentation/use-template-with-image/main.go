package main

import (
	"fmt"
	"os"
	"time"

	"github.com/unidoc/unioffice/common"
	"github.com/unidoc/unioffice/schema/soo/dml"
	"github.com/unidoc/unioffice/schema/soo/pml"

	"github.com/unidoc/unioffice/presentation"
)

func main() {
	startTime := time.Now()

	// Start building pptx
	ppt, err := presentation.OpenTemplate("template.potx")
	if err != nil {
		fmt.Println("presentation.OpenTemplate err ", err)
		os.Exit(1)
	}

	// Clear out example slides
	for _, s := range ppt.Slides() {
		if err = ppt.RemoveSlide(s); err != nil {
			fmt.Println("ppt.RemoveSlide err ", err)
			os.Exit(1)
		}
	}

	// Add new slide from template
	layout, err := ppt.GetLayoutByName("Picture with Caption")
	if err != nil {
		fmt.Println("ppt.GetLayoutByName err ", err)
		os.Exit(1)
	}

	// Add local image to pptx
	image, err := common.ImageFromFile("gophercolor.png")
	if err != nil {
		fmt.Println("common.ImageFromFile err ", err)
		os.Exit(1)
	}

	iRef, err := ppt.AddImage(image)
	if err != nil {
		fmt.Println("ppt.AddImage err ", err)
		os.Exit(1)
	}

	slide, err := ppt.AddDefaultSlideWithLayout(layout)
	if err != nil {
		fmt.Println("ppt.AddDefaultSlideWithLayout err ", err)
		os.Exit(1)
	}

	// Inject content into placeholders
	title, _ := slide.GetPlaceholder(pml.ST_PlaceholderTypeTitle)
	title.SetText("New title")

	body, _ := slide.GetPlaceholder(pml.ST_PlaceholderTypeBody)
	body.SetText("New body text")

	imageRelID := slide.AddImageToRels(iRef)

	pic, err := slide.GetPlaceholder(pml.ST_PlaceholderTypePic)
	if err != nil {
		fmt.Println("ppt.AddImage err ", err)
		os.Exit(1)
	}

	spPr := dml.NewCT_ShapeProperties()
	spPr.BlipFill = dml.NewCT_BlipFillProperties()
	spPr.BlipFill.Blip = dml.NewCT_Blip()
	spPr.BlipFill.Blip.EmbedAttr = &imageRelID
	spPr.BlipFill.Stretch = dml.NewCT_StretchInfoProperties() // stretch to parent block with default values

	pic.X().SpPr = spPr

	if err := ppt.Validate(); err != nil {
		fmt.Println("ppt.Validate err ", err)
	}

	if err := ppt.SaveToFile("mod.pptx"); err != nil {
		fmt.Println("ppt.SaveToFile err ", err)
	}

	duration := time.Now().Sub(startTime).Seconds()
	fmt.Println("success! took ", duration, " seconds")
}
