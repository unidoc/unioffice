// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"baliance.com/gooxml/spreadsheet/formula"

	"baliance.com/gooxml"
	"baliance.com/gooxml/algo"
	"baliance.com/gooxml/schema/soo/sml"
	"baliance.com/gooxml/spreadsheet/format"
)

// PivotTable is a pivot table including the cache and records.
type PivotTable struct {
	x *sml.PivotTableDefinition
	c *sml.PivotCacheDefinition
	r *sml.PivotCacheRecords
	w *Workbook
}

// X returns the inner wrapped XML type.
func (p PivotTable) X() *sml.PivotTableDefinition {
	return p.x
}

// Location is the location within a sheet that the pivot table wiil be
// displayed.
func (p PivotTable) Location() string {
	return p.x.Location.RefAttr
}

func (p PivotTable) SetLocation(ref string) {
	p.x.Location.RefAttr = ref
}

func (p PivotTable) AddPivotField() PivotField {
	if p.x.PivotFields == nil {
		p.x.PivotFields = sml.NewCT_PivotFields()
	}
	pf := sml.NewCT_PivotField()
	pf.ShowAllAttr = gooxml.Bool(false)
	p.x.PivotFields.PivotField = append(p.x.PivotFields.PivotField, pf)
	p.x.PivotFields.CountAttr = gooxml.Uint32(uint32(len(p.x.PivotFields.PivotField)))
	return PivotField{pf}
}

func (p PivotTable) AddRowField() PivotRowColField {
	if p.x.RowFields == nil {
		p.x.RowFields = sml.NewCT_RowFields()
	}
	pf := sml.NewCT_Field()
	p.x.RowFields.Field = append(p.x.RowFields.Field, pf)
	p.x.RowFields.CountAttr = gooxml.Uint32(uint32(len(p.x.RowFields.Field)))
	return PivotRowColField{pf}
}

func (p PivotTable) AddColumnFIeld() PivotRowColField {
	if p.x.ColFields == nil {
		p.x.ColFields = sml.NewCT_ColFields()
	}
	pf := sml.NewCT_Field()
	p.x.ColFields.Field = append(p.x.ColFields.Field, pf)
	p.x.ColFields.CountAttr = gooxml.Uint32(uint32(len(p.x.ColFields.Field)))
	return PivotRowColField{pf}
}

func (p PivotTable) AddDataField() PivotDataField {
	if p.x.DataFields == nil {
		p.x.DataFields = sml.NewCT_DataFields()
	}
	pf := sml.NewCT_DataField()
	p.x.DataFields.DataField = append(p.x.DataFields.DataField, pf)
	p.x.DataFields.CountAttr = gooxml.Uint32(uint32(len(p.x.DataFields.DataField)))
	return PivotDataField{pf}
}

// SetSource sets the source of data for the pivot table.
func (p PivotTable) SetSource(sheet Sheet, ref string) {
	p.c.CacheSource = sml.NewCT_CacheSource()
	p.c.CacheSource.TypeAttr = sml.ST_SourceTypeWorksheet
	p.c.CacheSource.WorksheetSource = sml.NewCT_WorksheetSource()
	p.c.CacheSource.WorksheetSource.SheetAttr = gooxml.String(sheet.Name())
	p.c.CacheSource.WorksheetSource.RefAttr = gooxml.String(ref)
}

// Name returns the pivot table name.
func (p PivotTable) Name() string {
	return p.x.NameAttr
}

// SetName sets the pivot table name.
func (p PivotTable) SetName(name string) {
	p.x.NameAttr = name
}

