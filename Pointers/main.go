package main

import "fmt"

func foo(y *int) {
	fmt.Println("before", y)
	fmt.Println("before", *y)
	*y = 43
	fmt.Println("after", y)
	fmt.Println("after", *y)
}

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	fmt.Println(*b)
	x := 0
	foo(&x)
	fmt.Println(x)
}
