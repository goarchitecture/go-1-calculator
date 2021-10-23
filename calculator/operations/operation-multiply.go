package operations

type Multiply struct {
	x, y float64
}

func NewMultiply() *Multiply {
	return &Multiply{}
}

const OperationMultiply = "*"

func (op *Multiply) Setup(operands ...float64) Operation {
	// todo: validate
	op.x, op.y = operands[0], operands[1]
	return op
}

func (op *Multiply) Match(userOperation string) bool {
	return userOperation == OperationMultiply
}

func (op *Multiply) Do() (float64, error) {
	return op.x * op.y, nil
}
