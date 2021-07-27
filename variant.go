package golisp

type enumVariantType uint8

const (
	VAR_UNKNOWN enumVariantType = iota
	VAR_NULL
	VAR_DATE
	VAR_BOOL
	VAR_FLOAT
	VAR_INT
	VAR_STRING
	VAR_MAX
)

type variant struct {
	valueType enumVariantType
	value     interface{}
}
