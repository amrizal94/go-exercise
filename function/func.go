package main

import "fmt"

var sum = func(x ...int) int {
	if len(x) < 2 {
		fmt.Println("minimal inputkan 2 angka dong")
		return -1
	}
	res := 0
	for _, y := range x {
		res += y
	}
	return res
}

func main() {
	fmt.Println(sum(1, 2, 3, 4))
}
