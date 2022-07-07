package main

import "fmt"

// docs says: dont use arrays, use slices
func arrayExamples() {
	var x [5]int
	fmt.Println(x)
	fmt.Println(x[0])
	x[0] = 12
	fmt.Println(x)

	fmt.Println(len(x))
}

// arrays : var x [5]int (specify the size)
// slices : x :=  []int{1,1,2} (dont specify the size)

func SlicesExamples() {
	//x := type{values}
	x := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[1])
	// for range
	for i, v := range x {
		fmt.Println(i, v)
	}
	//slice a slice
	fmt.Println(x[1:5])

	// exercise
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	// append

	x = append(x, 7, 8, 9)
	fmt.Println(x)

	y := []int{10, 11, 12}
	x = append(x, y...)
	fmt.Println(x)

	// delete

	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	//make (,len,capacity)
	a := make([]int, 10, 100)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(a)
	// cant do:
	// a[10] = 11
	// can do is:
	a = append(a, 12)
	fmt.Println(a)

	// slice multidimensional
	sl := []string{"ola", "adeus"}
	fmt.Println(sl)
	sl2 := []string{"bom dia", "boa tarde"}
	fmt.Println(sl2)

	xp := [][]string{sl, sl2}
	fmt.Println(xp)

}

func mapsExamples() {
	//key, value
	m := map[string]int{
		"Martim": 12,
		"Ines":   11,
	}

	fmt.Println(m)
	fmt.Println(m["Martim"])
	fmt.Println(m["NAOEXISTE"])

	v, ok := m["NAOEXISTE"]
	fmt.Println(v)
	fmt.Println(ok)
	if v, ok := m["Martim"]; ok {
		fmt.Println("existe e o valor é ", v)
	}

	// add and range map

	m["todd"] = 177
	fmt.Println(m["todd"])
	for i, v := range m {
		fmt.Println(i, v)
	}

	// delete

	delete(m, "todd")
	fmt.Println(m)
}

func main() {
	fmt.Println("Hello!")
	arrayExamples()
	SlicesExamples()
	mapsExamples()
}
