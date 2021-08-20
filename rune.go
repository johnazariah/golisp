package golisp

import "unicode"

func isLParen(r rune) bool {
	return r == '('
}

func isRParen(r rune) bool {
	return r == ')'
}

func isSeparator(r rune) bool {
	return unicode.IsSpace(r) || isLParen(r) || isRParen(r)
}

func isQuote(r rune) bool {
	return r == '"'
}
