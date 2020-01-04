// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var rnd *rand.Rand

func init() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	RegisterFunction("ABS", makeMathWrapper("ASIN", math.Abs))
	RegisterFunction("ACOS", makeMathWrapper("ASIN", math.Acos))
	RegisterFunction("ACOSH", makeMathWrapper("ASIN", math.Acosh))
	RegisterFunction("_xlfn.ACOT", makeMathWrapper("ACOT", func(v float64) float64 { return math.Pi/2 - math.Atan(v) }))
	RegisterFunction("_xlfn.ACOTH", makeMathWrapper("ACOTH", func(v float64) float64 { return math.Atanh(1 / v) }))
	RegisterFunction("_xlfn.ARABIC", Arabic)
	RegisterFunction("ASIN", makeMathWrapper("ASIN", math.Asin))
	RegisterFunction("ASINH", makeMathWrapper("ASINH", math.Asinh))
	RegisterFunction("ATAN", makeMathWrapper("ATAN", math.Atan))
	RegisterFunction("ATANH", makeMathWrapper("ATANH", math.Atanh))
	RegisterFunction("ATAN2", Atan2)
	RegisterFunction("_xlfn.BASE", Base)
	RegisterFunction("CEILING", Ceiling)
	RegisterFunction("_xlfn.CEILING.MATH", CeilingMath)
	RegisterFunction("_xlfn.CEILING.PRECISE", CeilingPrecise)
	RegisterFunction("COMBIN", Combin)
	RegisterFunction("_xlfn.COMBINA", Combina)
	RegisterFunction("COS", makeMathWrapper("COS", math.Cos))
	RegisterFunction("COSH", makeMathWrapper("COSH", math.Cosh))
	RegisterFunction("_xlfn.COT", makeMathWrapperInv("COT", math.Tan))
	RegisterFunction("_xlfn.COTH", makeMathWrapperInv("COTH", math.Tanh))
	RegisterFunction("_xlfn.CSC", makeMathWrapperInv("CSC", math.Sin))
	RegisterFunction("_xlfn.CSCH", makeMathWrapperInv("CSC", math.Sinh))
	RegisterFunction("_xlfn.DECIMAL", Decimal)
	RegisterFunction("DEGREES", Degrees)
	RegisterFunction("EVEN", Even)
	RegisterFunction("EXP", makeMathWrapper("EXP", math.Exp))
	RegisterFunction("FACT", Fact)
	RegisterFunction("FACTDOUBLE", FactDouble)
	RegisterFunction("FLOOR", Floor)
	RegisterFunction("_xlfn.FLOOR.MATH", FloorMath)
	RegisterFunction("_xlfn.FLOOR.PRECISE", FloorPrecise)
	RegisterFunction("GCD", GCD)
	RegisterFunction("INT", Int)
	RegisterFunction("ISO.CEILING", CeilingPrecise)
	RegisterFunction("LCM", LCM)
	RegisterFunction("LN", makeMathWrapper("LN", math.Log))
	RegisterFunction("LOG", Log)
	RegisterFunction("LOG10", makeMathWrapper("LOG10", math.Log10))
	RegisterFunction("MDETERM", MDeterm)
	RegisterFunction("MOD", Mod)
	RegisterFunction("MROUND", Mround)
	RegisterFunction("MULTINOMIAL", Multinomial)
	RegisterFunction("_xlfn.MUNIT", Munit)
	RegisterFunction("ODD", Odd)
	RegisterFunction("PI", Pi)
	RegisterFunction("POWER", Power)
	RegisterFunction("PRODUCT", Product)
	RegisterFunction("QUOTIENT", Quotient)
	RegisterFunction("RADIANS", Radians)
	RegisterFunction("RAND", Rand)
	RegisterFunction("RANDBETWEEN", RandBetween)
	RegisterFunction("ROMAN", Roman)
	RegisterFunction("ROUND", Round)
	RegisterFunction("ROUNDDOWN", RoundDown)
	RegisterFunction("ROUNDUP", RoundUp)
	RegisterFunction("_xlfn.SEC", makeMathWrapperInv("SEC", math.Cos))
	RegisterFunction("_xlfn.SECH", makeMathWrapperInv("SECH", math.Cosh))
	RegisterFunction("SERIESSUM", SeriesSum)
	RegisterFunction("SIGN", Sign)
	RegisterFunction("SIN", makeMathWrapper("SIN", math.Sin))
	RegisterFunction("SINH", makeMathWrapper("SINH", math.Sinh))
	RegisterFunction("SQRT", makeMathWrapper("SQRT", math.Sqrt))
	RegisterFunction("SQRTPI", makeMathWrapper("SQRTPI", func(v float64) float64 { return math.Sqrt(v * math.Pi) }))
	RegisterFunction("SUM", Sum)
	RegisterFunction("SUMIF", SumIf)
	RegisterFunction("SUMIFS", SumIfs)
	RegisterFunction("SUMPRODUCT", SumProduct)
	RegisterFunction("SUMSQ", SumSquares)
	RegisterFunction("TAN", makeMathWrapper("TAN", math.Tan))
	RegisterFunction("TANH", makeMathWrapper("TANH", math.Tanh))
	RegisterFunction("TRUNC", Trunc)
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
			if math.IsNaN(v) {
				return MakeErrorResult(name + " returned NaN")
			}
			if math.IsInf(v, 0) {
				return MakeErrorResult(name + " returned infinity")
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

func makeMathWrapperInv(name string, fn func(x float64) float64) Function {
	return func(args []Result) Result {
		if len(args) != 1 {
			return MakeErrorResult(name + " requires one argument")
		}

		arg := args[0].AsNumber()
		switch arg.Type {
		case ResultTypeNumber:
			v := fn(arg.ValueNumber)
			if math.IsNaN(v) {
				return MakeErrorResult(name + " returned NaN")
			}
			if math.IsInf(v, 0) {
				return MakeErrorResult(name + " returned infinity")
			}
			if v == 0 {
				return MakeErrorResultType(ErrorTypeDivideByZero, name+" divide by zero")
			}
			return MakeNumberResult(1 / v)
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
		return MakeErrorResult("first argument to CEILING.MATH() must be a number")
	}

	// significance
	significance := float64(1)
	if number.ValueNumber < 0 {
		significance = -1
	}
	if len(args) > 1 {
		sigArg := args[1].AsNumber()
		if sigArg.Type != ResultTypeNumber {
			return MakeErrorResult("second argument to CEILING.MATH() must be a number")
		}
		significance = sigArg.ValueNumber
	}

	// round direction
	direction := float64(1)
	if len(args) > 2 {
		dirArg := args[2].AsNumber()
		if dirArg.Type != ResultTypeNumber {
			return MakeErrorResult("third argument to CEILING.MATH() must be a number")
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

// Ceiling is an implementation of the CEILING function which
// returns the ceiling of a number.
func Ceiling(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("CEILING() requires at least one argument")
	}
	if len(args) > 2 {
		return MakeErrorResult("CEILING() allows at most two arguments")
	}
	// number to round
	number := args[0].AsNumber()
	if number.Type != ResultTypeNumber {
		return MakeErrorResult("first argument to CEILING() must be a number")
	}

	// significance
	significance := float64(1)
	if number.ValueNumber < 0 {
		significance = -1
	}
	if len(args) > 1 {
		sigArg := args[1].AsNumber()
		if sigArg.Type != ResultTypeNumber {
			return MakeErrorResult("second argument to CEILING() must be a number")
		}
		significance = sigArg.ValueNumber
	}

	if significance < 0 && number.ValueNumber > 0 {
		return MakeErrorResultType(ErrorTypeNum, "negative sig to CEILING() invalid")
	}

	if len(args) == 1 {
		return MakeNumberResult(math.Ceil(number.ValueNumber))
	}

	v := number.ValueNumber
	v, res := math.Modf(v / significance)
	if res > 0 {
		v++
	}
	return MakeNumberResult(v * significance)
}

// CeilingPrecise is an implementation of the CEILING.PRECISE function which
// returns the ceiling of a number.
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
		return MakeErrorResult("first argument to CEILING.PRECISE() must be a number")
	}

	// significance
	significance := float64(1)
	if number.ValueNumber < 0 {
		significance = -1
	}
	if len(args) > 1 {
		sigArg := args[1].AsNumber()
		if sigArg.Type != ResultTypeNumber {
			return MakeErrorResult("second argument to CEILING.PRECISE() must be a number")
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
		return MakeErrorResult("first argument to BASE() must be a number")
	}

	radixArg := args[1].AsNumber()
	if radixArg.Type != ResultTypeNumber {
		return MakeErrorResult("second argument to BASE() must be a number")
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
			return MakeErrorResult("third argument to BASE() must be a number")
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

// Even is an implementation of the Excel EVEN() that rounds a number to the
// nearest even integer.
func Even(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("EVEN() requires one argument")
	}
	vArg := args[0].AsNumber()
	if vArg.Type != ResultTypeNumber {
		return MakeErrorResult("EVEN() requires number argument")
	}

	sign := math.Signbit(vArg.ValueNumber)
	m, r := math.Modf(vArg.ValueNumber / 2)
	v := m * 2
	if r != 0 {
		if !sign {
			v += 2
		} else {
			v -= 2
		}
	}
	return MakeNumberResult(v)
}

func fact(f float64) float64 {
	res := float64(1)
	for i := float64(2); i <= f; i++ {
		res *= i
	}
	return res
}

// Fact is an implementation of the excel FACT function which returns the
// factorial of a positive numeric input.
func Fact(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("FACT() accepts a single numeric argument")
	}
	vArg := args[0].AsNumber()
	if vArg.Type != ResultTypeNumber {
		return MakeErrorResult("FACT() accepts a single numeric argument")
	}
	if vArg.ValueNumber < 0 {
		return MakeErrorResult("FACT() accepts only positive arguments")
	}
	return MakeNumberResult(fact(vArg.ValueNumber))
}

// FactDouble is an implementation of the excel FACTDOUBLE function which
// returns the double factorial of a positive numeric input.
func FactDouble(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("FACTDOUBLE() accepts a single numeric argument")
	}
	vArg := args[0].AsNumber()
	if vArg.Type != ResultTypeNumber {
		return MakeErrorResult("FACTDOUBLE() accepts a single numeric argument")
	}
	if vArg.ValueNumber < 0 {
		return MakeErrorResult("FACTDOUBLE() accepts only positive arguments")
	}

	res := float64(1)
	v := math.Trunc(vArg.ValueNumber)
	for i := v; i > 1; i -= 2 {
		res *= i
	}
	return MakeNumberResult(res)
}

// FloorMath implements _xlfn.FLOOR.MATH which rounds numbers down to the
// nearest multiple of the second argument, toward or away from zero as
// specified by the third argument.
func FloorMath(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("FLOOR.MATH() requires at least one argument")
	}
	if len(args) > 3 {
		return MakeErrorResult("FLOOR.MATH() allows at most three arguments")
	}
	// number to round
	number := args[0].AsNumber()
	if number.Type != ResultTypeNumber {
		return MakeErrorResult("first argument to FLOOR.MATH() must be a number")
	}

	// significance
	significance := float64(1)
	if number.ValueNumber < 0 {
		significance = -1
	}
	if len(args) > 1 {
		sigArg := args[1].AsNumber()
		if sigArg.Type != ResultTypeNumber {
			return MakeErrorResult("second argument to FLOOR.MATH() must be a number")
		}
		significance = sigArg.ValueNumber
	}

	// round direction
	direction := float64(1)
	if len(args) > 2 {
		dirArg := args[2].AsNumber()
		if dirArg.Type != ResultTypeNumber {
			return MakeErrorResult("third argument to FLOOR.MATH() must be a number")
		}
		direction = dirArg.ValueNumber
	}

	if len(args) == 1 {
		return MakeNumberResult(math.Floor(number.ValueNumber))
	}

	v := number.ValueNumber
	v, res := math.Modf(v / significance)
	if res != 0 && number.ValueNumber < 0 && direction > 0 {
		v++
	}
	return MakeNumberResult(v * significance)
}

// Floor is an implementation of the FlOOR function.
func Floor(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("FLOOR() requires two arguments")
	}
	// number to round
	number := args[0].AsNumber()
	if number.Type != ResultTypeNumber {
		return MakeErrorResult("first argument to FLOOR() must be a number")
	}

	// significance
	var significance float64
	sigArg := args[1].AsNumber()
	if sigArg.Type != ResultTypeNumber {
		return MakeErrorResult("second argument to FLOOR() must be a number")
	}

	significance = sigArg.ValueNumber
	if significance < 0 && number.ValueNumber >= 0 {
		return MakeErrorResultType(ErrorTypeNum, "invalid arguments to FLOOR")
	}

	v := number.ValueNumber
	v, res := math.Modf(v / significance)
	if res != 0 {
		if number.ValueNumber < 0 && res < 0 {
			v--
		}
	}
	return MakeNumberResult(v * significance)
}

