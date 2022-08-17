package main

import (
	"fmt"
	"sync"
)

type person struct{
	first string
	last string
	age int
}

func (p *person) speak(){
	fmt.Println("Hello my name is " , p.first)
}

type human interface{
	speak()
}

func saySomething(h human){
	h.speak()
}

func testPointers(){
	p := person{
		first: "Maria",
		last: "Joana",
		age: 12,
	}
	saySomething(&p)
	//ERROR saySomething(p)
	p.speak()
}

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	fmt.Println("Hello from main world")
	go printHelloToAnotherWorld()
	go printHelloToAnotherWorld()
	wg.Wait()
}

func printHelloToAnotherWorld() {
	fmt.Println("Hello from another  world")
	wg.Done()
}
