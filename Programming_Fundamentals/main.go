package main

import (
	"fmt"
)

var x bool

func boolValues() {
	x = true
	a := 7
	b := 42
	fmt.Println(x)
	fmt.Println(a == b)
	fmt.Println(a <= b)
}

func main() {
	fmt.Println("Hello!")
	boolValues()
}
