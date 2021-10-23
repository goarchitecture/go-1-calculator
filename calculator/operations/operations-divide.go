package operations

import "fmt"

type Divide struct {
	x, y float64
}

func NewDivide() *Divide {
	return &Divide{}
}

const OperationDivide = "/"

func (op *Divide) Setup(operands ...float64) Operation {
	// todo: validate
	op.x, op.y = operands[0], operands[1]
	return op
}

func (op *Divide) Match(userOperation string) bool {
	return userOperation == OperationDivide
}

func (op *Divide) Do() (float64, error) {
	if op.y == 0 {
		return 0, fmt.Errorf("can not dividy by zero")
	}

	return op.x / op.y, nil
}
