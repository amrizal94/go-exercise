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

var concat = func(s ...string) string {
	if len(s) <= 1 {
		return fmt.Sprintln("kasihlah 2 kata biar bisa digabungkan")
	}
	for i := 1; i < len(s); i++ {
		s[0] += s[i]
	}
	return s[0]
}

func main() {
	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(concat("hello", "world"))
}
