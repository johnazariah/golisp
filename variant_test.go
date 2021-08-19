package golisp

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_EdgeCases_GetTypeConsistentValue(t *testing.T) {
	tests := [...]struct {
		desc          string
		input         Variant
		expectedValue interface{}
		expectedError error
	}{
		{desc: "succeed: from VAR_UNKNOWN", input: Variant{VariantType: VAR_UNKNOWN, VariantValue: nil}, expectedValue: nil},
		{desc: "succeed: from VAR_NULL", input: Variant{VariantType: VAR_NULL, VariantValue: nil}, expectedValue: nil},
		{desc: "succeed: from VAR_MAX", input: Variant{VariantType: VAR_MAX, VariantValue: nil}, expectedError: buildUnhandledVariantTypeError()},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			if actualValue, actualError := test.input.GetTypeConsistentValue(); actualError == nil {
				assert.Equal(t, test.expectedValue, actualValue, "fail")
			} else {
				assert.EqualError(t, test.expectedError, actualError.Error(), actualError.Error())
			}
		})
	}
}

func TestToDebugString(t *testing.T) {
	now := time.Now()
	tests := [...]struct {
		desc          string
		input         Variant
		expectedValue string
	}{
		{desc: "VAR_UNKNOWN", input: Variant{VariantType: VAR_UNKNOWN, VariantValue: nil}, expectedValue: "UNKNOWN"},
		{desc: "VAR_NULL", input: Variant{VariantType: VAR_NULL, VariantValue: nil}, expectedValue: "NIL"},
		{desc: "VAR_DATE", input: Variant{VariantType: VAR_DATE, VariantValue: now}, expectedValue: now.Format(time.RFC3339)},
		{desc: "VAR_BOOL", input: Variant{VariantType: VAR_BOOL, VariantValue: true}, expectedValue: fmt.Sprintf("%t", true)},
		{desc: "VAR_FLOAT", input: Variant{VariantType: VAR_FLOAT, VariantValue: 3.14}, expectedValue: "3.140000e+00"},
		{desc: "VAR_INT", input: Variant{VariantType: VAR_INT, VariantValue: 97}, expectedValue: "97"},
		{desc: "VAR_STRING", input: Variant{VariantType: VAR_STRING, VariantValue: "Henlo!"}, expectedValue: "Henlo!"},
		{desc: "VAR_IDENT", input: Variant{VariantType: VAR_IDENT, VariantValue: "w"}, expectedValue: "w"},
		{desc: "VAR_ERROR", input: Variant{VariantType: VAR_ERROR, VariantValue: errRandom}, expectedValue: errRandom.Error()},
		{desc: "inconsistent", input: Variant{VariantType: VAR_DATE, VariantValue: "Some Random String"}, expectedValue: "type error: value [Some Random String] is inconsistent with type \"VAR_DATE\""},
		// {desc: "VAR_MAX", input: Variant{VariantType: VAR_MAX, VariantValue: nil}, expectedValue: "UNKNOWN"}, should panic
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actualValue := test.input.ToDebugString()
			assert.Equal(t, test.expectedValue, actualValue, "fail")
		})
	}
}

