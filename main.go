package main

import (
	"fmt"
	"github.com/goarchitecture/go-simple-calculator/calculator"
	sci_operations "github.com/goarchitecture/go-simple-calculator/sci-ops"
)

func main() {
	fmt.Println("Hello calculator")

	calc := calculator.NewStandardCalculator()
	calc.AppendOperation(sci_operations.NewSin())

	if err := calc.ReadInput(); err != nil {
		fmt.Printf("input failure: %s\n", err)
		return
	}

	result, err := calc.Do()
	if err != nil {
		fmt.Printf("calc failed: %s\n", err)
		return
	}

	if err := Output(result); err != nil {
		fmt.Printf("could not output: %s\n", err)
		return
	}

	fmt.Println("good bye")
}

func Output(result float64) error {
	if _, err := fmt.Printf("Result = %.2f\n", result); err != nil {
		return fmt.Errorf("could not printf result: %w", err)
	}

	return nil
}
