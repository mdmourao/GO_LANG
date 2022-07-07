package main

import "fmt"

func loops() {
	// no while :(

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i <= 10; i++ {
		for i2 := 0; i2 <= 2; i2++ {
			fmt.Println(i, i2)
		}
	}

	// we can use for as a while

	a := 1
	b := 4
	for a < b {
		fmt.Println(a)
		a++
	}

	// different way:

	c := 1
	d := 4
	for {
		if c > d {
			break
		}
		fmt.Println(c)
		c++
	}

	// continue

	x := 43 % 40
	y := 43 / 40
	fmt.Println(x, y)

	x = 0
	y = 100
	fmt.Println(x, y)

	for {
		// precisa de estar aqui pq o continue ia fazer skip dele
		x++
		if x > y {
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}

}

func loopsASCII() {
	for i := 33; i < 123; i++ {
		fmt.Printf("%#U", i)
	}
	fmt.Println("")
}

func conditional() {
	if true {
		fmt.Println("001")
	}
	if false {
		fmt.Println("002")
	}
	if !true {
		fmt.Println("003")
	}
	if !false {
		fmt.Println("004")
	}

	if 2 == 2 {
		fmt.Println("005")
	}
	if 2 < 1 {
		fmt.Println("0096")
	}
	// criei a variavel z no if e apenas posso usar dentro do if!!!! muito util
	if z := 12; z == 12 {
		fmt.Println("criei e usei")
	}

	if true != false {
		fmt.Println("Sempre")
	} else {
		fmt.Println("nunca")
	}

	x := 12

	if x == 10 {
		fmt.Println("era 10")
	} else if x == 11 {
		fmt.Println("era 11")
	} else {
		fmt.Println("não era 10 nem 11")
	}

	if (1 < 0 && 2 > 1) {
		fmt.Println("bla bla")
	}

	if (1 < 0 || 2 > 1) {
		fmt.Println("bla bla")
	}

}

func switchExamples() {
	switch {
	case false:
		fmt.Println("false")
	case 2 == 2:
		fmt.Println("prints")
	case 3 == 3:
		fmt.Println("also true, no print")
	}

	//PERIGOSO!!!!
	switch {
	case false:
		fmt.Println("false")
	case 2 == 2:
		fmt.Println("prints")
		fallthrough
	case 3 == 3:
		fmt.Println("also true, no print")
		fallthrough
	case (4 != 3):
		fmt.Println("falso, mas imprime")
	}

	// default case
	switch {
	case false:
		fmt.Println("false")
	case 2 == 2:
		fmt.Println("prints")
		fallthrough
	default:
		fmt.Println("default")
	}

	str := "ola"

	switch str {
	case "adeus":
		fmt.Println("disse adeus")
	case "ola":
		fmt.Println("disse ola")
	case "hello", "bye":
		fmt.Println("fala ingles e disse hello ou bye")
	default:
		fmt.Println("não disse nem ola, nem adeus")

	}
}

func main() {
	fmt.Println("Hello!")
	loops()
	loopsASCII()
	conditional()
	switchExamples()
}
