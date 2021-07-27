package golisp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	cases := [...]struct {
		desc    string
		input   string
		success string
		failure string
	}{
		{desc: "empty string", input: "", success: "NIL"},
		{desc: "whitespace string", input: " ", success: "NIL"},
		{desc: "numeric literal", input: "1", success: "1"},
		{desc: "identifier literal", input: "a", success: "a"},
		{desc: "quoted raw string", input: `"Now is the time"`, success: `"Now is the time"`},
		{desc: "quoted string", input: "\"Now is the time\"", success: "\"Now is the time\""},
		{desc: "valid list", input: "(+ 1 2)", success: "(+ 1 2)"},
		{desc: "valid list", input: "(+ 1 (2))", success: "(+ 1 (2))"},
		{desc: "valid list", input: "(+ (1) (2))", success: "(+ (1) (2))"},
		{desc: "valid list", input: "(+ (1) (2))", success: "(+ (1) (2))"},
		{desc: "valid list", input: "(+ (1) (+ 2 3))", success: "(+ (1) (+ 2 3))"},
		{desc: "valid list", input: "(+ (1) (+ 2 3) 4)", success: "(+ (1) (+ 2 3) 4)"},
		{desc: "valid list", input: "(+ (1) (+ 2 3) a)", success: "(+ (1) (+ 2 3) a)"},
		{desc: "parse error", input: "(", success: "NIL", failure: "unexpected end of string"},
		{desc: "parse error", input: "(+ (* a b)", success: "NIL", failure: "unexpected end of string"},
		{desc: "parse error", input: "(+ (1) (+ 2 3)", success: "NIL", failure: "unexpected end of string"},
		{desc: "unexpected rparen", input: ")", success: "NIL", failure: "unexpected close paren"},
		{desc: "trailing garbage", input: "(+ 1 (+ 2 3)))", success: "NIL", failure: "unexpected trailing text"},
		{desc: "trailing garbage", input: "())", success: "NIL", failure: "unexpected trailing text"},
		{desc: "trailing garbage", input: "(+ 1 2))", success: "NIL", failure: "unexpected trailing text"},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			sexpr, e := Parse(c.input)
			if c.failure == "" {
				assert.Nil(t, e, "parse error")
				assert.Equal(t, c.success, sexpr.String(), "success values did not match!")
			} else {
				assert.EqualError(t, e, c.failure)
			}
		})
	}
}
