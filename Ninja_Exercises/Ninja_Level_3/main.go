package main

import "fmt"

func Exercise1() {
	for i := 0; i < 100001; i++ {
		fmt.Println(i)
	}
}

func Exercise2() {
	for i := 65; i <= 90; i++ {
		for i2 := 0; i2 < 3; i2++ {
			fmt.Printf("%#U", i)
		}
	}
}

func Exercise3() {
	a := 2000
	for a <= 2022 {
		fmt.Println(a)
		a++
	}
}

func Exercise4() {
	a := 2000
	for {
		if a > 2022 {
			break
		}
		fmt.Println(a)
		a++
	}
}

func Exercise5() {
	for i := 0; i < 101; i++ {
		fmt.Println(i % 4)
	}
}

func Exercise6() {
	if true == false {
		fmt.Println("mundo acabou!")
	}
}

func Exercise7() {
	if true == false {
		fmt.Println("mundo acabou!")
	} else if true == true {
		fmt.Println("mundo esta ok")
	} else {
		fmt.Println("mundo acabou again")
	}
}

func Exercise8() {
	switch {
	case false == false:
		fmt.Println("expressão verdadeira")
	}
}

func Exercise9() {
	favSport := "futebol"
	switch favSport {
	case "tenis":
		fmt.Println("era tenis")
	case "futebol":
		fmt.Println("era futebol")
	}
}

func Exercise10() {
	//true
	// false
	// true
	// true
	// false
}

func main() {
	fmt.Println("Hello")
	//Exercise1()
	Exercise2()
	Exercise3()
	Exercise4()
	Exercise5()
	Exercise6()
	Exercise7()
	Exercise8()
	Exercise9()
}
