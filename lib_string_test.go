package golisp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var str = &(StringLibrary{})

func TestConcat(t *testing.T) {
	tests := [...]struct {
		desc     string
		input    []Variant
		expected Variant
	}{
		{
			desc: "single string value",
			input: []Variant{
				{VariantType: VAR_STRING, VariantValue: "Hello"},
			},
			expected: Variant{VariantType: VAR_STRING, VariantValue: "Hello"},
		},
		{
			desc: "two string values",
			input: []Variant{
				{VariantType: VAR_STRING, VariantValue: "Hello, "},
				{VariantType: VAR_STRING, VariantValue: "World!"},
			},
			expected: Variant{VariantType: VAR_STRING, VariantValue: "Hello, World!"},
		},
		{
			desc: "coerced string values",
			input: []Variant{
				{VariantType: VAR_STRING, VariantValue: "Hello, "},
				{VariantType: VAR_STRING, VariantValue: "Contestant #"},
				{VariantType: VAR_INT, VariantValue: int64(1)},
			},
			expected: Variant{VariantType: VAR_STRING, VariantValue: "Hello, Contestant #1"},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := str.concat(test.input)
			assert.Equal(t, test.expected, actual, "computation error")
		})
	}
}
