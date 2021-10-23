package calculator

import "fmt"

type Calculator struct {
	x, y float64
	op   string
}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (c *Calculator) ReadInput() (err error) {
	fmt.Println("Enter 2 operands")
	if _, err = fmt.Scanf("%f %f", &c.x, &c.y); err != nil {
		err = fmt.Errorf("could not parse operands: %w", err)
		return
	}

	fmt.Println("Enter operation")
	if _, err = fmt.Scan(&c.op); err != nil {
		err = fmt.Errorf("could not parse operator: %w", err)
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
