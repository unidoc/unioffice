// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"fmt"

	"baliance.com/gooxml/presentation"
)

func main() {
	ppt := presentation.New()
	ppt.AddSlide()

	fmt.Println(ppt.Validate())
	ppt.SaveToFile("simple.pptx")
}
