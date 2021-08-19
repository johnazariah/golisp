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
	strings := [...]string{
		"VAR_UNKNOWN",
		"VAR_NULL",
		"VAR_DATE",
		"VAR_BOOL",
		"VAR_FLOAT",
		"VAR_INT",
		"VAR_STRING",
		"VAR_IDENT",
		"VAR_ERROR",
		"VAR_MAX",
	}

	if t >= VAR_MAX {
		panic("unknown value for EnumVariantType")
	}

	return strings[t]
}

type Variant struct {
	VariantType  EnumVariantType
	VariantValue interface{}
}

func (b *Variant) GetTypeConsistentValue() (interface{}, error) {
	switch b.VariantType {
	case VAR_UNKNOWN:
		return nil, nil

	case VAR_NULL:
		return nil, nil

	case VAR_DATE:
		switch b.VariantValue.(type) {
		case time.Time:
			return b.VariantValue, nil
		default:
			return nil, buildInconsistentTypeError(b.VariantValue, b.VariantType)
		}

	case VAR_BOOL:
		switch b.VariantValue.(type) {
		case bool:
			return b.VariantValue.(bool), nil

		case int:
			return b.VariantValue.(int) != int(0), nil
		case int8:
			return b.VariantValue.(int8) != int8(0), nil
		case int16:
			return b.VariantValue.(int16) != int16(0), nil
		case int32:
			return b.VariantValue.(int32) != int32(0), nil
		case int64:
			return b.VariantValue.(int64) != int64(0), nil
		case uint:
			return b.VariantValue.(uint) != uint(0), nil
		case uint8:
			return b.VariantValue.(uint8) != uint8(0), nil
		case uint16:
			return b.VariantValue.(uint16) != uint16(0), nil
		case uint32:
			return b.VariantValue.(uint32) != uint32(0), nil
		case uint64:
			return b.VariantValue.(uint64) != uint64(0), nil

		default:
			return nil, buildInconsistentTypeError(b.VariantValue, b.VariantType)
		}

	case VAR_FLOAT:
		switch b.VariantValue.(type) {
		case float32:
			return float64(b.VariantValue.(float32)), nil
		case float64:
			return b.VariantValue.(float64), nil

		case int:
			return float64(b.VariantValue.(int)), nil
		case int8:
			return float64(b.VariantValue.(int8)), nil
		case int16:
			return float64(b.VariantValue.(int16)), nil
		case int32:
			return float64(b.VariantValue.(int32)), nil
		case int64:
			return float64(b.VariantValue.(int64)), nil
		case uint:
			return float64(b.VariantValue.(uint)), nil
		case uint8:
			return float64(b.VariantValue.(uint8)), nil
		case uint16:
			return float64(b.VariantValue.(uint16)), nil
		case uint32:
			return float64(b.VariantValue.(uint32)), nil
		case uint64:
			return float64(b.VariantValue.(uint64)), nil

		case bool:
			if b.VariantValue.(bool) {
				return float64(1.0), nil
			}
			return float64(0.0), nil

		default:
			return nil, buildInconsistentTypeError(b.VariantValue, b.VariantType)
		}
	case VAR_INT:
		switch b.VariantValue.(type) {
		case int:
			return int64(b.VariantValue.(int)), nil
		case int8:
			return int64(b.VariantValue.(int8)), nil
		case int16:
			return int64(b.VariantValue.(int16)), nil
		case int32:
			return int64(b.VariantValue.(int32)), nil
		case int64:
			return int64(b.VariantValue.(int64)), nil
		case uint:
			return int64(b.VariantValue.(uint)), nil
		case uint8:
			return int64(b.VariantValue.(uint8)), nil
		case uint16:
			return int64(b.VariantValue.(uint16)), nil
		case uint32:
			return int64(b.VariantValue.(uint32)), nil
		case uint64:
			return int64(b.VariantValue.(uint64)), nil

		case bool:
			if b.VariantValue.(bool) {
				return int64(1), nil
			}
			return int64(0), nil

		default:
			return nil, buildInconsistentTypeError(b.VariantValue, b.VariantType)
		}

	case VAR_STRING:
		switch b.VariantValue.(type) {
		case time.Time:
			return b.VariantValue.(time.Time).Format(time.RFC3339), nil

		case bool:
			return fmt.Sprintf("%t", b.VariantValue.(bool)), nil

		case int:
			return fmt.Sprintf("%d", b.VariantValue.(int)), nil
		case int8:
			return fmt.Sprintf("%d", b.VariantValue.(int8)), nil
		case int16:
			return fmt.Sprintf("%d", b.VariantValue.(int16)), nil
		case int32:
			return fmt.Sprintf("%d", b.VariantValue.(int32)), nil
		case int64:
			return fmt.Sprintf("%d", b.VariantValue.(int64)), nil
		case uint:
			return fmt.Sprintf("%d", b.VariantValue.(uint)), nil
		case uint8:
			return fmt.Sprintf("%d", b.VariantValue.(uint8)), nil
		case uint16:
			return fmt.Sprintf("%d", b.VariantValue.(uint16)), nil
		case uint32:
			return fmt.Sprintf("%d", b.VariantValue.(uint32)), nil
		case uint64:
			return fmt.Sprintf("%d", b.VariantValue.(uint64)), nil

		case float32:
			return fmt.Sprintf("%e", b.VariantValue.(float32)), nil
		case float64:
			return fmt.Sprintf("%e", b.VariantValue.(float64)), nil

		case string:
			return b.VariantValue, nil

		case error:
			return b.VariantValue.(error).Error(), nil

		default:
			return nil, buildInconsistentTypeError(b.VariantValue, b.VariantType)
		}

	case VAR_IDENT:
		switch b.VariantValue.(type) {
		case string:
			return b.VariantValue, nil

		default:
			return nil, buildInconsistentTypeError(b.VariantValue, b.VariantType)
		}

	case VAR_ERROR:
		switch b.VariantValue.(type) {
		case error:
			return b.VariantValue, nil
		default:
			return nil, buildInconsistentTypeError(b.VariantValue, b.VariantType)
		}

	case VAR_MAX:
		break
	}

	return nil, buildUnhandledVariantTypeError()
}

