// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package unioffice

import (
	"encoding/xml"
	"errors"
	"fmt"
	"reflect"
)

// Any is the interface used for marshaling/unmarshaling xsd:any
type Any interface {
	MarshalXML(e *xml.Encoder, start xml.StartElement) error
	UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
}

var creatorFns = map[string]interface{}{}

// RegisterConstructor registers a constructor function used for unmarshaling
// xsd:any elements.
func RegisterConstructor(ns, name string, fn interface{}) {
	creatorFns[ns+"/"+name] = fn
}

// CreateElement creates an element with the given namespace and name. It is
// used to unmarshal some xsd:any elements to the appropriate concrete type.
func CreateElement(start xml.StartElement) (Any, error) {
	fn, ok := creatorFns[start.Name.Space+"/"+start.Name.Local]
	if !ok {
		r := &XSDAny{}
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
