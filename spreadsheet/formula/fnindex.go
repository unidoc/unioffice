// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

import (
	"fmt"
	"strings"

	"github.com/unidoc/unioffice/internal/mergesort"
	"github.com/unidoc/unioffice/internal/wildcard"
	"github.com/unidoc/unioffice/spreadsheet/reference"
)

func init() {
	RegisterFunction("CHOOSE", Choose)
	RegisterFunction("COLUMN", Column)
	RegisterFunction("COLUMNS", Columns)
	RegisterFunction("INDEX", Index)
	RegisterFunctionComplex("INDIRECT", Indirect)
	RegisterFunctionComplex("OFFSET", Offset)
	RegisterFunction("MATCH", Match)
	RegisterFunction("HLOOKUP", HLookup)
	RegisterFunction("LARGE", Large)
	RegisterFunction("LOOKUP", Lookup)
	RegisterFunction("ROW", Row)
	RegisterFunction("ROWS", Rows)
	RegisterFunction("SMALL", Small)
	RegisterFunction("VLOOKUP", VLookup)
	RegisterFunction("TRANSPOSE", Transpose)
}

// Choose implements the Excel CHOOSE function.
func Choose(args []Result) Result {
	if len(args) < 2 {
		return MakeErrorResult("CHOOSE requires two arguments")
	}
	index := args[0]
	if index.Type != ResultTypeNumber {
		return MakeErrorResult("CHOOSE requires first argument of type number")
	}
	i := int(index.ValueNumber)
	if len(args) <= i {
		return MakeErrorResult("Index should be less or equal to the number of values")
	}
	return args[i]
}

// Column implements the Excel COLUMN function.
func Column(args []Result) Result {
	if len(args) < 1 {
		return MakeErrorResult("COLUMN requires one argument")
	}
	ref := args[0].Ref
	if ref.Type != ReferenceTypeCell {
		return MakeErrorResult("COLUMN requires an argument to be of type reference")
	}
	cr, err := reference.ParseCellReference(ref.Value)
	if err != nil {
		return MakeErrorResult("Incorrect reference: " + ref.Value)
	}
	return MakeNumberResult(float64(cr.ColumnIdx + 1))
}

// Columns implements the Excel COLUMNS function.
func Columns(args []Result) Result {
	if len(args) < 1 {
		return MakeErrorResult("COLUMNS requires one argument")
	}
	arrResult := args[0]
	if arrResult.Type != ResultTypeArray && arrResult.Type != ResultTypeList {
		return MakeErrorResult("COLUMNS requires first argument of type array")
	}
	arr := arrResult.ValueArray
	if len(arr) == 0 {
		return MakeErrorResult("COLUMNS requires array to contain at least 1 row")
	}
	return MakeNumberResult(float64(len(arr[0])))
}

