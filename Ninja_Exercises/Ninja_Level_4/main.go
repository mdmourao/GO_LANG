package main

import "fmt"

func Exercise1() {
	var arr [5]int
	arr[0] = 12
	arr[1] = 13
	arr[2] = 14
	arr[3] = 15
	arr[4] = 16
	fmt.Println(arr)
	fmt.Printf("%T\n", arr)

	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func Exercise2() {
	arr := []int{1, 2, 3, 4, 5}

	fmt.Println(arr)
	fmt.Printf("%T\n", arr)

	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func Exercise3() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	sl1 := x[:5]
	fmt.Println(sl1)
	sl2 := x[5:]
	fmt.Println(sl2)
	sl3 := x[2:7]
	fmt.Println(sl3)
	sl4 := x[1:6]
	fmt.Println(sl4)
}

func Exercise4() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	x1 := []int{56, 57, 58, 59, 60}
	x = append(x, x1...)
	fmt.Println(x)
}

func Exercise5() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}

func Exercise6() {
	x := make([]int, 10, 11)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 1)
	// vai duplicar a slice! (implica poder computacional a copiar para outra array)
	x = append(x, 2)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}

func Exercise7() {
	sl1 := []string{"j", "b", "sns"}
	sl2 := []string{"m", "m", "h"}
	x := [][]string{sl1, sl2}
	for i, v := range x {
		fmt.Println(i, v)
		for i1, v2 := range v {
			fmt.Println(i1, v2)
		}
	}
}

func Exercise8() {
	mp := map[string][]string{
		"joana": []string{"ola", "adeus"},
	}
	fmt.Println(mp)
	for k, v := range mp {
		fmt.Println("key", k)
		for _, v1 := range v {
			fmt.Println(v1)
		}
	}
}

func Exercise9() {
	mp := map[string][]string{
		"joana": []string{"ola", "adeus"},
	}
	mp["ines"] = []string{"ola2", "adeus2"}

	for k, v := range mp {
		fmt.Println("key", k)
		for _, v1 := range v {
			fmt.Println(v1)
		}
	}
}

func Exercise10() {
	mp := map[string][]string{
		"joana": []string{"ola", "adeus"},
	}
	mp["ines"] = []string{"ola2", "adeus2"}

	for k, v := range mp {
		fmt.Println("key", k)
		for _, v1 := range v {
			fmt.Println(v1)
		}
	}
	fmt.Println("DELETE")
	delete(mp, "ines")

	for k, v := range mp {
		fmt.Println("key", k)
		for _, v1 := range v {
			fmt.Println(v1)
		}
	}
}

func main() {
	fmt.Println("Hello")
	Exercise1()
	Exercise2()
	Exercise3()
	Exercise4()
	Exercise5()
	Exercise6()
	Exercise7()
	Exercise8()
	Exercise9()

	Exercise10()
}
