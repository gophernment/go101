package main

import "fmt"

func main() {
	fmt.Println("Hello.")
	fmt.Print("What is your name?: ")

	var name string
	fmt.Scanln(&name)

	fmt.Printf("Hi %s.\n", name)
}
