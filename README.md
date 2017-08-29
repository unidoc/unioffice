**gooxml** is a library for creation of OpenOffice documents (.docx, .xlsx and
.pptx).  

[![Build Status](https://travis-ci.org/baliance/gooxml.svg?branch=master)](https://travis-ci.org/baliance/gooxml)

## Status ##

The current focus is on documents (.docx). Some other code exists for
spreadsheets and presentations, but it is even less complete.  The current plan
is to get documents working well, then spreadsheets, and finally presentations.


## Installation ##
    
    go get baliance.com/gooxml/
    go build baliance.com/gooxml/...

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
