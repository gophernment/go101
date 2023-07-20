package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sum("1", "2"))
	fmt.Println(sum("a", "2"))
	fmt.Println(sum("1", "b"))
}

func sum(leftOperand, rightOperand string) int {
	var operand1, operand2 int = 0, 0

	if operand1, err := strconv.Atoi(leftOperand); err != nil {
		operand1 = 0
	} else {
		operand1 = operand1
		if operand2, err := strconv.Atoi(rightOperand); err != nil {
			operand2 = 0
		} else {
			operand2 = operand2
		}
	}

	return operand1 + operand2
}
