// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package spreadsheet

import (
	"github.com/unidoc/unioffice/schema/soo/sml"
)

// DVCompareType is a comparison type for a data validation rule. This restricts
// the input format of the cell.
type DVCompareType byte

const (
	DVCompareTypeWholeNumber = DVCompareType(sml.ST_DataValidationTypeWhole)
	DVCompareTypeDecimal     = DVCompareType(sml.ST_DataValidationTypeDecimal)
	DVCompareTypeDate        = DVCompareType(sml.ST_DataValidationTypeDate)
	DVCompareTypeTime        = DVCompareType(sml.ST_DataValidationTypeTime)
	DVompareTypeTextLength   = DVCompareType(sml.ST_DataValidationTypeTextLength)
)

// DVCompareOp is a comparison operator for a data validation rule.
type DVCompareOp byte

const (
	DVCompareOpEqual        = DVCompareOp(sml.ST_DataValidationOperatorEqual)
	DVCompareOpBetween      = DVCompareOp(sml.ST_DataValidationOperatorBetween)
	DVCompareOpNotBetween   = DVCompareOp(sml.ST_DataValidationOperatorNotBetween)
	DVCompareOpNotEqual     = DVCompareOp(sml.ST_DataValidationOperatorNotEqual)
	DVCompareOpGreater      = DVCompareOp(sml.ST_DataValidationOperatorGreaterThan)
	DVCompareOpGreaterEqual = DVCompareOp(sml.ST_DataValidationOperatorGreaterThanOrEqual)
	DVCompareOpLess         = DVCompareOp(sml.ST_DataValidationOperatorLessThan)
	DVCompareOpLessEqual    = DVCompareOp(sml.ST_DataValidationOperatorLessThanOrEqual)
)

const (
	DVOpGreater = sml.ST_DataValidationOperatorGreaterThanOrEqual
)

// DataValidationCompare is a view on a data validation rule that is oriented
// towards value comparisons.
type DataValidationCompare struct {
	x *sml.CT_DataValidation
}

// SetValue sets the first value to be used in the comparison.  For comparisons
// that need only one value, this is the only value used.  For comparisons like
// 'between' that require two values, SetValue2 must also be used.
func (d DataValidationCompare) SetValue(v string) {
	d.x.Formula1 = &v
}
func (d DataValidationCompare) SetValue2(v string) {
	d.x.Formula2 = &v
}
