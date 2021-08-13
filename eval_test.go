package golisp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {
	cases := [...]struct {
		desc             string
		input            string
		successValueType EnumVariantType
		successValue     interface{}
		failureMessage   string
	}{
		{desc: "atom: empty string", input: "", successValueType: VAR_NULL, successValue: "NIL"},
		{desc: "atom: whitespace string", input: " ", successValueType: VAR_NULL, successValue: "NIL"},
		{desc: "atom: numeric literal", input: "1", successValueType: VAR_INT, successValue: int64(1)},
		{desc: "atom: identifier literal", input: "a", successValueType: VAR_IDENT, successValue: "a"},
		{desc: "atom: quoted raw string", input: `"Now is the time"`, successValueType: VAR_STRING, successValue: "Now is the time"},
		{desc: "atom: quoted string", input: "\"Now is the time\"", successValueType: VAR_STRING, successValue: "Now is the time"},
		{desc: "add two ints", input: "(+ 1 2)", successValueType: VAR_INT, successValue: int64(3)},
		{desc: "add an int and a float", input: "(+ 1 2.0)", successValueType: VAR_FLOAT, successValue: float64(3.0)},
		// {desc: "add two strings", input: "(+ \"Hello, \" \"World!\")", successValueType: VAR_STRING, successValue: "Hello, World!"},
		// {desc: "add two strings and an int", input: "(+ \"Hello, \" \"Competitor \" 27 \"!\")", successValueType: VAR_STRING, successValue: "Hello, Competitor 27!"},
		{desc: "or two bools - false", input: "(or false 0)", successValueType: VAR_BOOL, successValue: false},
		{desc: "nor two bools - true", input: "(nor false 0)", successValueType: VAR_BOOL, successValue: true},
		{desc: "and two bools - true", input: "(and true t)", successValueType: VAR_BOOL, successValue: true},
		{desc: "nand two bools - false", input: "(nand true t)", successValueType: VAR_BOOL, successValue: false},
		{desc: "xor two bools - true", input: "(xor true false)", successValueType: VAR_BOOL, successValue: true},
		{desc: "xnor two bools - true", input: "(xnor f 0)", successValueType: VAR_BOOL, successValue: true},
		{desc: "&& many bools", input: "(&& 1 t T TRUE true True)", successValueType: VAR_BOOL, successValue: true},
		{desc: "|| many bools", input: "(|| 0 f F FALSE false False)", successValueType: VAR_BOOL, successValue: false},
		{desc: "! single false", input: "(! false)", successValueType: VAR_BOOL, successValue: true},
		{desc: "or many bools - true", input: "(or 1 t T TRUE true True)", successValueType: VAR_BOOL, successValue: true},
		{desc: "nand many bools - true", input: "(nand 0 f F FALSE false False)", successValueType: VAR_BOOL, successValue: true},
		{desc: "or nested - true", input: "(or (or 1 t) (or T TRUE (or true True)))", successValueType: VAR_BOOL, successValue: true},
		{desc: "invalid type", input: "(or 3.1415 today)", failureMessage: fmt.Sprintf("type error: argument of unacceptable type %q passed to \"or\"", VAR_FLOAT)},
		{desc: "unresolved identifier", input: "(or yesterday today)", failureMessage: fmt.Sprintf("scope error: unresolved identifier %q", "yesterday")},
		{desc: "invalid function", input: "(+ (1) (2))", failureMessage: "requested function \"1\" not found"},
		{desc: "unknown symbol", input: "(+ 1 2 a)", failureMessage: "scope error: unresolved identifier \"a\""},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			sexpr, e := Parse(c.input)
			assert.Nil(t, e, "parse error")

			ctx := sexpr.Eval(NewEvaluationContext(nil))
			switch ctx.EvaluatedValue.VariantType {
			case VAR_MAX:
				assert.Fail(t, "should never return VAR_MAX")

			case VAR_NULL, VAR_UNKNOWN:
				assert.Equal(t, c.successValueType, ctx.EvaluatedValue.VariantType, "unexpected success value type")

			case VAR_ERROR:
				assert.NotNil(t, c.failureMessage, "encountered unexpected failure")
				assert.Equal(t, c.failureMessage, ctx.EvaluatedValue.ExtractString(), "unexpected failure message")

			default:
				assert.Equal(t, c.successValueType, ctx.EvaluatedValue.VariantType, "unexpected success value type")
				assert.Equal(t, c.successValue, ctx.EvaluatedValue.VariantValue, "success values do not match")
			}
		})
	}
}
