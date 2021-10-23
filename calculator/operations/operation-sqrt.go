package operations

import "math"

type Sqrt struct {
	x float64
}

func NewSqrt() *Sqrt {
	return &Sqrt{}
}

const OperationSqrt = "S"

func (op *Sqrt) Setup(operands ...float64) Operation {
	// todo: validate
	op.x = operands[0]
	return op
}

func (op *Sqrt) Match(userOperation string) bool {
	return userOperation == OperationSqrt
}

func (op *Sqrt) Do() (float64, error) {
	return math.Sqrt(op.x), nil
}
