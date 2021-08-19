package golisp

import (
	"math"
)

type ArithmeticLibrary struct {
}

func (l *ArithmeticLibrary) add(args []Variant) Variant {
	return foldNumbers(
		args,
		func(a int64, b int64) (int64, error) { return a + b, nil },
		func(a float64, b float64) (float64, error) { return a + b, nil },
		"add")
}

func (l *ArithmeticLibrary) subtract(args []Variant) Variant {
	functionName := "sub"
	switch len(args) {
	case 1:
		return unaryOpNumber(
			args,
			func(a int64) (int64, error) {
				return -a, nil
			},
			func(a float64) (float64, error) {
				return -a, nil
			},
			functionName,
		)

	case 2:
		return binaryOpNumbers(
			args,
			func(a int64, b int64) (int64, error) {
				return a - b, nil
			},
			func(a float64, b float64) (float64, error) {
				return a - b, nil
			},
			functionName)

	default:
		return Variant{VariantType: VAR_ERROR, VariantValue: buildArityError_1or2(functionName)}
	}
}

func (l *ArithmeticLibrary) multiply(args []Variant) Variant {
	return foldNumbers(
		args,
		func(a int64, b int64) (int64, error) { return a * b, nil },
		func(a float64, b float64) (float64, error) { return a * b, nil },
		"mul")
}

func (l *ArithmeticLibrary) divide(args []Variant) Variant {
	return binaryOpFloats(
		args,
		func(a float64, b float64) (float64, error) {
			if b == float64(0) {
				return 0, buildDivideByZeroError()
			}
			return a / b, nil
		},
		"div")
}

func (l *ArithmeticLibrary) power(args []Variant) Variant {
	return binaryOpFloats(
		args,
		func(a float64, b float64) (float64, error) {
			return math.Pow(a, b), nil
		},
		"pow")
}

func (l *ArithmeticLibrary) InjectFunctions(functions FunctionTable) FunctionTable {
	functions["add"] = l.add
	functions["sub"] = l.subtract
	functions["mul"] = l.multiply
	functions["div"] = l.divide
	functions["pow"] = l.power
	functions["+"] = l.add
	functions["-"] = l.subtract
	functions["*"] = l.multiply
	functions["/"] = l.divide
	functions["^"] = l.power
	return functions
}
