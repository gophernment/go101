package main

import "fmt"

func main() {
	prime(100)
}

func prime(max int) {
	for i := 2; i <= max; i++ {
		count := 0
		for j := i; j > 0; j-- {
			if i%j == 0 {
				count++
			}
		}
		if count == 2 {
			fmt.Print(i)
		}
	}
}
