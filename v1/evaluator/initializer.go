package evaluator

import (
	"calculus/v1/core"
	"calculus/v1/std/num"
)

const (
	lg = "lg"
	pow = "pow"
)

func InitializeLibraries(e *evaluator) core.EvaluatorInitializer {
	return func(library ...core.Library) {
		// operators
		e.binaryHandlers[core.ADD] = num.Sum
		e.binaryHandlers[core.MUL] = num.Mul
		e.binaryHandlers[core.SUB] = num.Sub
		e.binaryHandlers[core.DIV] = num.Div
		e.unaryHandler[core.INC] = num.Inc
		e.unaryHandler[core.DEC] = num.Dec

		// functions
		e.callableHandlers[lg] = num.Lg
		e.callableHandlers[pow] = num.Pow
	}
}
