// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

import (
	"fmt"
	"math"

	"github.com/unidoc/unioffice/spreadsheet/update"
)

// BinOpType is the binary operation operator type
//go:generate stringer -type=BinOpType
type BinOpType byte

// Operator type constants
const (
	BinOpTypeUnknown BinOpType = iota
	BinOpTypePlus
	BinOpTypeMinus
	BinOpTypeMult
	BinOpTypeDiv
	BinOpTypeExp
	BinOpTypeLT
	BinOpTypeGT
	BinOpTypeEQ
	BinOpTypeLEQ
	BinOpTypeGEQ
	BinOpTypeNE
	BinOpTypeConcat // '&' in Excel
)

// BinaryExpr is a binary expression.
type BinaryExpr struct {
	lhs, rhs Expression
	op       BinOpType
}

// NewBinaryExpr constructs a new binary expression with a given operator.
func NewBinaryExpr(lhs Expression, op BinOpType, rhs Expression) Expression {
	return BinaryExpr{lhs, rhs, op}
}

func equalsToEmpty(res Result) bool {
	if res.Type == ResultTypeString {
		return res.ValueString == ""
	}
	return res.ValueNumber == 0
}

// Eval evaluates the binary expression using the context given.
func (b BinaryExpr) Eval(ctx Context, ev Evaluator) Result {
	lhs := b.lhs.Eval(ctx, ev)
	if lhs.Type == ResultTypeError {
		return lhs
	}
	rhs := b.rhs.Eval(ctx, ev)
	if rhs.Type == ResultTypeError {
		return rhs
	}

	// peel off array/list ops first
	if lhs.Type == rhs.Type {
		if lhs.Type == ResultTypeArray {
			if !sameDim(lhs.ValueArray, rhs.ValueArray) {
				return MakeErrorResult("lhs/rhs should have same dimensions")
			}
			return arrayOp(b.op, lhs.ValueArray, rhs.ValueArray)
		} else if lhs.Type == ResultTypeList {
			if len(lhs.ValueList) != len(rhs.ValueList) {
				return MakeErrorResult("lhs/rhs should have same dimensions")
			}
			return listOp(b.op, lhs.ValueList, rhs.ValueList)
		}

	} else if lhs.Type == ResultTypeArray && (rhs.Type == ResultTypeNumber || rhs.Type == ResultTypeString) {
		return arrayValueOp(b.op, lhs.ValueArray, rhs)
	} else if lhs.Type == ResultTypeList && (rhs.Type == ResultTypeNumber || rhs.Type == ResultTypeString) {
		return listValueOp(b.op, lhs.ValueList, rhs)
	}

	// TODO: check for and add support for binary operators on boolean values
	switch b.op {
	case BinOpTypePlus:
		if lhs.Type == rhs.Type {
			if lhs.Type == ResultTypeNumber {
				return MakeNumberResult(lhs.ValueNumber + rhs.ValueNumber)
			}
		}
	case BinOpTypeMinus:
		if lhs.Type == rhs.Type {
			if lhs.Type == ResultTypeNumber {
				return MakeNumberResult(lhs.ValueNumber - rhs.ValueNumber)
			}
		}
	case BinOpTypeMult:
		if lhs.Type == rhs.Type {
			if lhs.Type == ResultTypeNumber {
				return MakeNumberResult(lhs.ValueNumber * rhs.ValueNumber)
			}
		}
	case BinOpTypeDiv:
		if lhs.Type == rhs.Type {
			if lhs.Type == ResultTypeNumber {
				if rhs.ValueNumber == 0 {
					return MakeErrorResultType(ErrorTypeDivideByZero, "divide by zero")
				}
				return MakeNumberResult(lhs.ValueNumber / rhs.ValueNumber)
			}
		}
	case BinOpTypeExp:
		if lhs.Type == rhs.Type {
			if lhs.Type == ResultTypeNumber {
				return MakeNumberResult(math.Pow(lhs.ValueNumber, rhs.ValueNumber))
			}
		}
	case BinOpTypeLT:
		if lhs.Type == rhs.Type {
			if lhs.Type == ResultTypeNumber {
				return MakeBoolResult(lhs.ValueNumber < rhs.ValueNumber)
			}
			if lhs.Type == ResultTypeString {
				return MakeBoolResult(lhs.ValueString < rhs.ValueString)
			}
			if lhs.Type == ResultTypeEmpty { // empty always equals to empty but always less than string or number
				return MakeBoolResult(false)
			}
		} else if lhs.Type == ResultTypeString && rhs.Type == ResultTypeNumber { // when comparing string to number, string is always greater than number
			return MakeBoolResult(false)
		} else if lhs.Type == ResultTypeNumber && rhs.Type == ResultTypeString {
			return MakeBoolResult(true)
		} else if lhs.Type == ResultTypeEmpty && (rhs.Type == ResultTypeNumber || rhs.Type == ResultTypeString) {
			return MakeBoolResult(true)
		} else if (lhs.Type == ResultTypeNumber || lhs.Type == ResultTypeString) && rhs.Type == ResultTypeEmpty {
			return MakeBoolResult(false)
		}
	case BinOpTypeGT:
		if lhs.Type == rhs.Type {
			if lhs.Type == ResultTypeNumber {
				return MakeBoolResult(lhs.ValueNumber > rhs.ValueNumber)
			}
			if lhs.Type == ResultTypeString {
				return MakeBoolResult(lhs.ValueString > rhs.ValueString)
			}
			if lhs.Type == ResultTypeEmpty {
				return MakeBoolResult(false)
			}
		} else if lhs.Type == ResultTypeString && rhs.Type == ResultTypeNumber {
			return MakeBoolResult(true)
		} else if lhs.Type == ResultTypeNumber && rhs.Type == ResultTypeString {
			return MakeBoolResult(false)
		} else if lhs.Type == ResultTypeEmpty && (rhs.Type == ResultTypeNumber || rhs.Type == ResultTypeString) {
			return MakeBoolResult(false)
		} else if (lhs.Type == ResultTypeNumber || lhs.Type == ResultTypeString) && rhs.Type == ResultTypeEmpty {
			return MakeBoolResult(true)
		}
	case BinOpTypeEQ:
		if lhs.Type == rhs.Type {
			if lhs.Type == ResultTypeNumber {
				// TODO: see what Excel does regarding floating point comparison
				return MakeBoolResult(lhs.ValueNumber == rhs.ValueNumber)
			}
			if lhs.Type == ResultTypeString {
				return MakeBoolResult(lhs.ValueString == rhs.ValueString)
			}
			if lhs.Type == ResultTypeEmpty {
				return MakeBoolResult(true)
			}
		} else if (lhs.Type == ResultTypeString && rhs.Type == ResultTypeNumber) || (lhs.Type == ResultTypeNumber && rhs.Type == ResultTypeString) { // string never equals to number
			return MakeBoolResult(false)
		} else if lhs.Type == ResultTypeEmpty && (rhs.Type == ResultTypeNumber || rhs.Type == ResultTypeString) {
			return MakeBoolResult(equalsToEmpty(rhs))
		} else if (lhs.Type == ResultTypeNumber || lhs.Type == ResultTypeString) && rhs.Type == ResultTypeEmpty {
			return MakeBoolResult(equalsToEmpty(lhs))
		}
	case BinOpTypeNE:
		if lhs.Type == rhs.Type {
			if lhs.Type == ResultTypeNumber {
				return MakeBoolResult(lhs.ValueNumber != rhs.ValueNumber)
			}
			if lhs.Type == ResultTypeString {
				return MakeBoolResult(lhs.ValueString != rhs.ValueString)
			}
			if lhs.Type == ResultTypeEmpty {
				return MakeBoolResult(false)
			}
		} else if (lhs.Type == ResultTypeString && rhs.Type == ResultTypeNumber) || (lhs.Type == ResultTypeNumber && rhs.Type == ResultTypeString) {
			return MakeBoolResult(true)
		} else if lhs.Type == ResultTypeEmpty && (rhs.Type == ResultTypeNumber || rhs.Type == ResultTypeString) {
			return MakeBoolResult(!equalsToEmpty(rhs))
		} else if (lhs.Type == ResultTypeNumber || lhs.Type == ResultTypeString) && rhs.Type == ResultTypeEmpty {
			return MakeBoolResult(!equalsToEmpty(lhs))
		}
	case BinOpTypeLEQ:
		if lhs.Type == rhs.Type {
			if lhs.Type == ResultTypeNumber {
				return MakeBoolResult(lhs.ValueNumber <= rhs.ValueNumber)
			}
			if lhs.Type == ResultTypeString {
				return MakeBoolResult(lhs.ValueString <= rhs.ValueString)
			}
			if lhs.Type == ResultTypeEmpty {
				return MakeBoolResult(true)
			}
		} else if lhs.Type == ResultTypeString && rhs.Type == ResultTypeNumber {
			return MakeBoolResult(false)
		} else if lhs.Type == ResultTypeNumber && rhs.Type == ResultTypeString {
			return MakeBoolResult(true)
		} else if lhs.Type == ResultTypeEmpty && (rhs.Type == ResultTypeNumber || rhs.Type == ResultTypeString) {
			return MakeBoolResult(equalsToEmpty(rhs))
		} else if (lhs.Type == ResultTypeNumber || lhs.Type == ResultTypeString) && rhs.Type == ResultTypeEmpty {
			return MakeBoolResult(equalsToEmpty(lhs))
		}
	case BinOpTypeGEQ:
		if lhs.Type == rhs.Type {
			if lhs.Type == ResultTypeNumber {
				return MakeBoolResult(lhs.ValueNumber >= rhs.ValueNumber)
			}
			if lhs.Type == ResultTypeString {
				return MakeBoolResult(lhs.ValueString >= rhs.ValueString)
			}
			if lhs.Type == ResultTypeEmpty {
				return MakeBoolResult(true)
			}
		} else if lhs.Type == ResultTypeString && rhs.Type == ResultTypeNumber {
			return MakeBoolResult(true)
		} else if lhs.Type == ResultTypeNumber && rhs.Type == ResultTypeString {
			return MakeBoolResult(false)
		} else if lhs.Type == ResultTypeEmpty && (rhs.Type == ResultTypeNumber || rhs.Type == ResultTypeString) {
			return MakeBoolResult(equalsToEmpty(rhs))
		} else if (lhs.Type == ResultTypeNumber || lhs.Type == ResultTypeString) && rhs.Type == ResultTypeEmpty {
			return MakeBoolResult(equalsToEmpty(lhs))
		}
	case BinOpTypeConcat:
		return MakeStringResult(lhs.Value() + rhs.Value())
	}

	return MakeErrorResult("unsupported binary op")
}

