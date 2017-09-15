// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func init() {
	RegisterFunction("ABS", makeMathWrapper("ASIN", math.Abs))
	RegisterFunction("ACOS", makeMathWrapper("ASIN", math.Acos))
	RegisterFunction("ACOSH", makeMathWrapper("ASIN", math.Acosh))
	// TODO: RegisterFunction("ACOT", Acot) /// Excel 2013+
	// TODO: RegisterFunction("ACOTH", Acoth) /// Excel 2013+
	// TODO: RegisterFunction("_xlfn.AGGREGATE", Aggregate) // lots of dependencies
	RegisterFunction("_xlfn.ARABIC", Arabic)
	RegisterFunction("ASIN", makeMathWrapper("ASIN", math.Asin))
	RegisterFunction("ASINH", makeMathWrapper("ASINH", math.Asinh))
	RegisterFunction("ATAN", makeMathWrapper("ATAN", math.Atan))
	RegisterFunction("ATANH", makeMathWrapper("ATANH", math.Atanh))
	RegisterFunction("ATAN2", Atan2)
	RegisterFunction("_xlfn.BASE", Base)
	// RegisterFunction("CEILING", ) // TODO: figure out how this acts, Libre doesn't use it
	RegisterFunction("_xlfn.CEILING.MATH", CeilingMath)
	RegisterFunction("_xlfn.CEILING.PRECISE", CeilingPrecise)
	RegisterFunction("COMBIN", Combin)
	RegisterFunction("_xlfn.COMBINA", Combina)
	RegisterFunction("COS", makeMathWrapper("COS", math.Cos))
	RegisterFunction("COSH", makeMathWrapper("COSH", math.Cosh))
	//RegisterFunction("COT",
	//RegisterFunction("COTH"
	//RegisterFunction("CSC"
	//RegisterFunction("CSCH"
	RegisterFunction("_xlfn.DECIMAL", Decimal)
	RegisterFunction("DEGREES", Degrees)
	RegisterFunction("PI", Pi)
}

// makeMathWrapper is used to wrap single argument math functions from the Go
// standard library and present them as a spreadsheet function.
func makeMathWrapper(name string, fn func(x float64) float64) Function {
	return func(args []Result) Result {
		if len(args) != 1 {
			return MakeErrorResult(name + " requires one argument")
		}

		arg := args[0].AsNumber()
		switch arg.Type {
		case ResultTypeNumber:
			v := fn(arg.ValueNumber)
			if v != v {
				return MakeErrorResult(name + " returned NaN")
			}
			return MakeNumberResult(v)
		case ResultTypeList, ResultTypeString:
			return MakeErrorResult(name + " requires a numeric argument")
		case ResultTypeError:
			return arg
		default:
			return MakeErrorResult(fmt.Sprintf("unhandled %s() argument type %s", name, arg.Type))
		}
	}
}

// Atan2 implements the Excel ATAN2 function.  It accepts two numeric arguments,
// and the arguments are (x,y), reversed from normal to match Excel's behaviour.
func Atan2(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("ATAN2 requires two arguments")
	}

	arg1 := args[0].AsNumber()
	arg2 := args[1].AsNumber()
	if arg1.Type == ResultTypeNumber && arg2.Type == ResultTypeNumber {
		// args are swapped here
		v := math.Atan2(arg2.ValueNumber, arg1.ValueNumber)
		if v != v {
			return MakeErrorResult("ATAN2 returned NaN")
		}
		return MakeNumberResult(v)
	}

	for _, t := range []ResultType{arg1.Type, arg2.Type} {
		switch t {
		case ResultTypeList, ResultTypeString:
			return MakeErrorResult("ATAN2 requires a numeric argument")
		case ResultTypeError:
			return arg1
		default:
			return MakeErrorResult(fmt.Sprintf("unhandled ATAN2() argument type %s", t))
		}
	}
	return MakeErrorResult("unhandled error for ATAN2()")
}

