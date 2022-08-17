package main

import (
	"encoding/json"
	"fmt"
	"sort"

	"golang.org/x/crypto/bcrypt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func (p person) String() string {
	return fmt.Sprintf("%s: %d", p.Last, p.Age)
}

type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func marchal() {
	p1 := person{
		First: "Maria",
		Last:  "Joana",
		Age:   12,
	}
	p2 := person{
		First: "Maria",
		Last:  "Joana",
		Age:   12,
	}
	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(bs))
	}

}

type personType struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func unmarchal() {
	s := `[{"First":"Maria","Last":"Joana","Age":12},{"First":"Maria","Last":"Joana","Age":12}]`
	bs := []byte(s)

	var people []personType
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(people[0])
	}

}

func sortExample() {
	xi := []int{12, 32, 1, 3, 5, 6, 13, 0, 4, 7, 33}
	xs := []string{"James", "Q", "M", "Money", "Dr. No"}

	fmt.Println(xi)
	fmt.Println(xs)

	sort.Ints(xi)
	sort.Strings(xs)

	fmt.Println("SORT:")
	fmt.Println(xi)
	fmt.Println(xs)

}

func customSort() {
	p1 := person{
		First: "Maria",
		Last:  "Ines",
		Age:   12,
	}
	p2 := person{
		First: "Maria",
		Last:  "Tiago",
		Age:   10,
	}
	p3 := person{
		First: "Maria",
		Last:  "Pedro",
		Age:   22,
	}
	p4 := person{
		First: "Maria",
		Last:  "Joao",
		Age:   20,
	}

	people := []person{p1, p2, p3, p4}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
}

func passwordExample() {
	fmt.Println("Password:")
	pass := `password1234`
	bs, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(bs))
	}

	//login example
	fakePassword := "admin"

	err = bcrypt.CompareHashAndPassword(bs, []byte(fakePassword))
	if err != nil {
		fmt.Println("Nice try - error")
	} else {
		fmt.Println("Login OK")
	}
}

func main() {
	fmt.Println("Hello")
	marchal()
	unmarchal()
	sortExample()
	customSort()
	passwordExample()
}
