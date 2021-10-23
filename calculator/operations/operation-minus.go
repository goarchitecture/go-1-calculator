package operations

type Minus struct {
	x, y float64
}

func NewMinus() *Minus {
	return &Minus{}
}

const OperationMinus = "-"

func (op *Minus) Setup(operands ...float64) Operation {
	// todo: validate
	op.x, op.y = operands[0], operands[1]
	return op
}

func (op *Minus) Match(userOperation string) bool {
	return userOperation == OperationMinus
}

func (op *Minus) Do() (float64, error) {
	return op.x - op.y, nil
}
