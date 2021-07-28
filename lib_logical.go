package golisp

type LogicalLibrary struct {
}

func (l *LogicalLibrary) ensureBooleanArgs(args []Variant, functionName string) error {
	return ensureArgumentTypesMatch(args, []EnumVariantType{VAR_BOOL, VAR_INT}, []EnumVariantType{}, functionName)
}

func (l *LogicalLibrary) xor(args []Variant) Variant {
	functionName := "xor"

	if e := ensureMinimimArity(args, 2, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	if e := ensureMaximumArity(args, 2, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	if e := l.ensureBooleanArgs(args, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	v0, e := args[0].ExtractBool()
	if e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}
	v1, e := args[1].ExtractBool()
	if e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	res := v1 != v0
	return Variant{VariantType: VAR_BOOL, VariantValue: res}
}

func (l *LogicalLibrary) and(args []Variant) Variant {
	functionName := "and"

	if e := ensureMinimimArity(args, 2, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	if e := l.ensureBooleanArgs(args, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	v, e := args[0].ExtractBool()
	if e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}
	res := v
	for _, a := range args[1:] {
		v, e := a.ExtractBool()
		if e != nil {
			return Variant{VariantType: VAR_ERROR, VariantValue: e}
		}
		res = res && v
	}
	return Variant{VariantType: VAR_BOOL, VariantValue: res}
}

func (l *LogicalLibrary) or(args []Variant) Variant {
	functionName := "or"

	if e := ensureMinimimArity(args, 2, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	if e := l.ensureBooleanArgs(args, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	v, e := args[0].ExtractBool()
	if e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}
	res := v
	for _, a := range args[1:] {
		v, e := a.ExtractBool()
		if e != nil {
			return Variant{VariantType: VAR_ERROR, VariantValue: e}
		}
		res = res || v
	}
	return Variant{VariantType: VAR_BOOL, VariantValue: res}
}

func (l *LogicalLibrary) not(args []Variant) Variant {
	functionName := "not"

	if e := ensureExactArity(args, 1, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	if e := l.ensureBooleanArgs(args, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	v, e := args[0].ExtractBool()
	if e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}
	res := !v

	return Variant{VariantType: VAR_BOOL, VariantValue: res}
}

func (l *LogicalLibrary) nor(args []Variant) Variant {
	functionName := "nor"

	if e := ensureMinimimArity(args, 2, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	if e := l.ensureBooleanArgs(args, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}
	return l.not([]Variant{l.or(args)})
}

func (l *LogicalLibrary) nand(args []Variant) Variant {
	functionName := "nand"

	if e := ensureMinimimArity(args, 2, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	if e := l.ensureBooleanArgs(args, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	return l.not([]Variant{l.and(args)})
}

func (l *LogicalLibrary) xnor(args []Variant) Variant {
	functionName := "xnor"

	if e := ensureMinimimArity(args, 2, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	if e := l.ensureBooleanArgs(args, functionName); e != nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: e}
	}

	return l.not([]Variant{l.xor(args)})
}

func (l *LogicalLibrary) InjectFunctions(functions FunctionTable) FunctionTable {
	functions["or"] = l.or
	functions["nor"] = l.nor
	functions["and"] = l.and
	functions["nand"] = l.nand
	functions["xor"] = l.xor
	functions["xnor"] = l.xnor
	functions["not"] = l.not
	return functions
}
