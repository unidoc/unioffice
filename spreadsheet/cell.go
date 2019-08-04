// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/common"
	"github.com/unidoc/unioffice/schema/soo/sml"
	"github.com/unidoc/unioffice/spreadsheet/format"
	"github.com/unidoc/unioffice/spreadsheet/reference"
)

const iso8601Format = "2006-01-02T15:04:05Z07:00"

// Cell is a single cell within a sheet.
type Cell struct {
	w *Workbook
	s *sml.Worksheet
	r *sml.CT_Row
	x *sml.CT_Cell
}

// X returns the inner wrapped XML type.
func (c Cell) X() *sml.CT_Cell {
	return c.x
}

// Reference returns the cell reference (e.g. "A4"). This is not required,
// however both gooxml and Excel will always set it.
func (c Cell) Reference() string {
	if c.x.RAttr != nil {
		return *c.x.RAttr
	}
	return ""
}

// Clear clears the cell's value and type.
func (c Cell) Clear() {
	c.clearValue()
	c.x.TAttr = sml.ST_CellTypeUnset
}

func (c Cell) clearValue() {
	c.x.F = nil
	c.x.Is = nil
	c.x.V = nil
	c.x.TAttr = sml.ST_CellTypeUnset
}

// SetInlineString adds a string inline instead of in the shared strings table.
func (c Cell) SetInlineString(s string) {
	c.clearValue()
	c.x.Is = sml.NewCT_Rst()
	c.x.Is.T = unioffice.String(s)
	c.x.TAttr = sml.ST_CellTypeInlineStr
}

// SetRichTextString sets the cell to rich string mode and returns a struct that
// can be used to add formatted text to the cell.
func (c Cell) SetRichTextString() RichText {
	c.clearValue()
	c.x.Is = sml.NewCT_Rst()
	c.x.TAttr = sml.ST_CellTypeInlineStr
	return RichText{c.x.Is}
}

// SetFormulaRaw sets the cell type to formula, and the raw formula to the given string
func (c Cell) SetFormulaRaw(s string) {
	c.clearValue()
	c.x.TAttr = sml.ST_CellTypeStr
	c.x.F = sml.NewCT_CellFormula()
	c.x.F.Content = s
}

// SetFormulaArray sets the cell type to formula array, and the raw formula to
// the given string. This is equivlent to entering a formula and pressing
// Ctrl+Shift+Enter in Excel.
func (c Cell) SetFormulaArray(s string) {
	c.clearValue()
	c.x.TAttr = sml.ST_CellTypeStr
	c.x.F = sml.NewCT_CellFormula()
	c.x.F.TAttr = sml.ST_CellFormulaTypeArray
	c.x.F.Content = s
}

// SetFormulaShared sets the cell type to formula shared, and the raw formula to
// the given string. The range is the range of cells that the formula applies
// to, and is used to conserve disk space.
func (c Cell) SetFormulaShared(formula string, rows, cols uint32) error {
	c.clearValue()
	c.x.TAttr = sml.ST_CellTypeStr
	c.x.F = sml.NewCT_CellFormula()
	c.x.F.TAttr = sml.ST_CellFormulaTypeShared
	c.x.F.Content = formula
	cref, err := reference.ParseCellReference(c.Reference())
	if err != nil {
		return err
	}

	sid := uint32(0)
	for _, r := range c.s.SheetData.Row {
		for _, c := range r.C {
			if c.F != nil && c.F.SiAttr != nil && *c.F.SiAttr >= sid {
				sid = *c.F.SiAttr
			}
		}
	}
	sid++

	ref := fmt.Sprintf("%s%d:%s%d", cref.Column, cref.RowIdx, reference.IndexToColumn(cref.ColumnIdx+cols), cref.RowIdx+rows)
	c.x.F.RefAttr = unioffice.String(ref)
	c.x.F.SiAttr = unioffice.Uint32(sid)
	sheet := Sheet{c.w, nil, c.s}
	for row := cref.RowIdx; row <= cref.RowIdx+rows; row++ {
		for col := cref.ColumnIdx; col <= cref.ColumnIdx+cols; col++ {
			if row == cref.RowIdx && col == cref.ColumnIdx {
				continue
			}
			ref := fmt.Sprintf("%s%d", reference.IndexToColumn(col), row)
			sheet.Cell(ref).Clear()
			sheet.Cell(ref).X().F = sml.NewCT_CellFormula()
			sheet.Cell(ref).X().F.TAttr = sml.ST_CellFormulaTypeShared
			sheet.Cell(ref).X().F.SiAttr = unioffice.Uint32(sid)
		}
	}
	return nil
}

// SetString sets the cell type to string, and the value to the given string,
// returning an ID from the shared strings table. To reuse a string, call
// SetStringByID with the ID returned.
func (c Cell) SetString(s string) int {
	c.clearValue()
	id := c.w.SharedStrings.AddString(s)
	c.x.V = unioffice.String(strconv.Itoa(id))
	c.x.TAttr = sml.ST_CellTypeS
	return id
}

