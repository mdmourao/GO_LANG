package main

import "testing"

// go test

func TestMySumTable(t *testing.T) {
	type test struct{
		data []int
		answer int
	}

	tests := []test{
		test{[]int{1,2,3},6},
		test{[]int{1,1,1},3},
		test{[]int{1,2},3},
	}

	for _, v := range tests{
		x := mySum(v.data...)
		if(x != v.answer){
			t.Error("error with values: ", v.data)
		}
	}

	if mySum(2, 3) != 5 {
		t.Error("Error 2+3 result")
	}
}
