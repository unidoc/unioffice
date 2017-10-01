// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

import (
	"fmt"
	"strings"

	"baliance.com/gooxml/spreadsheet/reference"
)

func init() {
	RegisterFunction("INDEX", Index)
	RegisterFunctionComplex("INDIRECT", Indirect)
	RegisterFunctionComplex("OFFSET", Offset)
	RegisterFunction("HLOOKUP", HLookup)
	RegisterFunction("LOOKUP", Lookup)
	RegisterFunction("VLOOKUP", VLookup)
	RegisterFunction("TRANSPOSE", Transpose)
}

// Index implements the Excel INDEX function
func Index(args []Result) Result {
	if len(args) < 3 {
		return MakeErrorResult("INDEX requires three arguments")
	}
	arr := args[0]
	if arr.Type != ResultTypeArray && arr.Type != ResultTypeList {
		return MakeErrorResult("INDEX requires first argument of type array")
	}
	rowArg := args[1].AsNumber()
	if rowArg.Type != ResultTypeNumber {
		return MakeErrorResult("INDEX requires numeric row argument")
	}
	colArg := args[2].AsNumber()
	if colArg.Type != ResultTypeNumber {
		return MakeErrorResult("INDEX requires numeric col argument")
	}
	row := int(rowArg.ValueNumber) - 1
	col := int(colArg.ValueNumber) - 1
	var rowVal []Result
	if arr.Type == ResultTypeArray {
		if row < 0 || row >= len(arr.ValueArray) {
			return MakeErrorResult("INDEX has row out of range")
		}
		rowVal = arr.ValueArray[row]
	} else {
		if row < 0 || row >= 1 {
			return MakeErrorResult("INDEX has row out of range")
		}
		rowVal = arr.ValueList
	}

	if col < 0 || col > len(rowVal) {
		return MakeErrorResult("INDEX has col out of range")
	}

	rv := rowVal[col]
	// empty cell returns a zero
	if rv.Type == ResultTypeEmpty {
		return MakeNumberResult(0)
	}
	return rv
}

// Indirect is an implementation of the Excel INDIRECT function that returns the
// contents of a cell.
func Indirect(ctx Context, ev Evaluator, args []Result) Result {
	if len(args) != 1 && len(args) != 2 {
		return MakeErrorResult("INDIRECT requires one or two arguments")
	}
	sarg := args[0].AsString()
	if sarg.Type != ResultTypeString {
		return MakeErrorResult("INDIRECT requires first argument to be of type string")
	}
	return ctx.Cell(sarg.ValueString, ev)
}

func Offset(ctx Context, ev Evaluator, args []Result) Result {
	if len(args) != 5 {
		return MakeErrorResult("OFFSET requires one or two arguments")
	}
	ref := args[0].Ref
	// resolve a named range
	for ref.Type == ReferenceTypeNamedRange {
		ref = ctx.NamedRange(ref.Value)
	}

	origin := ""
	switch ref.Type {
	case ReferenceTypeCell:
		origin = ref.Value
	case ReferenceTypeRange:
		sp := strings.Split(ref.Value, ":")
		if len(sp) == 2 {
			origin = sp[0]
		}
	default:
		return MakeErrorResult(fmt.Sprintf("Invalid range in OFFSET(): %s", ref.Type))
	}

	col, rowIdx, err := ParseCellReference(origin)
	if err != nil {
		return MakeErrorResult(fmt.Sprintf("parse origin error OFFSET(): %s", err))
	}

	rOff := args[1].AsNumber()
	if rOff.Type != ResultTypeNumber {
		return MakeErrorResult("OFFSET requires numeric row offset")
	}
	cOff := args[2].AsNumber()
	if cOff.Type != ResultTypeNumber {
		return MakeErrorResult("OFFSET requires numeric col offset")
	}

	height := args[3].AsNumber()
	if height.Type != ResultTypeNumber {
		return MakeErrorResult("OFFSET requires numeric height")
	}
	width := args[4].AsNumber()
	if width.Type != ResultTypeNumber {
		return MakeErrorResult("OFFSET requires numeric width")
	}
	colIdx := reference.ColumnToIndex(col)
	origRow := rowIdx + uint32(rOff.ValueNumber)
	origCol := colIdx + uint32(cOff.ValueNumber)
	endRow := origRow + uint32(height.ValueNumber) - 1
	endCol := origCol + uint32(width.ValueNumber) - 1

	beg := fmt.Sprintf("%s%d", reference.IndexToColumn(origCol), origRow)
	end := fmt.Sprintf("%s%d", reference.IndexToColumn(endCol), endRow)
	return resultFromCellRange(ctx, ev, beg, end)
}

