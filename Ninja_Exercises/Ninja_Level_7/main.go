package main

import "fmt"

func Exercise1() {
	x := 12
	fmt.Println(x)
	fmt.Println(&x)
}

type person struct {
	first string
	last  string
}

func changeMe(p *person) {
	(*p).first = "Joana New"
	p.first = "Joana New 2"
}

func main() {
	fmt.Println("Ex 7")
	Exercise1()
	p := person{
		"ines",
		"maria",
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}
