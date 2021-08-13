package golisp

type LogicalLibrary struct {
}

func (l *LogicalLibrary) not(args []Variant) Variant {
	return unaryOpBoolean(
		args,
		func(a bool) bool { return !a },
		"not",
	)
}

func (l *LogicalLibrary) and(args []Variant) Variant {
	return foldBooleans(
		args,
		func(a bool, b bool) bool { return a && b },
		"and",
	)
}

func (l *LogicalLibrary) nand(args []Variant) Variant {
	temp := foldBooleans(
		args,
		func(a bool, b bool) bool { return a && b },
		"nand",
	)

	return l.not([]Variant{temp})
}

func (l *LogicalLibrary) or(args []Variant) Variant {
	return foldBooleans(
		args,
		func(a bool, b bool) bool { return a || b },
		"or",
	)
}

func (l *LogicalLibrary) nor(args []Variant) Variant {
	temp := foldBooleans(
		args,
		func(a bool, b bool) bool { return a || b },
		"nor",
	)

	return l.not([]Variant{temp})
}

func (l *LogicalLibrary) xor(args []Variant) Variant {
	return binaryOpBoolean(
		args,
		func(a bool, b bool) bool { return a != b },
		"xor",
	)
}

func (l *LogicalLibrary) xnor(args []Variant) Variant {
	temp := binaryOpBoolean(
		args,
		func(a bool, b bool) bool { return a != b },
		"xnor",
	)

	return l.not([]Variant{temp})
}

func (l *LogicalLibrary) InjectFunctions(functions FunctionTable) FunctionTable {
	functions["or"] = l.or
	functions["nor"] = l.nor
	functions["and"] = l.and
	functions["nand"] = l.nand
	functions["xor"] = l.xor
	functions["xnor"] = l.xnor
	functions["not"] = l.not
	functions["||"] = l.or
	functions["&&"] = l.and
	functions["^^"] = l.xor
	functions["!"] = l.not
	return functions
}
