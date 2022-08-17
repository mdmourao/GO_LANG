package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	n, err := fmt.Println("Error maybe?")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
	openFile()
}

func openFile() {
	f, err := os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}

//Panic - run defer functions (we can use recover inside defer functions)
//Fatal - 100% dead
//log.println() - we can choose the  log file!

func erros(i int) (int, error){
	if i < 0 {
		return 0 , errors.New("Error <0 number")
	}
	return 77, nil
}