func TestCoerceToBool(t *testing.T) {
	targetType := VAR_BOOL
	tests := [...]struct {
		desc          string
		input         Variant
		expectedValue bool
		expectedError error
	}{
		{desc: "succeed: from bool - true", input: Variant{VariantType: VAR_BOOL, VariantValue: true}, expectedValue: true},
		{desc: "succeed: from bool - false", input: Variant{VariantType: VAR_BOOL, VariantValue: false}, expectedValue: false},

		{desc: "succeed: from int - true", input: Variant{VariantType: VAR_INT, VariantValue: int(1)}, expectedValue: true},
		{desc: "succeed: from int8 - true", input: Variant{VariantType: VAR_INT, VariantValue: int8(1)}, expectedValue: true},
		{desc: "succeed: from int16 - true", input: Variant{VariantType: VAR_INT, VariantValue: int16(1)}, expectedValue: true},
		{desc: "succeed: from int32 - true", input: Variant{VariantType: VAR_INT, VariantValue: int32(1)}, expectedValue: true},
		{desc: "succeed: from int64 - true", input: Variant{VariantType: VAR_INT, VariantValue: int64(1)}, expectedValue: true},
		{desc: "succeed: from uint - true", input: Variant{VariantType: VAR_INT, VariantValue: uint(1)}, expectedValue: true},
		{desc: "succeed: from uint8 - true", input: Variant{VariantType: VAR_INT, VariantValue: uint8(1)}, expectedValue: true},
		{desc: "succeed: from uint16 - true", input: Variant{VariantType: VAR_INT, VariantValue: uint16(1)}, expectedValue: true},
		{desc: "succeed: from uint32 - true", input: Variant{VariantType: VAR_INT, VariantValue: uint32(1)}, expectedValue: true},
		{desc: "succeed: from uint64 - true", input: Variant{VariantType: VAR_INT, VariantValue: uint64(1)}, expectedValue: true},

		{desc: "succeed: from int - false", input: Variant{VariantType: VAR_INT, VariantValue: int(0)}, expectedValue: false},
		{desc: "succeed: from int8 - false", input: Variant{VariantType: VAR_INT, VariantValue: int8(0)}, expectedValue: false},
		{desc: "succeed: from int16 - false", input: Variant{VariantType: VAR_INT, VariantValue: int16(0)}, expectedValue: false},
		{desc: "succeed: from int32 - false", input: Variant{VariantType: VAR_INT, VariantValue: int32(0)}, expectedValue: false},
		{desc: "succeed: from int64 - false", input: Variant{VariantType: VAR_INT, VariantValue: int64(0)}, expectedValue: false},
		{desc: "succeed: from uint - false", input: Variant{VariantType: VAR_INT, VariantValue: uint(0)}, expectedValue: false},
		{desc: "succeed: from uint8 - false", input: Variant{VariantType: VAR_INT, VariantValue: uint8(0)}, expectedValue: false},
		{desc: "succeed: from uint16 - false", input: Variant{VariantType: VAR_INT, VariantValue: uint16(0)}, expectedValue: false},
		{desc: "succeed: from uint32 - false", input: Variant{VariantType: VAR_INT, VariantValue: uint32(0)}, expectedValue: false},
		{desc: "succeed: from uint64 - false", input: Variant{VariantType: VAR_INT, VariantValue: uint64(0)}, expectedValue: false},

		{
			desc:          "fail: from string",
			input:         Variant{VariantType: VAR_STRING, VariantValue: "Some Random String"},
			expectedError: buildTypeError(VAR_STRING, targetType),
		},
		{
			desc:          "fail: inconsistent type",
			input:         Variant{VariantType: targetType, VariantValue: "Some Random String"},
			expectedError: buildInconsistentTypeError("Some Random String", targetType),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			if actualValue, actualError := test.input.CoerceToBool(); actualError == nil {
				assert.Equal(t, test.expectedValue, actualValue, "fail")
			} else {
				assert.EqualError(t, test.expectedError, actualError.Error(), actualError.Error())
			}
		})
	}
}

func TestCoerceToFloat(t *testing.T) {
	targetType := VAR_FLOAT
	tests := [...]struct {
		desc          string
		input         Variant
		expectedValue float64
		expectedError error
	}{
		{desc: "succeed: from float32", input: Variant{VariantType: VAR_FLOAT, VariantValue: float32(3.14159268979323846)}, expectedValue: float64(3.1415927410125732)},
		{desc: "succeed: from float64", input: Variant{VariantType: VAR_FLOAT, VariantValue: float64(3.14159268979323846)}, expectedValue: float64(3.14159268979323846)},

		{desc: "succeed: from int", input: Variant{VariantType: VAR_INT, VariantValue: int(99)}, expectedValue: float64(99)},
		{desc: "succeed: from int8", input: Variant{VariantType: VAR_INT, VariantValue: int8(99)}, expectedValue: float64(99)},
		{desc: "succeed: from int16", input: Variant{VariantType: VAR_INT, VariantValue: int16(99)}, expectedValue: float64(99)},
		{desc: "succeed: from int32", input: Variant{VariantType: VAR_INT, VariantValue: int32(99)}, expectedValue: float64(99)},
		{desc: "succeed: from int64", input: Variant{VariantType: VAR_INT, VariantValue: int64(99)}, expectedValue: float64(99)},
		{desc: "succeed: from uint", input: Variant{VariantType: VAR_INT, VariantValue: uint(99)}, expectedValue: float64(99)},
		{desc: "succeed: from uint8", input: Variant{VariantType: VAR_INT, VariantValue: uint8(99)}, expectedValue: float64(99)},
		{desc: "succeed: from uint16", input: Variant{VariantType: VAR_INT, VariantValue: uint16(99)}, expectedValue: float64(99)},
		{desc: "succeed: from uint32", input: Variant{VariantType: VAR_INT, VariantValue: uint32(99)}, expectedValue: float64(99)},
		{desc: "succeed: from uint64", input: Variant{VariantType: VAR_INT, VariantValue: uint64(99)}, expectedValue: float64(99)},

		{desc: "succeed: from bool - true", input: Variant{VariantType: VAR_BOOL, VariantValue: true}, expectedValue: float64(1)},
		{desc: "succeed: from bool - false", input: Variant{VariantType: VAR_BOOL, VariantValue: false}, expectedValue: float64(0)},

		{
			desc:          "fail: from string",
			input:         Variant{VariantType: VAR_STRING, VariantValue: "Some Random String"},
			expectedError: buildTypeError(VAR_STRING, targetType),
		},
		{
			desc:          "fail: inconsistent type",
			input:         Variant{VariantType: targetType, VariantValue: "Some Random String"},
			expectedError: buildInconsistentTypeError("Some Random String", targetType),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			if actualValue, actualError := test.input.CoerceToFloat(); actualError == nil {
				assert.Equal(t, test.expectedValue, actualValue, "fail")
			} else {
				assert.EqualError(t, test.expectedError, actualError.Error(), actualError.Error())
			}
		})
	}
}

