package golisp

import "fmt"

func buildUnexpectedEndOfStringError() error {
	return fmt.Errorf("parse error: unexpected end of string")
}

func buildUnexpectedCloseParenError() error {
	return fmt.Errorf("parse error: unexpected close paren")
}

func buildUnexpectedTrailingTextError() error {
	return fmt.Errorf("parse error: unexpected trailing text")
}

func buildUnhandledVariantTypeError() error {
	return fmt.Errorf("dev error: unhandled variant type")
}

func buildVariantShouldNotBeMaxError() error {
	return fmt.Errorf("dev error: should never have type VAR_MAX")
}

func buildMaximumArityError(arity int, functionName string) error {
	return fmt.Errorf("arity error: expected at most %d arguments for %q", arity, functionName)
}

func buildMinimumArityError(arity int, functionName string) error {
	return fmt.Errorf("arity error: expected at least %d arguments for %q", arity, functionName)
}

func buildExactArityError(arity int, functionName string) error {
	return fmt.Errorf("arity error: expected exactly %d arguments for %q", arity, functionName)
}

func buildArityError_1or2(functionName string) error {
	return fmt.Errorf("arity error: expected exactly 1 or 2 arguments for %q", functionName)
}

func buildForbiddenTypeError(functionName string, forbiddenType EnumVariantType) error {
	return fmt.Errorf("type error: argument to %q can never be of type %q", functionName, forbiddenType)
}

func buildUnacceptableTypeError(variantType EnumVariantType, functionName string) error {
	return fmt.Errorf("type error: argument of unacceptable type %q passed to %q", variantType, functionName)
}

func buildInconsistentTypeError(variantValue interface{}, variantType EnumVariantType) error {
	return fmt.Errorf("type error: value [%v] is inconsistent with type %q", variantValue, variantType)
}

func buildTypeError(variantType EnumVariantType, expectedType EnumVariantType) error {
	return fmt.Errorf("type error: cannot represent variant of type %q as %q", variantType, expectedType)
}

func buildUnresolvedIdentifierError(identifier string) error {
	return fmt.Errorf("scope error: unresolved identifier %q", identifier)
}

func buildFunctionNameNotFoundError(identifier string) error {
	return fmt.Errorf("scope error: requested function %q not found", identifier)
}

func buildDivideByZeroError() error {
	return fmt.Errorf("math error: attempt to divide by zero")
}

func buildGetPromotedNumberTypeReturnedInvalidType() error {
	return fmt.Errorf("panic: getPromotedNumberType returned something other than VAR_INT and VAR_FLOAT")
}
