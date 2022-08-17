package main

import "testing"

// go test
// golint
//  go test -cover
// we can also do bench

//go test -coverprofile c.out
//go tool cover -html=c.out

func TestMySum(t *testing.T) {
	if mySum(2, 3) != 5 {
		t.Error("Error 2+3 result")
	}
}