// Reference returns an invalid reference for BinaryExpr.
func (b BinaryExpr) Reference(ctx Context, ev Evaluator) Reference {
	return ReferenceInvalid
}

// sameDim returns true if the arrays have the same dimensions.
func sameDim(lhs, rhs [][]Result) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for i := range lhs {
		if len(lhs[i]) != len(rhs[i]) {
			return false
		}
	}
	return true
}

func arrayOp(op BinOpType, lhs, rhs [][]Result) Result {
	// we can assume the arrays are the same size here
	res := [][]Result{}
	for i := range lhs {
		lst := listOp(op, lhs[i], rhs[i])
		if lst.Type == ResultTypeError {
			return lst
		}
		res = append(res, lst.ValueList)
	}
	return MakeArrayResult(res)
}

func listOp(op BinOpType, lhs, rhs []Result) Result {
	res := []Result{}
	// we can assume the arrays are the same size here
	for i := range lhs {
		l := lhs[i].AsNumber()
		r := rhs[i].AsNumber()
		if l.Type != ResultTypeNumber || r.Type != ResultTypeNumber {
			return MakeErrorResult("non-nunmeric value in binary operation")
		}
		switch op {
		case BinOpTypePlus:
			res = append(res, MakeNumberResult(l.ValueNumber+r.ValueNumber))
		case BinOpTypeMinus:
			res = append(res, MakeNumberResult(l.ValueNumber-r.ValueNumber))
		case BinOpTypeMult:
			res = append(res, MakeNumberResult(l.ValueNumber*r.ValueNumber))
		case BinOpTypeDiv:
			if r.ValueNumber == 0 {
				return MakeErrorResultType(ErrorTypeDivideByZero, "")
			}
			res = append(res, MakeNumberResult(l.ValueNumber/r.ValueNumber))
		case BinOpTypeExp:
			res = append(res, MakeNumberResult(math.Pow(l.ValueNumber, r.ValueNumber)))
		case BinOpTypeLT:
			res = append(res, MakeBoolResult(l.ValueNumber < r.ValueNumber))
		case BinOpTypeGT:
			res = append(res, MakeBoolResult(l.ValueNumber > r.ValueNumber))
		case BinOpTypeEQ:
			res = append(res, MakeBoolResult(l.ValueNumber == r.ValueNumber))
		case BinOpTypeLEQ:
			res = append(res, MakeBoolResult(l.ValueNumber <= r.ValueNumber))
		case BinOpTypeGEQ:
			res = append(res, MakeBoolResult(l.ValueNumber >= r.ValueNumber))
		case BinOpTypeNE:
			res = append(res, MakeBoolResult(l.ValueNumber != r.ValueNumber))
		// TODO: support concat here
		// case BinOpTypeConcat:
		default:
			return MakeErrorResult(fmt.Sprintf("unsupported list binary op %s", op))
		}
	}
	return MakeListResult(res)
}

