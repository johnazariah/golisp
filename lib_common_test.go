package golisp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumArity(t *testing.T) {
	tests := [...]struct {
		desc     string
		input    []Variant
		arity    int
		expected error
	}{
		{
			desc: "less than",
			input: []Variant{
				{VariantType: VAR_UNKNOWN},
			},
			arity:    2,
			expected: nil,
		},
		{
			desc: "exact",
			input: []Variant{
				{VariantType: VAR_UNKNOWN},
				{VariantType: VAR_UNKNOWN},
			},
			arity:    2,
			expected: nil,
		},
		{
			desc: "more than",
			input: []Variant{
				{VariantType: VAR_UNKNOWN},
				{VariantType: VAR_UNKNOWN},
				{VariantType: VAR_UNKNOWN},
				{VariantType: VAR_UNKNOWN},
				{VariantType: VAR_UNKNOWN},
			},
			arity:    2,
			expected: fmt.Errorf("arity error: expected at most %d arguments for %q", 2, "test"),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := ensureMaximumArity(test.input, test.arity, "test")
			assert.Equal(t, test.expected, actual, "fail")
		})
	}
}

func TestMinimumArity(t *testing.T) {
	tests := [...]struct {
		desc     string
		input    []Variant
		arity    int
		expected error
	}{
		{
			desc: "less than",
			input: []Variant{
				{VariantType: VAR_UNKNOWN},
			},
			arity:    2,
			expected: fmt.Errorf("arity error: expected at least %d arguments for %q", 2, "test"),
		},
		{
			desc: "exact",
			input: []Variant{
				{VariantType: VAR_UNKNOWN},
				{VariantType: VAR_UNKNOWN},
			},
			arity:    2,
			expected: nil,
		},
		{
			desc: "more than",
			input: []Variant{
				{VariantType: VAR_UNKNOWN},
				{VariantType: VAR_UNKNOWN},
				{VariantType: VAR_UNKNOWN},
				{VariantType: VAR_UNKNOWN},
				{VariantType: VAR_UNKNOWN},
			},
			arity:    2,
			expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := ensureMinimimArity(test.input, test.arity, "test")
			assert.Equal(t, test.expected, actual, "fail")
		})
	}
}

func TestExactArity(t *testing.T) {
	tests := [...]struct {
		desc     string
		input    []Variant
		arity    int
		expected error
	}{
		{
			desc: "less than",
			input: []Variant{
				{VariantType: VAR_UNKNOWN},
			},
			arity:    2,
			expected: fmt.Errorf("arity error: expected exactly %d arguments for %q", 2, "test"),
		},
		{
			desc: "exact",
			input: []Variant{
				{VariantType: VAR_UNKNOWN},
				{VariantType: VAR_UNKNOWN},
			},
			arity:    2,
			expected: nil,
		},
		{
			desc: "more than",
			input: []Variant{
				{VariantType: VAR_UNKNOWN},
				{VariantType: VAR_UNKNOWN},
				{VariantType: VAR_UNKNOWN},
				{VariantType: VAR_UNKNOWN},
				{VariantType: VAR_UNKNOWN},
			},
			arity:    2,
			expected: fmt.Errorf("arity error: expected exactly %d arguments for %q", 2, "test"),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := ensureExactArity(test.input, test.arity, "test")
			assert.Equal(t, test.expected, actual, "fail")
		})
	}
}

func TestEnsureTypeIsNotInvalid(t *testing.T) {
	tests := [...]struct {
		desc     string
		input    Variant
		expected error
	}{
		{
			desc:     "VAR_ERROR",
			input:    Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("an error")},
			expected: fmt.Errorf("an error"),
		},
		{
			desc:     "VAR_MAX",
			input:    Variant{VariantType: VAR_MAX},
			expected: fmt.Errorf("dev error: should never have type VAR_MAX"),
		},
		{
			desc:     "VAR_IDENT",
			input:    Variant{VariantType: VAR_IDENT, VariantValue: "a"},
			expected: fmt.Errorf("scope error: unresolved identifier %q", "a"),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := ensureTypeIsNotInvalid(test.input)
			assert.Equal(t, test.expected, actual, "fail")
		})
	}
}

