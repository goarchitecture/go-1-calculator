package calculator

import (
	"testing"
)

func TestNewCalculator(t *testing.T) {
	calc := NewCalculator()
	if calc == nil {
		t.Errorf("NewCalculator must not return nil")
	}
}

func TestCalculator_DoValid(t *testing.T) {
	calc := NewCalculator()

	type ValidExample struct {
		x, y     float64
		op       string
		expected float64
	}
	validExamples := []ValidExample{
		{0, 0, "+", 0},
		{0, 0, "-", 0},
		{100, 100, "+", 200},
		{100, 2, "*", 200},
		{100, 200, "/", 0.5},
		// todo: more cases
	}

	for _, ex := range validExamples {
		calc.x = ex.x
		calc.y = ex.y
		calc.op = ex.op

		result, err := calc.Do()
		if err != nil {
			t.Errorf("do must not return err")
		}
		if result != ex.expected {
			t.Errorf("%#v %s %#v: result expected to be %#v, got %#v", ex.x, ex.op, ex.y, ex.expected, result)
		}
	}
}

func TestCalculator_DoInvalid(t *testing.T) {
	calc := NewCalculator()

	type InvalidExample struct {
		x, y float64
		op   string
	}
	invalidExamples := []InvalidExample{
		{100, 0, "/"},
		// todo: more cases
	}

	for _, ex := range invalidExamples {
		calc.x = ex.x
		calc.y = ex.y
		calc.op = ex.op

		_, err := calc.Do()
		if err == nil {
			t.Errorf("%#v %s %#v: do must return err!", ex.x, ex.op, ex.y)
		}
	}
}