// Recalculate reconstrucst the pivot cache and records from pivot table data.
func (p PivotTable) Recalculate() {
	if p.c.CacheSource.WorksheetSource == nil ||
		p.c.CacheSource.WorksheetSource.SheetAttr == nil ||
		p.c.CacheSource.WorksheetSource.RefAttr == nil {
		return
	}
	from, to, err := ParseRangeReference(*p.c.CacheSource.WorksheetSource.RefAttr)
	if err != nil {
		log.Printf("error recomputing pivot table: %s", err)
		return
	}
	fc, frIdx, err := ParseCellReference(from)
	if err != nil {
		log.Printf("error recomputing pivot table: %s", err)
		return
	}

	tc, trIdx, err := ParseCellReference(to)
	if err != nil {
		log.Printf("error recomputing pivot table: %s", err)
		return
	}

	sheet := p.w.GetSheet(*p.c.CacheSource.WorksheetSource.SheetAttr)
	if !sheet.IsValid() {
		return
	}

	hasHeaderRow := true

	fcIdx := ColumnToIndex(fc)
	tcIdx := ColumnToIndex(tc)
	p.c.CacheFields = sml.NewCT_CacheFields()
	numCols := tcIdx - fcIdx + 1
	p.c.CacheFields.CountAttr = gooxml.Uint32(numCols)
	shared := []*sml.CT_SharedItems{}

	// create a cache field for each column
	for i := uint32(0); i < numCols; i++ {
		fld := sml.NewCT_CacheField()
		if hasHeaderRow {
			ref := fmt.Sprintf("%s%d", IndexToColumn(fcIdx+i), frIdx)
			fld.NameAttr, _ = sheet.Cell(ref).GetRawValue()
		} else {
			fld.NameAttr = IndexToColumn(i)
		}
		p.c.CacheFields.CacheField = append(p.c.CacheFields.CacheField, fld)
		fld.SharedItems = sml.NewCT_SharedItems()
		shared = append(shared, fld.SharedItems)
	}

	var fev = formula.NewEvaluator()
	// then determine the type of each column of data, either numbers or strings
	allNumbers := make([]bool, fcIdx+numCols)
	allIntegers := make([]bool, fcIdx+numCols)
	for c := fcIdx; c <= tcIdx; c++ {
		allNumbers[c] = true
		allIntegers[c] = true
		for r := frIdx; r <= trIdx; r++ {
			if hasHeaderRow && r == frIdx {
				continue
			}
			ref := fmt.Sprintf("%s%d", IndexToColumn(c), r)
			cell := sheet.Cell(ref)
			value, _ := cell.GetRawValue()
			if cell.HasFormula() {
				res := fev.Eval(sheet.FormulaContext(), value)
				value = res.Value()
			}
			allNumbers[c] = allNumbers[c] && format.IsNumber(value)
			allIntegers[c] = allIntegers[c] && allNumbers[c]
			if strings.IndexByte(value, '.') != -1 {
				allIntegers[c] = false
			}
		}
	}

	type index struct {
		strings map[string]int
		numbers map[float64]int
	}
	indices := make([]index, fcIdx+numCols)
	for i := uint32(0); i < numCols; i++ {
		if allNumbers[i+fcIdx] {
			indices[i].numbers = make(map[float64]int)
		} else {
			indices[i].strings = make(map[string]int)
		}
	}

	// construct the sorted/uniq'd column data
	for c := fcIdx; c <= tcIdx; c++ {
		for r := frIdx; r <= trIdx; r++ {
			if hasHeaderRow && r == frIdx {
				continue
			}
			ref := fmt.Sprintf("%s%d", IndexToColumn(c), r)
			cell := sheet.Cell(ref)
			value, _ := cell.GetRawValue()
			if cell.HasFormula() {
				res := fev.Eval(sheet.FormulaContext(), value)
				value = res.Value()
			}

			if allNumbers[c] {
				num := sml.NewCT_Number()
				num.VAttr, _ = strconv.ParseFloat(value, 64)
				shared[c].N = append(shared[c].N, num)
			} else {
				str := sml.NewCT_String()
				str.VAttr = cell.GetFormattedValue()
				shared[c].S = append(shared[c].S, str)
			}
		}
		if allNumbers[c] {
			sort.Slice(shared[c].N, func(i, j int) bool {
				return shared[c].N[i].VAttr < shared[c].N[j].VAttr
			})
			n := algo.Unique(shared[c].N, func(i, j int) bool {
				return shared[c].N[i].VAttr == shared[c].N[j].VAttr
			})
			shared[c].N = shared[c].N[0:n]
			shared[c].CountAttr = gooxml.Uint32(uint32(n))

			for i := 0; i < n; i++ {
				v := shared[c].N[i].VAttr
				indices[c].numbers[v] = i
			}
		} else {
			sort.Slice(shared[c].S, func(i, j int) bool {
				return shared[c].S[i].VAttr < shared[c].S[j].VAttr
			})
			n := algo.Unique(shared[c].S, func(i, j int) bool {
				return shared[c].S[i].VAttr == shared[c].S[j].VAttr
			})
			shared[c].S = shared[c].S[0:n]
			shared[c].CountAttr = gooxml.Uint32(uint32(n))

			for i := 0; i < n; i++ {
				v := shared[c].S[i].VAttr
				indices[c].strings[v] = i
			}
		}
	}

	p.r.CountAttr = gooxml.Uint32(trIdx - frIdx)
	p.c.RecordCountAttr = gooxml.Uint32(trIdx - frIdx)

	for r := frIdx; r <= trIdx; r++ {
		if hasHeaderRow && r == frIdx {
			continue
		}
		rec := sml.NewCT_Record()
		p.r.R = append(p.r.R, rec)

		for c := fcIdx; c <= tcIdx; c++ {
			ref := fmt.Sprintf("%s%d", IndexToColumn(c), r)
			if allNumbers[c] {
				shared[c].ContainsNumberAttr = gooxml.Bool(true)
				shared[c].ContainsSemiMixedTypesAttr = gooxml.Bool(false)
				shared[c].ContainsStringAttr = gooxml.Bool(false)
				shared[c].ContainsIntegerAttr = gooxml.Bool(allIntegers[c])
				v, _ := sheet.Cell(ref).GetValueAsNumber()
				idx := indices[c].numbers[v]
				x := sml.NewCT_Index()
				x.VAttr = uint32(idx)
				rec.X = append(rec.X, x)
			} else {
				shared[c].ContainsStringAttr = gooxml.Bool(true)
				v, _ := sheet.Cell(ref).GetRawValue()
				idx := indices[c].strings[v]
				x := sml.NewCT_Index()
				x.VAttr = uint32(idx)
				rec.X = append(rec.X, x)
			}
		}
	}
}
