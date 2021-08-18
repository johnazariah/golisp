package golisp

import (
	"testing"
)

type TokenizerTestResult struct {
	tokenType enumTokenType
	value     string
}

func TestTokenizer(t *testing.T) {
	tests := [...]struct {
		input    string
		expected []TokenizerTestResult
	}{
		{input: "", expected: []TokenizerTestResult{}},
		{input: " ", expected: []TokenizerTestResult{}},
		{input: "( )", expected: []TokenizerTestResult{{tokenType: TOK_LPAREN}, {tokenType: TOK_RPAREN}}},
		{input: " ( ) ", expected: []TokenizerTestResult{{tokenType: TOK_LPAREN}, {tokenType: TOK_RPAREN}}},
		{input: "(123", expected: []TokenizerTestResult{{tokenType: TOK_LPAREN}, {tokenType: TOK_SYMBOL, value: "123"}}},
		{input: "123", expected: []TokenizerTestResult{{tokenType: TOK_SYMBOL, value: "123"}}},
		{input: "1", expected: []TokenizerTestResult{{tokenType: TOK_SYMBOL, value: "1"}}},
		{input: "123x", expected: []TokenizerTestResult{{tokenType: TOK_SYMBOL, value: "123x"}}},
		{input: "a123x", expected: []TokenizerTestResult{{tokenType: TOK_SYMBOL, value: "a123x"}}},
		{input: " ( 1 2 3 4 ) ", expected: []TokenizerTestResult{{tokenType: TOK_LPAREN}, {tokenType: TOK_SYMBOL, value: "1"}, {tokenType: TOK_SYMBOL, value: "2"}, {tokenType: TOK_SYMBOL, value: "3"}, {tokenType: TOK_SYMBOL, value: "4"}, {tokenType: TOK_RPAREN}}},
		{input: "(1 2 3 4)", expected: []TokenizerTestResult{{tokenType: TOK_LPAREN}, {tokenType: TOK_SYMBOL, value: "1"}, {tokenType: TOK_SYMBOL, value: "2"}, {tokenType: TOK_SYMBOL, value: "3"}, {tokenType: TOK_SYMBOL, value: "4"}, {tokenType: TOK_RPAREN}}},
		{input: "(a a a a)", expected: []TokenizerTestResult{{tokenType: TOK_LPAREN}, {tokenType: TOK_SYMBOL, value: "a"}, {tokenType: TOK_SYMBOL, value: "a"}, {tokenType: TOK_SYMBOL, value: "a"}, {tokenType: TOK_SYMBOL, value: "a"}, {tokenType: TOK_RPAREN}}},
		{input: "(() )", expected: []TokenizerTestResult{{tokenType: TOK_LPAREN}, {tokenType: TOK_LPAREN}, {tokenType: TOK_RPAREN}, {tokenType: TOK_RPAREN}}},
		{input: "(\"this and that\")", expected: []TokenizerTestResult{{tokenType: TOK_LPAREN}, {tokenType: TOK_QUOTEDSTRING, value: "\"this and that\""}, {tokenType: TOK_RPAREN}}},
		{input: "(/* this is a comment */)", expected: []TokenizerTestResult{{tokenType: TOK_LPAREN}, {tokenType: TOK_COMMENT, value: "/* this is a comment */"}, {tokenType: TOK_RPAREN}}},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			ti := newTokenizerContext(test.input)

			i := 0

			for token := ti.NextToken(); token != nil; token = ti.NextToken() {

				if test.expected[i].tokenType != token.tokenType {
					t.Errorf("Expected Token Type: [%v]; Actual [%v]", test.expected[i].tokenType, token.tokenType)
					break
				}

				tokenValue := token.rawValue(ti)

				if test.expected[i].value != tokenValue {
					t.Errorf("Expected Value: [%v]; Actual [%v]", test.expected[i].value, tokenValue)
					break
				}

				i++
			}

			if i != len(test.expected) {
				t.Errorf("Token count mismatch. Found [%v] tokens but expecting [%v]", i, len(test.expected))
			}
		})
	}
}