// Index implements the Excel INDEX function.
func Index(args []Result) Result {
	argsNum := len(args)
	if argsNum < 2 || argsNum > 3 {
		return MakeErrorResult("INDEX requires from one to three arguments")
	}
	arr := args[0]
	if arr.Type != ResultTypeArray && arr.Type != ResultTypeList {
		return MakeErrorResult("INDEX requires first argument of type array")
	}
	rowArg := args[1].AsNumber()
	if rowArg.Type != ResultTypeNumber {
		return MakeErrorResult("INDEX requires numeric row argument")
	}
	row := int(rowArg.ValueNumber) - 1
	col := -1
	if argsNum == 3 && args[2].Type != ResultTypeEmpty {
		colArg := args[2].AsNumber()
		if colArg.Type != ResultTypeNumber {
			return MakeErrorResult("INDEX requires numeric col argument")
		}
		col = int(colArg.ValueNumber) - 1
	}
	if row == -1 && col == -1 {
		return MakeErrorResult("INDEX requires row or col argument")
	}
	var rowVal []Result
	if arr.Type == ResultTypeArray {
		valueArray := arr.ValueArray
		if row < -1 || row >= len(valueArray) {
			return MakeErrorResult("INDEX has row out of range")
		}
		if row == -1 {
			if col >= len(valueArray[0]) {
				return MakeErrorResult("INDEX has col out of range")
			}
			oneColumnArray := [][]Result{}
			for _, row := range valueArray {
				v := row[col]
				if v.Type == ResultTypeEmpty {
					v = MakeNumberResult(0)
				}
				oneColumnArray = append(oneColumnArray, []Result{v})
			}
			return MakeArrayResult(oneColumnArray)
		}
		rowVal = valueArray[row]
	} else {
		valueList := arr.ValueList
		if row < -1 || row >= 1 {
			return MakeErrorResult("INDEX has row out of range")
		}
		if row == -1 {
			if col >= len(valueList) {
				return MakeErrorResult("INDEX has col out of range")
			}
			v := valueList[col]
			if v.Type == ResultTypeEmpty {
				v = MakeNumberResult(0)
			}
			return v
		}
		rowVal = valueList
	}

	if col < -1 || col > len(rowVal) {
		return MakeErrorResult("INDEX has col out of range")
	}

	if col == -1 {
		listResult := []Result{}
		for _, v := range rowVal {
			if v.Type == ResultTypeEmpty {
				listResult = append(listResult, MakeNumberResult(0))
			} else {
				listResult = append(listResult, v)
			}
		}
		return MakeArrayResult([][]Result{listResult})
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

// Match implements the MATCH function.
func Match(args []Result) Result {
	argsNum := len(args)
	if argsNum != 2 && argsNum != 3 {
		return MakeErrorResult("MATCH requires two or three arguments")
	}

	matchType := 1
	if argsNum == 3 && args[2].Type != ResultTypeEmpty {
		if args[2].Type != ResultTypeNumber {
			return MakeErrorResult("MATCH requires the third argument to be a number")
		}
		typeArg := args[2].ValueNumber
		if typeArg == -1 || typeArg == 0 {
			matchType = int(typeArg)
		}
	}

	arrResult := args[1]
	var values []Result

	switch arrResult.Type {
	case ResultTypeList:
		values = arrResult.ValueList
	case ResultTypeArray:
		arr := arrResult.ValueArray
		if len(arr[0]) != 1 {
			return MakeErrorResult("MATCH requires the second argument to be a one-dimensional range")
		}
		for _, list := range arr {
			values = append(values, list[0])
		}
	default:
		return MakeErrorResult("MATCH requires the second argument to be a one-dimensional range")
	}

	criteria := parseCriteria(args[0])

	switch matchType {
	case 0:
		for i, value := range values {
			if compareForMatch(value, criteria) {
				return MakeNumberResult(float64(i + 1))
			}
		}
	case -1:
		for i := 0; i < len(values); i++ {
			if compareForMatch(values[i], criteria) {
				return MakeNumberResult(float64(i + 1))
			}
			if criteria.isNumber && (values[i].ValueNumber < criteria.cNum) {
				if i == 0 {
					return MakeErrorResultType(ErrorTypeNA, "")
				}
				return MakeNumberResult(float64(i))
			}
		}
	case 1:
		for i := 0; i < len(values); i++ {
			if compareForMatch(values[i], criteria) {
				return MakeNumberResult(float64(i + 1))
			}
			if criteria.isNumber && (values[i].ValueNumber > criteria.cNum) {
				if i == 0 {
					return MakeErrorResultType(ErrorTypeNA, "")
				}
				return MakeNumberResult(float64(i))
			}
		}
	}
	return MakeErrorResultType(ErrorTypeNA, "")
}

func compareForMatch(value Result, criteria *criteriaParsed) bool {
	if value.Type == ResultTypeEmpty {
		return false
	}
	if criteria.isNumber {
		return value.ValueNumber == criteria.cNum
	} else {
		valueStr := strings.ToLower(value.ValueString)
		return criteria.cStr == valueStr || wildcard.Match(criteria.cStr, valueStr)
	}
}

// Offset is an implementation of the Excel OFFSET function.
func Offset(ctx Context, ev Evaluator, args []Result) Result {
	if len(args) != 3 && len(args) != 5 {
		return MakeErrorResult("OFFSET requires three or five arguments")
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

	parsedRef, parseErr := reference.ParseCellReference(origin)
	if parseErr != nil {
		return MakeErrorResult(fmt.Sprintf("parse origin error OFFSET(): %s", parseErr.Error()))
	}
	col, rowIdx, sheetName := parsedRef.Column, parsedRef.RowIdx, parsedRef.SheetName

	rOff := args[1].AsNumber()
	if rOff.Type != ResultTypeNumber {
		return MakeErrorResult("OFFSET requires numeric row offset")
	}
	cOff := args[2].AsNumber()
	if cOff.Type != ResultTypeNumber {
		return MakeErrorResult("OFFSET requires numeric col offset")
	}

	var height, width Result

	if len(args) == 3 {
		height = MakeNumberResult(1)
		width = MakeNumberResult(1)
	} else {
		height = args[3].AsNumber()
		if height.Type != ResultTypeNumber {
			return MakeErrorResult("OFFSET requires numeric height")
		}
		if height.ValueNumber == 0 {
			return MakeErrorResultType(ErrorTypeRef, "")
		}
		width = args[4].AsNumber()
		if width.Type != ResultTypeNumber {
			return MakeErrorResult("OFFSET requires numeric width")
		}
		if width.ValueNumber == 0 {
			return MakeErrorResultType(ErrorTypeRef, "")
		}
	}
	colIdx := reference.ColumnToIndex(col)
	origRow := rowIdx + uint32(rOff.ValueNumber)
	origCol := colIdx + uint32(cOff.ValueNumber)
	endRow := origRow + uint32(height.ValueNumber)
	endCol := origCol + uint32(width.ValueNumber)
	if height.ValueNumber > 0 {
		endRow--
	} else {
		endRow++
		origRow, endRow = endRow, origRow
	}
	if width.ValueNumber > 0 {
		endCol--
	} else {
		endCol++
		origCol, endCol = endCol, origCol
	}

	beg := fmt.Sprintf("%s%d", reference.IndexToColumn(origCol), origRow)
	end := fmt.Sprintf("%s%d", reference.IndexToColumn(endCol), endRow)
	if sheetName == "" {
		return resultFromCellRange(ctx, ev, beg, end)
	} else {
		return resultFromCellRange(ctx.Sheet(sheetName), ev, beg, end)
	}
}

// VLookup implements the VLOOKUP function that returns a matching value from a
// column in an array.
func VLookup(args []Result) Result {
	argsNum := len(args)
	if argsNum < 3 {
		return MakeErrorResult("VLOOKUP requires at least three arguments")
	}
	if argsNum > 4 {
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
	if argsNum == 4 && args[3].Type != ResultTypeEmpty {
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
		switch compareResults(rval, lookupValue, false, exactMatch) {
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

func compareResults(lhs, rhs Result, caseSensitive, exactMatch bool) cmpResult {
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
		ls := lhs.ValueString
		rs := rhs.ValueString
		if !caseSensitive {
			ls = strings.ToLower(ls)
			rs = strings.ToLower(rs)
		}
		if exactMatch {
			match := wildcard.Match(rs, ls)
			if match {
				return cmpResultEqual
			} else {
				return cmpResultGreater
			}
		}
		return cmpResult(strings.Compare(ls, rs))
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
			cmp := compareResults(lhs.ValueList[i], rhs.ValueList[i], caseSensitive, exactMatch)
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
				cmp := compareResults(lrow[c], rrow[c], caseSensitive, exactMatch)
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
		if compareResults(lookupValue, v, false, false) == cmpResultEqual {
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
				col = append(col, empty)
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
		switch compareResults(val, lookupValue, false, exactMatch) {
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

// Row implements the Excel ROW function.
func Row(args []Result) Result {
	if len(args) < 1 {
		return MakeErrorResult("ROW requires one argument")
	}
	ref := args[0].Ref
	if ref.Type != ReferenceTypeCell {
		return MakeErrorResult("ROW requires an argument to be of type reference")
	}
	cr, err := reference.ParseCellReference(ref.Value)
	if err != nil {
		return MakeErrorResult("Incorrect reference: " + ref.Value)
	}
	return MakeNumberResult(float64(cr.RowIdx))
}

// Rows implements the Excel ROWS function.
func Rows(args []Result) Result {
	if len(args) < 1 {
		return MakeErrorResult("ROWS requires one argument")
	}
	arrResult := args[0]
	if arrResult.Type != ResultTypeArray && arrResult.Type != ResultTypeList {
		return MakeErrorResult("ROWS requires first argument of type array")
	}
	arr := arrResult.ValueArray
	if len(arr) == 0 {
		return MakeErrorResult("ROWS requires array to contain at least 1 row")
	}
	return MakeNumberResult(float64(len(arr)))
}

// Large implements the Excel LARGE function.
func Large(args []Result) Result {
	return kth(args, true)
}

// Small implements the Excel SMALL function.
func Small(args []Result) Result {
	return kth(args, false)
}

func kth(args []Result, large bool) Result {
	var funcName string
	if large {
		funcName = "LARGE"
	} else {
		funcName = "SMALL"
	}
	if len(args) != 2 {
		return MakeErrorResult(funcName + " requires two arguments")
	}
	arrResult := args[0]
	var arr [][]Result
	switch arrResult.Type {
	case ResultTypeArray:
		arr = arrResult.ValueArray
	case ResultTypeList:
		arr = [][]Result{arrResult.ValueList}
	default:
		return MakeErrorResult(funcName + " requires first argument of type array")
	}
	if len(arr) == 0 {
		return MakeErrorResult(funcName + " requires array to contain at least 1 row")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult(funcName + " requires second argument of type number")
	}
	kfloat := args[1].ValueNumber
	if kfloat < 1 {
		return MakeErrorResultType(ErrorTypeNum, funcName+" requires second argument of type number more than 0")
	}
	k := int(kfloat)
	if float64(k) != kfloat {
		return MakeErrorResultType(ErrorTypeNum, funcName+" requires second argument of type number more than 0")
	}
	unsorted := []float64{}
	for _, row := range arr {
		for _, v := range row {
			if v.Type == ResultTypeNumber {
				unsorted = append(unsorted, v.ValueNumber)
			}
		}
	}
	if k > len(unsorted) {
		return MakeErrorResultType(ErrorTypeNum, funcName+" requires second argument of type number less or equal than the number of numbers in the array")
	}
	sorted := mergesort.MergeSort(unsorted)
	if large {
		return MakeNumberResult(sorted[len(sorted)-k])
	} else {
		return MakeNumberResult(sorted[k-1])
	}
}
