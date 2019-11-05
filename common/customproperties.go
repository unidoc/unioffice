// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package common

import (
	"time"

	"github.com/unidoc/unioffice/schema/soo/ofc/custom_properties"
	"github.com/unidoc/unioffice/schema/soo/ofc/docPropsVTypes"
)

// CustomProperties contains document specific properties.
type CustomProperties struct {
	x *custom_properties.Properties
}

// CustomProperty contains document specific property
type CustomProperty *custom_properties.CT_Property

// NewCustomProperties constructs a new CustomProperties.
func NewCustomProperties() CustomProperties {
	return CustomProperties{x: custom_properties.NewProperties()}
}

// X returns the inner wrapped XML type.
func (c CustomProperties) X() *custom_properties.Properties {
	return c.x
}

func (c CustomProperties) PropertiesList() []*custom_properties.CT_Property {
	return c.x.Property
}

func (c CustomProperties) GetPropertyByName(name string) CustomProperty {
	propsList := c.x.Property
	for _, property := range propsList {
		if *property.NameAttr == name {
			return CustomProperty(property)
		}
	}
	return nil
}

func (c CustomProperties) getNewProperty(name string) *custom_properties.CT_Property {
	list := c.x.Property
	maxPid := int32(1)
	for _, p := range list {
		if p.PidAttr > maxPid {
			maxPid = p.PidAttr
		}
	}
	newProperty := custom_properties.NewCT_Property()
	newProperty.NameAttr = &name
	newProperty.PidAttr = maxPid + 1
	return newProperty
}

func (c CustomProperties) setProperty(newProperty *custom_properties.CT_Property) {
	existingProperty := c.GetPropertyByName(*newProperty.NameAttr)
	if existingProperty == nil {
		c.x.Property = append(c.x.Property, newProperty)
	} else {
		newProperty.FmtidAttr = existingProperty.FmtidAttr
		if existingProperty.PidAttr == 0 {
			newProperty.PidAttr = existingProperty.PidAttr
		}
		newProperty.LinkTargetAttr = existingProperty.LinkTargetAttr
		*existingProperty = *newProperty
	}
}

func (c CustomProperties) SetPropertyAsVector(name string, vector *docPropsVTypes.Vector) {
	property := c.getNewProperty(name)
	property.Vector = vector
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsArray(name string, array *docPropsVTypes.Array) {
	property := c.getNewProperty(name)
	property.Array = array
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsBlob(name, blob string) {
	property := c.getNewProperty(name)
	property.Blob = &blob
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsOblob(name, oblob string) {
	property := c.getNewProperty(name)
	property.Oblob = &oblob
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsEmpty(name string) {
	property := c.getNewProperty(name)
	property.Empty = docPropsVTypes.NewEmpty()
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsNull(name string) {
	property := c.getNewProperty(name)
	property.Null = docPropsVTypes.NewNull()
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsI1(name string, i1 int8) {
	property := c.getNewProperty(name)
	property.I1 = &i1
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsI2(name string, i2 int16) {
	property := c.getNewProperty(name)
	property.I2 = &i2
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsI4(name string, i4 int32) {
	property := c.getNewProperty(name)
	property.I4 = &i4
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsI8(name string, i8 int64) {
	property := c.getNewProperty(name)
	property.I8 = &i8
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsInt(name string, i int) {
	property := c.getNewProperty(name)
	newValue := int32(i)
	property.Int = &newValue
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsUi1(name string, ui1 uint8) {
	property := c.getNewProperty(name)
	property.Ui1 = &ui1
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsUi2(name string, ui2 uint16) {
	property := c.getNewProperty(name)
	property.Ui2 = &ui2
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsUi4(name string, ui4 uint32) {
	property := c.getNewProperty(name)
	property.Ui4 = &ui4
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsUi8(name string, ui8 uint64) {
	property := c.getNewProperty(name)
	property.Ui8 = &ui8
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsUint(name string, ui uint) {
	property := c.getNewProperty(name)
	newValue := uint32(ui)
	property.Uint = &newValue
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsR4(name string, r4 float32) {
	property := c.getNewProperty(name)
	property.R4 = &r4
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsR8(name string, r8 float64) {
	property := c.getNewProperty(name)
	property.R8 = &r8
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsDecimal(name string, decimal float64) {
	property := c.getNewProperty(name)
	property.Decimal = &decimal
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsLpstr(name string, lpstr string) {
	property := c.getNewProperty(name)
	property.Lpstr = &lpstr
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsLpwstr(name string, lpwstr string) {
	property := c.getNewProperty(name)
	property.Lpwstr = &lpwstr
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsBstr(name string, bstr string) {
	property := c.getNewProperty(name)
	property.Bstr = &bstr
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsCy(name string, cy string) {
	property := c.getNewProperty(name)
	property.Cy = &cy
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsError(name string, error string) {
	property := c.getNewProperty(name)
	property.Error = &error
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsStream(name string, stream string) {
	property := c.getNewProperty(name)
	property.Stream = &stream
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsOstream(name string, ostream string) {
	property := c.getNewProperty(name)
	property.Ostream = &ostream
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsStorage(name string, storage string) {
	property := c.getNewProperty(name)
	property.Storage = &storage
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsOstorage(name string, ostorage string) {
	property := c.getNewProperty(name)
	property.Ostorage = &ostorage
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsClsid(name string, clsid string) {
	property := c.getNewProperty(name)
	property.Clsid = &clsid
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsDate(name string, date time.Time) {
	property := c.getNewProperty(name)
	property.Date = &date
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsFiletime(name string, filetime time.Time) {
	property := c.getNewProperty(name)
	property.Filetime = &filetime
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsBool(name string, b bool) {
	property := c.getNewProperty(name)
	property.Bool = &b
	c.setProperty(property)
}

func (c CustomProperties) SetPropertyAsVstream(name string, vstream *docPropsVTypes.Vstream) {
	property := c.getNewProperty(name)
	property.Vstream = vstream
	c.setProperty(property)
}
