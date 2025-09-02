package main

import (
	"fmt"
	"os"
)

var history []string // Global to persist between calls

func main() {
	for {
		calculator()
	}
}

func calculator() {
	var num1, num2 int
	var choice int

	fmt.Println("\nSelect the operation you want to do:")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Modulus")
	fmt.Println("6. Exit")
	fmt.Println("7. View History")

	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)

	// Handle exit
	if choice == 6 {
		fmt.Println("Exiting calculator. Goodbye!")
		os.Exit(0) // properly exit program
	}

	// Show history
	if choice == 7 {
		fmt.Println("Calculation History:")
		if len(history) == 0 {
			fmt.Println("No history available.")
		} else {
			for _, record := range history {
				fmt.Println(record)
			}
		}
		return // skip rest of function
	}

	// Ask for numbers only if not exit/history
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	switch choice {
	case 1:
		result := num1 + num2
		saveAndPrint(num1, num2, "+", result)
	case 2:
		result := num1 - num2
		saveAndPrint(num1, num2, "-", result)
	case 3:
		result := num1 * num2
		saveAndPrint(num1, num2, "*", result)
	case 4:
		if num2 == 0 {
			fmt.Println("Error: Division by zero")
			return
		}
		result := num1 / num2
		saveAndPrint(num1, num2, "/", result)
	case 5:
		if num2 == 0 {
			fmt.Println("Error: Modulus by zero")
			return
		}
		result := num1 % num2
		saveAndPrint(num1, num2, "%", result)
	default:
		fmt.Println("Invalid choice.")
	}
}

func saveAndPrint(a, b int, op string, result int) {
	record := fmt.Sprintf("%d %s %d = %d", a, op, b, result)
	history = append(history, record)
	fmt.Println("The result is:", result)
}
