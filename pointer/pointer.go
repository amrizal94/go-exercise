package main

import "fmt"

var noPointer = func(y byte) {
	fmt.Println(y)
	y = 90
	fmt.Println(y)
}

func main() {
	var x uint8
	fmt.Println("No Pointer")
	noPointer(x)
	fmt.Println(x)
	fmt.Println("Pointer")
	pointer(&x)
	fmt.Println(x)
}

var pointer = func(y *byte) {
	fmt.Println(*y)
	*y = 90
	fmt.Println(*y)
}
