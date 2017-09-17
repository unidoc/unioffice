// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

func init() {
	RegisterFunction("INDEX", Index)
}

// Index implements the Excel INDEX function
func Index(args []Result) Result {
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
