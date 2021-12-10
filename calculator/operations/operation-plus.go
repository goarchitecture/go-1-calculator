package operations

import "fmt"

const OperationTypePlus OperationType = "+"

type OperationPlus struct {
	x, y float64
}

func NewOperationPlus() *OperationPlus {
	return &OperationPlus{}
}

func (o *OperationPlus) Calc() (float64, error) {
	return o.x + o.y, nil
}

func (o *OperationPlus) Match(opType OperationType) bool {
	return opType == OperationTypePlus
}

func (o *OperationPlus) Init(arguments ...float64) error {
	if len(arguments) != 2 {
		return fmt.Errorf("OperationPlus requires 2 arguments")
	}

	o.x = arguments[0]
	o.y = arguments[1]
	return nil
}
