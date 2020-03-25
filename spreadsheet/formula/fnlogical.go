// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

func init() {
	RegisterFunction("AND", And)
	RegisterFunction("FALSE", False)
	RegisterFunction("IF", If)
	RegisterFunction("IFERROR", IfError)
	RegisterFunction("_xlfn.IFNA", IfNA) // Only in Excel 2013+
	RegisterFunction("IFS", Ifs)
	RegisterFunction("_xlfn.IFS", Ifs)
	RegisterFunction("NOT", Not)
	RegisterFunction("OR", Or)
	RegisterFunction("TRUE", True)
	RegisterFunction("_xlfn.XOR", Xor)
}

// And is an implementation of the Excel AND() function.
func And(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("AND requires at least one argument")
	}
	result := true
	for _, a := range args {
		a = a.AsNumber()
		switch a.Type {
		case ResultTypeList, ResultTypeArray:
			res := And(a.ListValues())
			if res.Type == ResultTypeError {
				return res
			}
			if res.ValueNumber == 0 {
				result = false
			}
		case ResultTypeNumber:
			if a.ValueNumber == 0 {
				result = false
			}
		case ResultTypeString:
			return MakeErrorResult("AND doesn't operate on strings")
		case ResultTypeError:
			return a
		default:
			return MakeErrorResult("unsupported argument type in AND")
		}
	}
	return MakeBoolResult(result)
}

// False is an implementation of the Excel FALSE() function. It takes no
// arguments.
func False(args []Result) Result {
	if len(args) != 0 {
		return MakeErrorResult("FALSE takes no arguments")
	}
	return MakeBoolResult(false)
}

// If is an implementation of the Excel IF() function. It takes one, two or
// three arguments.
func If(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("IF requires at least one argument")
	}
	if len(args) > 3 {
		return MakeErrorResult("IF accepts at most three arguments")
	}

	cond := args[0]
	switch cond.Type {
	case ResultTypeError:
		return cond
	case ResultTypeNumber:
		// single argument just returns the condition value
		if len(args) == 1 {
			return MakeBoolResult(cond.ValueNumber != 0)
		}

		// true case
		if cond.ValueNumber != 0 {
			return args[1]
		}

		// false case
		if len(args) == 3 {
			return args[2]
		} else {
			return MakeBoolResult(false)
		}
	case ResultTypeList:
		return ifList(args)
	case ResultTypeArray:
		return ifArray(args)

	default:
		return MakeErrorResult("IF initial argument must be numeric or array")
	}
}

func fillArray(arg Result, rows, cols int) [][]Result {
	result := [][]Result{}
	switch arg.Type {
	case ResultTypeArray:
		for ir, row := range arg.ValueArray {
			if ir < rows {
				result = append(result, fillList(MakeListResult(row), cols))
			} else {
				result = append(result, fillList(MakeErrorResultType(ErrorTypeNA, ""), cols))
			}
		}
	case ResultTypeList:
		list := fillList(arg, cols)
		for r := 0; r < rows; r++ {
			result = append(result, list)
		}
	case ResultTypeNumber, ResultTypeString, ResultTypeError, ResultTypeEmpty:
		for r := 0; r < rows; r++ {
			list := fillList(arg, cols)
			result = append(result, list)
		}
	}
	return result
}

func fillList(arg Result, cols int) []Result {
	list := []Result{}
	switch arg.Type {
	case ResultTypeList:
		valueList := arg.ValueList
		maxCols := len(valueList)
		for ic := 0; ic < cols; ic++ {
			if ic < maxCols {
				list = append(list, valueList[ic])
			} else {
				list = append(list, MakeErrorResultType(ErrorTypeNA, ""))
			}
		}
	case ResultTypeNumber, ResultTypeString, ResultTypeError, ResultTypeEmpty:
		for ic := 0; ic < cols; ic++ {
			list = append(list, arg)
		}
	}
	return list
}