// VLookup implements the VLOOKUP function that returns a matching value from a
// column in an array.
func VLookup(args []Result) Result {
	if len(args) < 3 {
		return MakeErrorResult("VLOOKUP requires at least three arguments")
	}
	if len(args) > 4 {
		return MakeErrorResult("VLOOKUP requires at most four arguments")
	}
	lookupValue := args[0]
	arr := args[1]
	if arr.Type != ResultTypeArray {
		return MakeErrorResult("VLOOKUP requires second argument of type array")
	}
	colArg := args[2].AsNumber()
	if colArg.Type != ResultTypeNumber {
		return MakeErrorResult("VLOOKUP requires numeric col argument")
	}
	exactMatch := false
	if len(args) == 4 {
		em := args[3].AsNumber()
		if em.Type != ResultTypeNumber {
			return MakeErrorResult("VLOOKUP requires numeric match argument")
		}
		if em.ValueNumber == 0 {
			exactMatch = true
		}
	}

	// 1 indexed in input, so change to zero indexed
	col := int(colArg.ValueNumber) - 1

	matchIdx := -1
	wasExact := false
lfor:
	for i, row := range arr.ValueArray {
		if len(row) == 0 {
			continue
		}
		rval := row[0]
		switch compareResults(rval, lookupValue, false) {
		case cmpResultLess:
			// less than
			matchIdx = i
		case cmpResultEqual:
			// exact match
			matchIdx = i
			wasExact = true
			break lfor
		}
	}
	if matchIdx == -1 {
		return MakeErrorResultType(ErrorTypeNA, "VLOOKUP no result found")
	}
	row := arr.ValueArray[matchIdx]
	if col < 0 || col >= len(row) {
		return MakeErrorResult("VLOOKUP has invalid column index")
	}
	if wasExact || !exactMatch {
		return row[col]
	}
	return MakeErrorResultType(ErrorTypeNA, "VLOOKUP no result found")
}

type cmpResult int8

const (
	cmpResultEqual   cmpResult = 0
	cmpResultLess    cmpResult = -1
	cmpResultGreater cmpResult = 1
	cmpResultInvalid cmpResult = 2
)

func compareResults(lhs, rhs Result, caseSensitive bool) cmpResult {
	lhs = lhs.AsNumber()
	rhs = rhs.AsNumber()
	// differing types
	if lhs.Type != rhs.Type {
		return cmpResultInvalid
	}

	// both numbers
	if lhs.Type == ResultTypeNumber {
		if lhs.ValueNumber == rhs.ValueNumber {
			return cmpResultEqual
		}
		if lhs.ValueNumber < rhs.ValueNumber {
			return cmpResultLess
		}
		return cmpResultGreater
	}

	// both strings
	if lhs.Type == ResultTypeString {
		if !caseSensitive {
			return cmpResult(strings.Compare(strings.ToLower(lhs.ValueString),
				strings.ToLower(rhs.ValueString)))
		}
		return cmpResult(strings.Compare(lhs.ValueString, rhs.ValueString))
	}

	// empty cells are equal
	if lhs.Type == ResultTypeEmpty {
		return cmpResultEqual
	}

	// compare lists recursively
	if lhs.Type == ResultTypeList {
		if len(lhs.ValueList) < len(rhs.ValueList) {
			return cmpResultLess
		}
		if len(lhs.ValueList) > len(rhs.ValueList) {
			return cmpResultGreater
		}
		for i := range lhs.ValueList {
			cmp := compareResults(lhs.ValueList[i], rhs.ValueList[i], caseSensitive)
			if cmp != cmpResultEqual {
				return cmp
			}
		}
		return cmpResultEqual
	}

	// compare arrays recursively
	if lhs.Type == ResultTypeList {
		if len(lhs.ValueArray) < len(rhs.ValueArray) {
			return cmpResultLess
		}
		if len(lhs.ValueArray) > len(rhs.ValueArray) {
			return cmpResultGreater
		}
		for i := range lhs.ValueArray {
			lrow := lhs.ValueArray[i]
			rrow := lhs.ValueArray[i]
			if len(lrow) < len(rrow) {
				return cmpResultLess
			}
			if len(lrow) > len(rrow) {
				return cmpResultGreater
			}
			for c := range lrow {
				cmp := compareResults(lrow[c], rrow[c], caseSensitive)
				if cmp != cmpResultEqual {
					return cmp
				}
			}
		}
		return cmpResultEqual
	}

	return cmpResultInvalid
}

