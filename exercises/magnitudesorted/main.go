package main

import "fmt"

func main() {
	arr := [5]int{3, 1, 7, 4, -3}
	arr = bubbleSort(arr)
	fmt.Println(arr)
}

func bubbleSort(arr [5]int) [5]int {
	var swapped bool
	for i := 0; i < len(arr); i++ {
		swapped = false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	return arr
}
