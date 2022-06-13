package main

import "fmt"

func exercise1() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}

//var x int
//var y string
//var z bool

var x = 42
var y = "James Bond"
var z = true

func exercise2() {
	// the zero-values
	fmt.Println(x, y, z)
	// output: 0 false ""
}

func exercise3() {
	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(s)
}

type myNewType int

var x4 myNewType

func exercise4() {
	fmt.Println(x4)
	fmt.Printf("%T\n", x4)
	x4 = 42
	fmt.Println(x4)
}

var y5 int

func exercise5() {
	y5 = int(x4)
	fmt.Println(y5)
	fmt.Printf("%T\n", y5)
}

func main() {
	//exercise1()
	//exercise2()
	//exercise3()
	//exercise4()
	exercise5()
}
