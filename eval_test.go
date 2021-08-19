package golisp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {
	tests := [...]struct {
		desc     string
		input    string
		expected Variant
	}{
		{desc: "atom: empty string", input: "", expected: Variant{VariantType: VAR_NULL}},
		{desc: "atom: whitespace string", input: " ", expected: Variant{VariantType: VAR_NULL}},
		{desc: "atom: numeric literal", input: "1", expected: Variant{VariantType: VAR_INT, VariantValue: int64(1)}},
		{desc: "atom: identifier literal", input: "a", expected: Variant{VariantType: VAR_IDENT, VariantValue: "a"}},
		{desc: "atom: quoted raw string", input: `"Now is the time"`, expected: Variant{VariantType: VAR_STRING, VariantValue: "Now is the time"}},
		{desc: "atom: quoted string", input: "\"Now is the time\"", expected: Variant{VariantType: VAR_STRING, VariantValue: "Now is the time"}},
		{desc: "add two ints", input: "(+ 1 2)", expected: Variant{VariantType: VAR_INT, VariantValue: int64(3)}},
		{desc: "add an int and a float", input: "(+ 1 2.0)", expected: Variant{VariantType: VAR_FLOAT, VariantValue: float64(3.0)}},
		{desc: "or two bools - false", input: "(or false 0)", expected: Variant{VariantType: VAR_BOOL, VariantValue: false}},
		{desc: "nor two bools - true", input: "(nor false 0)", expected: Variant{VariantType: VAR_BOOL, VariantValue: true}},
		{desc: "and two bools - true", input: "(and true t)", expected: Variant{VariantType: VAR_BOOL, VariantValue: true}},
		{desc: "nand two bools - false", input: "(nand true t)", expected: Variant{VariantType: VAR_BOOL, VariantValue: false}},
		{desc: "xor two bools - true", input: "(xor true false)", expected: Variant{VariantType: VAR_BOOL, VariantValue: true}},
		{desc: "xnor two bools - true", input: "(xnor f 0)", expected: Variant{VariantType: VAR_BOOL, VariantValue: true}},
		{desc: "&& many bools", input: "(&& 1 t T TRUE true True)", expected: Variant{VariantType: VAR_BOOL, VariantValue: true}},
		{desc: "|| many bools", input: "(|| 0 f F FALSE false False)", expected: Variant{VariantType: VAR_BOOL, VariantValue: false}},
		{desc: "! single false", input: "(! false)", expected: Variant{VariantType: VAR_BOOL, VariantValue: true}},
		{desc: "or many bools - true", input: "(or 1 t T TRUE true True)", expected: Variant{VariantType: VAR_BOOL, VariantValue: true}},
		{desc: "nand many bools - true", input: "(nand 0 f F FALSE false False)", expected: Variant{VariantType: VAR_BOOL, VariantValue: true}},
		{desc: "or nested - true", input: "(or (or 1 t) (or T TRUE (or true True)))", expected: Variant{VariantType: VAR_BOOL, VariantValue: true}},
		{desc: "concat two strings", input: "(concat \"Hello, \" \"World!\")", expected: Variant{VariantType: VAR_STRING, VariantValue: "Hello, World!"}},
		{desc: "concat two strings and an int", input: "(++ \"Hello, \" \"Competitor \" 27 \"!\")", expected: Variant{VariantType: VAR_STRING, VariantValue: "Hello, Competitor 27!"}},
		{desc: "invalid type", input: "(or 3.1415 today)", expected: Variant{VariantType: VAR_ERROR, VariantValue: buildUnacceptableTypeError(VAR_FLOAT, "or")}},
		{desc: "unresolved identifier", input: "(or yesterday today)", expected: Variant{VariantType: VAR_ERROR, VariantValue: buildUnresolvedIdentifierError("yesterday")}},
		{desc: "invalid function", input: "(+ (1) (2))", expected: Variant{VariantType: VAR_ERROR, VariantValue: buildInvalidFunctionNameError("1")}},
		{desc: "unknown symbol", input: "(+ 1 2 a)", expected: Variant{VariantType: VAR_ERROR, VariantValue: buildUnresolvedIdentifierError("a")}},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			sexpr, e := Parse(test.input)
			assert.Nil(t, e, "parse error")

			ctx := sexpr.Eval(NewEvaluationContext(nil))
			assert.Equal(t, test.expected, ctx.EvaluatedValue)
		})
	}
}
