package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	data := [][]int{
		{5, 2, 8, 4, 6, 9},
		{1, 5, 9, 3, 6},
		{3, 4, 2, 1, 9, 0},
		{7, 8, 9, 6, 5, 4, 1, 2},
		{11, 6, 3, 99, 5, 8},
	}

	for i := range data {
		data[i] = slowBubbleSort(data[i])
	}

	for _, v := range data {
		fmt.Println(v)
	}

	fmt.Println(time.Since(start))
}

func slowBubbleSort(arr []int) []int {
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
		time.Sleep(time.Millisecond)
	}

	return arr
}
