// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

import "fmt"

func init() {
	RegisterFunction("INDEX", Index)
	RegisterFunctionComplex("INDIRECT", Indirect)
	RegisterFunctionComplex("OFFSET", Offset)
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

	return rowVal[col]
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
	origin := "A1"
	switch ref.Type {
	case ReferenceTypeCell:
		origin = ref.Value
	case ReferenceTypeRange:
	case ReferenceTypeNamedRange:
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
	cOff := args[1].AsNumber()
	if cOff.Type != ResultTypeNumber {
		return MakeErrorResult("OFFSET requires numeric col offset")
	}

	height := args[1].AsNumber()
	if height.Type != ResultTypeNumber {
		return MakeErrorResult("OFFSET requires numeric height")
	}
	width := args[1].AsNumber()
	if width.Type != ResultTypeNumber {
		return MakeErrorResult("OFFSET requires numeric width")
	}
	colIdx := ColumnToIndex(col)
	origRow := rowIdx + uint32(rOff.ValueNumber)
	origCol := colIdx + uint32(cOff.ValueNumber)
	endRow := origRow + uint32(height.ValueNumber)
	endCol := origCol + uint32(width.ValueNumber)

	beg := fmt.Sprintf("%s%d", IndexToColumn(origCol), origRow)
	end := fmt.Sprintf("%s%d", IndexToColumn(endCol), endRow)
	return resultFromCellRange(ctx, ev, beg, end)

}
