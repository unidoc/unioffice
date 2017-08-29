// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"baliance.com/gooxml/presentation"
)

func main() {
	ppt := presentation.New()
	ppt.AddSlide()

	ppt.SaveToFile("simple.pptx")
}