func TestEnsureArgumentTypesMatch(t *testing.T) {
	tests := [...]struct {
		desc            string
		input           []Variant
		acceptableTypes []EnumVariantType
		forbiddenTypes  []EnumVariantType
		expected        error
	}{
		{
			desc: "invalid type",
			input: []Variant{
				{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("an error")},
			},
			expected: fmt.Errorf("an error"),
		},
		{
			desc: "forbidden type",
			input: []Variant{
				{VariantType: VAR_STRING, VariantValue: "a forbidden string"},
			},
			forbiddenTypes: []EnumVariantType{
				VAR_STRING,
			},
			expected: fmt.Errorf("type error: argument to %q can never be of type %q", "test", "VAR_STRING"),
		},
		{
			desc: "acceptable type",
			input: []Variant{
				{VariantType: VAR_INT, VariantValue: 42},
			},
			forbiddenTypes: []EnumVariantType{
				VAR_STRING,
				VAR_DATE,
			},
			acceptableTypes: []EnumVariantType{
				VAR_BOOL,
				VAR_INT,
				VAR_FLOAT,
			},
			expected: nil,
		},
		{
			desc: "no acceptable type",
			input: []Variant{
				{VariantType: VAR_INT, VariantValue: 42},
			},
			forbiddenTypes: []EnumVariantType{
				VAR_STRING,
				VAR_DATE,
			},
			acceptableTypes: []EnumVariantType{
				VAR_BOOL,
				VAR_FLOAT,
			},
			expected: fmt.Errorf("type error: argument of unacceptable type %q passed to %q", "VAR_INT", "test"),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := ensureArgumentTypesMatch(test.input, test.acceptableTypes, test.forbiddenTypes, "test")
			assert.Equal(t, test.expected, actual, "fail")
		})
	}
}

func TestGetPromotedNumberType(t *testing.T) {
	tests := [...]struct {
		desc                string
		input               []Variant
		expectedVariantType EnumVariantType
		expectedError       error
	}{
		{
			desc: "invalid type",
			input: []Variant{
				{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("an error")},
			},
			expectedError: fmt.Errorf("an error"),
		},
		{
			desc: "promote to int",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: true},
			},
			expectedVariantType: VAR_INT,
		},
		{
			desc: "resolve to int",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_INT, VariantValue: 27},
			},
			expectedVariantType: VAR_INT,
		},
		{
			desc: "resolve to float",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_INT, VariantValue: 27},
				{VariantType: VAR_FLOAT, VariantValue: 3.14},
			},
			expectedVariantType: VAR_FLOAT,
		},
		{
			desc: "unacceptable type",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_INT, VariantValue: 27},
				{VariantType: VAR_STRING, VariantValue: "pi"},
			},
			expectedError: fmt.Errorf("type error: argument of unacceptable type %q passed to %q", "VAR_STRING", "test"),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			v, e := getPromotedNumberType(test.input, "test")
			if e == nil {
				assert.Equal(t, test.expectedVariantType, v, "fail")
			} else {
				assert.Equal(t, test.expectedError, e, "fail")
			}
		})
	}
}

func TestUnaryOpNumber(t *testing.T) {
	tests := [...]struct {
		desc     string
		input    []Variant
		expected Variant
	}{
		{
			desc: "incorrect arity",
			input: []Variant{
				{VariantType: VAR_INT, VariantValue: 2},
				{VariantType: VAR_INT, VariantValue: 2},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("arity error: expected exactly 1 arguments for %q", "test")},
		},
		{
			desc: "invalid type - error",
			input: []Variant{
				{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("an error")},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("an error")},
		},
		{
			desc: "invalid type - string",
			input: []Variant{
				{VariantType: VAR_STRING, VariantValue: "pi"},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("type error: argument of unacceptable type %q passed to %q", "VAR_STRING", "test")},
		},
		{
			desc: "success - int",
			input: []Variant{
				{VariantType: VAR_INT, VariantValue: int64(2)},
			},
			expected: Variant{VariantType: VAR_INT, VariantValue: int64(2)},
		},
		{
			desc: "success - float",
			input: []Variant{
				{VariantType: VAR_FLOAT, VariantValue: float64(2.71828)},
			},
			expected: Variant{VariantType: VAR_FLOAT, VariantValue: float64(2.71828)},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			v := unaryOpNumber(test.input, func(a int64) (int64, error) { return a, nil }, func(a float64) (float64, error) { return a, nil }, "test")
			assert.Equal(t, test.expected, v, "fail")
		})
	}
}

