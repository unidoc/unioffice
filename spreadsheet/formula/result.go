// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

import (
	"math"
	"strconv"
)

// ResultType is the type of the result
//go:generate stringer -type=ResultType
type ResultType byte

// ResultType constants.
const (
	ResultTypeUnknown ResultType = iota
	ResultTypeNumber
	ResultTypeString
	ResultTypeList
	ResultTypeArray
	ResultTypeError
	ResultTypeEmpty
)

// Result is the result of a formula or cell evaluation .
type Result struct {
	ValueNumber  float64
	ValueString  string
	ValueList    []Result
	ValueArray   [][]Result
	IsBoolean    bool
	ErrorMessage string
	Type         ResultType

	Ref Reference
}

func (r Result) String() string {
	return r.Value()
}

// Value returns a string version of the result.
func (r Result) Value() string {
	switch r.Type {
	case ResultTypeNumber:
		n := strconv.FormatFloat(r.ValueNumber, 'f', -1, 64)
		// HACK: currently only used for testing, need to write a better general
		// number format function
		if len(n) > 12 {
			end := 12
			for i := end; i > 0 && n[i] == '0'; i-- {
				end--
			}
			n = n[0 : end+1]
		}
		return n
	case ResultTypeError:
		return r.ValueString
	case ResultTypeString:
		return r.ValueString
	case ResultTypeList:
		if len(r.ValueList) == 0 {
			return ""
		}
		return r.ValueList[0].Value()
	case ResultTypeArray:
		if len(r.ValueArray) == 0 || len(r.ValueArray[0]) == 0 {
			return ""
		}
		return r.ValueArray[0][0].Value()
	case ResultTypeEmpty:
		return ""
	default:
		return "unhandled result value"
	}
}

// AsNumber attempts to intepret a string cell value as a number. Upon success,
// it returns a new number result, upon  failure it returns the original result.
// This is used as functions return strings that can then act like number (e.g.
// LEFT(1.2345,3) + LEFT(1.2345,3) = 2.4)
func (r Result) AsNumber() Result {
	if r.Type == ResultTypeString {
		f, err := strconv.ParseFloat(r.ValueString, 64)
		if err == nil {
			return MakeNumberResult(f)
		}
	}
	if r.Type == ResultTypeEmpty {
		return MakeNumberResult(0)
	}
	return r
}

// ListValues converts an array to a list or returns a lists values. This is used
// for functions that can accept an array, but don't care about ordering to
// reuse the list function logic.
func (r Result) ListValues() []Result {
	if r.Type == ResultTypeArray {
		res := []Result{}
		for _, row := range r.ValueArray {
			for _, col := range row {
				res = append(res, col)
			}
		}
		return res
	}
	if r.Type == ResultTypeList {
		return r.ValueList
	}
	return nil
}
func (r Result) AsString() Result {
	switch r.Type {
	case ResultTypeNumber:
		return MakeStringResult(r.Value())
	default:
		return r
	}
}

// MakeNumberResult constructs a number result.
func MakeNumberResult(v float64) Result {
	// Excel doesn't use negative zero
	if v == math.Copysign(0, -1) {
		v = 0
	}
	return Result{Type: ResultTypeNumber, ValueNumber: v}
}

// MakeBoolResult constructs a boolean result (internally a number).
func MakeBoolResult(b bool) Result {
	if b {
		return Result{Type: ResultTypeNumber, ValueNumber: 1, IsBoolean: true}
	}
	return Result{Type: ResultTypeNumber, ValueNumber: 0, IsBoolean: true}
}

// MakeErrorResult constructs a #VALUE! error with a given extra error message.
// The error message is for debugging formula evaluation only and is not stored
// in the sheet.
func MakeErrorResult(msg string) Result {
	return MakeErrorResultType(ErrorTypeValue, msg)
}

// ErrorType is a formula evaluation error type.
type ErrorType byte

// ErrorType constants.
const (
	ErrorTypeValue ErrorType = iota
	ErrorTypeNull
	ErrorTypeRef
	ErrorTypeName
	ErrorTypeNum
	ErrorTypeSpill
	ErrorTypeNA
	ErrorTypeDivideByZero
)

// MakeErrorResultType makes an error result of a given type with a specified
// debug message
func MakeErrorResultType(t ErrorType, msg string) Result {
	switch t {
	case ErrorTypeNull:
		return Result{Type: ResultTypeError, ValueString: "#NULL!", ErrorMessage: msg}
	case ErrorTypeValue:
		return Result{Type: ResultTypeError, ValueString: "#VALUE!", ErrorMessage: msg}
	case ErrorTypeRef:
		return Result{Type: ResultTypeError, ValueString: "#REF!", ErrorMessage: msg}
	case ErrorTypeName:
		return Result{Type: ResultTypeError, ValueString: "#NAME?", ErrorMessage: msg}
	case ErrorTypeNum:
		return Result{Type: ResultTypeError, ValueString: "#NUM!", ErrorMessage: msg}
	case ErrorTypeSpill:
		return Result{Type: ResultTypeError, ValueString: "#SPILL!", ErrorMessage: msg}
	case ErrorTypeNA:
		return Result{Type: ResultTypeError, ValueString: "#N/A", ErrorMessage: msg}
	case ErrorTypeDivideByZero:
		return Result{Type: ResultTypeError, ValueString: "#DIV/0!", ErrorMessage: msg}
	default:
		return Result{Type: ResultTypeError, ValueString: "#VALUE!", ErrorMessage: msg}
	}
}

// MakeStringResult constructs a string result.
func MakeStringResult(s string) Result {
	return Result{Type: ResultTypeString, ValueString: s}
}

// MakeEmptyResult is ued when parsing an empty argument.
func MakeEmptyResult() Result {
	return Result{Type: ResultTypeEmpty}
}

// MakeArrayResult constructs an array result (matrix).
func MakeArrayResult(arr [][]Result) Result {
	return Result{Type: ResultTypeArray, ValueArray: arr}
}

// MakeListResult constructs a list result.
func MakeListResult(list []Result) Result {
	return Result{Type: ResultTypeList, ValueList: list}
}