// SetStringByID sets the cell type to string, and the value a string in the
// shared strings table.
func (c Cell) SetStringByID(id int) {
	c.clearValue()
	c.x.V = unioffice.String(strconv.Itoa(id))
	c.x.TAttr = sml.ST_CellTypeS
}

// SetNumber sets the cell type to number, and the value to the given number
func (c Cell) SetNumber(v float64) {
	c.clearValue()
	// NaN / Infinity
	if math.IsNaN(v) || math.IsInf(v, 0) {
		c.x.TAttr = sml.ST_CellTypeE
		c.x.V = unioffice.String("#NUM!")
		return
	}

	// cell type number
	c.x.TAttr = sml.ST_CellTypeN
	c.x.V = unioffice.String(strconv.FormatFloat(v, 'g', -1, 64))
}

// Column returns the cell column
func (c Cell) Column() (string, error) {
	cref, err := reference.ParseCellReference(c.Reference())
	if err != nil {
		return "", err
	}
	return cref.Column, nil
}

func (c Cell) getFormat() string {
	if c.x.SAttr == nil {
		return "General"
	}
	sid := *c.x.SAttr
	f := c.w.StyleSheet.GetCellStyle(sid)
	nf := c.w.StyleSheet.GetNumberFormat(f.NumberFormat())
	return nf.GetFormat()
}

// GetFormattedValue returns the formatted cell value as it would appear in
// Excel. This involves determining the format string to apply, parsing it, and
// then formatting the value according to the format string.  This should only
// be used if you care about replicating what Excel would show, otherwise
// GetValueAsNumber()/GetValueAsTime
func (c Cell) GetFormattedValue() string {
	f := c.getFormat()
	switch c.x.TAttr {
	// boolean
	case sml.ST_CellTypeB:
		b, _ := c.GetValueAsBool()
		if b {
			return "TRUE"
		}
		return "FALSE"
	// number
	case sml.ST_CellTypeN:
		v, _ := c.GetValueAsNumber()
		return format.Number(v, f)
	// error
	case sml.ST_CellTypeE:
		if c.x.V != nil {
			return *c.x.V
		}
		return ""
	// string / inline string
	case sml.ST_CellTypeS, sml.ST_CellTypeInlineStr:
		return format.String(c.GetString(), f)
	case sml.ST_CellTypeStr:
		s := c.GetString()
		if format.IsNumber(s) {
			v, _ := strconv.ParseFloat(s, 64)
			return format.Number(v, f)
		}
		return format.String(s, f)
	case sml.ST_CellTypeUnset:
		fallthrough
	default:
		s, _ := c.GetRawValue()
		// avoid returning zero for an empty cell
		if len(s) == 0 {
			return ""
		}

		v, err := c.GetValueAsNumber()
		if err == nil {
			return format.Number(v, f)
		}
		return format.String(s, f)
	}
}

// GetValueAsNumber retrieves the cell's value as a number
func (c Cell) GetValueAsNumber() (float64, error) {
	if c.x.V == nil && c.x.Is == nil {
		// empty cells have an implicit zero value
		return 0, nil
	}
	if c.x.TAttr == sml.ST_CellTypeS || !format.IsNumber(*c.x.V) {
		return math.NaN(), errors.New("cell is not of number type")
	}
	return strconv.ParseFloat(*c.x.V, 64)
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
	c.x.V = unioffice.String(strconv.Itoa(b2i(v)))
	c.x.TAttr = sml.ST_CellTypeB
}

// GetValueAsBool retrieves the cell's value as a boolean
func (c Cell) GetValueAsBool() (bool, error) {
	if c.x.TAttr != sml.ST_CellTypeB {
		return false, errors.New("cell is not of bool type")
	}
	if c.x.V == nil {
		return false, errors.New("cell has no value")
	}
	return strconv.ParseBool(*c.x.V)
}

func asLocal(d time.Time) time.Time {
	d = d.UTC()
	return time.Date(d.Year(), d.Month(), d.Day(), d.Hour(),
		d.Minute(), d.Second(), d.Nanosecond(), time.Local)
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
		unioffice.Log("times before 1900 are not supported")
		return
	}

	delta := d.Sub(epoch)

	result := new(big.Float)

	deltaNs := new(big.Float)
	deltaNs.SetPrec(128)
	deltaNs.SetUint64(uint64(delta))

	nsPerDay := new(big.Float)
	nsPerDay.SetUint64(24 * 60 * 60 * 1e9)
	result.Quo(deltaNs, nsPerDay)

	c.x.V = unioffice.String(result.Text('g', 20))
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
		unioffice.Log("dates before 1900 are not supported")
		return
	}
	delta := d.Sub(epoch)

	result := new(big.Float)

	deltaNs := new(big.Float)
	deltaNs.SetPrec(128)
	deltaNs.SetUint64(uint64(delta))

	nsPerDay := new(big.Float)
	nsPerDay.SetUint64(24 * 60 * 60 * 1e9)
	result.Quo(deltaNs, nsPerDay)

	hrs, _ := result.Uint64()

	c.x.V = unioffice.Stringf("%d", hrs)
}

