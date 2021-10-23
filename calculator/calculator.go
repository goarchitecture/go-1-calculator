package calculator

import (
	"fmt"
	"strconv"
)

type Calculator struct {
	x, y float64
	op   string
}

func NewCalculator() *Calculator {
	return &Calculator{}
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
	switch c.op {
	case "+":
		result = c.x + c.y
	case "-":
		result = c.x - c.y
	case "/":
		if c.y == 0 {
			err = fmt.Errorf("can no divide by zero")
			return
		}
		result = c.x / c.y
	case "*":
		result = c.x * c.y
	default:
		err = fmt.Errorf("unsupported operation")
		return
	}

	return
}