func ifArray(args []Result) Result {
	condArray := args[0].ValueArray
	// single argument returns list of contitions
	if len(args) == 1 {
		result := [][]Result{}
		for _, v := range condArray {
			result = append(result, ifList([]Result{MakeListResult(v)}).ValueList)
		}
		return MakeArrayResult(result)
	} else if len(args) == 2 {
		rows := len(condArray)
		cols := len(condArray[0])
		truesArray := fillArray(args[1], rows, cols)
		tl := len(truesArray)
		result := [][]Result{}
		var truesList []Result
		for i, v := range condArray {
			if i < tl {
				truesList = truesArray[i]
			} else {
				truesList = fillList(MakeErrorResultType(ErrorTypeNA, ""), cols)
			}
			result = append(result, ifList([]Result{MakeListResult(v), MakeListResult(truesList)}).ValueList)
		}
		return MakeArrayResult(result)
	} else if len(args) == 3 {
		rows := len(condArray)
		cols := len(condArray[0])
		truesArray := fillArray(args[1], rows, cols)
		falsesArray := fillArray(args[2], rows, cols)
		tl := len(truesArray)
		fl := len(falsesArray)
		result := [][]Result{}
		var truesList, falsesList []Result
		for i, v := range condArray {
			if i < tl {
				truesList = truesArray[i]
			} else {
				truesList = fillList(MakeErrorResultType(ErrorTypeNA, ""), cols)
			}
			if i < fl {
				falsesList = falsesArray[i]
			} else {
				falsesList = fillList(MakeErrorResultType(ErrorTypeNA, ""), cols)
			}
			result = append(result, ifList([]Result{MakeListResult(v), MakeListResult(truesList), MakeListResult(falsesList)}).ValueList)
		}
		return MakeArrayResult(result)
	}
	return MakeErrorResultType(ErrorTypeValue, "")
}

func ifList(args []Result) Result {
	condList := args[0].ValueList
	cols := len(condList)
	switch len(args) {
	// single argument returns list of contitions
	case 1:
		result := []Result{}
		for _, v := range condList {
			result = append(result, MakeBoolResult(v.ValueNumber != 0))
		}
		return MakeListResult(result)
	// two arguments case
	case 2:
		trues := args[1]
		switch trues.Type {
		case ResultTypeNumber, ResultTypeString, ResultTypeEmpty:
			result := []Result{}
			for _, v := range condList {
				var newValue Result
				if v.ValueNumber == 0 {
					newValue = MakeBoolResult(false)
				} else {
					newValue = trues
				}
				result = append(result, newValue)
			}
			return MakeListResult(result)
		case ResultTypeList:
			truesList := fillList(trues, cols)
			result := []Result{}
			for i, v := range condList {
				var newValue Result
				if v.ValueNumber == 0 {
					newValue = MakeBoolResult(false)
				} else {
					newValue = truesList[i]
				}
				result = append(result, newValue)
			}
			return MakeListResult(result)
		case ResultTypeArray:
			truesArray := fillArray(trues, len(trues.ValueArray), cols)
			result := [][]Result{}
			for _, row := range truesArray {
				rowResult := []Result{}
				for ic, v := range condList {
					var newValue Result
					if v.ValueNumber == 0 {
						newValue = MakeBoolResult(false)
					} else {
						newValue = row[ic]
					}
					rowResult = append(rowResult, newValue)
				}
				result = append(result, rowResult)
			}
			return MakeArrayResult(result)
		}
	// false case
	case 3:
		trues := args[1]
		falses := args[2]
		truesSingleValue := checkSingleValue(trues)
		falsesSingleValue := checkSingleValue(falses)
		if truesSingleValue && falsesSingleValue {
			result := []Result{}
			for _, v := range condList {
				var newValue Result
				if v.ValueNumber == 0 {
					newValue = falses
				} else {
					newValue = trues
				}
				result = append(result, newValue)
			}
			return MakeListResult(result)
		}

		if trues.Type != ResultTypeArray && falses.Type != ResultTypeArray {
			truesList := fillList(trues, cols)
			falsesList := fillList(falses, cols)
			result := []Result{}
			for i, v := range condList {
				var newValue Result
				if v.ValueNumber == 0 {
					newValue = falsesList[i]
				} else {
					newValue = truesList[i]
				}
				result = append(result, newValue)
			}
			return MakeListResult(result)
		}
		maxRows := len(trues.ValueArray)
		if len(falses.ValueArray) > maxRows {
			maxRows = len(falses.ValueArray)
		}
		truesArray := fillArray(trues, maxRows, cols)
		falsesArray := fillArray(falses, maxRows, cols)
		result := [][]Result{}
		for ir := 0; ir < maxRows; ir++ {
			rowResult := []Result{}
			for ic, v := range condList {
				var newValue Result
				if v.ValueNumber == 0 {
					newValue = falsesArray[ir][ic]
				} else {
					newValue = truesArray[ir][ic]
				}
				rowResult = append(rowResult, newValue)
			}
			result = append(result, rowResult)
		}
		return MakeArrayResult(result)
	}
	return MakeErrorResult("")
}