// Lookup implements the LOOKUP function that returns a matching value from a
// column, or from the same index in a second column.
func Lookup(args []Result) Result {
	if len(args) < 2 {
		return MakeErrorResult("LOOKUP requires at least two arguments")
	}
	if len(args) > 3 {
		return MakeErrorResult("LOOKUP requires at most three arguments")
	}
	lookupValue := args[0]
	arr := args[1]
	if arr.Type != ResultTypeArray && arr.Type != ResultTypeList {
		return MakeErrorResult("VLOOKUP requires second argument of type array")
	}
	col := extractCol(arr)

	idx := -1
	for i, v := range col {
		if compareResults(lookupValue, v, false) == cmpResultEqual {
			idx = i
		}
	}
	if idx == -1 {
		return MakeErrorResultType(ErrorTypeNA, "LOOKUP no result found")
	}
	resultCol := col
	if len(args) == 3 {
		resultCol = extractCol(args[2])
	}
	if idx < 0 || idx >= len(resultCol) {
		return MakeErrorResultType(ErrorTypeNA, "LOOKUP no result found")
	}
	return resultCol[idx]
}

func extractCol(arr Result) []Result {
	col := arr.ValueList
	if arr.Type == ResultTypeArray {
		col = nil
		for _, r := range arr.ValueArray {
			if len(r) > 0 {
				col = append(col, r[0])
			} else {
				col = append(col, MakeEmptyResult())
			}
		}
	}
	return col
}

// HLookup implements the HLOOKUP function that returns a matching value from a
// row in an array.
func HLookup(args []Result) Result {
	if len(args) < 3 {
		return MakeErrorResult("HLOOKUP requires at least three arguments")
	}
	if len(args) > 4 {
		return MakeErrorResult("HLOOKUP requires at most four arguments")
	}
	lookupValue := args[0]
	arr := args[1]
	if arr.Type != ResultTypeArray {
		return MakeErrorResult("HLOOKUP requires second argument of type array")
	}
	rowArg := args[2].AsNumber()
	if rowArg.Type != ResultTypeNumber {
		return MakeErrorResult("HLOOKUP requires numeric row argument")
	}
	exactMatch := false
	if len(args) == 4 {
		em := args[3].AsNumber()
		if em.Type != ResultTypeNumber {
			return MakeErrorResult("HLOOKUP requires numeric match argument")
		}
		if em.ValueNumber == 0 {
			exactMatch = true
		}
	}

	matchIdx := -1
	wasExact := false
	if len(arr.ValueArray) == 0 {
		return MakeErrorResult("HLOOKUP requires non-empty array")
	}
	row := arr.ValueArray[0]
lfor:
	for i, val := range row {
		switch compareResults(val, lookupValue, false) {
		case cmpResultLess:
			// less than
			matchIdx = i
		case cmpResultEqual:
			// exact match
			matchIdx = i
			wasExact = true
			break lfor
		}
	}
	if matchIdx == -1 {
		return MakeErrorResultType(ErrorTypeNA, "HLOOKUP no result found")
	}

	// 1 indexed in input, so change to zero indexed
	rowIdx := int(rowArg.ValueNumber) - 1
	if rowIdx < 0 || rowIdx > len(arr.ValueArray) {
		return MakeErrorResult("HLOOKUP had invalid index")
	}

	row = arr.ValueArray[rowIdx]
	if matchIdx < 0 || matchIdx >= len(row) {
		return MakeErrorResult("VLOOKUP has invalid column index")
	}
	if wasExact || !exactMatch {
		return row[matchIdx]
	}

	return MakeErrorResultType(ErrorTypeNA, "VLOOKUP no result found")
}

// Transpose implements the TRANSPOSE function that transposes a cell range.
func Transpose(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("TRANSPOSE requires a single argument")
	}
	if args[0].Type != ResultTypeArray && args[0].Type != ResultTypeList {
		return MakeErrorResult("TRANSPOSE requires a range argument")
	}

	arg := args[0]
	// list to array of rows with one column
	if arg.Type == ResultTypeList {
		res := [][]Result{}
		for _, v := range arg.ValueList {
			res = append(res, []Result{v})
		}
		return MakeArrayResult(res)
	}

	// tranpose the array
	res := make([][]Result, len(arg.ValueArray[0]))
	for _, row := range arg.ValueArray {
		for j, v := range row {
			res[j] = append(res[j], v)
		}
	}
	return MakeArrayResult(res)
}
