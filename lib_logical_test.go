package golisp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var logical = &(LogicalLibrary{})

func TestOr(t *testing.T) {
	tests := [...]struct {
		desc     string
		input    []Variant
		expected Variant
	}{
		{
			desc: "single true value",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: true},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("arity error: expected at least 2 arguments for %q", "or")},
		},
		{
			desc: "single false value",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("arity error: expected at least 2 arguments for %q", "or")},
		},
		{
			desc: "all false values",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: false},
			},
			expected: Variant{VariantType: VAR_BOOL, VariantValue: false},
		},
		{
			desc: "mixed values",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: true},
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: false},
			},
			expected: Variant{VariantType: VAR_BOOL, VariantValue: true},
		},
		{
			desc: "incorrectly typed values",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_STRING, VariantValue: "true"},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("type error: argument of unacceptable type %q passed to %q", VAR_STRING, "or")},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := logical.or(test.input)
			assert.Equal(t, test.expected, actual, "computation error")
		})
	}
}

func TestNor(t *testing.T) {
	tests := [...]struct {
		desc     string
		input    []Variant
		expected Variant
	}{
		{
			desc: "single true value",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: true},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("arity error: expected at least 2 arguments for %q", "nor")},
		},
		{
			desc: "single false value",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("arity error: expected at least 2 arguments for %q", "nor")},
		},
		{
			desc: "all false values",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: false},
			},
			expected: Variant{VariantType: VAR_BOOL, VariantValue: true},
		},
		{
			desc: "mixed values",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: true},
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: false},
			},
			expected: Variant{VariantType: VAR_BOOL, VariantValue: false},
		},
		{
			desc: "incorrectly typed values",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_STRING, VariantValue: "true"},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("type error: argument of unacceptable type %q passed to %q", VAR_STRING, "nor")},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := logical.nor(test.input)
			assert.Equal(t, test.expected, actual, "computation error")
		})
	}
}

func TestAnd(t *testing.T) {
	tests := [...]struct {
		desc     string
		input    []Variant
		expected Variant
	}{
		{
			desc: "single true value",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: true},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("arity error: expected at least 2 arguments for %q", "and")},
		},
		{
			desc: "single false value",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("arity error: expected at least 2 arguments for %q", "and")},
		},
		{
			desc: "all true values",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: true},
				{VariantType: VAR_BOOL, VariantValue: true},
				{VariantType: VAR_BOOL, VariantValue: true},
				{VariantType: VAR_BOOL, VariantValue: true},
			},
			expected: Variant{VariantType: VAR_BOOL, VariantValue: true},
		},
		{
			desc: "mixed values",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: true},
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: false},
			},
			expected: Variant{VariantType: VAR_BOOL, VariantValue: false},
		},
		{
			desc: "incorrectly typed values",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_STRING, VariantValue: "true"},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("type error: argument of unacceptable type %q passed to %q", VAR_STRING, "and")},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := logical.and(test.input)
			assert.Equal(t, test.expected, actual, "computation error")
		})
	}
}

func TestNand(t *testing.T) {
	tests := [...]struct {
		desc     string
		input    []Variant
		expected Variant
	}{
		{
			desc: "single true value",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: true},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("arity error: expected at least 2 arguments for %q", "nand")},
		},
		{
			desc: "single false value",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("arity error: expected at least 2 arguments for %q", "nand")},
		},
		{
			desc: "all true values",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: true},
				{VariantType: VAR_BOOL, VariantValue: true},
				{VariantType: VAR_BOOL, VariantValue: true},
				{VariantType: VAR_BOOL, VariantValue: true},
			},
			expected: Variant{VariantType: VAR_BOOL, VariantValue: false},
		},
		{
			desc: "mixed values",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: true},
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: false},
			},
			expected: Variant{VariantType: VAR_BOOL, VariantValue: true},
		},
		{
			desc: "incorrectly typed values",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_STRING, VariantValue: "true"},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("type error: argument of unacceptable type %q passed to %q", VAR_STRING, "nand")},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := logical.nand(test.input)
			assert.Equal(t, test.expected, actual, "computation error")
		})
	}
}

func TestNot(t *testing.T) {
	tests := [...]struct {
		desc     string
		input    []Variant
		expected Variant
	}{
		{
			desc: "single true value",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: true},
			},
			expected: Variant{VariantType: VAR_BOOL, VariantValue: false},
		},
		{
			desc: "single false value",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
			},
			expected: Variant{VariantType: VAR_BOOL, VariantValue: true},
		},
		{
			desc: "all false values",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: false},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("arity error: expected exactly 1 arguments for %q", "not")},
		},
		{
			desc: "mixed values",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: true},
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: false},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("arity error: expected exactly 1 arguments for %q", "not")},
		},
		{
			desc: "incorrectly typed values",
			input: []Variant{
				{VariantType: VAR_STRING, VariantValue: "true"},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("type error: argument of unacceptable type %q passed to %q", VAR_STRING, "not")},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := logical.not(test.input)
			assert.Equal(t, test.expected, actual, "computation error")
		})
	}
}

func TestXor(t *testing.T) {
	tests := [...]struct {
		desc     string
		input    []Variant
		expected Variant
	}{
		{
			desc: "single true value",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: true},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("arity error: expected exactly 2 arguments for %q", "xor")},
		},
		{
			desc: "single false value",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("arity error: expected exactly 2 arguments for %q", "xor")},
		},
		{
			desc: "all false values",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: false},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("arity error: expected exactly 2 arguments for %q", "xor")},
		},
		{
			desc: "mixed values",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: true},
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: false},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("arity error: expected exactly 2 arguments for %q", "xor")},
		},
		{
			desc: "same values",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: false},
			},
			expected: Variant{VariantType: VAR_BOOL, VariantValue: false},
		},
		{
			desc: "different values",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: true},
			},
			expected: Variant{VariantType: VAR_BOOL, VariantValue: true},
		},
		{
			desc: "incorrectly typed values",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_STRING, VariantValue: "true"},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("type error: argument of unacceptable type %q passed to %q", VAR_STRING, "xor")},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := logical.xor(test.input)
			assert.Equal(t, test.expected, actual, "computation error")
		})
	}
}

func TestXnor(t *testing.T) {
	tests := [...]struct {
		desc     string
		input    []Variant
		expected Variant
	}{
		{
			desc: "single true value",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: true},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("arity error: expected exactly 2 arguments for %q", "xnor")},
		},
		{
			desc: "single false value",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("arity error: expected exactly 2 arguments for %q", "xnor")},
		},
		{
			desc: "all false values",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: false},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("arity error: expected exactly 2 arguments for %q", "xnor")},
		},
		{
			desc: "mixed values",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: true},
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: false},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("arity error: expected exactly 2 arguments for %q", "xnor")},
		},
		{
			desc: "same values",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: false},
			},
			expected: Variant{VariantType: VAR_BOOL, VariantValue: true},
		},
		{
			desc: "different values",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_BOOL, VariantValue: true},
			},
			expected: Variant{VariantType: VAR_BOOL, VariantValue: false},
		},
		{
			desc: "incorrectly typed values",
			input: []Variant{
				{VariantType: VAR_BOOL, VariantValue: false},
				{VariantType: VAR_STRING, VariantValue: "true"},
			},
			expected: Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("type error: argument of unacceptable type %q passed to %q", VAR_STRING, "xnor")},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := logical.xnor(test.input)
			assert.Equal(t, test.expected, actual, "computation error")
		})
	}
}
