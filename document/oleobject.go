//name: zhexiao(肖哲)
//date: 2019-10-10
//一些与公式相关的数据结构体
//================================start
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
