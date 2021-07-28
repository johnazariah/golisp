package golisp

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/araddon/dateparse"
)

func parseSExpr(tokenizer *tokenizerContext, tok *token, into *list) (SExpr, error) {
	if tok == nil {
		into.children = append(into.children, &null{})
		return into, nil
	}

	switch tok.tokenType {
	case TOK_COMMENT:
		break

	case TOK_QUOTEDSTRING:
		a := &atom{rawValue: tok.rawValue(tokenizer)}
		a.typedValue = Variant{VariantType: VAR_STRING, VariantValue: strings.Trim(a.rawValue, "\"")}
		into.children = append(into.children, a)

	case TOK_SYMBOL:
		a := &atom{rawValue: tok.rawValue(tokenizer)}

		// date in any format - dd/mm and mm/dd are both parsed as mm/dd because USA! :)
		if d, e := dateparse.ParseAny(a.rawValue); e == nil {
			a.typedValue = Variant{VariantType: VAR_DATE, VariantValue: d}
		} else if i, e := strconv.ParseInt(a.rawValue, 0, 64); e == nil {
			// int64
			a.typedValue = Variant{VariantType: VAR_INT, VariantValue: i}
		} else if f, e := strconv.ParseFloat(a.rawValue, 64); e == nil {
			// float64
			a.typedValue = Variant{VariantType: VAR_FLOAT, VariantValue: f}
		} else if b, e := strconv.ParseBool(a.rawValue); e == nil {
			// bool
			a.typedValue = Variant{VariantType: VAR_BOOL, VariantValue: b}
		} else {
			// identifier
			a.typedValue = Variant{VariantType: VAR_IDENT, VariantValue: a.rawValue}
		}

		into.children = append(into.children, a)

	case TOK_LPAREN:
		child := &list{children: []SExpr{}}
		var t *token = tokenizer.NextToken()

		for t != nil && t.tokenType != TOK_RPAREN {
			parseSExpr(tokenizer, t, child)
			t = tokenizer.NextToken()
		}

		if t == nil {
			return into, fmt.Errorf("unexpected end of string")
		}

		into.children = append(into.children, child)

	case TOK_RPAREN:
		return into, fmt.Errorf("unexpected close paren")
	}

	return into, nil
}

func Parse(s string) (SExpr, error) {
	tokenizer := newTokenizerContext(s)
	token := tokenizer.NextToken()

	r, e := parseSExpr(tokenizer, token, &list{children: []SExpr{}})

	if tokenizer.hasMoreText() {
		return &null{}, fmt.Errorf("unexpected trailing text")
	}

	if e == nil {
		return r.(*list).children[0], e
	}
	return &null{}, e
}
