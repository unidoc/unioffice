// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"log"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
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
	hdrID := s.d.docRels.FindRIDForN(h.Index(), unioffice.HeaderType)
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
	hdrID := s.d.docRels.FindRIDForN(f.Index(), unioffice.FooterType)
	if hdrID == "" {
		log.Print("unable to determine footer ID")
	}
	ftrRef.FooterReference.IdAttr = hdrID
}

// SetPageMargins sets the page margins for a section
func (s Section) SetPageMargins(top, right, bottom, left, header, footer, gutter measurement.Distance) {

	margins := wml.NewCT_PageMar()
	margins.TopAttr.Int64 = unioffice.Int64(int64(top / measurement.Twips))
	margins.BottomAttr.Int64 = unioffice.Int64(int64(bottom / measurement.Twips))
	margins.RightAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(right / measurement.Twips))
	margins.LeftAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(left / measurement.Twips))
	margins.HeaderAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(header / measurement.Twips))
	margins.FooterAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(footer / measurement.Twips))
	margins.GutterAttr.ST_UnsignedDecimalNumber = unioffice.Uint64(uint64(gutter / measurement.Twips))

	s.x.PgMar = margins
}
