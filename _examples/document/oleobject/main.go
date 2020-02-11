// oleObject using demo

// $ cd unioffice/_examples/document/oleobject
// $ go run main.go

// print:
//=============oleObject info===========
//Data path = C:\Users\ADMINI~1\AppData\Local\Temp\gooxml-docx131058703\zz524008729
//Wmf path = C:\Users\ADMINI~1\AppData\Local\Temp\gooxml-docx131058703\zz731151394
//===============end=============
//=============oleObject info===========
//Data path = C:\Users\ADMINI~1\AppData\Local\Temp\gooxml-docx131058703\zz304145587
//Wmf path = C:\Users\ADMINI~1\AppData\Local\Temp\gooxml-docx131058703\zz059105188
//===============end=============
//=============oleObject info===========
//Data path = C:\Users\ADMINI~1\AppData\Local\Temp\gooxml-docx131058703\zz462161757
//Wmf path = C:\Users\ADMINI~1\AppData\Local\Temp\gooxml-docx131058703\zz101524598
//===============end=============
//

// Comments:
// I developed a project can do convert the oleObject file to LaTeX data
// project path: https://github.com/zhexiao/mtef-go

// code Demo:
//import "github.com/zhexiao/mtef-go/eqn"
//latexData := eqn.Convert(path)

package main

import (
	"fmt"
	"github.com/unidoc/unioffice/document"
	"log"
)

func main() {
	filepath := "oleobject.docx"

	doc, err := document.Open(filepath)
	if err != nil {
		log.Panicf("Open file failed, err=%s", err)
	}

	//save oleObject data rid/filepath relationship
	oleObjectDataMap := make(map[string]string)
	for _, oleData := range doc.OleObjectPaths {
		path := oleData.Path()
		rid := oleData.Rid()

		oleObjectDataMap[rid] = path
	}

	//save oleObject wmf rid/filepath relationship
	oleObjectWmfMap := make(map[string]string)
	for _, oleData := range doc.OleObjectWmfPath {
		path := oleData.Path()
		rid := oleData.Rid()

		oleObjectWmfMap[rid] = path
	}

	// read the oleObject file information from the word data
	for _, paragraph := range doc.Paragraphs() {
		for _, run := range paragraph.Runs() {
			if run.OleObjects() != nil {
				for _, ole := range run.OleObjects() {
					dataRid := ole.OleRid()
					wmfRid := ole.ImagedataRid()

					fmt.Println("=============oleObject info===========")
					fmt.Printf("Data path = %s \n", oleObjectDataMap[dataRid])
					fmt.Printf("Wmf path = %s \n", oleObjectWmfMap[wmfRid])
					fmt.Println("===============end=============")
				}
			}
		}
	}
}
