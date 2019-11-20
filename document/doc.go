// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

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
