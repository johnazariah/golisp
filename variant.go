package golisp

import (
	"fmt"
	"time"
)

type EnumVariantType uint8

const (
	VAR_UNKNOWN EnumVariantType = iota
	VAR_NULL
	VAR_DATE
	VAR_BOOL
	VAR_FLOAT
	VAR_INT
	VAR_STRING
	VAR_IDENT
	VAR_ERROR
	VAR_MAX
)

func (t EnumVariantType) String() string {
	switch t {
	case VAR_NULL:
		return "VAR_NULL"
	case VAR_DATE:
		return "VAR_DATE"
	case VAR_BOOL:
		return "VAR_BOOL"
	case VAR_FLOAT:
		return "VAR_FLOAT"
	case VAR_INT:
		return "VAR_INT"
	case VAR_STRING:
		return "VAR_STRING"
	case VAR_IDENT:
		return "VAR_IDENT"
	case VAR_ERROR:
		return "VAR_ERROR"
	case VAR_MAX:
		return "VAR_MAX"
	case VAR_UNKNOWN:
		return "VAR_UNKNOWN"
	}

	panic("unknown value for EnumVariantType")
}

type Variant struct {
	VariantType  EnumVariantType
	VariantValue interface{}
}

func (v *Variant) ExtractString() string {
	switch v.VariantType {
	case VAR_UNKNOWN:
		return "UNKNOWN"
	case VAR_NULL:
		return "NIL"
	case VAR_DATE:
		return v.VariantValue.(time.Time).String()
	case VAR_BOOL:
		return fmt.Sprintf("%t", v.VariantValue.(bool))
	case VAR_FLOAT:
		return fmt.Sprintf("%e", v.VariantValue.(float64))
	case VAR_INT:
		return fmt.Sprintf("%d", v.VariantValue.(int64))
	case VAR_STRING:
		return v.VariantValue.(string)
	case VAR_IDENT:
		return v.VariantValue.(string)
	case VAR_ERROR:
		return v.VariantValue.(error).Error()
	case VAR_MAX:
		panic("MAX")
	}
	panic("should never get here")
}

func (b *Variant) ExtractBool() (bool, error) {
	switch b.VariantValue.(type) {
	case bool:
		return b.VariantValue.(bool), nil

	case int, int8, int16, int32, int64:
		return b.VariantValue.(int64) != int64(0), nil

	default:
		return false, fmt.Errorf("coerce error: unable to coerce %q of type %q to %q", b.ExtractString(), b.VariantType, "bool")
	}
}

func (b *Variant) ExtractInt() (int64, error) {
	switch b.VariantValue.(type) {
	case int, int8, int16, int32, int64:
		return b.VariantValue.(int64), nil
	case bool:
		if b.VariantValue.(bool) {
			return int64(1), nil
		}
		return int64(0), nil
	default:
		return 0, fmt.Errorf("coerce error: unable to coerce %q of type %q to %q", b.ExtractString(), b.VariantType, "int")
	}
}

func (b *Variant) ExtractFloat() (float64, error) {
	switch b.VariantValue.(type) {
	case float32, float64:
		return b.VariantValue.(float64), nil

	case int, int8, int16, int32, int64:
		return float64(b.VariantValue.(int64)), nil

	case bool:
		if b.VariantValue.(bool) {
			return float64(1.0), nil
		}
		return float64(0.0), nil

	default:
		return 0, fmt.Errorf("coerce error: unable to coerce %q of type %q to %q", b.ExtractString(), b.VariantType, "float")
	}
}
