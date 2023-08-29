package main

import "fmt"

var reverse = func(numbers ...byte) {
	for _, v := range numbers {
		defer fmt.Println(v)
	}
}

func main() {
	fmt.Println(1)
	fmt.Println(2)
	defer fmt.Println(3) // dijalakan paling akhir pada function
	fmt.Println(4)
	fmt.Println(5)
	fmt.Println(6)
	fmt.Println(7)
	numbers := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Reverse starting")
	reverse(numbers...)
}
