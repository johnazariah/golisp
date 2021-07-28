package golisp

import (
	"fmt"
	"strings"
)

// SExpr ::= null | atom | list
// atom ::= bool | int | float | "string"
// list ::= ( SExpr+ )

type SExpr interface {
	Eval(*EvaluationContext) *EvaluationContext
	String() string
}

/* null */
type null struct {
}

func (p *null) String() string {
	return "NIL"
}

/* atom */
type atom struct {
	rawValue   string
	typedValue Variant
}

func (p *atom) String() string {
	return p.rawValue
}

/* list */
type list struct {
	children []SExpr
}

func (p *list) String() string {
	builder := strings.Builder{}
	for _, c := range p.children {
		builder.WriteString(fmt.Sprintf("%s ", c.String()))
	}

	return fmt.Sprintf("(%s)", strings.TrimSpace(builder.String()))
}
