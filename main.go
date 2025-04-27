package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	for {
		var a, b float64
		var op string

		fmt.Println("\nEnter the first number (or 'exit' to quit):")
		var input string
		fmt.Scanln(&input)

		if strings.ToLower(input) == "exit" {
			fmt.Println("Exiting calculator.")
			break
		}

		_, err := fmt.Sscanf(input, "%f", &a)
		if err != nil {
			fmt.Println("Error: enter a valid number or 'exit'!")
			continue
		}

		fmt.Println("Enter operation (+, -, *, /, ^, sqrt):")
		fmt.Scanln(&op)

		if op != "sqrt" {
			fmt.Println("Enter the second number:")
			fmt.Scanln(&b)
		}

		switch op {
		case "+":
			fmt.Printf("Result: %.2f\n", a+b)
		case "-":
			fmt.Printf("Result: %.2f\n", a-b)
		case "*":
			fmt.Printf("Result: %.2f\n", a*b)
		case "/":
			if b == 0 {
				fmt.Println("Error: division by zero!")
				continue
			}
			fmt.Printf("Result: %.2f\n", a/b)
		case "^":
			fmt.Printf("Result: %.2f\n", math.Pow(a, b))
		case "sqrt":
			if a < 0 {
				fmt.Println("Error: square root of a negative number!")
				continue
			}
			fmt.Printf("Result: %.2f\n", math.Sqrt(a))
		default:
			fmt.Println("Invalid operation! Supported: +, -, *, /, ^, sqrt")
		}
	}
}
