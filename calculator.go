package main

import (
	"fmt"
)

func main() {
	fmt.Println("Please enter operation, either addition, subtraction, multiplication, or division")
	var operation string
	fmt.Scanln(&operation)
	switch operation {
	case "addition", "+":
		fmt.Println("Enter First Number:")
		var firstNumber int
		fmt.Scanln(&firstNumber)
		fmt.Println("Enter Second Number")
		var secondNumber int
		fmt.Scanln(&secondNumber)
		var (
			sum = firstNumber + secondNumber
		)
		fmt.Print(sum)
		fmt.Println("\n Press enter to close the calculator")
		fmt.Scanln()
	case "subtraction", "-":
		fmt.Println("Enter First Number:")
		var firstNumber int
		fmt.Scanln(&firstNumber)
		fmt.Println("Enter Second Number")
		var secondNumber int
		fmt.Scanln(&secondNumber)
		var (
			sum = firstNumber - secondNumber
		)
		fmt.Print(sum)
		fmt.Println("\n Press enter to close the calculator")
		fmt.Scanln()
	case "multiplication", "*":
		fmt.Println("Enter First Number:")
		var firstNumber int
		fmt.Scanln(&firstNumber)
		fmt.Println("Enter Second Number")
		var secondNumber int
		fmt.Scanln(&secondNumber)
		var (
			sum = firstNumber * secondNumber
		)
		fmt.Print(sum)
		fmt.Println("\n Press enter to close the calculator")
		fmt.Scanln()
	case "division", "/":
		fmt.Println("Enter First Number:")
		var firstNumber int
		fmt.Scanln(&firstNumber)
		fmt.Println("Enter Second Number")
		var secondNumber int
		fmt.Scanln(&secondNumber)
		var (
			sum = firstNumber / secondNumber
		)
		fmt.Print(sum)
		fmt.Println("\n Press enter to close the calculator")
		fmt.Scanln()
	default:
		fmt.Println("not a valid operation, please press enter to close the calculator because you don't deserve it")
		fmt.Scanln()
	}

}