func checkSingleValue(result Result) bool {
	t := result.Type
	return t != ResultTypeArray && t != ResultTypeList
}

// IfError is an implementation of the Excel IFERROR() function. It takes two arguments.
func IfError(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("IFERROR requires two arguments")
	}

	if args[0].Type != ResultTypeError {
		// empty cell returns a zero
		if args[0].Type == ResultTypeEmpty {
			return MakeNumberResult(0)
		}
		return args[0]
	}

	return args[1]
}

// IfNA is an implementation of the Excel IFNA() function. It takes two arguments.
func IfNA(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("IFNA requires two arguments")
	}

	if args[0].Type == ResultTypeError && args[0].ValueString == "#N/A" {
		return args[1]
	}
	return args[0]
}

// Ifs is an implementation of the Excel IFS() function.
func Ifs(args []Result) Result {
	if len(args) < 2 {
		return MakeErrorResult("IFS requires at least two arguments")
	}
	for i := 0; i < len(args)-1; i += 2 {
		if args[i].ValueNumber == 1 {
			return args[i+1]
		}
	}
	return MakeErrorResultType(ErrorTypeNA, "")
}

// Not is an implementation of the Excel NOT() function and takes a single
// argument.
func Not(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("NOT requires one argument")
	}
	switch args[0].Type {
	case ResultTypeError:
		return args[0]
	case ResultTypeString, ResultTypeList:
		return MakeErrorResult("NOT expects a numeric argument")
	case ResultTypeNumber:
		return MakeBoolResult(!(args[0].ValueNumber != 0))
	default:
		return MakeErrorResult("unhandled NOT argument type")
	}
}

// Or is an implementation of the Excel OR() function and takes a variable
// number of arguments.
func Or(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("OR requires at least one argument")
	}
	result := false
	for _, a := range args {
		switch a.Type {
		case ResultTypeList, ResultTypeArray:
			res := Or(a.ListValues())
			if res.Type == ResultTypeError {
				return res
			}
			if res.ValueNumber != 0 {
				result = true
			}
		case ResultTypeNumber:
			if a.ValueNumber != 0 {
				result = true
			}
		case ResultTypeString:
			return MakeErrorResult("OR doesn't operate on strings")
		case ResultTypeError:
			return a
		default:
			return MakeErrorResult("unsupported argument type in OR")
		}
	}
	return MakeBoolResult(result)
}

// True is an implementation of the Excel TRUE() function.  It takes no
// arguments.
func True(args []Result) Result {
	if len(args) != 0 {
		return MakeErrorResult("TRUE takes no arguments")
	}
	return MakeBoolResult(true)
}

// Xor is an implementation of the Excel XOR() function and takes a variable
// number of arguments. It's odd to say the least.  If any argument is numeric,
// it returns true if the number of non-zero numeric arguments is odd and false
// otherwise.  If no argument is numeric, it returns an error.
func Xor(args []Result) Result {
	if len(args) < 1 {
		return MakeErrorResult("XOR requires at least one argument")
	}
	cnt := 0
	hasNum := false
	for _, arg := range args {
		switch arg.Type {
		case ResultTypeList, ResultTypeArray:
			res := Xor(arg.ListValues())
			if res.Type == ResultTypeError {
				return res
			}
			if res.ValueNumber != 0 {
				cnt++
			}
			hasNum = true
		case ResultTypeNumber:
			if arg.ValueNumber != 0 {
				cnt++
			}
			hasNum = true
		case ResultTypeString:
			// XOR appers to treat strings as zero values.
		case ResultTypeError:
			return arg
		default:
			return MakeErrorResult("unsupported argument type in XOR")
		}
	}
	if !hasNum {
		return MakeErrorResult("XOR requires numeric input")
	}
	return MakeBoolResult(cnt%2 != 0)
}
