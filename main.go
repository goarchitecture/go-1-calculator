package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello calculator")

	fmt.Println("Enter 2 operands")
	var x, y float64
	if _, err := fmt.Scanf("%f %f", &x, &y); err != nil {
		fmt.Printf("invalid operands: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Enter operation")
	var op string
	if _, err := fmt.Scan(&op); err != nil {
		fmt.Printf("invalid operation: %s\n", err)
		os.Exit(1)
	}

	var result float64
	switch op {
	case "+":
		result = x + y
	case "-":
		result = x - y
	case "/":
		if y == 0 {
			fmt.Println("can not divide by zero")
			os.Exit(1)
		}
		result = x / y
	case "*":
		result = x * y
	}

	fmt.Printf("Result = %.2f\n", result)
}
