// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package gooxml

import (
	"encoding/xml"
	"strings"
	"unicode"
)

// XSDAny  is used to marshal/unmarshal xsd:any types in the OOXML schema.
type XSDAny struct {
	XMLName xml.Name
	Attrs   []xml.Attr
	Data    []byte
	Nodes   []*XSDAny
}

var wellKnownSchemas = map[string]string{
	"a":       "http://schemas.openxmlformats.org/drawingml/2006/main",
	"dc":      "http://purl.org/dc/elements/1.1/",
	"dcterms": "http://purl.org/dc/terms/",
	"mc":      "http://schemas.openxmlformats.org/markup-compatibility/2006",
	"mo":      "http://schemas.microsoft.com/office/mac/office/2008/main",
	"w":       "http://schemas.openxmlformats.org/wordprocessingml/2006/main",
	"w10":     "urn:schemas-microsoft-com:office:word",
	"w14":     "http://schemas.microsoft.com/office/word/2010/wordml",
	"w15":     "http://schemas.microsoft.com/office/word/2012/wordml",
	"wne":     "http://schemas.microsoft.com/office/word/2006/wordml",
	"wp":      "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing",
	"wp14":    "http://schemas.microsoft.com/office/word/2010/wordprocessingDrawing",
	"wpc":     "http://schemas.microsoft.com/office/word/2010/wordprocessingCanvas",
	"wpg":     "http://schemas.microsoft.com/office/word/2010/wordprocessingGroup",
	"wpi":     "http://schemas.microsoft.com/office/word/2010/wordprocessingInk",
	"wps":     "http://schemas.microsoft.com/office/word/2010/wordprocessingShape",
	"xsi":     "http://www.w3.org/2001/XMLSchema-instance",
	"x15ac":   "http://schemas.microsoft.com/office/spreadsheetml/2010/11/ac",
}

var wellKnownSchemasInv = func() map[string]string {
	r := map[string]string{}
	for pfx, ns := range wellKnownSchemas {
		r[ns] = pfx
	}
	return r
}()

type any struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:",any,attr"`
	Nodes   []*any     `xml:",any"`
	Data    []byte     `xml:",chardata"`
}

func dd(a *any) {
	for _, n := range a.Nodes {
		dd(n)
	}
}

// UnmarshalXML implements the xml.Unmarshaler interface.
func (x *XSDAny) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	a := any{}
	if err := d.DecodeElement(&a, &start); err != nil {
		return err
	}
	dd(&a)
	x.XMLName = a.XMLName
	x.Attrs = a.Attrs
	x.Data = a.Data
	x.Nodes = convertToXNodes(a.Nodes)
	return nil
}

type nsSet struct {
	urlToPrefix map[string]string
	prefixToURL map[string]string
	prefixes    []string //required for deterministic output
}

func (n *nsSet) getPrefix(ns string) string {
	// Common namespaces are used in these 'any' elements and some versions
	// of Word really want to the prefix to match what they write out.  This
	// occurred primarily with docProps/core.xml
	if pfx, ok := wellKnownSchemasInv[ns]; ok {
		if _, ok := n.prefixToURL[pfx]; !ok {
			n.prefixToURL[pfx] = ns
			n.urlToPrefix[ns] = pfx
			n.prefixes = append(n.prefixes, pfx)
		}
		return pfx
	}

	// trying to construct a decent looking valid prefix
	ns = strings.TrimFunc(ns, func(r rune) bool {
		return !unicode.IsLetter(r)
	})

	// do we have a prefix for this ns?
	if sc, ok := n.urlToPrefix[ns]; ok {
		return sc
	}

	// determine the last path portion of the namespace
	// "urn:schemas-microsoft-com:office:office" = "office"
	// "http://schemas.microsoft.com/office/word/2012/wordml" = "wordml"
	split := strings.Split(ns, "/")
	split = strings.Split(split[len(split)-1], ":")
	// last segment of the namesapce
	last := split[len(split)-1]
	lng := 0
	pfx := []byte{}
	for {
		if lng < len(last) {
			pfx = append(pfx, last[lng])
		} else {
			pfx = append(pfx, '_')
		}
		lng++
		// is this prefix unused?
		if _, ok := n.prefixToURL[string(pfx)]; !ok {
			n.prefixToURL[string(pfx)] = ns
			n.urlToPrefix[ns] = string(pfx)
			n.prefixes = append(n.prefixes, string(pfx))
			return string(pfx)
		}
	}
}

func (n nsSet) applyToNode(a *any) {
	if a.XMLName.Space == "" {
		return
	}
	pfx := n.getPrefix(a.XMLName.Space)
	a.XMLName.Space = ""
	a.XMLName.Local = pfx + ":" + a.XMLName.Local
	tmpAttr := a.Attrs
	a.Attrs = nil
	for _, attr := range tmpAttr {
		// skip namespace prefix declaration atributes as we create them later
		if attr.Name.Space == "xmlns" {
			continue
		}
		if attr.Name.Space != "" {
			pfx := n.getPrefix(attr.Name.Space)
			attr.Name.Space = ""
			attr.Name.Local = pfx + ":" + attr.Name.Local
		}
		a.Attrs = append(a.Attrs, attr)
	}
	for _, cn := range a.Nodes {
		n.applyToNode(cn)
	}
}

// collectNS walks a tree of nodes finding any non-default namespace being used
func (x *XSDAny) collectNS(ns *nsSet) {
	if x.XMLName.Space != "" {
		ns.getPrefix(x.XMLName.Space)
	}
	for _, attr := range x.Attrs {
		if attr.Name.Space != "" && attr.Name.Space != "xmlns" {
			ns.getPrefix(attr.Name.Space)
		}
	}
	for _, n := range x.Nodes {
		n.collectNS(ns)
	}
}

func convertToXNodes(an []*any) []*XSDAny {
	ret := []*XSDAny{}
	for _, a := range an {
		x := &XSDAny{}
		x.XMLName = a.XMLName
		x.Attrs = a.Attrs
		x.Data = a.Data
		x.Nodes = convertToXNodes(a.Nodes)
		ret = append(ret, x)
	}
	return ret
}
func convertToNodes(xn []*XSDAny) []*any {
	ret := []*any{}
	for _, x := range xn {
		a := &any{}
		a.XMLName = x.XMLName
		a.Attrs = x.Attrs
		a.Data = x.Data
		a.Nodes = convertToNodes(x.Nodes)
		ret = append(ret, a)
	}
	return ret
}

// MarshalXML implements the xml.Marshaler interface.
func (x *XSDAny) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = x.XMLName
	start.Attr = x.Attrs
	a := any{}
	a.XMLName = x.XMLName
	a.Attrs = x.Attrs
	a.Data = x.Data
	a.Nodes = convertToNodes(x.Nodes)

	ns := nsSet{
		urlToPrefix: map[string]string{},
		prefixToURL: map[string]string{},
	}

	// collect any namespaces in use in the node tree
	x.collectNS(&ns)

	// apply our new namespaces to the node and its children
	ns.applyToNode(&a)

	// add our prefixes and namespaces to root element
	for _, pfx := range ns.prefixes {
		ns := ns.prefixToURL[pfx]
		a.Attrs = append(a.Attrs, xml.Attr{
			Name:  xml.Name{Local: "xmlns:" + pfx},
			Value: ns,
		})
	}

	// finally write out our new element
	return e.Encode(&a)
}
