package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ball struct {
	hits          int
	lastPlayerHit string
}

var player = func(name string, table chan *ball, done chan *ball) {
	for {
		select {
		case ball := <-table:
			v := rand.Intn(1000)
			if v%11 == 0 {
				done <- ball
				fmt.Println(name, "\tDrop the ball")
				return
			}
			ball.hits++
			ball.lastPlayerHit = name
			fmt.Println(name, "\thits the ball", ball.hits)
			time.Sleep(50 * time.Millisecond)
			table <- ball

		case <-time.After(2 * time.Second):
			return

		}
	}
}

var referre = func(table chan *ball, done chan *ball) {

	table <- new(ball)
	for {
		if ball, oke := <-done;oke{
			fmt.Println(ball.lastPlayerHit, "\tthe winner")
			return
		}
	}
}

func main() {
	table := make(chan *ball)
	done := make(chan *ball)
	go player("ari", table, done)
	go player("rizal", table, done)

	referre(table, done)
}
