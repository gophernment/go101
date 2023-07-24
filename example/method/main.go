package main

import (
	"fmt"
	"strconv"
)

type Int int

func (i Int) toString() string {
	return strconv.Itoa(int(i))
}

func main() {
	var fn func(Int) string
	fn = Int.toString

	fmt.Println(fn(Int(9)))
}
