package golisp

import (
	"fmt"
)

type FunctionType func([]Variant) Variant
type FunctionTable map[string]FunctionType

type EvaluationContext struct {
	EvaluatedValue Variant
	Parent         *EvaluationContext
	FunctionTable  FunctionTable
	SymbolTable    map[string]Variant
}

func loadDefaultLibraries(functions FunctionTable) FunctionTable {
	functions = (&AdditionLibrary{}).InjectFunctions(functions)
	functions = (&LogicalLibrary{}).InjectFunctions(functions)
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
	functionName := p.children[0].Eval(ctx).EvaluatedValue.ExtractString()
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
		ctx.EvaluatedValue = Variant{VariantType: VAR_ERROR, VariantValue: fmt.Errorf("requested function %q not found", functionName)}
	}

	return ctx
}
