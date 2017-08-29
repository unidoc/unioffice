// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package gooxml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"reflect"
	"sort"
)

// Any is the interface used for marshaling/unmarshaling xsd:any
type Any interface {
	MarshalXML(e *xml.Encoder, start xml.StartElement) error
	UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
}

// Raw is used to unmarshal raw XML when we see an unknown tag
type Raw struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:",any,attr"`
	Value   []byte     `xml:",innerxml"`
}

// MarshalXML allows raw to have the Any interface.
func (r *Raw) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	tmp := struct {
		Attrs []xml.Attr `xml:",any,attr"`
		Value []byte     `xml:",innerxml"`
	}{}

	tmp.Value = r.Value
	start.Attr = nil
	s := xml.StartElement{Name: r.XMLName}
	tmpAttrs := make([]xml.Attr, len(r.Attrs))
	copy(tmpAttrs, r.Attrs)

	// fix namespaces in the element we're about to write
	for i := 0; i < len(tmpAttrs); {
		attr := tmpAttrs[i]
		// we unmarshaled an xmlns:foo="http:/foo.com" attribute for a <foo:bar/> element
		if attr.Name.Space == "xmlns" && attr.Value == r.XMLName.Space {
			// add xmlns:foo="http://foo.com" to the element
			s.Attr = append(s.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:" + attr.Name.Local}, Value: s.Name.Space})
			s.Name.Local = attr.Name.Local + ":" + s.Name.Local
			// rewrite <bar xmlns:foo="http://foo.com"/> to <foo:bar xmlns:foo="http://foo.com"/>
			// nuke our namespace which would have been put as xmlns="http://foo.com"
			s.Name.Space = ""
			tmpAttrs[i] = tmpAttrs[len(tmpAttrs)-1]
			tmpAttrs = tmpAttrs[0 : len(tmpAttrs)-1]
		} else {
			s.Attr = append(s.Attr, attr)
			i++
		}
	}

	// ensure consistent output
	sort.Slice(s.Attr, func(i, j int) bool {
		if s.Attr[i].Name.Space != s.Attr[j].Name.Space {
			return s.Attr[i].Name.Space < s.Attr[j].Name.Space
		}
		return s.Attr[i].Name.Local < s.Attr[j].Name.Local
	})

	if err := e.EncodeElement(tmp, s); err != nil {
		return err
	}
	return nil
}

// UnmarshalXML allows raw to have the Any interface.
func (r *Raw) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tmp := struct {
		XMLName xml.Name
		Attrs   []xml.Attr `xml:",any,attr"`
		Value   []byte     `xml:",innerxml"`
	}{}
	if err := d.DecodeElement(&tmp, &start); err != nil {
		return err
	}
	r.XMLName = tmp.XMLName
	r.Attrs = tmp.Attrs
	r.Value = tmp.Value
	return nil
}

var creatorFns = map[string]interface{}{}

// RegisterConstructor registers a constructor function used for unmarshaling
// xsd:any elements.
func RegisterConstructor(ns, name string, fn interface{}) {
	creatorFns[ns+"/"+name] = fn
}

// CreateElement creates an element with the given namespace and name.
func CreateElement(start xml.StartElement) (Any, error) {
	fn, ok := creatorFns[start.Name.Space+"/"+start.Name.Local]
	if !ok {
		r := &Raw{}
		return r, nil
	}

	v := reflect.ValueOf(fn)
	res := v.Call(nil)
	if len(res) != 1 {
		return nil, fmt.Errorf("constructor function should return one value, got %d", len(res))
	}
	any, ok := res[0].Interface().(Any)
	if !ok {
		return nil, errors.New("constructor function should return any 'Any'")
	}
	return any, nil
}
