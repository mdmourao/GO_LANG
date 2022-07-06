package main

import "fmt"

func Exercise1() {
	a1 := 12
	// decimal
	fmt.Println(a1)
	// binary
	fmt.Printf("%b\n", a1)
	// hex
	fmt.Printf("%x\n", a1)
	fmt.Printf("%#x\n", a1)
}

func Exercise2() {
	b1 := 1 == 2
	b2 := 1 <= 2
	b3 := 1 >= 2
	b4 := 1 != 2
	b5 := 1 < 2
	b6 := 1 > 2
	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(b4)
	fmt.Println(b5)
	fmt.Println(b5)
	fmt.Println(b6)
}

const (
	a     = 12
	b int = 13
)

func Exercise3() {
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
}

func Exercise4() {
	d1 := 1
	fmt.Printf("%d\t%b\t%#x\n", d1, d1, d1)
	d2 := d1 << 1
	fmt.Printf("%d\t%b\t%#x\n", d2, d2, d2)
}

func Exercise5() {
	// raw string
	a1 := `string 
	ola`
	fmt.Println(a1)

}

const (
	y1 = 2040 + iota
	y2 = 2040 + iota
	y3 = 2040 + iota
	y4 = 2040 + iota
)

func Exercise6() {
	fmt.Println(y1, y2, y3, y4)
}

func main() {
	fmt.Println("Hello")
	Exercise1()
	Exercise2()
	Exercise3()
	Exercise4()
	Exercise5()
	Exercise6()
}