// Arabic implements the Excel ARABIC function which parses roman numerals.  It
// accepts one numeric argument.
func Arabic(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("ARABIC requires one argument")
	}
	arg := args[0]
	switch arg.Type {
	case ResultTypeNumber, ResultTypeList, ResultTypeEmpty:
		return MakeErrorResult("ARABIC requires a string argument argument")
	case ResultTypeString:
		res := 0.0
		last := 0.0
		for _, c := range arg.ValueString {
			digit := 0.0
			switch c {
			case 'I':
				digit = 1
			case 'V':
				digit = 5
			case 'X':
				digit = 10
			case 'L':
				digit = 50
			case 'C':
				digit = 100
			case 'D':
				digit = 500
			case 'M':
				digit = 1000
			}
			res += digit
			switch {
			// repeated digits
			case last == digit &&
				(last == 5 || last == 50 || last == 500):
				return MakeErrorResult("invalid ARABIC format")
				// simpler form
			case 2*last == digit:
				return MakeErrorResult("invalid ARABIC format")
			}
			if last < digit {
				res -= 2 * last
			}
			last = digit
		}

		return MakeNumberResult(res)
	case ResultTypeError:
		return arg
	default:
		return MakeErrorResult(fmt.Sprintf("unhandled ACOSH() argument type %s", arg.Type))
	}
}

// CeilingMath implements _xlfn.CEILING.MATH which rounds numbers to the nearest
// multiple of the second argument, toward or away from zero as specified by the
// third argument.
func CeilingMath(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("CEILING.MATH() requires at least one argument")
	}
	if len(args) > 3 {
		return MakeErrorResult("CEILING.MATH() allows at most three arguments")
	}
	// number to round
	number := args[0].AsNumber()
	if number.Type != ResultTypeNumber {
		return MakeErrorResult("first arugment to CEILING.MATH() must be a number")
	}

	// significance
	significance := float64(1)
	if number.ValueNumber < 0 {
		significance = -1
	}
	if len(args) > 1 {
		sigArg := args[1].AsNumber()
		if sigArg.Type != ResultTypeNumber {
			return MakeErrorResult("second arugment to CEILING.MATH() must be a number")
		}
		significance = sigArg.ValueNumber
	}

	// round direction
	direction := float64(1)
	if len(args) > 2 {
		dirArg := args[2].AsNumber()
		if dirArg.Type != ResultTypeNumber {
			return MakeErrorResult("third arugment to CEILING.MATH() must be a number")
		}
		direction = dirArg.ValueNumber
	}

	if len(args) == 1 {
		return MakeNumberResult(math.Ceil(number.ValueNumber))
	}

	v := number.ValueNumber
	v, res := math.Modf(v / significance)
	if res != 0 {
		if number.ValueNumber > 0 {
			v++
		} else if direction < 0 {
			v--
		}
	}
	return MakeNumberResult(v * significance)
}

func CeilingPrecise(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("CEILING.PRECISE() requires at least one argument")
	}
	if len(args) > 2 {
		return MakeErrorResult("CEILING.PRECISE() allows at most two arguments")
	}
	// number to round
	number := args[0].AsNumber()
	if number.Type != ResultTypeNumber {
		return MakeErrorResult("first arugment to CEILING.PRECISE() must be a number")
	}

	// significance
	significance := float64(1)
	if number.ValueNumber < 0 {
		significance = -1
	}
	if len(args) > 1 {
		sigArg := args[1].AsNumber()
		if sigArg.Type != ResultTypeNumber {
			return MakeErrorResult("second arugment to CEILING.MATH() must be a number")
		}
		// don't care about sign of significance
		significance = math.Abs(sigArg.ValueNumber)
	}

	if len(args) == 1 {
		return MakeNumberResult(math.Ceil(number.ValueNumber))
	}

	v := number.ValueNumber
	v, res := math.Modf(v / significance)
	if res != 0 {
		if number.ValueNumber > 0 {
			v++
		}
	}
	return MakeNumberResult(v * significance)
}

