package main

import (
	"fmt"
	"log"
	"strings"

	_ "mainpackage/subpackage"
)

type weekday int

const (
	_ weekday = iota
	sunday
	monday
	tuesday
)

func mainOfMap() {
	m := map[string]string{
		"a": "apple",
		"b": "banana",
		"c": "coconut",
		"d": "durian",
		"e": "elderberry",
		"f": "fig",
		"g": "guava",
	}

	if v, ok := m["z"]; ok {
		fmt.Println(v)
	}
}

type aaa bool

func (aaa) Error() string {
	return "bamboo"
}

type rectangle struct {
	w, h float64
}

func (r rectangle) area() float64 {
	return r.w * r.h
}

type triangle struct {
	b, h float64
}

func (r *triangle) area() float64 {
	return r.b * r.h * 0.5
}
func (r *triangle) String() string {
	return fmt.Sprintf("%f", r.b*r.h*0.5)
}

type areaer interface {
	area() float64
}

func PrintArea(rec areaer) {
	fmt.Printf("area = %v\n", rec.area())
}

func doSomething(n int) {
	defer fmt.Println(n)
	defer func() {
		fmt.Println(n)
	}()
	n = n * n
	fmt.Println(n)
}

func init() {
	fmt.Println("---")
}

func main() {
	fmt.Println("Hi")
}

func factory() (func(), func() int) {
	var i int
	return func() {
			i++
		},
		func() int {
			return i
		}
}

func newFunc() func() int {
	a := 0
	return func() int {
		a++
		return a
	}
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}

var y int

func xxx() {
	defer func() {
		y = 20
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	var s []int
	fmt.Println(s[0])
	fmt.Println("done")
}

type String string

func (s *String) Lower() {
	*s = String(strings.ToLower(string(*s)))
}

func varix(n ...int) {
	fmt.Printf("%T %v\n", n, n)
}

func count(i int) int {
	n := 0
	for i := 0; i < i; i++ {
		n += i
	}
	return n
}
