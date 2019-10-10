//name: zhexiao
//date: 2019-10-10
//OLEObject struct
//================================start
package document

import "github.com/unidoc/unioffice/schema/soo/wml"

// InlineDrawing is an inlined image within a run.
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
