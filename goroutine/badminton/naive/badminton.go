package main

import (
	"fmt"
	"time"
)

type cock struct{
	hits int
}

var player = func (name string, raket chan *cock)  {	
for {
 cock := <- raket
 cock.hits++
 fmt.Println(name, "\thits the cock", cock.hits)
 time.Sleep(50 * time.Millisecond)
 raket<-cock
}
}

func main() {
	raket:= make(chan *cock)
	go player("ari", raket)
	go player("rizal", raket)
	raket <- new(cock)
	time.Sleep(1*time.Second)
}
