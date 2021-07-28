package golisp

import "fmt"

type FunctionLibrary interface {
	InjectFunctions(*FunctionTable) *FunctionTable
}

func ensureMaximumArity(args []Variant, arity int, functionName string) error {
	if len(args) <= arity {
		return nil
	}

	return fmt.Errorf("arity error: expected at most %d arguments for %q", arity, functionName)
}

func ensureMinimimArity(args []Variant, arity int, functionName string) error {
	if len(args) >= arity {
		return nil
	}

	return fmt.Errorf("arity error: expected at least %d arguments for %q", arity, functionName)
}

func ensureExactArity(args []Variant, arity int, functionName string) error {
	if len(args) == arity {
		return nil
	}
	return fmt.Errorf("arity error: expected exactly %d arguments for %q", arity, functionName)
}

func ensureTypeIsNotInvalid(a Variant) error {
	switch a.VariantType {
	case VAR_ERROR:
		return fmt.Errorf(a.ExtractString())

	case VAR_MAX:
		return fmt.Errorf("dev error: should never have type VAR_MAX")

	case VAR_IDENT:
		return fmt.Errorf("scope error: unresolved identifier %q", a.ExtractString())
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
				return fmt.Errorf("type error: argument to %q can never be of type %q", functionName, t)
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
			return fmt.Errorf("type error: argument of unacceptable type %q passed to %q", a.VariantType, functionName)
		}
	}

	return nil
}
