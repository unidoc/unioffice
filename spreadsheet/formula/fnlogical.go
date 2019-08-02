// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

func init() {
	RegisterFunction("AND", And)
	RegisterFunction("FALSE", False)
	RegisterFunction("IF", If)
	RegisterFunction("IFERROR", IfError)
	RegisterFunction("_xlfn.IFNA", IfNA) // Only in Excel 2013+
	// TODO:  RegisterFunction("IFS", Ifs) // Only in Excel 2016+
	RegisterFunction("NOT", Not)
	RegisterFunction("OR", Or)
	// TODO: RegisterFunction("SWITCH", Switch) // Only in Excel 2016+
	RegisterFunction("TRUE", True) // yup, TRUE()/FALSE() are functions in Excel, news to me...
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
	case ResultTypeNumber:
	case ResultTypeError:
		return cond
	default:
		return MakeErrorResult("IF initial argument must be numeric")
	}

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
	}
	return MakeBoolResult(false)

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
