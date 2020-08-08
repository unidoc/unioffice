// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

/*

Package unioffice provides creation, reading, and writing of ECMA 376 Office Open
XML documents, spreadsheets and presentations.  It is still early in
development, but is progressing quickly.  This library takes a slightly
different approach from others, in that it starts by trying to support all of
the ECMA-376 standard when marshaling/unmarshaling XML documents.  From there it
adds wrappers around the ECMA-376 derived types that provide a more convenient
interface.

The raw XML based types reside in the `schema/`` directory. These types are
always accessible from the wrapper types via a `X() method that returns the
raw type.  Except for the base documents (document.Document,
spreadsheet.Workbook and presentation.Presentation), the other wrapper types are
value types with non-pointer methods.  They exist solely to modify and return
data from one or more XML types.

The packages of interest are github.com/unidoc/unioffice/document,
unidoc/unioffice/spreadsheet and github.com/unidoc/unioffice/presentation.

*/
package unioffice // import "github.com/unidoc/unioffice"
