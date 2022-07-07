package main

import "fmt"

//create type
type person struct {
	first string
	last  string
	age   int
}

// podemos ter colisões usamos secretAgent.person.first
type secretAgent struct {
	person
	canKill bool
}

func Structs() {
	//value of type person
	p1 := person{
		first: "James",
		last:  "Bond",
	}
	pS := secretAgent{
		person: person{
			first: "Joao",
			last:  "Mario",
		},
		canKill: true,
	}
	fmt.Println(p1)
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	fmt.Println(pS.last, pS.canKill)
}

func anonymous() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Ines",
		last:  "Joana",
		age:   12,
	}
	fmt.Println(p1)
}

func main() {
	fmt.Println("Olá")
	Structs()
	anonymous()
}