func arrayValueOp(op BinOpType, lhs [][]Result, rhs Result) Result {
	// compare every item of array with a value
	res := [][]Result{}
	for i := range lhs {
		lst := listValueOp(op, lhs[i], rhs)
		if lst.Type == ResultTypeError {
			return lst
		}
		res = append(res, lst.ValueList)
	}
	return MakeArrayResult(res)
}

func listValueOp(op BinOpType, lhs []Result, rhs Result) Result {
	res := []Result{}
	// we can assume the arrays are the same size here
	switch rhs.Type {
	case ResultTypeNumber:
		rv := rhs.ValueNumber
		for i := range lhs {
			l := lhs[i].AsNumber()
			if l.Type != ResultTypeNumber {
				return MakeErrorResult("non-nunmeric value in binary operation")
			}
			switch op {
			case BinOpTypePlus:
				res = append(res, MakeNumberResult(l.ValueNumber+rv))
			case BinOpTypeMinus:
				res = append(res, MakeNumberResult(l.ValueNumber-rv))
			case BinOpTypeMult:
				res = append(res, MakeNumberResult(l.ValueNumber*rv))
			case BinOpTypeDiv:
				if rv == 0 {
					return MakeErrorResultType(ErrorTypeDivideByZero, "")
				}
				res = append(res, MakeNumberResult(l.ValueNumber/rv))
			case BinOpTypeExp:
				res = append(res, MakeNumberResult(math.Pow(l.ValueNumber, rv)))
			case BinOpTypeLT:
				res = append(res, MakeBoolResult(l.ValueNumber < rv))
			case BinOpTypeGT:
				res = append(res, MakeBoolResult(l.ValueNumber > rv))
			case BinOpTypeEQ:
				res = append(res, MakeBoolResult(l.ValueNumber == rv))
			case BinOpTypeLEQ:
				res = append(res, MakeBoolResult(l.ValueNumber <= rv))
			case BinOpTypeGEQ:
				res = append(res, MakeBoolResult(l.ValueNumber >= rv))
			case BinOpTypeNE:
				res = append(res, MakeBoolResult(l.ValueNumber != rv))
			// TODO: support concat here
			// case BinOpTypeConcat:
			default:
				return MakeErrorResult(fmt.Sprintf("unsupported list binary op %s", op))
			}
		}
	case ResultTypeString:
		rv := rhs.ValueString
		for i := range lhs {
			l := lhs[i].AsString()
			if l.Type != ResultTypeString {
				return MakeErrorResult("non-nunmeric value in binary operation")
			}
			switch op {
			case BinOpTypeLT:
				res = append(res, MakeBoolResult(l.ValueString < rv))
			case BinOpTypeGT:
				res = append(res, MakeBoolResult(l.ValueString > rv))
			case BinOpTypeEQ:
				res = append(res, MakeBoolResult(l.ValueString == rv))
			case BinOpTypeLEQ:
				res = append(res, MakeBoolResult(l.ValueString <= rv))
			case BinOpTypeGEQ:
				res = append(res, MakeBoolResult(l.ValueString >= rv))
			case BinOpTypeNE:
				res = append(res, MakeBoolResult(l.ValueString != rv))
			// TODO: support concat here
			// case BinOpTypeConcat:
			default:
				return MakeErrorResult(fmt.Sprintf("unsupported list binary op %s", op))
			}
		}
	default:
		return MakeErrorResult("non-nunmeric and non-string value in binary operation")
	}
	return MakeListResult(res)
}

// Eval evaluates the binary expression using the context given.
func (b BinaryExpr) String() string {
	opStr := ""
	switch b.op {
	case BinOpTypePlus:
		opStr = "+"
	case BinOpTypeMinus:
		opStr = "-"
	case BinOpTypeMult:
		opStr = "*"
	case BinOpTypeDiv:
		opStr = "/"
	case BinOpTypeExp:
		opStr = "^"
	case BinOpTypeLT:
		opStr = "<"
	case BinOpTypeGT:
		opStr = ">"
	case BinOpTypeEQ:
		opStr = "="
	case BinOpTypeLEQ:
		opStr = "<="
	case BinOpTypeGEQ:
		opStr = ">="
	case BinOpTypeNE:
		opStr = "<>"
	case BinOpTypeConcat:
		opStr = "&"
	}
	return b.lhs.String() + opStr + b.rhs.String()
}

// Update updates references in the BinaryExpr after removing a row/column.
func (b BinaryExpr) Update(q *update.UpdateQuery) Expression {
	new := b
	new.lhs = b.lhs.Update(q)
	new.rhs = b.rhs.Update(q)
	return new
}
