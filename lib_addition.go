package golisp

type AdditionLibrary struct {
}

func (l *AdditionLibrary) getBestResultType(args []Variant) (EnumVariantType, error) {
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
			// coerce everything else to string
			return VAR_STRING, nil
		}
	}

	return resultValueType, nil
}

func (l *AdditionLibrary) add(args []Variant) Variant {
	resultValueType, e := l.getBestResultType(args)
	if e != nil {
		return Variant{VariantType: resultValueType, VariantValue: e}
	}

	switch resultValueType {
	case VAR_FLOAT:
		var res float64 = 0.0
		for _, a := range args {
			v, e := a.ExtractFloat()
			if e != nil {
				return Variant{VariantType: VAR_ERROR, VariantValue: e}
			}
			res += v
		}
		return Variant{VariantType: VAR_FLOAT, VariantValue: res}

	case VAR_INT:
		var res int64 = 0
		for _, a := range args {
			v, e := a.ExtractInt()
			if e != nil {
				return Variant{VariantType: VAR_ERROR, VariantValue: e}
			}
			res += v
		}
		return Variant{VariantType: VAR_INT, VariantValue: res}

	default:
		var res = ""
		for _, a := range args {
			res += a.ExtractString()
		}
		return Variant{VariantType: VAR_STRING, VariantValue: res}
	}
}

func (l *AdditionLibrary) InjectFunctions(functions FunctionTable) FunctionTable {
	functions["+"] = l.add
	return functions
}
