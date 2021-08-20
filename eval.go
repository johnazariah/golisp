package golisp

type FunctionType func([]Variant) Variant
type FunctionTable map[string]FunctionType

type SymbolTable map[string]Variant

type EvaluationContext struct {
	EvaluatedValue Variant
	Parent         *EvaluationContext
	FunctionTable  FunctionTable
	SymbolTable    SymbolTable
}

func (ctx *EvaluationContext) resolveIdentifier(identifierName string) Variant {
	if ctx == nil {
		return Variant{VariantType: VAR_ERROR, VariantValue: buildUnresolvedIdentifierError(identifierName)}
	}
	if v, e := ctx.SymbolTable[identifierName]; e {
		return v
	}
	if v, e := ctx.FunctionTable[identifierName]; e {
		return Variant{VariantType: VAR_FUNCTION, VariantValue: v}
	}
	return ctx.Parent.resolveIdentifier(identifierName)
}

func loadDefaultLibraries(functions FunctionTable) FunctionTable {
	functions = (&ArithmeticLibrary{}).InjectFunctions(functions)
	functions = (&LogicalLibrary{}).InjectFunctions(functions)
	functions = (&StringLibrary{}).InjectFunctions(functions)
	return functions
}

func loadDefaultSymbols(symbols SymbolTable) SymbolTable {
	return symbols
}

func NewEvaluationContext(parent *EvaluationContext) *EvaluationContext {
	return &EvaluationContext{
		Parent:         parent,
		FunctionTable:  loadDefaultLibraries(FunctionTable{}),
		SymbolTable:    loadDefaultSymbols(SymbolTable{}),
		EvaluatedValue: Variant{VariantType: VAR_UNKNOWN},
	}
}

func (p *null) Eval(ctx *EvaluationContext) *EvaluationContext {
	ctx.EvaluatedValue = Variant{VariantType: VAR_NULL, VariantValue: nil}
	return ctx
}

func (p *atom) Eval(ctx *EvaluationContext) *EvaluationContext {
	switch p.typedValue.VariantType {
	case VAR_IDENT:
		if v, e := p.typedValue.GetIdentifierValue(); e != nil {
			ctx.EvaluatedValue = Variant{VariantType: VAR_ERROR, VariantValue: e}
		} else {
			ctx.EvaluatedValue = ctx.resolveIdentifier(v)
		}
	default:
		ctx.EvaluatedValue = p.typedValue.MakeConsistent()
	}
	return ctx
}

func (p *list) Eval(ctx *EvaluationContext) *EvaluationContext {
	if len(p.children) == 0 {
		ctx.EvaluatedValue = Variant{VariantType: VAR_NULL}
		return ctx
	}

	v := p.children[0].Eval(ctx).EvaluatedValue

	switch v.VariantType {
	case VAR_FUNCTION:
		functionArgs := []Variant{}

		for _, v := range p.children[1:] {
			switch v.(type) {
			case *null, *atom:
				functionArgs = append(functionArgs, v.Eval(ctx).EvaluatedValue)
			case *list:
				functionArgs = append(functionArgs, v.Eval(NewEvaluationContext(ctx)).EvaluatedValue)
			}
		}

		function := v.VariantValue.(FunctionType)
		ctx.EvaluatedValue = function(functionArgs)
	default:
		ctx.EvaluatedValue = Variant{VariantType: VAR_ERROR, VariantValue: buildFunctionNameNotFoundError(v.ToDebugString())}
	}

	return ctx
}