// FloorPrecise is an implementation of the FlOOR.PRECISE function.
func FloorPrecise(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("FLOOR.PRECISE() requires at least one argument")
	}
	if len(args) > 2 {
		return MakeErrorResult("FLOOR.PRECISE() allows at most two arguments")
	}
	// number to round
	number := args[0].AsNumber()
	if number.Type != ResultTypeNumber {
		return MakeErrorResult("first argument to FLOOR.PRECISE() must be a number")
	}

	// significance
	significance := float64(1)
	if number.ValueNumber < 0 {
		significance = -1
	}
	if len(args) > 1 {
		sigArg := args[1].AsNumber()
		if sigArg.Type != ResultTypeNumber {
			return MakeErrorResult("second argument to FLOOR.PRECISE() must be a number")
		}
		// don't care about sign of significance
		significance = math.Abs(sigArg.ValueNumber)
	}

	if len(args) == 1 {
		return MakeNumberResult(math.Floor(number.ValueNumber))
	}

	v := number.ValueNumber
	v, res := math.Modf(v / significance)
	if res != 0 {
		if number.ValueNumber < 0 {
			v--
		}
	}
	return MakeNumberResult(v * significance)
}

func gcd(a, b float64) float64 {
	a = math.Trunc(a)
	b = math.Trunc(b)
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}

