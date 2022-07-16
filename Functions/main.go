package main

import "fmt"

func foo() {
	fmt.Println("foo func")
}

//PASS by VALUE
func bar(s string) {
	fmt.Println(s)
}

func woo() string {
	return "woooo"
}

//QUE UTIL!!!!!!!!!!
func saa() (string, bool) {
	return "ola", true
}

// slice of int (0 or more values)
func variadicParameter(x ...int) {
	fmt.Printf("%T\n", x)
	fmt.Println(x)
}

func deferExample() {
	fmt.Println("deferExample")
}

func deferExample2() {
	fmt.Println("deferExample2")
}

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	licenceToKill bool
}

func (s secretAgent) speak() {
	fmt.Println("Over Over can i kill??")
	fmt.Println(s.licenceToKill)
}

func (p person) speak() {
	fmt.Println("just a normal person")
	fmt.Println(p.first)
}

// if a type has(implements) speaks - its human
type human interface {
	speak()
}

func testHuman(h human) {
	fmt.Println("Hello humans")
}

func testType(h human) {
	switch h.(type) {
	case person:
		fmt.Println(h.(person).first)
	case secretAgent:
		fmt.Println(h.(secretAgent).licenceToKill)
	}
}

// return func
func returnString() string {
	return "funcReturn"
}

func returnFunc() func() int {
	return func() int {
		return 12222
	}
}

//func
func even(fSUM func(xi ...int) int, vi ...int) int {
	return fSUM(vi...)
}

//scope example

func increm() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

//recursion (i prefer loops)
func recur(x int) int {
	if x == 0 {
		return x
	} else {
		a := x - 1
		fmt.Println(a)
		return recur(a)
	}
}

func main() {
	fmt.Println("Olá")
	foo()
	bar("bar string")
	fmt.Println(woo())
	fmt.Println(saa())
	xi := []int{1, 2, 3, 4, 5, 6, 7, 12}
	// unpack a slice
	variadicParameter(xi...)
	variadicParameter()
	// run when func is over (maybe useful to close files)
	defer deferExample()
	deferExample2()
	//methods
	sa1 := secretAgent{
		person: person{
			first: "Ines",
			last:  "Maria",
		},
		licenceToKill: false,
	}
	fmt.Println(sa1)
	sa1.speak()
	// interfaces
	p1 := person{
		"ines",
		"joana",
	}
	fmt.Println(p1)
	testHuman(p1)
	testHuman(sa1)
	testType(p1)
	testType(sa1)
	// anonymous func
	func() {
		fmt.Println("secret")
	}()
	func(x int) {
		fmt.Println(x)
	}(12)
	//func expression
	x := func() {
		fmt.Println("Hello")
	}
	x()
	// return func
	s1 := returnString()
	fmt.Println(s1)
	f1 := returnFunc()
	fmt.Println(f1())
	fmt.Println(returnFunc()())
	//(passing func as argument)
	fSUM := func(vi ...int) int {
		total := 0
		for _, v := range vi {
			total += v
		}
		return total
	}
	fmt.Println(even(fSUM, 1, 2, 2, 2, 4))
	{
		y := 12
		fmt.Println(y)
	}
	a := increm()
	fmt.Println(a())
	// output = 1
	fmt.Println(a())
	// output = 2
	fmt.Println(a())
	// output = 3
	b := increm()
	fmt.Println(b())
	// output = 1

	// recursion
	recur(12)
}
