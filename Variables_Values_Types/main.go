// we always need a package main to start the app
package main

import (
	"fmt"
)

// to declare variables outside func we need to use var
var y = 12
var z int

func declareVariables() {
	// declare and assigning (short declaration operator)
	x := 52
	fmt.Println(x)
	// assign a new value
	x = 2
	// declare and assigning
	fmt.Println(y)
	fmt.Println(z)
}

func types() {
	var y int = 7
	var str string = "ola!"
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", str)
	// we can't change y (int) to string
}

func useVariables() {
	// every value is the type of "interface{}"
	fmt.Println("Hello world V2.0")
	fmt.Println(true)
	fmt.Println(123)
	// we need to use all the variables in go! (very interesting...)
	// we can use '_' to not get the return
	n, err := fmt.Println("oi")
	fmt.Println(n, err)
}

func packagefmt() {
	var number = 12
	fmt.Printf("%b\n", number)
	fmt.Printf("%x\n", number)
	fmt.Printf("%#x\n", number)
	fmt.Print(number)
	fmt.Println(number)
	fmt.Printf("%b\t%x\t%v\n", number, number, number)
	s := fmt.Sprintf("%b\t%x\t%v\n", number, number, number)
	fmt.Println(s)
}

type newTypeInt int

var b newTypeInt

func createMyOwnType() {
	b = 1
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}

func conversion() {
	// int
	a := 1
	//main.newTypeInt
	b = 2
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	// conversion
	var c int
	c = int(b)
	fmt.Println(c)
}

func main() {
	fmt.Println("Hello World!")
	for i := 0; i <= 2; i++ {
		if i == 0 {
			fmt.Println(i)
		}
	}
	//useVariables()
	//declareVariables()
	//types()
	//packagefmt()
	//createMyOwnType()
	conversion()
}
