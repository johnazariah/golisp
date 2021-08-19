package golisp

type FunctionLibrary interface {
	InjectFunctions(*FunctionTable) *FunctionTable
}

func ensureMaximumArity(args []Variant, arity int, functionName string) error {
	if len(args) <= arity {
		return nil
	}
	return buildMaximumArityError(arity, functionName)
}

func ensureMinimimArity(args []Variant, arity int, functionName string) error {
	if len(args) >= arity {
		return nil
	}
	return buildMinimumArityError(arity, functionName)
}

func ensureExactArity(args []Variant, arity int, functionName string) error {
	if len(args) == arity {
		return nil
	}
	return buildExactArityError(arity, functionName)
}

func ensureTypeIsNotInvalid(a Variant) error {
	switch a.VariantType {
	case VAR_MAX:
		return buildVariantShouldNotBeMaxError()

	case VAR_ERROR:
		if v, e := a.GetErrorValue(); e != nil {
			return e
		} else {
			return v
		}

	case VAR_IDENT:
		return buildUnresolvedIdentifierError(a.ToDebugString())
	}
	return nil
}

func ensureArgumentTypesMatch(args []Variant, acceptableTypes []EnumVariantType, forbiddenTypes []EnumVariantType, functionName string) error {
	// check for forbidden types
	for _, a := range args {
		if e := ensureTypeIsNotInvalid(a); e != nil {
			return e
		}

		for _, t := range forbiddenTypes {
			if a.VariantType == t {
				return buildForbiddenTypeError(functionName, t)
			}
		}

		found := false
		for _, t := range acceptableTypes {
			if a.VariantType == t {
				found = true
				break
			}
		}

		if !found {
			return buildUnacceptableTypeError(a.VariantType, functionName)
		}
	}

	return nil
}

func ensureBooleanArgs(args []Variant, functionName string) error {
	return ensureArgumentTypesMatch(args, []EnumVariantType{VAR_BOOL, VAR_INT}, []EnumVariantType{}, functionName)
}