// GetValueAsTime retrieves the cell's value as a time.  There is no difference
// in SpreadsheetML between a time/date cell other than formatting, and that
// typically a date cell won't have a fractional component. GetValueAsTime will
// work for date cells as well.
func (c Cell) GetValueAsTime() (time.Time, error) {
	if c.x.TAttr != sml.ST_CellTypeUnset {
		return time.Time{}, errors.New("cell type should be unset")
	}
	if c.x.V == nil {
		return time.Time{}, errors.New("cell has no value")
	}
	f, _, err := big.ParseFloat(*c.x.V, 10, 128, big.ToNearestEven)
	if err != nil {
		return time.Time{}, err
	}

	day := new(big.Float)
	day.SetUint64(uint64(24 * time.Hour))
	f.Mul(f, day)
	ns, _ := f.Uint64()
	t := c.w.Epoch().Add(time.Duration(ns))
	return asLocal(t), nil
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
	c.x.SAttr = unioffice.Uint32(idx)
}

// GetString returns the string in a cell if it's an inline or string table
// string. Otherwise it returns an empty string.
func (c Cell) GetString() string {
	switch c.x.TAttr {
	case sml.ST_CellTypeInlineStr:
		if c.x.Is != nil && c.x.Is.T != nil {
			return *c.x.Is.T
		}
		if c.x.V != nil {
			return *c.x.V
		}
	case sml.ST_CellTypeS:
		if c.x.V == nil {
			return ""
		}
		id, err := strconv.Atoi(*c.x.V)
		if err != nil {
			return ""
		}
		s, err := c.w.SharedStrings.GetString(id)
		if err != nil {
			return ""
		}
		return s
	}
	if c.x.V == nil {
		return ""
	}
	return *c.x.V
}

func (c Cell) GetRawValue() (string, error) {
	switch c.x.TAttr {
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
	case sml.ST_CellTypeStr:
		if c.x.F != nil {
			return c.x.F.Content, nil
		}
	}
	if c.x.V == nil {
		return "", nil
	}
	return *c.x.V, nil
}

// SetHyperlink sets a hyperlink on a cell.
func (c Cell) SetHyperlink(hl common.Hyperlink) {
	if c.s.Hyperlinks == nil {
		c.s.Hyperlinks = sml.NewCT_Hyperlinks()
	}
	rel := common.Relationship(hl)

	hle := sml.NewCT_Hyperlink()
	hle.RefAttr = c.Reference()
	hle.IdAttr = unioffice.String(rel.ID())
	c.s.Hyperlinks.Hyperlink = append(c.s.Hyperlinks.Hyperlink, hle)
}

// AddHyperlink creates and sets a hyperlink on a cell.
func (c Cell) AddHyperlink(url string) {
	// store the relationships so we don't need to do a lookup here?
	for i, ws := range c.w.xws {
		if ws == c.s {
			// add a hyperlink relationship in the worksheet relationships file
			c.SetHyperlink(c.w.xwsRels[i].AddHyperlink(url))
			return
		}
	}
}

// IsNumber returns true if the cell is a number type cell.
func (c Cell) IsNumber() bool {
	switch c.x.TAttr {
	case sml.ST_CellTypeN:
		return true
	case sml.ST_CellTypeS, sml.ST_CellTypeB:
		return false
	}
	return c.x.V != nil && format.IsNumber(*c.x.V)
}

// IsEmpty returns true if the cell is empty.
func (c Cell) IsEmpty() bool {
	return c.x.TAttr == sml.ST_CellTypeUnset && c.x.V == nil && c.x.F == nil
}

// IsBool returns true if the cell is a boolean type cell.
func (c Cell) IsBool() bool {
	return c.x.TAttr == sml.ST_CellTypeB
}

// HasFormula returns true if the cell has an asoociated formula.
func (c Cell) HasFormula() bool {
	return c.x.F != nil
}

// GetFormula returns the formula for a cell.
func (c Cell) GetFormula() string {
	if c.x.F != nil {
		return c.x.F.Content
	}
	return ""
}

// GetCachedFormulaResult returns the cached formula result if it exists. If the
// cell type is not a formula cell, the result will be the cell value if it's a
// string/number/bool cell.
func (c Cell) GetCachedFormulaResult() string {
	if c.x.V != nil {
		return *c.x.V
	}
	return ""
}

func (c Cell) getRawSortValue() (string, bool) {
	if c.HasFormula() {
		v := c.GetCachedFormulaResult()
		return v, format.IsNumber(v)
	}
	v, _ := c.GetRawValue()
	return v, format.IsNumber(v)
}

// SetCachedFormulaResult sets the cached result of a formula. This is normally
// not needed but is used internally when expanding an array formula.
func (c Cell) SetCachedFormulaResult(s string) {
	c.x.V = &s
}

func b2i(v bool) int {
	if v {
		return 1
	}
	return 0
}
