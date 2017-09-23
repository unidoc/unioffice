// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"log"

	"baliance.com/gooxml"
	"baliance.com/gooxml/schema/soo/wml"
)

// Section is the beginning of a new section.
type Section struct {
	d *Document
	x *wml.CT_SectPr
}

// X returns the internally wrapped *wml.CT_SectPr.
func (s Section) X() *wml.CT_SectPr {
	return s.x
}

// SetHeader sets a section header.
func (s Section) SetHeader(h Header, t wml.ST_HdrFtr) {
	hdrRef := wml.NewEG_HdrFtrReferences()
	s.x.EG_HdrFtrReferences = append(s.x.EG_HdrFtrReferences, hdrRef)
	hdrRef.HeaderReference = wml.NewCT_HdrFtrRef()
	hdrRef.HeaderReference.TypeAttr = t
	hdrID := s.d.docRels.FindRIDForN(h.Index(), gooxml.HeaderType)
	if hdrID == "" {
		log.Print("unable to determine header ID")
	}
	hdrRef.HeaderReference.IdAttr = hdrID
}

// SetFooter sets a section footer.
func (s Section) SetFooter(f Footer, t wml.ST_HdrFtr) {
	ftrRef := wml.NewEG_HdrFtrReferences()
	s.x.EG_HdrFtrReferences = append(s.x.EG_HdrFtrReferences, ftrRef)
	ftrRef.FooterReference = wml.NewCT_HdrFtrRef()
	ftrRef.FooterReference.TypeAttr = t
	hdrID := s.d.docRels.FindRIDForN(f.Index(), gooxml.FooterType)
	if hdrID == "" {
		log.Print("unable to determine footer ID")
	}
	ftrRef.FooterReference.IdAttr = hdrID
}