func unaryOpBoolean(args []Variant, unaryOp func(bool) bool, functionName string) Variant {
	if e := ensureExactArity(args, 1, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	if e := ensureBooleanArgs(args, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	v0, e := args[0].CoerceToBool()
	if e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	res := unaryOp(v0)

	return Variant{VariantType: VAR_BOOL, VariantValue: res}
}

func binaryOpBoolean(args []Variant, binaryOp func(bool, bool) bool, functionName string) Variant {
	if e := ensureExactArity(args, 2, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	if e := ensureBooleanArgs(args, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	return foldBooleans(args, binaryOp, functionName)
}

func foldBooleans(args []Variant, bool_folder func(bool, bool) bool, functionName string) Variant {
	if e := ensureMinimimArity(args, 2, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	if e := ensureBooleanArgs(args, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	v, e := args[0].CoerceToBool()
	if e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}
	res := v
	for _, a := range args[1:] {
		v, e := a.CoerceToBool()
		if e != nil {
			return Variant{VariantType: VAR_ERROR, VariantValue: e}
		}
		res = bool_folder(res, v)
	}

	return Variant{VariantType: VAR_BOOL, VariantValue: res}
}

func ensureNumberArgs(args []Variant, functionName string) error {
	return ensureArgumentTypesMatch(args, []EnumVariantType{VAR_BOOL, VAR_INT, VAR_FLOAT}, []EnumVariantType{}, functionName)
}

func getPromotedNumberType(args []Variant, functionName string) (EnumVariantType, error) {
	resultValueType := VAR_UNKNOWN
	for _, a := range args {
		if e := ensureTypeIsNotInvalid(a); e != nil {
			return VAR_ERROR, e
		}

		switch a.VariantType {
		case VAR_BOOL, VAR_INT:
			if (resultValueType == VAR_UNKNOWN) || (resultValueType == VAR_INT) {
				resultValueType = VAR_INT
				continue
			}

		case VAR_FLOAT:
			if (resultValueType == VAR_UNKNOWN) || (resultValueType == VAR_INT) || (resultValueType == VAR_BOOL) || (resultValueType == VAR_FLOAT) {
				resultValueType = VAR_FLOAT
				continue
			}

		default:
			return VAR_ERROR, buildUnacceptableTypeError(a.VariantType, functionName)
		}
	}

	return resultValueType, nil
}

func unaryOpNumber(args []Variant, unaryOpInt func(int64) (int64, error), unaryOpFloat func(float64) (float64, error), functionName string) Variant {
	if e := ensureExactArity(args, 1, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	if e := ensureNumberArgs(args, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	resultValueType, e := getPromotedNumberType(args, functionName)
	if e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	switch resultValueType {
	case VAR_FLOAT:
		res0, e := args[0].CoerceToFloat()
		if e != nil {
			return Variant{VariantType: VAR_ERROR, VariantValue: e}
		}
		res, e := unaryOpFloat(res0)
		if e != nil {
			return Variant{VariantType: VAR_ERROR, VariantValue: e}
		}
		return Variant{VariantType: VAR_FLOAT, VariantValue: res}

	case VAR_INT:
		res0, e := args[0].CoerceToInt()
		if e != nil {
			return Variant{VariantType: VAR_ERROR, VariantValue: e}
		}
		res, e := unaryOpInt(res0)
		if e != nil {
			return Variant{VariantType: VAR_ERROR, VariantValue: e}
		}
		return Variant{VariantType: VAR_INT, VariantValue: res}

	default:
		return Variant{VariantType: VAR_ERROR, VariantValue: buildGetPromotedNumberTypeReturnedInvalidType()}
	}
}

func binaryOpNumbers(args []Variant, int_folder func(int64, int64) (int64, error), float_folder func(float64, float64) (float64, error), functionName string) Variant {
	if e := ensureNumberArgs(args, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	if e := ensureExactArity(args, 2, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	resultValueType, e := getPromotedNumberType(args, functionName)
	if e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	switch resultValueType {
	case VAR_FLOAT:
		return binaryOpFloats(args, float_folder, functionName)

	case VAR_INT:
		return binaryOpInts(args, int_folder, functionName)

	default:
		return Variant{VariantType: VAR_ERROR, VariantValue: buildGetPromotedNumberTypeReturnedInvalidType()}
	}
}

func binaryOpFloats(args []Variant, float_folder func(float64, float64) (float64, error), functionName string) Variant {
	if e := ensureNumberArgs(args, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	if e := ensureExactArity(args, 2, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	return foldFloats(args, float_folder, functionName)
}

func binaryOpInts(args []Variant, int_folder func(int64, int64) (int64, error), functionName string) Variant {
	if e := ensureNumberArgs(args, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	if e := ensureExactArity(args, 2, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	return foldInts(args, int_folder, functionName)
}

func foldNumbers(args []Variant, int_folder func(int64, int64) (int64, error), float_folder func(float64, float64) (float64, error), functionName string) Variant {
	if e := ensureNumberArgs(args, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	resultValueType, e := getPromotedNumberType(args, functionName)
	if e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	switch resultValueType {
	case VAR_FLOAT:
		return foldFloats(args, float_folder, functionName)

	case VAR_INT:
		return foldInts(args, int_folder, functionName)

	default:
		return Variant{VariantType: VAR_ERROR, VariantValue: buildGetPromotedNumberTypeReturnedInvalidType()}
	}
}

func foldFloats(args []Variant, float_folder func(float64, float64) (float64, error), functionName string) Variant {
	v, e := args[0].CoerceToFloat()
	if e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}
	res := v

	for _, a := range args[1:] {
		v, e := a.CoerceToFloat()
		if e != nil {
			return Variant{VariantType: VAR_ERROR, VariantValue: e}
		}

		if res, e = float_folder(res, v); e != nil {
			return Variant{VariantType: VAR_ERROR, VariantValue: e}
		}
	}

	return Variant{VariantType: VAR_FLOAT, VariantValue: res}
}

func foldInts(args []Variant, int_folder func(int64, int64) (int64, error), functionName string) Variant {
	v, e := args[0].CoerceToInt()
	if e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}
	res := v

	for _, a := range args[1:] {
		v, e := a.CoerceToInt()
		if e != nil {
			return Variant{VariantType: VAR_ERROR, VariantValue: e}
		}
		if res, e = int_folder(res, v); e != nil {
			return Variant{VariantType: VAR_ERROR, VariantValue: e}
		}
	}

	return Variant{VariantType: VAR_INT, VariantValue: res}
}

func foldStrings(args []Variant, string_folder func(string, string) (string, error), functionName string) Variant {
	v, e := args[0].CoerceToString()
	if e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}
	res := v

	for _, a := range args[1:] {
		v, e := a.CoerceToString()
		if e != nil {
			return Variant{VariantType: VAR_ERROR, VariantValue: e}
		}

		if res, e = string_folder(res, v); e != nil {
			return Variant{VariantType: VAR_ERROR, VariantValue: e}
		}
	}

	return Variant{VariantType: VAR_STRING, VariantValue: res}
}