// GCD implements the Excel GCD() function which returns the greatest common
// divisor of a range of numbers.
func GCD(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("GCD() requires at least one argument")
	}

	numbers := []float64{}
	for _, arg := range args {
		switch arg.Type {
		case ResultTypeString:
			na := arg.AsNumber()
			if na.Type != ResultTypeNumber {
				return MakeErrorResult("GCD() only accepts numeric arguments")
			}
			numbers = append(numbers, na.ValueNumber)
		case ResultTypeList, ResultTypeArray:
			res := GCD(arg.ListValues())
			if res.Type != ResultTypeNumber {
				return res
			}
			numbers = append(numbers, res.ValueNumber)
		case ResultTypeNumber:
			numbers = append(numbers, arg.ValueNumber)
		case ResultTypeError:
			return arg
		default:
			return MakeErrorResult(fmt.Sprintf("GCD() unsupported argument type %s", arg.Type))
		}
	}
	if numbers[0] < 0 {
		return MakeErrorResult("GCD() only accepts positive arguments")
	}

	if len(numbers) == 1 {
		return MakeNumberResult(numbers[0])
	}
	res := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < 0 {
			return MakeErrorResult("GCD() only accepts positive arguments")
		}
		res = gcd(res, numbers[i])
	}
	return MakeNumberResult(res)
}