func TestCoerceToInt(t *testing.T) {
	targetType := VAR_INT
	tests := [...]struct {
		desc          string
		input         Variant
		expectedValue int64
		expectedError error
	}{
		{desc: "succeed: from int", input: Variant{VariantType: VAR_INT, VariantValue: int(99)}, expectedValue: int64(99)},
		{desc: "succeed: from int8", input: Variant{VariantType: VAR_INT, VariantValue: int8(99)}, expectedValue: int64(99)},
		{desc: "succeed: from int16", input: Variant{VariantType: VAR_INT, VariantValue: int16(99)}, expectedValue: int64(99)},
		{desc: "succeed: from int32", input: Variant{VariantType: VAR_INT, VariantValue: int32(99)}, expectedValue: int64(99)},
		{desc: "succeed: from int64", input: Variant{VariantType: VAR_INT, VariantValue: int64(99)}, expectedValue: int64(99)},
		{desc: "succeed: from uint", input: Variant{VariantType: VAR_INT, VariantValue: uint(99)}, expectedValue: int64(99)},
		{desc: "succeed: from uint8", input: Variant{VariantType: VAR_INT, VariantValue: uint8(99)}, expectedValue: int64(99)},
		{desc: "succeed: from uint16", input: Variant{VariantType: VAR_INT, VariantValue: uint16(99)}, expectedValue: int64(99)},
		{desc: "succeed: from uint32", input: Variant{VariantType: VAR_INT, VariantValue: uint32(99)}, expectedValue: int64(99)},
		{desc: "succeed: from uint64", input: Variant{VariantType: VAR_INT, VariantValue: uint64(99)}, expectedValue: int64(99)},

		{desc: "succeed: from bool - true", input: Variant{VariantType: VAR_BOOL, VariantValue: true}, expectedValue: int64(1)},
		{desc: "succeed: from bool - false", input: Variant{VariantType: VAR_BOOL, VariantValue: false}, expectedValue: int64(0)},

		{
			desc:          "fail: from string",
			input:         Variant{VariantType: VAR_STRING, VariantValue: "Some Random String"},
			expectedError: buildTypeError(VAR_STRING, targetType),
		},
		{
			desc:          "fail: inconsistent type",
			input:         Variant{VariantType: targetType, VariantValue: "Some Random String"},
			expectedError: buildInconsistentTypeError("Some Random String", targetType),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			if actualValue, actualError := test.input.CoerceToInt(); actualError == nil {
				assert.Equal(t, test.expectedValue, actualValue, "fail")
			} else {
				assert.EqualError(t, test.expectedError, actualError.Error(), actualError.Error())
			}
		})
	}
}

func TestGetDateValue(t *testing.T) {
	sentinel := time.Now()
	targetType := VAR_DATE
	tests := [...]struct {
		desc          string
		input         Variant
		expectedValue time.Time
		expectedError error
	}{
		{desc: "succeed: from date", input: Variant{VariantType: VAR_DATE, VariantValue: sentinel}, expectedValue: sentinel},
		{
			desc:          "fail: from string",
			input:         Variant{VariantType: VAR_STRING, VariantValue: "Some Random String"},
			expectedError: buildTypeError(VAR_STRING, targetType),
		},
		{
			desc:          "fail: inconsistent type",
			input:         Variant{VariantType: targetType, VariantValue: "Some Random String"},
			expectedError: buildInconsistentTypeError("Some Random String", targetType),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			if actualValue, actualError := test.input.GetDateValue(); actualError == nil {
				assert.Equal(t, test.expectedValue, actualValue, "fail")
			} else {
				assert.EqualError(t, test.expectedError, actualError.Error(), actualError.Error())
			}
		})
	}
}

