package main

import (
	"fmt"
	"github.com/goarchitecture/go-simple-calculator/calculator"
	sci_operations "github.com/goarchitecture/go-simple-calculator/sci-ops"
	"github.com/namsral/flag"
)

var isScientific = flag.Bool("sci", false, "is calculator scientific")
var inline = flag.String("inline", "", "inline command")

func main() {
	flag.Parse()

	fmt.Println("Hello calculator")

	calc := calculator.NewStandardCalculator()
	if *isScientific {
		calc.AppendOperation(sci_operations.All()...)
	}

	if err := calc.ReadInput(*inline); err != nil {
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
