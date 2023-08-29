package main

import (
	"fmt"
	"time"
)

var say = func(word string) {
	for i := 0; i < 5; i++ {
		fmt.Println(word)
		time.Sleep(time.Millisecond)

	}
}

func main() {
	go say("Word")
	say("Hello")
}