func TestUnaryOpNumber_ErrorPassback(t *testing.T) {
	tests := [...]struct {
		desc     string
		input    []Variant
		expected Variant
	}{
		{
			desc: "incorrect arity",
			input: []Variant{
				{VariantType: VAR_INT, VariantValue: 2},
				{VariantType: VAR_INT, VariantValue: 2},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("arity error: expected exactly 1 arguments for %q", "test")},
		},
		{
			desc: "invalid type - error",
			input: []Variant{
				{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("an error")},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("an error")},
		},
		{
			desc: "invalid type - string",
			input: []Variant{
				{VariantType: VAR_STRING, VariantValue: "pi"},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("type error: argument of unacceptable type %q passed to %q", "VAR_STRING", "test")},
		},
		{
			desc: "failure - int",
			input: []Variant{
				{VariantType: VAR_INT, VariantValue: int64(2)},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("domain error")},
		},
		{
			desc: "failure - float",
			input: []Variant{
				{VariantType: VAR_FLOAT, VariantValue: float64(2.71828)},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("domain error")},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			v := unaryOpNumber(
				test.input,
				func(a int64) (int64, error) { return a, fmt.Errorf("domain error") },
				func(a float64) (float64, error) { return a, fmt.Errorf("domain error") },
				"test")
			assert.Equal(t, test.expected, v, "fail")
		})
	}
}

func TestBinaryOpNumber(t *testing.T) {
	tests := [...]struct {
		desc     string
		input    []Variant
		expected Variant
	}{
		{
			desc: "incorrect arity - 1",
			input: []Variant{
				{VariantType: VAR_INT, VariantValue: 2},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("arity error: expected exactly 2 arguments for %q", "test")},
		},
		{
			desc: "incorrect arity - 3",
			input: []Variant{
				{VariantType: VAR_INT, VariantValue: 2},
				{VariantType: VAR_INT, VariantValue: 2},
				{VariantType: VAR_INT, VariantValue: 2},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("arity error: expected exactly 2 arguments for %q", "test")},
		},
		{
			desc: "invalid type - error",
			input: []Variant{
				{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("an error")},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("an error")},
		},
		{
			desc: "invalid type - string",
			input: []Variant{
				{VariantType: VAR_STRING, VariantValue: "pi"},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("type error: argument of unacceptable type %q passed to %q", "VAR_STRING", "test")},
		},
		{
			desc: "success - int",
			input: []Variant{
				{VariantType: VAR_INT, VariantValue: int64(2)},
				{VariantType: VAR_INT, VariantValue: int64(2)},
			},
			expected: Variant{VariantType: VAR_INT, VariantValue: int64(4)},
		},
		{
			desc: "success - float",
			input: []Variant{
				{VariantType: VAR_FLOAT, VariantValue: float64(2.71828)},
				{VariantType: VAR_FLOAT, VariantValue: float64(2.71828)},
			},
			expected: Variant{VariantType: VAR_FLOAT, VariantValue: float64(5.43656)},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			v := binaryOpNumbers(
				test.input,
				func(a int64, b int64) (int64, error) { return a + b, nil },
				func(a float64, b float64) (float64, error) { return a + b, nil },
				"test")
			assert.Equal(t, test.expected, v, "fail")
		})
	}
}

func TestBinaryOpNumber_ErrorPassback(t *testing.T) {
	tests := [...]struct {
		desc     string
		input    []Variant
		expected Variant
	}{
		{
			desc: "incorrect arity",
			input: []Variant{
				{VariantType: VAR_INT, VariantValue: 2},
				{VariantType: VAR_INT, VariantValue: 2},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("arity error: expected exactly 1 arguments for %q", "test")},
		},
		{
			desc: "invalid type - error",
			input: []Variant{
				{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("an error")},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("an error")},
		},
		{
			desc: "invalid type - string",
			input: []Variant{
				{VariantType: VAR_STRING, VariantValue: "pi"},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("type error: argument of unacceptable type %q passed to %q", "VAR_STRING", "test")},
		},
		{
			desc: "failure - int",
			input: []Variant{
				{VariantType: VAR_INT, VariantValue: int64(2)},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("domain error")},
		},
		{
			desc: "failure - float",
			input: []Variant{
				{VariantType: VAR_FLOAT, VariantValue: float64(2.71828)},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("domain error")},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			v := unaryOpNumber(
				test.input,
				func(a int64) (int64, error) { return a, fmt.Errorf("domain error") },
				func(a float64) (float64, error) { return a, fmt.Errorf("domain error") },
				"test")
			assert.Equal(t, test.expected, v, "fail")
		})
	}
}
