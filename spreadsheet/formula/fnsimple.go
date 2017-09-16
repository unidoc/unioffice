// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

import (
	"fmt"
	"log"
	"math"
)

func init() {
	RegisterFunction("SUM", Sum)
	RegisterFunction("MAX", Max)
	RegisterFunction("MIN", Min)
}

// Sum is an implementation of the Excel SUM() function.
func Sum(args []Result) Result {
	// Sum returns zero with no arguments
	res := MakeNumberResult(0)
	for _, a := range args {
		a = a.AsNumber()
		switch a.Type {
		case ResultTypeNumber:
			res.ValueNumber += a.ValueNumber
		case ResultTypeList, ResultTypeArray:
			subSum := Sum(a.ListValues())
			// error as sum returns only numbers and errors
			if subSum.Type != ResultTypeNumber {
				return subSum
			}
			res.ValueNumber += subSum.ValueNumber
		case ResultTypeString:
			// treated as zero by Excel
		case ResultTypeError:
			return a
		case ResultTypeEmpty:
			// skip
		default:
			return MakeErrorResult(fmt.Sprintf("unhandled SUM() argument type %s", a.Type))
		}
	}
	return res
}

// Min is an implementation of the Excel MIN() function.
func Min(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("MIN requires at least one argument")
	}

	res := MakeNumberResult(math.MaxFloat64)
	for _, a := range args {
		switch a.Type {
		case ResultTypeNumber:
			if a.ValueNumber < res.ValueNumber {
				res.ValueNumber = a.ValueNumber
			}
		case ResultTypeList:
			subMin := Min(a.ValueList)
			if subMin.ValueNumber < res.ValueNumber {
				res.ValueNumber = subMin.ValueNumber
			}
		case ResultTypeString:
			// treated as zero by Excel
			if 0 < res.ValueNumber {
				res.ValueNumber = 0
			}
		case ResultTypeEmpty:
		// skip
		case ResultTypeError:
			return a
		default:
			log.Printf("unhandled MIN() argument type %s", a.Type)
		}
	}

	return res
}

// Max is an implementation of the Excel MAX() function.
func Max(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("MAX requires at least one argument")
	}
	res := MakeNumberResult(-math.MaxFloat64)
	for _, a := range args {
		switch a.Type {
		case ResultTypeNumber:
			if a.ValueNumber > res.ValueNumber {
				res.ValueNumber = a.ValueNumber
			}
		case ResultTypeList:
			subMax := Max(a.ValueList)
			if subMax.ValueNumber > res.ValueNumber {
				res.ValueNumber = subMax.ValueNumber
			}
		case ResultTypeEmpty:
			// skip
		case ResultTypeString:
			// treated as zero by Excel
			if 0 > res.ValueNumber {
				res.ValueNumber = 0
			}
		default:
			log.Printf("unhandled MAX() argument type %s", a.Type)
		}
	}
	return res
}
