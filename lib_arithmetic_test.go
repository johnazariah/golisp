package golisp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var arithmetic = &(ArithmeticLibrary{})

func TestAdd(t *testing.T) {
	tests := [...]struct {
		desc     string
		input    []Variant
		expected Variant
	}{
		{
			desc: "single int value",
			input: []Variant{
				{VariantType: VAR_INT, VariantValue: int64(1)},
			},
			expected: Variant{VariantType: VAR_INT, VariantValue: int64(1)},
		},
		{
			desc: "single float value",
			input: []Variant{
				{VariantType: VAR_FLOAT, VariantValue: float64(3.14)},
			},
			expected: Variant{VariantType: VAR_FLOAT, VariantValue: float64(3.14)},
		},
		{
			desc: "int + int value",
			input: []Variant{
				{VariantType: VAR_INT, VariantValue: int64(1)},
				{VariantType: VAR_INT, VariantValue: int64(2)},
				{VariantType: VAR_INT, VariantValue: int64(3)},
				{VariantType: VAR_INT, VariantValue: int64(4)},
			},
			expected: Variant{VariantType: VAR_INT, VariantValue: int64(10)},
		},
		{
			desc: "float + float value",
			input: []Variant{
				{VariantType: VAR_FLOAT, VariantValue: float64(3.14)},
				{VariantType: VAR_FLOAT, VariantValue: float64(3.14)},
			},
			expected: Variant{VariantType: VAR_FLOAT, VariantValue: float64(6.28)},
		},
		{
			desc: "float + int value",
			input: []Variant{
				{VariantType: VAR_INT, VariantValue: int64(1)},
				{VariantType: VAR_FLOAT, VariantValue: float64(3.14)},
			},
			expected: Variant{VariantType: VAR_FLOAT, VariantValue: float64(4.140000000000001)},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := arithmetic.add(test.input)
			assert.Equal(t, test.expected, actual, "computation error")
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := [...]struct {
		desc     string
		input    []Variant
		expected Variant
	}{
		{
			desc: "single int value",
			input: []Variant{
				{VariantType: VAR_INT, VariantValue: int64(1)},
			},
			expected: Variant{VariantType: VAR_INT, VariantValue: int64(-1)},
		},
		{
			desc: "single float value",
			input: []Variant{
				{VariantType: VAR_FLOAT, VariantValue: float64(3.14)},
			},
			expected: Variant{VariantType: VAR_FLOAT, VariantValue: float64(-3.14)},
		},
		{
			desc: "int - int invalid",
			input: []Variant{
				{VariantType: VAR_INT, VariantValue: int64(1)},
				{VariantType: VAR_INT, VariantValue: int64(2)},
				{VariantType: VAR_INT, VariantValue: int64(3)},
				{VariantType: VAR_INT, VariantValue: int64(4)},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: buildArityError_1or2("sub")},
		},
		{
			desc: "int - int valid",
			input: []Variant{
				{VariantType: VAR_INT, VariantValue: int64(4)},
				{VariantType: VAR_INT, VariantValue: int64(2)},
			},
			expected: Variant{VariantType: VAR_INT, VariantValue: int64(2)},
		},
		{
			desc: "float - float value",
			input: []Variant{
				{VariantType: VAR_FLOAT, VariantValue: float64(3.14)},
				{VariantType: VAR_FLOAT, VariantValue: float64(3.14)},
			},
			expected: Variant{VariantType: VAR_FLOAT, VariantValue: float64(0.0)},
		},
		{
			desc: "float - int value",
			input: []Variant{
				{VariantType: VAR_FLOAT, VariantValue: float64(3.14)},
				{VariantType: VAR_INT, VariantValue: int64(1)},
			},
			expected: Variant{VariantType: VAR_FLOAT, VariantValue: float64(2.14)},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := arithmetic.subtract(test.input)
			assert.Equal(t, test.expected, actual, "computation error")
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := [...]struct {
		desc     string
		input    []Variant
		expected Variant
	}{
		{
			desc: "single int value",
			input: []Variant{
				{VariantType: VAR_INT, VariantValue: int64(1)},
			},
			expected: Variant{VariantType: VAR_INT, VariantValue: int64(1)},
		},
		{
			desc: "single float value",
			input: []Variant{
				{VariantType: VAR_FLOAT, VariantValue: float64(3.14)},
			},
			expected: Variant{VariantType: VAR_FLOAT, VariantValue: float64(3.14)},
		},
		{
			desc: "int * int value",
			input: []Variant{
				{VariantType: VAR_INT, VariantValue: int64(1)},
				{VariantType: VAR_INT, VariantValue: int64(2)},
				{VariantType: VAR_INT, VariantValue: int64(3)},
				{VariantType: VAR_INT, VariantValue: int64(4)},
			},
			expected: Variant{VariantType: VAR_INT, VariantValue: int64(24)},
		},
		{
			desc: "float * float value",
			input: []Variant{
				{VariantType: VAR_FLOAT, VariantValue: float64(3.14)},
				{VariantType: VAR_FLOAT, VariantValue: float64(3.14)},
			},
			expected: Variant{VariantType: VAR_FLOAT, VariantValue: float64(9.8596)},
		},
		{
			desc: "float * int value",
			input: []Variant{
				{VariantType: VAR_INT, VariantValue: int64(2)},
				{VariantType: VAR_FLOAT, VariantValue: float64(3.14)},
			},
			expected: Variant{VariantType: VAR_FLOAT, VariantValue: float64(6.28)},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := arithmetic.multiply(test.input)
			assert.Equal(t, test.expected, actual, "computation error")
		})
	}
}

func TestDivide(t *testing.T) {
	tests := [...]struct {
		desc     string
		input    []Variant
		expected Variant
	}{
		{
			desc: "single int value",
			input: []Variant{
				{VariantType: VAR_INT, VariantValue: int64(1)},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: buildExactArityError(2, "div")},
		},
		{
			desc: "single float value",
			input: []Variant{
				{VariantType: VAR_FLOAT, VariantValue: float64(3.14)},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: buildExactArityError(2, "div")},
		},
		{
			desc: "int - int invalid",
			input: []Variant{
				{VariantType: VAR_INT, VariantValue: int64(1)},
				{VariantType: VAR_INT, VariantValue: int64(2)},
				{VariantType: VAR_INT, VariantValue: int64(3)},
				{VariantType: VAR_INT, VariantValue: int64(4)},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: buildExactArityError(2, "div")},
		},
		{
			desc: "float / float value",
			input: []Variant{
				{VariantType: VAR_FLOAT, VariantValue: float64(3.14)},
				{VariantType: VAR_FLOAT, VariantValue: float64(3.14)},
			},
			expected: Variant{VariantType: VAR_FLOAT, VariantValue: float64(1.0)},
		},
		{
			desc: "float / int value",
			input: []Variant{
				{VariantType: VAR_FLOAT, VariantValue: float64(6.28)},
				{VariantType: VAR_INT, VariantValue: int64(2)},
			},
			expected: Variant{VariantType: VAR_FLOAT, VariantValue: float64(3.14)},
		},
		{
			desc: "div by zero",
			input: []Variant{
				{VariantType: VAR_FLOAT, VariantValue: float64(6.28)},
				{VariantType: VAR_INT, VariantValue: int64(0)},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: buildDivideByZeroError()},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := arithmetic.divide(test.input)
			assert.Equal(t, test.expected, actual, "computation error")
		})
	}
}

func TestPower(t *testing.T) {
	tests := [...]struct {
		desc     string
		input    []Variant
		expected Variant
	}{
		{
			desc: "single int value",
			input: []Variant{
				{VariantType: VAR_INT, VariantValue: int64(1)},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: buildExactArityError(2, "pow")},
		},
		{
			desc: "single float value",
			input: []Variant{
				{VariantType: VAR_FLOAT, VariantValue: float64(3.14)},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: buildExactArityError(2, "pow")},
		},
		{
			desc: "int[] invalid",
			input: []Variant{
				{VariantType: VAR_INT, VariantValue: int64(1)},
				{VariantType: VAR_INT, VariantValue: int64(2)},
				{VariantType: VAR_INT, VariantValue: int64(3)},
				{VariantType: VAR_INT, VariantValue: int64(4)},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: buildExactArityError(2, "pow")},
		},
		{
			desc: "float ^ float value",
			input: []Variant{
				{VariantType: VAR_FLOAT, VariantValue: float64(3.14)},
				{VariantType: VAR_FLOAT, VariantValue: float64(2.0)},
			},
			expected: Variant{VariantType: VAR_FLOAT, VariantValue: float64(9.8596)},
		},
		{
			desc: "float ^ int value",
			input: []Variant{
				{VariantType: VAR_FLOAT, VariantValue: float64(3.14)},
				{VariantType: VAR_INT, VariantValue: int64(3)},
			},
			expected: Variant{VariantType: VAR_FLOAT, VariantValue: float64(30.959144000000002)},
		},
		{
			desc: "int ^ int value",
			input: []Variant{
				{VariantType: VAR_FLOAT, VariantValue: int64(3)},
				{VariantType: VAR_INT, VariantValue: int64(2)},
			},
			expected: Variant{VariantType: VAR_FLOAT, VariantValue: float64(9.00)},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := arithmetic.power(test.input)
			assert.Equal(t, test.expected, actual, "computation error")
		})
	}
}
