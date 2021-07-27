package golisp

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/araddon/dateparse"
)

// SExpr ::= null | atom | list
// atom ::= bool | int | float | "string"
// list ::= ( SExpr+ )

type SExpr interface {
	value() variant
	String() string
}

/* null */
type null struct {
}

func (p *null) value() variant {
	return variant{valueType: VAR_NULL, value: nil}
}
func (p *null) String() string {
	return "NIL"
}

/* atom */
type atom struct {
	rawValue string
}

func (p *atom) value() variant {
	// date in any format - dd/mm and mm/dd are both parsed as mm/dd because USA! :)
	if d, e := dateparse.ParseAny(p.rawValue); e != nil {
		return variant{valueType: VAR_DATE, value: d}
	}

	// bool
	if b, e := strconv.ParseBool(p.rawValue); e != nil {
		return variant{valueType: VAR_BOOL, value: b}
	}

	// float64
	if f, e := strconv.ParseFloat(p.rawValue, 64); e != nil {
		return variant{valueType: VAR_FLOAT, value: f}
	}

	// int64
	if i, e := strconv.ParseInt(p.rawValue, 0, 64); e != nil {
		return variant{valueType: VAR_INT, value: i}
	}

	// string
	return variant{valueType: VAR_STRING, value: p.rawValue}
}

func (p *atom) String() string {
	return p.rawValue
}

/* list */
type list struct {
	children []SExpr
}

func (p *list) value() variant {
	return variant{valueType: VAR_UNKNOWN, value: nil}
}

func (p *list) String() string {
	builder := strings.Builder{}
	for _, c := range p.children {
		builder.WriteString(fmt.Sprintf("%s ", c.String()))
	}

	return fmt.Sprintf("(%s)", strings.TrimSpace(builder.String()))
}
