// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

/*

Package document provides creation, reading, and writing of ECMA 376 Open
Office XML documents.

Example:

	doc := document.New()
	para := doc.AddParagraph()
	run := para.AddRun()
	run.SetText("foo")
	doc.SaveToFile("foo.docx")
*/
package document