func TestCoerceToString(t *testing.T) {
	sentinel := "Now is the time for all to come to the aid of humanity!"
	tests := [...]struct {
		desc          string
		input         Variant
		expectedValue string
		expectedError error
	}{
		{desc: "succeed: from date", input: Variant{VariantType: VAR_DATE, VariantValue: time.Date(2021, time.August, 19, 10, 13, 15, 0, time.UTC)}, expectedValue: "2021-08-19T10:13:15Z"},

		{desc: "succeed: from float32", input: Variant{VariantType: VAR_FLOAT, VariantValue: float32(3.14159268979323846)}, expectedValue: "3.141593e+00"},
		{desc: "succeed: from float64", input: Variant{VariantType: VAR_FLOAT, VariantValue: float64(3.14159268979323846)}, expectedValue: "3.141593e+00"},

		{desc: "succeed: from int", input: Variant{VariantType: VAR_INT, VariantValue: int(99)}, expectedValue: "99"},
		{desc: "succeed: from int8", input: Variant{VariantType: VAR_INT, VariantValue: int8(99)}, expectedValue: "99"},
		{desc: "succeed: from int16", input: Variant{VariantType: VAR_INT, VariantValue: int16(99)}, expectedValue: "99"},
		{desc: "succeed: from int32", input: Variant{VariantType: VAR_INT, VariantValue: int32(99)}, expectedValue: "99"},
		{desc: "succeed: from int64", input: Variant{VariantType: VAR_INT, VariantValue: int64(99)}, expectedValue: "99"},
		{desc: "succeed: from uint", input: Variant{VariantType: VAR_INT, VariantValue: uint(99)}, expectedValue: "99"},
		{desc: "succeed: from uint8", input: Variant{VariantType: VAR_INT, VariantValue: uint8(99)}, expectedValue: "99"},
		{desc: "succeed: from uint16", input: Variant{VariantType: VAR_INT, VariantValue: uint16(99)}, expectedValue: "99"},
		{desc: "succeed: from uint32", input: Variant{VariantType: VAR_INT, VariantValue: uint32(99)}, expectedValue: "99"},
		{desc: "succeed: from uint64", input: Variant{VariantType: VAR_INT, VariantValue: uint64(99)}, expectedValue: "99"},

		{desc: "succeed: from bool - true", input: Variant{VariantType: VAR_BOOL, VariantValue: true}, expectedValue: "true"},
		{desc: "succeed: from bool - false", input: Variant{VariantType: VAR_BOOL, VariantValue: false}, expectedValue: "false"},

		{desc: "succeed: from string", input: Variant{VariantType: VAR_STRING, VariantValue: sentinel}, expectedValue: sentinel},

		{desc: "succeed: from error", input: Variant{VariantType: VAR_ERROR, VariantValue: errRandom}, expectedValue: errRandom.Error()},
		{
			desc:          "fail: inconsistent type",
			input:         Variant{VariantType: VAR_FLOAT, VariantValue: complex(1.0, 1.0)},
			expectedError: buildInconsistentTypeError("(1+1i)", VAR_STRING),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			if actualValue, actualError := test.input.CoerceToString(); actualError == nil {
				assert.Equal(t, test.expectedValue, actualValue, "fail")
			} else {
				assert.EqualError(t, test.expectedError, actualError.Error(), actualError.Error())
			}
		})
	}
}

func TestGetIdentifierValue(t *testing.T) {
	sentinel := "Now is the time for all to come to the aid of humanity!"
	targetType := VAR_IDENT
	tests := [...]struct {
		desc          string
		input         Variant
		expectedValue string
		expectedError error
	}{
		{desc: "succeed: from ident", input: Variant{VariantType: VAR_IDENT, VariantValue: sentinel}, expectedValue: sentinel},
		{
			desc:          "fail: from date",
			input:         Variant{VariantType: VAR_DATE, VariantValue: time.Now()},
			expectedError: buildTypeError(VAR_DATE, targetType),
		},
		{
			desc:          "fail: inconsistent type",
			input:         Variant{VariantType: targetType, VariantValue: 87},
			expectedError: buildInconsistentTypeError(87, targetType),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			if actualValue, actualError := test.input.GetIdentifierValue(); actualError == nil {
				assert.Equal(t, test.expectedValue, actualValue, "fail")
			} else {
				assert.EqualError(t, test.expectedError, actualError.Error(), actualError.Error())
			}
		})
	}
}

func TestGetErrorValue(t *testing.T) {
	targetType := VAR_ERROR
	tests := [...]struct {
		desc          string
		input         Variant
		expectedValue error
		expectedError error
	}{
		{desc: "succeed: from error", input: Variant{VariantType: VAR_ERROR, VariantValue: errRandom}, expectedValue: errRandom},
		{
			desc:          "fail: from date",
			input:         Variant{VariantType: VAR_DATE, VariantValue: time.Now()},
			expectedError: buildTypeError(VAR_DATE, targetType),
		},
		{
			desc:          "fail: inconsistent type",
			input:         Variant{VariantType: targetType, VariantValue: 87},
			expectedError: buildInconsistentTypeError(87, targetType),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			if actualValue, actualError := test.input.GetErrorValue(); actualError == nil {
				assert.Equal(t, test.expectedValue, actualValue, "fail")
			} else {
				assert.EqualError(t, test.expectedError, actualError.Error(), actualError.Error())
			}
		})
	}
}
