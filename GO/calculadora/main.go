package main

import (
	"fmt"
)

func main() {
	var op string
	var num1, num2 float64

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)

	fmt.Print("Digite a operação (+, -, *, /): ")
	fmt.Scanln(&op)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)

	var result float64

	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		fmt.Println("Operação inválida")
		return
	}

	fmt.Printf("%v %v %v = %v", num1, op, num2, result)
}