func lcm(a, b float64) float64 {
	a = math.Trunc(a)
	b = math.Trunc(b)
	if a == 0 && b == 0 {
		return 0
	}
	return a * b / gcd(a, b)
}

// LCM implements the Excel LCM() function which returns the least common
// multiple of a range of numbers.
func LCM(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("LCM() requires at least one argument")
	}

	numbers := []float64{}
	for _, arg := range args {
		switch arg.Type {
		case ResultTypeString:
			na := arg.AsNumber()
			if na.Type != ResultTypeNumber {
				return MakeErrorResult("LCM() only accepts numeric arguments")
			}
			numbers = append(numbers, na.ValueNumber)
		case ResultTypeList:
			res := LCM(arg.ValueList)
			if res.Type != ResultTypeNumber {
				return res
			}
			numbers = append(numbers, res.ValueNumber)
		case ResultTypeNumber:
			numbers = append(numbers, arg.ValueNumber)
		case ResultTypeError:
			return arg
		}
	}
	if numbers[0] < 0 {
		return MakeErrorResult("LCM() only accepts positive arguments")
	}

	if len(numbers) == 1 {
		return MakeNumberResult(numbers[0])
	}
	res := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < 0 {
			return MakeErrorResult("LCM() only accepts positive arguments")
		}
		res = lcm(res, numbers[i])
	}
	return MakeNumberResult(res)
}

// Int is an implementation of the Excel INT() function that rounds a number
// down to an integer.
func Int(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("INT() requires a single numeric argument")
	}
	nArg := args[0].AsNumber()
	if nArg.Type != ResultTypeNumber {
		return MakeErrorResult("INT() requires a single numeric argument")
	}
	trunc, rem := math.Modf(nArg.ValueNumber)
	if rem < 0 {
		trunc--
	}
	return MakeNumberResult(trunc)
}

// Log implements the Excel LOG function which returns the log of a number. By
// default the result is base 10, however the second argument to the function
// can specify a different base.
func Log(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("LOG() requires at least one numeric argument")
	}
	if len(args) > 2 {
		return MakeErrorResult("LOG() accepts a maximum of two arguments")
	}
	nArg := args[0].AsNumber()
	if nArg.Type != ResultTypeNumber {
		return MakeErrorResult("LOG() requires at least one numeric argument")
	}
	base := 10.0
	if len(args) > 1 {
		bArg := args[1].AsNumber()
		if bArg.Type != ResultTypeNumber {
			return MakeErrorResult("LOG() requires second argument to be numeric")
		}
		base = args[1].ValueNumber
	}
	if nArg.ValueNumber == 0 {
		return MakeErrorResult("LOG() requires first argument to be non-zero")
	}
	if base == 0 {
		return MakeErrorResult("LOG() requires second argument to be non-zero")
	}

	return MakeNumberResult(math.Log(nArg.ValueNumber) / math.Log(base))

}

