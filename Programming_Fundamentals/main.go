package main

import (
	"fmt"
	"runtime"
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

var a3 int8
var a4 float64

func numberVars() {
	// you cant:
	// a3 = 21.11111
	a1 := 12
	a2 := 12.222222
	fmt.Printf("%T\n", a1)
	fmt.Printf("%T\n", a2)
}

func runtimeExamples() {
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
}

func stringExamples() {
	s := "examplooo string"
	s = "podes mudar a string"
	// estranho:
	s2 := `"examplooo 
	string"`
	fmt.Printf("%T\n", s)
	fmt.Println(s2)
	// strings to bytes
	bs := []byte(s)
	fmt.Println(bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}

	for i, v := range s {
		fmt.Println(i, v)
	}

}

func numeralSystems() {
	//base 10
	// 0123456789
	//base 2
	// 01
	s := "H"
	fmt.Println(s)
	bs := []byte(s)
	fmt.Println(bs)
}

const a = 42
const str = "ola"

const (
	b1 int = 12
	b2     = 21
)

func constants() {
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", str)
	fmt.Printf("%T\n", b1)
}

// incrementa a variavel c1 = 0, c2 = 1
const (
	c1 = iota
	c2
)

func iotaExample() {
	fmt.Println(c1)
	fmt.Println(c2)
}

func bitShifting() {
	d1 := 3
	fmt.Println("")
	fmt.Printf("%d\t%b", d1, d1)

	d2 := d1 << 1
	fmt.Println("")
	fmt.Printf("%d\t%b", d2, d2)

	kb := 1024
	mb := kb * 1024
	gb := mb * 1024

	fmt.Println()
	fmt.Println(kb)
	fmt.Println(mb)
	fmt.Println(gb)

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)

}

func main() {
	fmt.Println("Hello!")
	boolValues()
	numberVars()
	runtimeExamples()
	stringExamples()
	numeralSystems()
	constants()
	iotaExample()
	bitShifting()
}