// Base is an implementation of the Excel BASE function that returns a string
// form of an integer in a specified base and of a minimum length with padded
// zeros.
func Base(args []Result) Result {
	if len(args) < 2 {
		return MakeErrorResult("BASE() requires at least two arguments")
	}
	if len(args) > 3 {
		return MakeErrorResult("BASE() allows at most three arguments")
	}
	// number to convert
	number := args[0].AsNumber()
	if number.Type != ResultTypeNumber {
		return MakeErrorResult("first arugment to BASE() must be a number")
	}

	radixArg := args[1].AsNumber()
	if radixArg.Type != ResultTypeNumber {
		return MakeErrorResult("second arugment to BASE() must be a number")
	}
	radix := int(radixArg.ValueNumber)
	if radix < 0 || radix > 36 {
		return MakeErrorResult("radix must be in the range [0,36]")
	}

	// min length of result
	minLength := 0
	if len(args) > 2 {
		lenArg := args[2].AsNumber()
		if lenArg.Type != ResultTypeNumber {
			return MakeErrorResult("third arugment to BASE() must be a number")
		}
		minLength = int(lenArg.ValueNumber)
	}

	s := strconv.FormatInt(int64(number.ValueNumber), radix)
	if len(s) < minLength {
		s = strings.Repeat("0", minLength-len(s)) + s
	}
	return MakeStringResult(s)
}

// Combin is an implementation of the Excel COMBINA function whic returns the
// number of combinations.
func Combin(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("COMBIN() requires two argument")
	}
	nArg := args[0].AsNumber()
	kArg := args[1].AsNumber()
	if nArg.Type != ResultTypeNumber || kArg.Type != ResultTypeNumber {
		return MakeErrorResult("COMBIN() requires numeric arguments")
	}
	n := math.Trunc(nArg.ValueNumber)
	k := math.Trunc(kArg.ValueNumber)
	if k > n {
		return MakeErrorResult("COMBIN() requires k <= n")
	}
	if k == n || k == 0 {
		return MakeNumberResult(1)
	}

	res := float64(1)
	for i := float64(1); i <= k; i++ {
		res *= (n + 1 - i) / i
	}

	return MakeNumberResult(res)
}

// Combina is an implementation of the Excel COMBINA function whic returns the
// number of combinations with repetitions.
func Combina(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("COMBINA() requires two argument")
	}
	nArg := args[0].AsNumber()
	kArg := args[1].AsNumber()
	if nArg.Type != ResultTypeNumber || kArg.Type != ResultTypeNumber {
		return MakeErrorResult("COMBINA() requires numeric arguments")
	}
	n := math.Trunc(nArg.ValueNumber)
	k := math.Trunc(kArg.ValueNumber)
	if n < k {
		return MakeErrorResult("COMBINA() requires n > k")
	}
	if n == 0 {
		return MakeNumberResult(0)
	}
	args[0] = MakeNumberResult(n + k - 1)
	args[1] = MakeNumberResult(n - 1)
	return Combin(args)
}

func fact(f float64) float64 {
	res := float64(1)
	for i := float64(2); i <= f; i++ {
		res *= i
	}
	return res
}

// Decimal is an implementation of the Excel function DECIMAL() that parses a string
// in a given base and returns the numeric result.
func Decimal(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("DECIMAL() requires two arguments")
	}
	sArg := args[0].AsString()
	if sArg.Type != ResultTypeString {
		return MakeErrorResult("DECIMAL() requires string first argument")
	}
	baseArg := args[1].AsNumber()
	if baseArg.Type != ResultTypeNumber {
		return MakeErrorResult("DECIMAL() requires number second argument")
	}

	sv := sArg.ValueString
	if len(sv) > 2 && (strings.HasPrefix(sv, "0x") || strings.HasPrefix(sv, "0X")) {
		sv = sv[2:]
	}
	i, err := strconv.ParseInt(sv, int(baseArg.ValueNumber), 64)
	if err != nil {
		return MakeErrorResult("DECIMAL() error in conversion")
	}
	return MakeNumberResult(float64(i))
}

// Degrees is an implementation of the Excel function DEGREES() that converts
// radians to degrees.
func Degrees(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("DEGREES() requires one argument")
	}
	vArg := args[0].AsNumber()
	if vArg.Type != ResultTypeNumber {
		return MakeErrorResult("DEGREES() requires number argument")
	}

	return MakeNumberResult(180.0 / math.Pi * vArg.ValueNumber)
}

// Pi is an implementation of the Excel Pi() function that just returns the Pi
// constant.
func Pi(args []Result) Result {
	if len(args) != 0 {
		return MakeErrorResult("PI() accepts no arguments")
	}
	return MakeNumberResult(math.Pi)
}
