package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(sum("1+2"))
	fmt.Println(sum("a+2"))
	fmt.Println(sum("1+b"))
}

func sum(s string) (count int, err error) {
	var operand1, operand2 int
	operands := strings.Split(s, "+")
	if operand1, err := strconv.Atoi(operands[0]); err != nil {
		return operand1, err
	}
	operand2, err = strconv.Atoi(operands[0])

	return operand1 + operand2, err
}
