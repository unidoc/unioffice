**gooxml** is a library for creation of Office Open XML documents (.docx, .xlsx
and .pptx).  

[![Build Status](https://travis-ci.org/baliance/gooxml.svg?branch=master)](https://travis-ci.org/baliance/gooxml)
[![GitHub (pre-)release](https://img.shields.io/github/release/baliance/gooxml/all.svg)]()
[![License: AGPL v3](https://img.shields.io/badge/License-Dual%20AGPL%20v3/Commercial-blue.svg)](https://www.gnu.org/licenses/agpl-3.0)
[![GoDoc](https://godoc.org/baliance.com/gooxml?status.svg)](https://godoc.org/baliance.com/gooxml)
[![Go Report Card](https://goreportcard.com/badge/baliance.com/gooxml)](https://goreportcard.com/report/baliance.com/gooxml)

![https://baliance.com/gooxml/](./_examples/preview.png "gooxml")

## Status ##

- Documents (docx) work well, and there are no known issues.
- Spreadsheets (xlsx) support is new and the API may change.  
- PowerPoint (pptx) is unsupported at the moment, the XML types exist and some
  prototype code is checked in but it will be reworked once docx/xlsx are
  'finished'.

## Installation ##
    
    go get baliance.com/gooxml/
    go build -i baliance.com/gooxml/...

## Document Examples ##

- [Simple Text Formatting](https://github.com/baliance/gooxml/tree/master/_examples/document/simple) Text font colors, sizes, highlighting, etc.
- [Auto Generated Table of Contents](https://github.com/baliance/gooxml/tree/master/_examples/document/toc) Creating document headings with an auto generated TOC based off of the headingds
- [Floating Image](https://github.com/baliance/gooxml/tree/master/_examples/document/image) Placing an image somewhere on a page, absolutely positioned with different text wrapping.
- [Header & Footer](https://github.com/baliance/gooxml/tree/master/_examples/document/header-footer) Creating headers and footers including page numbering.
- [Multiple Headers & Footers](https://github.com/baliance/gooxml/tree/master/_examples/document/header-footer-multiple) Using different headers and footers depending on document section.
- [Inline Tables](https://github.com/baliance/gooxml/tree/master/_examples/document/tables) Adding an table with and without borders.
- [Using Existing Word Document as a Template](https://github.com/baliance/gooxml/tree/master/_examples/document/use-template) Opening a document as a template to re-use the styles created in the document.
- [Filling out Form Fields](https://github.com/baliance/gooxml/tree/master/_examples/document/fill-out-form) Opening a document with embedded form fields, filling out the fields and saving the result as  a new filled form.

## Spreadsheet Examples ##
- [Simple](https://github.com/baliance/gooxml/tree/master/_examples/spreadsheet/simple) A simple sheet with a few cells
- [Named Cells](https://github.com/baliance/gooxml/tree/master/_examples/spreadsheet/named-cells) Different ways of referencing rows and cells
- [Cell Number/Date/Time Formats](https://github.com/baliance/gooxml/tree/master/_examples/spreadsheet/number-date-time-formats) Creating cells with various number/date/time formats
- [Line Chart](https://github.com/baliance/gooxml/tree/master/_examples/spreadsheet/line-chart)/[Line Chart 3D](https://github.com/baliance/gooxml/tree/master/_examples/spreadsheet/line-chart-3d) Line Charts
- [Bar Chart](https://github.com/baliance/gooxml/tree/master/_examples/spreadsheet/bar-chart) Bar Charts
- [Mutiple Charts](https://github.com/baliance/gooxml/tree/master/_examples/spreadsheet/multiple-charts) Multiple charts on a single sheet
- [Named Cell Ranges](https://github.com/baliance/gooxml/tree/master/_examples/spreadsheet/named-ranges) Naming cell ranges
- [Merged Cells](https://github.com/baliance/gooxml/tree/master/_examples/spreadsheet/merged) Merge and unmerge cells
- [Conditional Formatting](https://github.com/baliance/gooxml/tree/master/_examples/spreadsheet/conditional-formatting) Conditionally formatting cells, styling, gradients, icons, data bar
- [Complex](https://github.com/baliance/gooxml/tree/master/_examples/spreadsheet/complex) Multiple charts, auto filtering and conditional formatting
- [Borders](https://github.com/baliance/gooxml/tree/master/_examples/spreadsheet/borders) Individual cell borders and rectangular borders around a range of cells.
- [Validation](https://github.com/baliance/gooxml/tree/master/_examples/spreadsheet/validation) Data validation including combo box dropdowns.
- [Frozen Rows/Cols](https://github.com/baliance/gooxml/tree/master/_examples/spreadsheet/freeze-rows-cols) A sheet with a frozen header column and row

## Raw Types ##

The OOXML specification is large and creating a friendly API to cover the entire
specification is a very time consuming endeavor.  This library attempts to
provide an easy to use API for common use cases in creating OOXML documents
while allowing users to fall back to raw document manipulation should the
library's API not cover a specific use case.

The raw XML based types reside in the ```schema/``` directory. These types are
accessible from the wrapper types via a ```X()``` method that returns the raw
type. 

For example, the library currently doesn't have an API for setting a document
background color. However it's easy to do manually via editing the
```CT_Background``` element of the document.

    dox := document.New()
    doc.X().Background = wordprocessingml.NewCT_Background()
	doc.X().Background.ColorAttr = &wordprocessingml.ST_HexColor{}
	doc.X().Background.ColorAttr.ST_HexColorRGB = color.RGB(50, 50, 50).AsRGBString()

### Contribution guidelines ###

[![CLA assistant](https://cla-assistant.io/readme/badge/baliance/gooxml)](https://cla-assistant.io/baliance/gooxml)

All contributors are must sign a contributor license agreement before their code
will be reviewed and merged.


### Licensing ###

This library is offered under a dual license. It is freely available for use
under the terms of AGPLv3. If you would like to use this library for a closed
source project, please contact sales@baliance.com.

There are no differences in functionality between the open source and commercial 
versions. You are encouraged to use the open source version to evaluate the library
before purchasing a commercial license.

### Consulting ###

Baliance also offers consulting services related to enhancing the gooxml library
on a case by case basis. Please contact consulting@baliance.com if interested.