func minor(sqMtx [][]Result, idx int) [][]Result {
	ret := [][]Result{}
	for i := range sqMtx {
		if i == 0 {
			continue
		}
		row := []Result{}
		for j := range sqMtx {
			if j == idx {
				continue
			}
			row = append(row, sqMtx[i][j])
		}
		ret = append(ret, row)
	}
	return ret
}
func det(sqMtx [][]Result) float64 {
	// two by two
	if len(sqMtx) == 2 {
		m00 := sqMtx[0][0].AsNumber()
		m01 := sqMtx[0][1].AsNumber()
		m10 := sqMtx[1][0].AsNumber()
		m11 := sqMtx[1][1].AsNumber()
		if m00.Type != ResultTypeNumber || m01.Type != ResultTypeNumber ||
			m10.Type != ResultTypeNumber || m11.Type != ResultTypeNumber {
			return math.NaN()
		}
		return m00.ValueNumber*m11.ValueNumber -
			m10.ValueNumber*m01.ValueNumber
	}

	res := float64(0)
	sgn := float64(1)
	for j := range sqMtx {
		res += sgn * sqMtx[0][j].ValueNumber * det(minor(sqMtx, j))
		sgn *= -1
	}
	return res
}

// MDeterm is an implementation of the Excel MDETERM which finds the determinant
// of a matrix.
func MDeterm(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("MDETERM() requires a single array argument")
	}

	mtx := args[0]
	if mtx.Type != ResultTypeArray {
		return MakeErrorResult("MDETERM() requires a single array argument")
	}

	numRows := len(mtx.ValueArray)
	for _, row := range mtx.ValueArray {
		if len(row) != numRows {
			return MakeErrorResult("MDETERM() requires a square matrix")
		}
	}
	return MakeNumberResult(det(mtx.ValueArray))
}

// Mod is an implementation of the Excel MOD function which returns the
// remainder after division. It requires two numeric argumnts.
func Mod(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("MOD() requires two numeric arguments")
	}
	n := args[0].AsNumber()
	d := args[1].AsNumber()
	if n.Type != ResultTypeNumber || d.Type != ResultTypeNumber {
		return MakeErrorResult("MOD() requires two numeric arguments")
	}
	if d.ValueNumber == 0 {
		return MakeErrorResultType(ErrorTypeDivideByZero, "MOD() divide by zero")
	}

	// Per MS page, MOD(n, d) = n - d*INT(n/d)
	// where INT is trunc in:
	trunc, rem := math.Modf(n.ValueNumber / d.ValueNumber)
	if rem < 0 {
		trunc--
	}
	return MakeNumberResult(n.ValueNumber - d.ValueNumber*trunc)
}

// Mround is an implementation of the Excel MROUND function.  It is not a
// generic rounding function and has some oddities to match Excel's behavior.
func Mround(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("MROUND() requires two numeric arguments")
	}
	// number to round
	number := args[0].AsNumber()
	if number.Type != ResultTypeNumber {
		return MakeErrorResult("first argument to MROUND() must be a number")
	}

	// significance
	significance := float64(1)
	sigArg := args[1].AsNumber()
	if sigArg.Type != ResultTypeNumber {
		return MakeErrorResult("second argument to MROUND() must be a number")
	}
	significance = sigArg.ValueNumber

	if significance < 0 && number.ValueNumber > 0 ||
		significance > 0 && number.ValueNumber < 0 {
		return MakeErrorResult("MROUND() argument signs must match")
	}

	v := number.ValueNumber
	v, res := math.Modf(v / significance)
	if math.Trunc(res+0.5) > 0 {
		v++
	}
	return MakeNumberResult(v * significance)
}

func multinomial(args []Result) (float64, float64, Result) {
	num := 0.0
	denom := 1.0
	for _, arg := range args {
		switch arg.Type {
		case ResultTypeNumber:
			num += arg.ValueNumber
			denom *= fact(arg.ValueNumber)
		case ResultTypeList, ResultTypeArray:
			n, d, e := multinomial(arg.ListValues())
			num += n
			denom *= fact(d)
			if e.Type == ResultTypeError {
				return 0, 0, e
			}
		case ResultTypeString:
			return 0, 0, MakeErrorResult("MULTINOMIAL() requires numeric arguments")
		case ResultTypeError:
			return 0, 0, arg
		}
	}
	return num, denom, empty
}

// Multinomial implements the excel MULTINOMIAL function.
func Multinomial(args []Result) Result {
	if len(args) < 1 {
		return MakeErrorResult("MULTINOMIAL() requires at least one numeric input")
	}
	num, denom, err := multinomial(args)
	if err.Type == ResultTypeError {
		return err
	}
	return MakeNumberResult(fact(num) / denom)
}

// Munit is an implementation of the Excel MUNIT function that returns an
// identity matrix.
func Munit(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("MUNIT() requires one numeric input")
	}
	dim := args[0].AsNumber()
	if dim.Type != ResultTypeNumber {
		return MakeErrorResult("MUNIT() requires one numeric input")
	}
	dimInt := int(dim.ValueNumber)
	mtx := make([][]Result, 0, dimInt)
	for i := 0; i < dimInt; i++ {
		row := make([]Result, dimInt)
		for j := 0; j < dimInt; j++ {
			if i == j {
				row[j] = MakeNumberResult(1.0)
			} else {
				row[j] = MakeNumberResult(0.0)
			}
		}
		mtx = append(mtx, row)
	}
	return MakeArrayResult(mtx)
}

