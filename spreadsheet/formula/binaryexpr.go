// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

import "math"

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

// Eval evaluates the binary expression using the context given.
func (b BinaryExpr) Eval(ctx Context, ev Evaluator) Result {
	lhs := b.lhs.Eval(ctx, ev)
	rhs := b.rhs.Eval(ctx, ev)

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
		}
	case BinOpTypeGT:
		if lhs.Type == rhs.Type {
			if lhs.Type == ResultTypeNumber {
				return MakeBoolResult(lhs.ValueNumber > rhs.ValueNumber)
			}
		}
	case BinOpTypeEQ:
		if lhs.Type == rhs.Type {
			if lhs.Type == ResultTypeNumber {
				// TODO: see what Excel does regarding floating point comparison
				return MakeBoolResult(lhs.ValueNumber == rhs.ValueNumber)
			}
		}
	case BinOpTypeNE:
		if lhs.Type == rhs.Type {
			if lhs.Type == ResultTypeNumber {
				return MakeBoolResult(lhs.ValueNumber != rhs.ValueNumber)
			}
		}
	case BinOpTypeLEQ:
		if lhs.Type == rhs.Type {
			if lhs.Type == ResultTypeNumber {
				return MakeBoolResult(lhs.ValueNumber <= rhs.ValueNumber)
			}
		}
	case BinOpTypeGEQ:
		if lhs.Type == rhs.Type {
			if lhs.Type == ResultTypeNumber {
				return MakeBoolResult(lhs.ValueNumber >= rhs.ValueNumber)
			}
		}
	case BinOpTypeConcat:
		return MakeStringResult(lhs.Value() + rhs.Value())
	}

	return MakeErrorResult("unsupported binary op")
}

func (b BinaryExpr) Reference(ctx Context, ev Evaluator) Reference {
	return ReferenceInvalid
}
