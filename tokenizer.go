package golisp

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

type enumTokenType uint8

const (
	TOK_BEGIN enumTokenType = iota
	TOK_LPAREN
	TOK_RPAREN
	TOK_COMMENT
	TOK_QUOTEDSTRING
	TOK_SYMBOL
	TOK_END
	// put new tokens between BEGIN and END, and ensure you implement `String()` correctly!
	TOK_UNKNOWN
)

func (t enumTokenType) String() string {
	names := [...]string{
		"BEGIN",
		"LPAREN",
		"RPAREN",
		"COMMENT",
		"QUOTEDSTRING",
		"SYMBOL",
		"END",
		"UNKNOWN",
	}

	if t < TOK_BEGIN || t > TOK_END {
		return names[TOK_UNKNOWN]
	}

	return names[t]
}

type token struct {
	start     int
	finish    int
	tokenType enumTokenType
}

type tokenizerContext struct {
	code string
	idx  int
}

func (t *token) rawValue(ctx *tokenizerContext) string {
	switch t.tokenType {
	case TOK_SYMBOL, TOK_QUOTEDSTRING, TOK_COMMENT:
		return ctx.code[t.start:t.finish]
	default:
		return ""
	}
}

func (ctx *tokenizerContext) hasMoreText() bool {
	runecount := utf8.RuneCountInString(ctx.code)
	return (ctx.idx < runecount)
}

func (ctx *tokenizerContext) currentRune() (rune, int) {
	return utf8.DecodeRuneInString(ctx.code[ctx.idx:])
}

func newTokenizerContext(code string) *tokenizerContext {
	return &tokenizerContext{code: strings.TrimSpace(code), idx: 0}
}

func (ctx *tokenizerContext) skipWhitespace() {
	runecount := utf8.RuneCountInString(ctx.code)

	for ctx.idx < runecount {
		r, width := ctx.currentRune()

		if !unicode.IsSpace(r) {
			break
		}

		ctx.idx += width
	}
}

func (ctx *tokenizerContext) readChar(ch rune, tokenType enumTokenType) *token {
	r, width := ctx.currentRune()
	if r == ch {
		start := ctx.idx

		ctx.idx += width
		finish := ctx.idx

		return &token{start: start, finish: finish, tokenType: tokenType}
	}

	return nil
}

func (ctx *tokenizerContext) read_LPARAM() *token {
	return ctx.readChar('(', TOK_LPAREN)
}

func (ctx *tokenizerContext) read_RPARAM() *token {
	return ctx.readChar(')', TOK_RPAREN)
}

func (ctx *tokenizerContext) read_QUOTEDSTRING() *token {
	runeValue, width := ctx.currentRune()
	if !isQuote(runeValue) {
		return nil
	}

	start := ctx.idx
	ctx.idx += width

	runecount := utf8.RuneCountInString(ctx.code)
	for ctx.idx < runecount {
		runeValue, width = ctx.currentRune()
		ctx.idx += width

		if isQuote(runeValue) {
			return &token{start: start, finish: ctx.idx, tokenType: TOK_QUOTEDSTRING}
		}
	}

	ctx.idx = start
	return nil
}

func (ctx *tokenizerContext) read_COMMENT() *token {
	if !(strings.HasPrefix(ctx.code[ctx.idx:], "/*")) {
		return nil
	}

	start := ctx.idx
	runecount := utf8.RuneCountInString(ctx.code)

	for curr := start; curr < runecount; curr++ {
		if strings.HasSuffix(ctx.code[start:curr], "*/") {
			ctx.idx = curr
			return &token{start: start, finish: curr, tokenType: TOK_COMMENT}
		}
	}

	return nil
}

func (ctx *tokenizerContext) read_SYMBOL() *token {
	start := ctx.idx
	runecount := utf8.RuneCountInString(ctx.code)

	for ctx.idx < runecount {
		runeValue, width := ctx.currentRune()

		if isSeparator(runeValue) {
			break
		}

		ctx.idx += width
	}

	if start == ctx.idx {
		return nil
	}

	return &token{start: start, finish: ctx.idx, tokenType: TOK_SYMBOL}
}

func (ctx *tokenizerContext) NextToken() *token {
	tokenizerFuncs := [...](func() *token){
		ctx.read_LPARAM,
		ctx.read_RPARAM,
		ctx.read_COMMENT,
		ctx.read_QUOTEDSTRING,
		ctx.read_SYMBOL,
	}

	ctx.skipWhitespace()

	runecount := utf8.RuneCountInString(ctx.code)
	if ctx.idx < runecount {
		for _, tokenizerFunc := range tokenizerFuncs {
			if token := tokenizerFunc(); token != nil {
				return token
			}
		}
	}
	return nil
}
