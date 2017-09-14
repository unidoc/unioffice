// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

import "fmt"

// ResultType is the type of the result
//go:generate stringer -type=ResultType
type ResultType byte

// ResultType constants.
const (
	ResultTypeUnknown ResultType = iota
	ResultTypeNumber
	ResultTypeString
	ResultTypeList
	ResultTypeError
)

// Result is the result of a formula or cell evaluation .
type Result struct {
	ValueNumber  float64
	ValueString  string
	ValueList    []Result
	ErrorMessage string
	Type         ResultType
}

// Value returns a string version of the formula.
func (r Result) Value() string {
	switch r.Type {
	case ResultTypeNumber:
		return fmt.Sprintf("%g", r.ValueNumber)
	case ResultTypeError:
		return r.ValueString
	case ResultTypeString:
		return r.ValueString
	default:
		return "unhandled result value"
	}
}

// MakeNumberResult constructs a number result.
func MakeNumberResult(v float64) Result {
	return Result{Type: ResultTypeNumber, ValueNumber: v}
}

// MakeBoolResult constructs a boolean result (internally a number).
func MakeBoolResult(b bool) Result {
	if b {
		return Result{Type: ResultTypeNumber, ValueNumber: 1}
	}
	return Result{Type: ResultTypeNumber, ValueNumber: 0}
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
