package sci_operations

import (
	"github.com/goarchitecture/go-simple-calculator/calculator/operations"
	"math"
)

type Sin struct {
	x float64
}

func NewSin() *Sin {
	return &Sin{}
}

const OperationSin = "sin"

func (op *Sin) Setup(operands ...float64) operations.Operation {
	// todo: validate
	op.x = operands[0]
	return op
}

func (op *Sin) Match(userOperation string) bool {
	return userOperation == OperationSin
}

func (op *Sin) Do() (float64, error) {
	return math.Sin(op.x), nil
}
