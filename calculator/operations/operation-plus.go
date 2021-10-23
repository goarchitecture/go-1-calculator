package operations

import "fmt"

type Plus struct {
	x, y float64

	setupErr error
}

func NewPlus() *Plus {
	return &Plus{}
}

const OperationPlus = "+"

func (op *Plus) Setup(operands ...float64) Operation {
	if len(operands) < 2 {
		op.setupErr = fmt.Errorf("invalid setup: inputs must be = 2")
		return op
	}
	op.x, op.y = operands[0], operands[1]
	return op
}

func (op *Plus) Match(userOperation string) bool {
	return userOperation == OperationPlus
}

func (op *Plus) Do() (float64, error) {
	if op.setupErr != nil {
		return 0, op.setupErr
	}

	return op.x + op.y, nil
}
