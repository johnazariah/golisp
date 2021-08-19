package golisp

type FunctionType func([]Variant) Variant
type FunctionTable map[string]FunctionType

type EvaluationContext struct {
	EvaluatedValue Variant
	Parent         *EvaluationContext
	FunctionTable  FunctionTable
	SymbolTable    map[string]Variant
}

func loadDefaultLibraries(functions FunctionTable) FunctionTable {
	functions = (&ArithmeticLibrary{}).InjectFunctions(functions)
	functions = (&LogicalLibrary{}).InjectFunctions(functions)
	functions = (&StringLibrary{}).InjectFunctions(functions)
	return functions
}

func NewEvaluationContext(parent *EvaluationContext) *EvaluationContext {
	functions := loadDefaultLibraries(FunctionTable{})

	return &EvaluationContext{
		Parent:         parent,
		FunctionTable:  functions,
		SymbolTable:    map[string]Variant{},
		EvaluatedValue: Variant{VariantType: VAR_UNKNOWN},
	}
}

func (p *null) Eval(ctx *EvaluationContext) *EvaluationContext {
	ctx.EvaluatedValue = Variant{VariantType: VAR_NULL, VariantValue: nil}
	return ctx
}

func (p *atom) Eval(ctx *EvaluationContext) *EvaluationContext {
	ctx.EvaluatedValue = p.typedValue
	return ctx
}

func (p *list) Eval(ctx *EvaluationContext) *EvaluationContext {
	v := p.children[0].Eval(ctx).EvaluatedValue
	if functionName, err := v.GetIdentifierValue(); err != nil {
		ctx.EvaluatedValue = Variant{VariantType: VAR_ERROR, VariantValue: buildInvalidFunctionNameError(v.ToDebugString())}
	} else {
		functionArgs := []Variant{}

		for _, v := range p.children[1:] {
			switch v.(type) {
			case *null, *atom:
				functionArgs = append(functionArgs, v.Eval(ctx).EvaluatedValue)
			case *list:
				functionArgs = append(functionArgs, v.Eval(NewEvaluationContext(ctx)).EvaluatedValue)
			}
		}

		function, found := ctx.FunctionTable[functionName]
		if found {
			ctx.EvaluatedValue = function(functionArgs)
		} else {
			ctx.EvaluatedValue = Variant{VariantType: VAR_ERROR, VariantValue: buildFunctionNameNotFoundError(v.ToDebugString())}
		}
	}

	return ctx
}
