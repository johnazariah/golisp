package golisp

type StringLibrary struct {
}

func (l *StringLibrary) concat(args []Variant) Variant {
	return foldStrings(
		args,
		func(a string, b string) (string, error) { return a + b, nil },
		"concat")
}

func (l *StringLibrary) InjectFunctions(functions FunctionTable) FunctionTable {
	functions["concat"] = l.concat
	functions["++"] = l.concat
	return functions
}
