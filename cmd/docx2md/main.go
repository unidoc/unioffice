// Copyright 2017 Baliance. All rights reserved.

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"baliance.com/gooxml/document"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatalf("pass a single document as a parameter")
	}
	doc, err := document.Open(flag.Arg(0))
	if err != nil {
		log.Fatalf("error opening: %s", err)
	}
	imgCnt := 0
	for _, para := range doc.Paragraphs() {
		if s := para.Style(); s != "" {
			// pull out particular headings
			switch s {
			case "Heading1":
				fmt.Print("# ")
			case "Heading2":
				fmt.Print("## ")
			case "Heading3":
				fmt.Print("### ")
			case "Heading4":
				fmt.Print("#### ")
			}
		}
		for _, run := range para.Runs() {
			// print paragraph text
			if s := run.Text(); s != "" {
				switch {
				case run.IsBold():
					fmt.Printf("**%s**", s)
				case run.IsItalic():
					fmt.Printf("*%s*", s)
				default:
					fmt.Print(s)
				}
			}

			// check for any anchored images
			for _, anc := range run.DrawingAnchored() {
				img, ok := anc.GetImage()
				if ok {
					imgCnt++
					fn := fmt.Sprintf("image%d.%s", imgCnt, img.Format())
					// copy image to disk
					dst, err := os.Create(fn)
					if err != nil {
						log.Fatalf("error creating %s: %s", fn, err)
					}
					f, err := os.Open(img.Path())
					if err != nil {
						log.Fatalf("error reading %s: %s", img.Path(), err)
					}
					_, err = io.Copy(dst, f)
					if err != nil {
						log.Fatalf("error copying image: %s", err)
					}
					f.Close()
					dst.Close()
					// and add a markdown reference to it
					fmt.Printf("\n![%s](%s)\n", fn, fn)
				}
			}
		}
		fmt.Println()
		fmt.Println()
	}
}
