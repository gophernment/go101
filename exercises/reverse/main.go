package main

import "fmt"

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	answer := "god yzal eht revo depmuj xof nworb kciuq ehT"

	output := Reverse(input)

	fmt.Println(answer)
	fmt.Println(output)
}

func Reverse(s string) string {
	return s
}
