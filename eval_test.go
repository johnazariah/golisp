package golisp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {

	context := NewEvaluationContext(nil)
	context.SymbolTable["known_symbol_string"] = Variant{VariantType: VAR_STRING, VariantValue: "A Known Symbol"}
	context.SymbolTable["x"] = Variant{VariantType: VAR_INT, VariantValue: int64(22)}

	tests := [...]struct {
		desc     string
		input    string
		expected Variant
	}{
		{desc: "atom: empty string", input: "", expected: Variant{VariantType: VAR_NULL}},
		{desc: "atom: whitespace string", input: " ", expected: Variant{VariantType: VAR_NULL}},
		{desc: "atom: numeric literal", input: "1", expected: Variant{VariantType: VAR_INT, VariantValue: int64(1)}},
		{desc: "atom: unknown identifier literal", input: "a", expected: Variant{VariantType: VAR_ERROR, VariantValue: buildUnresolvedIdentifierError("a")}},
		{desc: "atom: known identifier literal", input: "x", expected: Variant{VariantType: VAR_INT, VariantValue: int64(22)}},
		{desc: "atom: known identifier literal", input: "known_symbol_string", expected: Variant{VariantType: VAR_STRING, VariantValue: "A Known Symbol"}},
		{desc: "atom: quoted raw string", input: `"Now is the time"`, expected: Variant{VariantType: VAR_STRING, VariantValue: "Now is the time"}},
		{desc: "atom: quoted string", input: "\"Now is the time\"", expected: Variant{VariantType: VAR_STRING, VariantValue: "Now is the time"}},
		{desc: "empty list", input: "()", expected: Variant{VariantType: VAR_NULL}},
		{desc: "empty list with comment", input: "((*yee haw*))", expected: Variant{VariantType: VAR_NULL}},
		{desc: "add two ints with embedded comment", input: "(+ (* this is a comment *) 1 (* this is another comment*) 2)", expected: Variant{VariantType: VAR_INT, VariantValue: int64(3)}},
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
		{desc: "unknown symbol", input: "(+ 1 2 a)", expected: Variant{VariantType: VAR_ERROR, VariantValue: buildUnresolvedIdentifierError("a")}},
		{desc: "known symbol", input: "(+ 1 2 x)", expected: Variant{VariantType: VAR_INT, VariantValue: int64(25)}},
		{desc: "known symbol - nested", input: "(+ 1 2 (+ 5 x))", expected: Variant{VariantType: VAR_INT, VariantValue: int64(30)}},
		{desc: "invalid type", input: "(or 3.1415 today)", expected: Variant{VariantType: VAR_ERROR, VariantValue: buildUnacceptableTypeError(VAR_FLOAT, "or")}},
		{desc: "invalid function", input: "(+ (1) (2))", expected: Variant{VariantType: VAR_ERROR, VariantValue: buildFunctionNameNotFoundError("1")}},
		{desc: "unresolved identifier", input: "(or yesterday today)", expected: Variant{VariantType: VAR_ERROR, VariantValue: buildUnresolvedIdentifierError("yesterday")}},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			sexpr, e := Parse(test.input)
			assert.Nil(t, e, "parse error")

			ctx := sexpr.Eval(context)
			assert.Equal(t, test.expected, ctx.EvaluatedValue)
		})
	}
}