// Odd is an implementation of the Excel ODD() that rounds a number to the
// nearest odd integer.
func Odd(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("ODD() requires one argument")
	}
	vArg := args[0].AsNumber()
	if vArg.Type != ResultTypeNumber {
		return MakeErrorResult("ODD() requires number argument")
	}

	sign := math.Signbit(vArg.ValueNumber)
	m, r := math.Modf((vArg.ValueNumber - 1) / 2)
	v := m*2 + 1
	if r != 0 {
		if !sign {
			v += 2
		} else {
			v -= 2
		}
	}
	return MakeNumberResult(v)
}

// Pi is an implementation of the Excel Pi() function that just returns the Pi
// constant.
func Pi(args []Result) Result {
	if len(args) != 0 {
		return MakeErrorResult("PI() accepts no arguments")
	}
	return MakeNumberResult(math.Pi)
}

// Power is an implementation of the Excel POWER function that raises a number
// to a power. It requires two numeric arguments.
func Power(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("POWER() requires two numeric arguments")
	}

	number := args[0].AsNumber()
	if number.Type != ResultTypeNumber {
		return MakeErrorResult("first argument to POWER() must be a number")
	}

	exp := args[1].AsNumber()
	if exp.Type != ResultTypeNumber {
		return MakeErrorResult("second argument to POWER() must be a number")
	}

	return MakeNumberResult(math.Pow(number.ValueNumber, exp.ValueNumber))
}

// Product is an implementation of the Excel PRODUCT() function.
func Product(args []Result) Result {
	res := 1.0
	for _, a := range args {
		a = a.AsNumber()
		switch a.Type {
		case ResultTypeNumber:
			res *= a.ValueNumber
		case ResultTypeList, ResultTypeArray:
			subSum := Product(a.ListValues())
			if subSum.Type != ResultTypeNumber {
				return subSum
			}
			res *= subSum.ValueNumber
		case ResultTypeString:
			// treated as zero by Excel
		case ResultTypeError:
			return a
		case ResultTypeEmpty:
			// skip
		default:
			return MakeErrorResult(fmt.Sprintf("unhandled PRODUCT() argument type %s", a.Type))
		}
	}
	return MakeNumberResult(res)
}

// Quotient is an implementation of the Excel QUOTIENT function that returns the
// integer portion of division.
func Quotient(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("QUOTIENT() requires two numeric arguments")
	}
	arg1 := args[0].AsNumber()
	arg2 := args[1].AsNumber()
	if arg1.Type != ResultTypeNumber || arg2.Type != ResultTypeNumber {
		return MakeErrorResult("QUOTIENT() requires two numeric arguments")
	}
	if arg2.ValueNumber == 0 {
		return MakeErrorResultType(ErrorTypeDivideByZero, "QUOTIENT() divide by zero")
	}

	return MakeNumberResult(math.Trunc(arg1.ValueNumber / arg2.ValueNumber))
}

// Radians is an implementation of the Excel function RADIANS() that converts
// degrees to radians.
func Radians(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("RADIANS() requires one argument")
	}
	vArg := args[0].AsNumber()
	if vArg.Type != ResultTypeNumber {
		return MakeErrorResult("RADIANS() requires number argument")
	}

	return MakeNumberResult(math.Pi / 180.0 * vArg.ValueNumber)
}

// Rand is an implementation of the Excel RAND() function that returns random
// numbers in the range [0,1).
func Rand(args []Result) Result {
	if len(args) != 0 {
		return MakeErrorResult("RAND() accepts no arguments")
	}
	return MakeNumberResult(rnd.Float64())
}

// RandBetween is an implementation of the Excel RANDBETWEEN() function that returns a random
// integer in the range specified.
func RandBetween(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("RANDBETWEEN() requires two numeric arguments")
	}
	arg1 := args[0].AsNumber()
	arg2 := args[1].AsNumber()
	if arg1.Type != ResultTypeNumber || arg2.Type != ResultTypeNumber {
		return MakeErrorResult("RANDBETWEEN() requires two numeric arguments")
	}
	if arg2.ValueNumber < arg1.ValueNumber {
		return MakeErrorResult("RANDBETWEEN() requires second argument to be larger")
	}
	bottom := int64(arg1.ValueNumber)
	top := int64(arg2.ValueNumber)
	return MakeNumberResult(float64(rnd.Int63n(top-bottom+1) + bottom))
}

type ri struct {
	n float64
	s string
}

