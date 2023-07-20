package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(generation(40))
}

func generation(age int) string {
	thisYear := time.Now().Year()
	birthYear := thisYear - age

	return "X"
}