func (b *Variant) ToDebugString() string {
	v, e := b.GetTypeConsistentValue()
	if e != nil {
		return e.Error()
	}

	switch b.VariantType {
	case VAR_UNKNOWN:
		return "UNKNOWN"
	case VAR_NULL:
		return "NIL"
	case VAR_DATE:
		return v.(time.Time).Format(time.RFC3339)
	case VAR_BOOL:
		return fmt.Sprintf("%t", v.(bool))
	case VAR_FLOAT:
		return fmt.Sprintf("%e", v.(float64))
	case VAR_INT:
		return fmt.Sprintf("%d", v.(int64))
	case VAR_STRING:
		return v.(string)
	case VAR_IDENT:
		return v.(string)
	case VAR_ERROR:
		return v.(error).Error()
	case VAR_MAX:
		break
	}

	panic("should never get here")
}

func (b *Variant) GetDateValue() (time.Time, error) {
	targetType := VAR_DATE
	errorValue := time.Time{}
	acceptableTypes := map[EnumVariantType]bool{
		VAR_DATE: true,
	}

	if _, t := acceptableTypes[b.VariantType]; !t {
		return errorValue, buildTypeError(b.VariantType, targetType)
	}

	if value, err := b.GetTypeConsistentValue(); err != nil {
		return errorValue, err
	} else {
		return value.(time.Time), nil
	}
}

func (b *Variant) CoerceToBool() (bool, error) {
	targetType := VAR_BOOL
	errorValue := false

	acceptableTypes := map[EnumVariantType]bool{
		VAR_INT:  true,
		VAR_BOOL: true,
	}

	if _, t := acceptableTypes[b.VariantType]; !t {
		return errorValue, buildTypeError(b.VariantType, targetType)
	}

	coerced := &Variant{VariantType: targetType, VariantValue: b.VariantValue}

	if value, err := coerced.GetTypeConsistentValue(); err != nil {
		return errorValue, err
	} else {
		return value.(bool), nil
	}
}

func (b *Variant) CoerceToFloat() (float64, error) {
	targetType := VAR_FLOAT
	errorValue := float64(0)

	acceptableTypes := map[EnumVariantType]bool{
		VAR_FLOAT: true,
		VAR_INT:   true,
		VAR_BOOL:  true,
	}
	if _, t := acceptableTypes[b.VariantType]; !t {
		return errorValue, buildTypeError(b.VariantType, targetType)
	}

	coerced := &Variant{VariantType: targetType, VariantValue: b.VariantValue}

	if value, err := coerced.GetTypeConsistentValue(); err != nil {
		return errorValue, err
	} else {
		return value.(float64), nil
	}
}

func (b *Variant) CoerceToInt() (int64, error) {
	targetType := VAR_INT
	errorValue := int64(0)

	acceptableTypes := map[EnumVariantType]bool{
		VAR_INT:  true,
		VAR_BOOL: true,
	}
	if _, t := acceptableTypes[b.VariantType]; !t {
		return errorValue, buildTypeError(b.VariantType, targetType)
	}

	coerced := &Variant{VariantType: targetType, VariantValue: b.VariantValue}

	if value, err := coerced.GetTypeConsistentValue(); err != nil {
		return errorValue, err
	} else {
		return value.(int64), nil
	}
}
func (b *Variant) CoerceToString() (string, error) {
	targetType := VAR_STRING
	errorValue := ""

	coerced := &Variant{VariantType: targetType, VariantValue: b.VariantValue}

	if value, err := coerced.GetTypeConsistentValue(); err != nil {
		return errorValue, err
	} else {
		return value.(string), nil
	}
}

func (b *Variant) GetIdentifierValue() (string, error) {
	targetType := VAR_IDENT
	errorValue := ""
	acceptableTypes := map[EnumVariantType]bool{
		VAR_IDENT: true,
	}

	if _, t := acceptableTypes[b.VariantType]; !t {
		return errorValue, buildTypeError(b.VariantType, targetType)
	}

	if value, err := b.GetTypeConsistentValue(); err != nil {
		return errorValue, err
	} else {
		return value.(string), nil
	}
}

func (b *Variant) GetErrorValue() (error, error) {
	targetType := VAR_ERROR
	errorValue := fmt.Errorf("")
	acceptableTypes := map[EnumVariantType]bool{
		VAR_ERROR: true,
	}

	if _, t := acceptableTypes[b.VariantType]; !t {
		return errorValue, buildTypeError(b.VariantType, targetType)
	}

	if value, err := b.GetTypeConsistentValue(); err != nil {
		return errorValue, err
	} else {
		return value.(error), nil
	}
}