var r1tables = []ri{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

var r2tables = []ri{
	{1000, "M"},
	{950, "LM"},
	{900, "CM"},
	{500, "D"},
	{450, "LD"},
	{400, "CD"},
	{100, "C"},
	{95, "VC"},
	{90, "XC"},
	{50, "L"},
	{45, "VL"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

var r3tables = []ri{
	{1000, "M"},
	{990, "XM"},
	{950, "LM"},
	{900, "CM"},
	{500, "D"},
	{490, "XD"},
	{450, "LD"},
	{400, "CD"},
	{100, "C"},
	{99, "IC"},
	{90, "XC"},
	{50, "L"},
	{45, "VL"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

var r4tables = []ri{
	{1000, "M"},
	{995, "VM"},
	{990, "XM"},
	{950, "LM"},
	{900, "CM"},
	{500, "D"},
	{495, "VD"},
	{490, "XD"},
	{450, "LD"},
	{400, "CD"},
	{100, "C"},
	{99, "IC"},
	{90, "XC"},
	{50, "L"},
	{45, "VL"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

var r5tables = []ri{
	{1000, "M"},
	{999, "IM"},
	{995, "VM"},
	{990, "XM"},
	{950, "LM"},
	{900, "CM"},
	{500, "D"},
	{499, "ID"},
	{495, "VD"},
	{490, "XD"},
	{450, "LD"},
	{400, "CD"},
	{100, "C"},
	{99, "IC"},
	{90, "XC"},
	{50, "L"},
	{45, "VL"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// Roman is an implementation of the Excel ROMAN function that convers numbers
// to roman numerals in one of 5 formats.
func Roman(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("ROMAN() requires at least one numeric argument")
	}
	if len(args) > 2 {
		return MakeErrorResult("ROMAN() requires at most two numeric arguments")
	}

	nArg := args[0].AsNumber()
	if nArg.Type != ResultTypeNumber {
		return MakeErrorResult("ROMAN() requires at least one numeric argument")
	}
	format := 0
	if len(args) > 1 {
		fmtArg := args[1]
		if fmtArg.Type != ResultTypeNumber {
			return MakeErrorResult("ROMAN() requires second argument to be numeric")
		}
		format = int(fmtArg.ValueNumber)
		if format < 0 {
			format = 0
		} else if format > 4 {
			format = 4
		}
	}
	dt := r1tables
	switch format {
	case 1:
		dt = r2tables
	case 2:
		dt = r3tables
	case 3:
		dt = r4tables
	case 4:
		dt = r5tables
	}
	v := math.Trunc(nArg.ValueNumber)
	buf := bytes.Buffer{}
	for _, r := range dt {
		for v >= r.n {
			buf.WriteString(r.s)
			v -= r.n
		}
	}
	return MakeStringResult(buf.String())
}

type rmode byte

const (
	closest rmode = iota
	down
	up
)

// Round is an implementation of the Excel ROUND function that rounds a number
// to a specified number of digits.
func round(args []Result, mode rmode) Result {
	if len(args) != 2 {
		return MakeErrorResult("ROUND() requires two numeric arguments")
	}
	// number to round
	number := args[0].AsNumber()
	if number.Type != ResultTypeNumber {
		return MakeErrorResult("first argument to ROUND() must be a number")
	}

	digitArg := args[1].AsNumber()
	if digitArg.Type != ResultTypeNumber {
		return MakeErrorResult("second argument to ROUND() must be a number")
	}
	digits := digitArg.ValueNumber

	v := number.ValueNumber

	significance := 1.0
	if digits > 0 {
		significance = math.Pow(1/10.0, digits)
	} else {
		significance = math.Pow(10.0, -digits)
	}

	v, res := math.Modf(v / significance)
	switch mode {
	case closest:
		const eps = 0.499999999
		if res >= eps {
			v++
		} else if res <= -eps {
			v--
		}
	case down:
		// do nothing, truncates
	case up:
		if res > 0 {
			v++
		} else if res < 0 {
			v--
		}
	}

	return MakeNumberResult(v * significance)
}

// Round is an implementation of the Excel ROUND function that rounds a number
// to a specified number of digits.
func Round(args []Result) Result {
	return round(args, closest)
}

// RoundDown is an implementation of the Excel ROUNDDOWN function that rounds a number
// down to a specified number of digits.
func RoundDown(args []Result) Result {
	return round(args, down)
}

// RoundUp is an implementation of the Excel ROUNDUP function that rounds a number
// up to a specified number of digits.
func RoundUp(args []Result) Result {
	return round(args, up)
}

// SeriesSum implements the Excel SERIESSUM function.
func SeriesSum(args []Result) Result {
	if len(args) != 4 {
		return MakeErrorResult("SERIESSUM() requires 4 arguments")
	}
	x := args[0].AsNumber()
	n := args[1].AsNumber()
	m := args[2].AsNumber()
	coeffs := args[3].ListValues()
	if x.Type != ResultTypeNumber || n.Type != ResultTypeNumber || m.Type != ResultTypeNumber {
		return MakeErrorResult("SERIESSUM() requires first three arguments to be numeric")
	}
	res := float64(0)
	for i, c := range coeffs {
		res += c.ValueNumber * math.Pow(x.ValueNumber, n.ValueNumber+float64(i)*m.ValueNumber)
	}
	return MakeNumberResult(res)
}

func Sign(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("SIGN() requires one argument")
	}
	vArg := args[0].AsNumber()
	if vArg.Type != ResultTypeNumber {
		return MakeErrorResult("SIGN() requires a numeric argument")
	}
	if vArg.ValueNumber < 0 {
		return MakeNumberResult(-1)
	} else if vArg.ValueNumber > 0 {
		return MakeNumberResult(1)
	}
	return MakeNumberResult(0)
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

// SumIf implements the SUMIF function.
func SumIf(args []Result) Result {
	if len(args) < 3 {
		return MakeErrorResult("SUMIF requires three arguments")
	}

	arrResult := args[0]
	if arrResult.Type != ResultTypeArray && arrResult.Type != ResultTypeList {
		return MakeErrorResult("SUMIF requires first argument of type array")
	}
	arr := arrayFromRange(arrResult)

	sumArrResult := args[2]
	if sumArrResult.Type != ResultTypeArray && sumArrResult.Type != ResultTypeList {
		return MakeErrorResult("SUMIF requires last argument of type array")
	}
	sumArr := arrayFromRange(sumArrResult)

	criteria := parseCriteria(args[1])
	sum := 0.0
	for ir, r := range arr {
		for ic, value := range r {
			if compare(value, criteria) {
				sum += sumArr[ir][ic].ValueNumber
			}
		}
	}
	return MakeNumberResult(sum)
}

// SumIfs implements the SUMIFS function.
func SumIfs(args []Result) Result {
	errorResult := checkIfsRanges(args, true, "SUMIFS")
	if errorResult.Type != ResultTypeEmpty {
		return errorResult
	}
	match := getIfsMatch(args[1:])
	sum := 0.0
	sumArr := arrayFromRange(args[0])
	for _, indexes := range match {
		sum += sumArr[indexes.rowIndex][indexes.colIndex].ValueNumber
	}
	return MakeNumberResult(float64(sum))
}

// SumProduct is an implementation of the Excel SUMPRODUCT() function.
func SumProduct(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("SUMPRODUCT() requires at least one argument")
	}
	t := args[0].Type
	for _, a := range args {
		if a.Type != t {
			return MakeErrorResult("SUMPRODUCT() requires all arguments of the same type")
		}
	}
	switch t {
	case ResultTypeNumber:
		return Product(args)
	case ResultTypeList, ResultTypeArray:
		n := len(args[0].ListValues())
		res := make([]float64, n)
		for i := range res {
			res[i] = 1.0
		}
		for _, a := range args {
			if len(a.ListValues()) != n {
				return MakeErrorResult("SUMPRODUCT() requires all arguments to have the same dimension")
			}
			for i, v := range a.ListValues() {
				v = v.AsNumber()
				if v.Type != ResultTypeNumber {
					return MakeErrorResult("SUMPRODUCT() requires all arguments to be numeric")
				}
				res[i] = res[i] * v.ValueNumber
			}
		}
		v := 0.0
		for _, r := range res {
			v += r
		}
		return MakeNumberResult(v)

	}
	return MakeNumberResult(1.0)
}

// SumSquares is an implementation of the Excel SUMSQ() function.
func SumSquares(args []Result) Result {
	// Sum returns zero with no arguments
	res := MakeNumberResult(0)
	for _, a := range args {
		a = a.AsNumber()
		switch a.Type {
		case ResultTypeNumber:
			res.ValueNumber += a.ValueNumber * a.ValueNumber
		case ResultTypeList, ResultTypeArray:
			subSum := SumSquares(a.ListValues())
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
			return MakeErrorResult(fmt.Sprintf("unhandled SUMSQUARES() argument type %s", a.Type))
		}
	}
	return res
}

func Trunc(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("TRUNC() requires at least one numeric arguments")
	}
	// number to truncate
	number := args[0].AsNumber()
	if number.Type != ResultTypeNumber {
		return MakeErrorResult("first argument to TRUNC() must be a number")
	}

	digits := float64(0)
	if len(args) > 1 {
		digitArg := args[1].AsNumber()
		if digitArg.Type != ResultTypeNumber {
			return MakeErrorResult("second argument to TRUNC() must be a number")
		}
		digits = digitArg.ValueNumber
	}

	v := number.ValueNumber

	significance := 1.0
	if digits >= 0 {
		significance = math.Pow(1/10.0, digits)
	} else {
		// Excel returns zero for this case
		return MakeNumberResult(0)
	}

	v, res := math.Modf(v / significance)

	eps := 0.99999
	if res > eps {
		v++
	} else if res < -eps {
		v--
	}
	_ = res
	return MakeNumberResult(v * significance)
}
