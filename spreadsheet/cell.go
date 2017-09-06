// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"errors"
	"log"
	"strconv"
	"time"

	"baliance.com/gooxml"
	sml "baliance.com/gooxml/schema/schemas.openxmlformats.org/spreadsheetml"
)

const iso8601Format = "2006-01-02T15:04:05Z07:00"

// Cell is a single cell within a sheet.
type Cell struct {
	w *Workbook
	x *sml.CT_Cell
}

// X returns the inner wrapped XML type.
func (c Cell) X() *sml.CT_Cell {
	return c.x
}

func (c Cell) clearValue() {
	c.x.F = nil
	c.x.Is = nil
	c.x.V = nil
}

// SetInlineString adds a string inline instead of in the shared strings table.
func (c Cell) SetInlineString(s string) {
	c.clearValue()
	c.x.Is = sml.NewCT_Rst()
	c.x.Is.T = gooxml.String(s)
	c.x.TAttr = sml.ST_CellTypeInlineStr
}

func (c Cell) SetRichTextString() RichText {
	c.clearValue()
	c.x.Is = sml.NewCT_Rst()
	c.x.TAttr = sml.ST_CellTypeInlineStr
	return RichText{c.x.Is}
}

// SetFormulaRaw sets the cell type to formula, and the raw formula to the given string
func (c Cell) SetFormulaRaw(s string) {
	c.x.TAttr = sml.ST_CellTypeStr
	c.x.F = sml.NewCT_CellFormula()
	c.x.F.Content = s
	c.x.V = nil
}

// SetString sets the cell type to string, and the value to the given string,
// returning an ID from the shared strings table. To reuse a string, call
// SetStringByID with the ID returned.
func (c Cell) SetString(s string) int {
	c.clearValue()
	id := c.w.SharedStrings.AddString(s)
	c.x.V = gooxml.String(strconv.Itoa(id))
	c.x.TAttr = sml.ST_CellTypeS
	return id
}

// SetStringByID sets the cell type to string, and the value a string in the
// shared strings table.
func (c Cell) SetStringByID(id int) {
	c.clearValue()
	c.x.V = gooxml.String(strconv.Itoa(id))
	c.x.TAttr = sml.ST_CellTypeS
}

// SetNumber sets the cell type to number, and the value to the given number
func (c Cell) SetNumber(v float64) {
	c.clearValue()
	c.x.V = gooxml.String(strconv.FormatFloat(v, 'g', -1, 64))
	// cell type number
	c.x.TAttr = sml.ST_CellTypeN
}

// SetNumberWithStyle sets a number and applies a standard format to the cell.
func (c Cell) SetNumberWithStyle(v float64, f StandardFormat) {
	c.SetNumber(v)
	c.SetStyle(c.w.StyleSheet.GetOrCreateStandardNumberFormat(f))
}

// SetBool sets the cell type to boolean and the value to the given boolean
// value.
func (c Cell) SetBool(v bool) {
	c.clearValue()
	c.x.V = gooxml.String(strconv.Itoa(b2i(v)))
	c.x.TAttr = sml.ST_CellTypeB
}

func asUTC(d time.Time) time.Time {
	// Excel appears to interpret and serial dates in the local timezone, so
	// first ensure the time is converted internally.
	d = d.Local()

	// Then to avoid any daylight savings differences showing up between our
	// epoch and the current time, we 'cast' the time to UTC and later subtract
	// from the epoch in UTC.
	return time.Date(d.Year(), d.Month(), d.Day(), d.Hour(),
		d.Minute(), d.Second(), d.Nanosecond(), time.UTC)
}

// SetTime sets the cell value to a date. It's stored as the number of days past
// th sheet epoch. When we support v5 strict, we can store an ISO 8601 date
// string directly, however that's not allowed with v5 transitional  (even
// though it works in Excel).
func (c Cell) SetTime(d time.Time) {
	c.clearValue()
	d = asUTC(d)
	epoch := c.w.Epoch()
	if d.Before(epoch) {
		// the ECMA 376 standard says these works, but Excel doesn't appear to
		// support negative serial dates
		log.Printf("times before 1900 are not supported")
		return
	}
	delta := d.Sub(epoch)
	c.x.V = gooxml.Stringf("%G", delta.Hours()/24)
}

// SetDate sets the cell value to a date. It's stored as the number of days past
// th sheet epoch. When we support v5 strict, we can store an ISO 8601 date
// string directly, however that's not allowed with v5 transitional  (even
// though it works in Excel). The cell is not styled via this method, so it will
// display as a number. SetDateWithStyle should normally be used instead.
func (c Cell) SetDate(d time.Time) {
	c.clearValue()
	d = asUTC(d)
	epoch := c.w.Epoch()
	if d.Before(epoch) {
		// the ECMA 376 standard says these works, but Excel doesn't appear to
		// support negative serial dates
		log.Printf("dates before 1900 are not supported")
		return
	}
	delta := d.Sub(epoch)
	c.x.V = gooxml.Stringf("%d", int(delta.Hours()/24))
}

// SetDateWithStyle sets a date with the default date style applied.
func (c Cell) SetDateWithStyle(d time.Time) {
	c.SetDate(d)
	for _, cs := range c.w.StyleSheet.CellStyles() {
		// found an existing number format
		if cs.HasNumberFormat() && cs.NumberFormat() == uint32(StandardFormatDate) {
			c.SetStyle(cs)
			return
		}
	}
	// need to create a new format
	cs := c.w.StyleSheet.AddCellStyle()
	cs.SetNumberFormatStandard(StandardFormatDate)
	c.SetStyle(cs)
}

// SetStyle applies a style to the cell.  This style is referenced in the
// generated XML via CellStyle.Index().
func (c Cell) SetStyle(cs CellStyle) {
	c.SetStyleIndex(cs.Index())
}

// SetStyleIndex directly sets a style index to the cell.  This should only be
// called with an index retrieved from CellStyle.Index()
func (c Cell) SetStyleIndex(idx uint32) {
	c.x.SAttr = gooxml.Uint32(idx)
}

func (c Cell) GetValue() (string, error) {
	switch c.x.TAttr {
	case sml.ST_CellTypeB:
		if c.x.V == nil {
			return "", nil
		}
		return *c.x.V, nil
	case sml.ST_CellTypeInlineStr:
		if c.x.Is == nil || c.x.Is.T == nil {
			return "", nil
		}
		return *c.x.Is.T, nil
	case sml.ST_CellTypeS:
		if c.x.V == nil {
			return "", nil
		}
		id, err := strconv.Atoi(*c.x.V)
		if err != nil {
			return "", err
		}
		return c.w.SharedStrings.GetString(id)
	case sml.ST_CellTypeE:
	case sml.ST_CellTypeN:
	case sml.ST_CellTypeStr:
	default:
	}
	return "", errors.New("unsupported cell type")
}

func b2i(v bool) int {
	if v {
		return 1
	}
	return 0
}
