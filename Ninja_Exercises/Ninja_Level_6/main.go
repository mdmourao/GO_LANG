package main

import (
	"fmt"
	"math"
)

//Exercise 1

func foo1() int {
	return 7
}

func bar1() (int, string) {
	return 7, "seven"
}

// Exercise 2 (variadic parameter int)

func foo2(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum

}

func bar2(xi []int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum

}

// Exercise 3

// Exercise 4
type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(p.first, p.age)
}

// Exercise 5
type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) calArea() float64 {
	return s.length * s.length
}

func (c circle) calArea() float64 {
	return c.radius * c.radius * math.Pi
}

type shape interface {
	calArea() float64
}

func info(s shape) {
	fmt.Println(s.calArea())
}

// Exercise 6
// Exercise 7
// Exercise 8
func returnsFuncs() func() int {
	return func() int {
		return 8
	}
}

// Exercise 9

func runFuncs(f func() int) int {
	x := f()
	return x
}

// Exercise 10

func encloseScope() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

// Exercise 11

func main() {
	fmt.Println("Exercise 6")
	// Exercise 1
	fmt.Println(foo1())
	fmt.Println(bar1())
	x, s := bar1()
	fmt.Println(x, s)
	// Exercise 2
	xi := []int{1, 3, 4, 5, 6, 67}
	result := foo2(xi...)
	fmt.Println(result)
	result2 := bar2(xi)
	fmt.Println(result2)
	// Exercise 3 (defer)
	xi2 := []int{1, 3, 4, 5, 6, 11}
	defer fmt.Println("defer", foo2(xi2...))
	// Exercise 4
	p := person{
		first: "ines",
		last:  "maria",
		age:   12,
	}
	p.speak()
	// Exercise 5
	sq := square{
		12.2,
	}
	cir := circle{
		3.1,
	}
	info(sq)
	info(cir)
	// Exercise 6/7 (anonymous func)
	f := func() int {
		return 7
	}
	func() {
		fmt.Println(7)
	}()
	fmt.Println(f())
	// Exercise 8
	f8 := returnsFuncs()
	fmt.Println(f8())
	// Exercise 9
	f9 := func() int {
		return 7
	}
	fmt.Println(runFuncs(f9))
	// Exercise 10
	{
		y := 0
		fmt.Println(y)
	}
	a1 := encloseScope()
	a2 := encloseScope()
	fmt.Println("Ex 10")
	fmt.Println(a1())
	fmt.Println(a2())
	fmt.Println(a1())
	fmt.Println(a2())
	fmt.Println(a2())
	// Exercise 11
}
