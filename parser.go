package golisp

import (
	"fmt"
)

func parseSExpr(tokenizer *tokenizerContext, tok *token, into *list) (SExpr, error) {
	if tok == nil {
		into.children = append(into.children, &null{})
		return into, nil
	}

	switch tok.tokenType {
	case TOK_COMMENT:
		break

	case TOK_SYMBOL, TOK_QUOTEDSTRING:
		into.children = append(into.children, &atom{rawValue: tok.rawValue(tokenizer)})

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
