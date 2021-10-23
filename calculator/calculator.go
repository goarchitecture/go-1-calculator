package calculator

import (
	"fmt"
	"github.com/goarchitecture/go-simple-calculator/calculator/operations"
	"strconv"
)

type Calculator struct {
	x, y float64
	op   string

	operations []operations.Operation
}

func NewCalculator(op ...operations.Operation) *Calculator {
	return &Calculator{operations: op}
}

func NewStandardCalculator() *Calculator {
	return NewCalculator(
		operations.NewPlus(), operations.NewMinus(),
		operations.NewMultiply(), operations.NewSqrt(), operations.NewDivide(),
	)
}

func (c *Calculator) AppendOperation(op operations.Operation) *Calculator {
	c.operations = append(c.operations, op)
	return c
}

func (c *Calculator) ReadInput() (err error) {
	fmt.Println("Enter expression (a * b): ")

	var xStr, yStr string
	if _, err = fmt.Scanf("%s %s %s", &xStr, &c.op, &yStr); err != nil {
		err = fmt.Errorf("не верный формат ввода")
		return
	}

	if c.x, err = strconv.ParseFloat(xStr, 64); err != nil {
		err = fmt.Errorf("first operand is not a number")
		return
	}

	if c.y, err = strconv.ParseFloat(yStr, 64); err != nil {
		err = fmt.Errorf("second operand is not a number")
		return
	}

	return
}

func (c *Calculator) Do() (result float64, err error) {
	for _, op := range c.operations {
		if op.Match(c.op) {
			return op.Setup(c.x, c.y).Do()
		}
	}

	err = fmt.Errorf("operation is not supported")
	return
}
