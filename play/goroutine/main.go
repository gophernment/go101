package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int)
	go fibonacciRoutine(ch)
	for i := 0; i < 10; i++ {
		fmt.Print(<-ch, " ")
	}
}

func fibonacciRoutine(ch chan<- int) {
	a, b := 0, 1
	for {
		ch <- a
		a, b = b, a+b
	}
}

func fibonacciFunc() func() int {
	a, b := 0, 1
	return func() int {
		defer func() {
			a, b = b, a+b
		}()
		return a
	}
}

func fibonacci(n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
}

func slow(i int) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	fmt.Println(i)
}

type resource struct {
	s   string
	mux sync.Mutex
}

func (r *resource) Set(s string) {
	r.mux.Lock()
	defer r.mux.Unlock()
	r.s = s
}
