// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.
package document

import "github.com/unidoc/unioffice/schema/soo/wml"

type OleObjectPath struct {
	rid  string
	path string
}

type OleObjectWmfPath struct {
	rid  string
	path string
}

type OleObject struct {
	oleobject *wml.CT_OleObject
	shape     *wml.CT_Shape
}

func (o OleObject) Shape() *wml.CT_Shape {
	return o.shape
}

func (o OleObject) OleObject() *wml.CT_OleObject {
	return o.oleobject
}

func (o OleObject) OleRid() string {
	return *o.oleobject.IdAttr
}

func (o OleObject) ImagedataRid() string {
	return *o.shape.Imagedata.IdAttr
}

func (o OleObjectPath) Rid() string {
	return o.rid
}

func (o OleObjectPath) Path() string {
	return o.path
}

func (o OleObjectWmfPath) Rid() string {
	return o.rid
}

func (o OleObjectWmfPath) Path() string {
	return o.path
}
