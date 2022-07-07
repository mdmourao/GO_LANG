package main

import "fmt"

type person struct {
	first            string
	last             string
	favoriteIceCream []string
}

func Exercise1() {
	p1 := person{
		first:            "Joana",
		last:             "Ines",
		favoriteIceCream: []string{"morango", "chocolate"},
	}
	fmt.Println(p1)
	for i, v := range p1.favoriteIceCream {
		fmt.Println(i, v)
	}
}

func Exercise2() {
	p1 := person{
		first:            "Joana",
		last:             "Ines",
		favoriteIceCream: []string{"morango", "chocolate"},
	}
	mp := map[string]person{}
	mp[p1.last] = p1
	fmt.Println(mp)
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func Exercise3() {
	tk := truck{
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		fourWheel: false,
	}
	fmt.Println(tk)
	sed := sedan{
		vehicle: vehicle{
			doors: 3,
			color: "blue",
		},
		luxury: true,
	}
	fmt.Println(sed)
}

func Exercise4() {
	p1 := struct {
		first            string
		last             string
		favoriteIceCream []string
	}{
		first:            "Joana",
		last:             "Ines",
		favoriteIceCream: []string{"morango", "chocolate"},
	}
	fmt.Println(p1)
}

func main() {
	fmt.Println("Olá")
	Exercise1()
	Exercise2()
	Exercise3()
	Exercise4()

}
