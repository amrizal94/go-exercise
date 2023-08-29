package main

import (
	"fmt"
	"math/rand"
	"time"
)

var sum = func(ch chan int, numbers ...int) {
	total := 0
	for _, number := range numbers {
		total += number
	}
	ch <- total
}
var berhitung = func(numbers ...int) {
	for _, number := range numbers {
		go func(i int) {
			fmt.Println(i)
		}(number)
	}
}

func main() {
	tick := time.Tick(100 * time.Millisecond)
	door := time.After(500 * time.Millisecond)

	ch := make(chan int, 4)
	numbers := []int{1, 2, 3, 4, 5, 6}
	a := numbers[:len(numbers)/2]
	b := numbers[len(numbers)/2:]
	sum(ch, a...)
	sum(ch, b...)
	sum(ch, numbers...)

	for i := 0; i < 3; i++ {
		fmt.Print(<-ch)

	}
	println()
	fmt.Println("Mulai berhitung...")
	berhitung(numbers...)
	time.Sleep(time.Millisecond)

	go func(numbers ...int) {
		ch <- numbers[0]
	}(numbers[rand.Intn(len(numbers))])

	select {
	case v := <-ch:
		fmt.Println("Select memilih value...", v)
	case <-time.After(time.Duration(rand.Intn(2)) * time.Second):
		fmt.Println("timeout")
	}

	for {
		select {
		case <-tick:
			fmt.Println("tick...")
		case <-door:
			fmt.Println("door...!!!")
			return
		default:
			fmt.Println("	.")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
