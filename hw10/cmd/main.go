package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type player struct {
	name  string
	score int
}

func gameplay(game chan string) {
	var p1, p2 player
	p1.name = "Mike"
	p2.name = "Max"
	p1.score, p2.score = 0, 0
	var randomHit = 0
	for {
		if (p1.score == 11) || (p2.score == 11) {
			break
		}
		game <- p1.name + " pong"
		randomHit = rand.Intn(4)
		if randomHit == 0 || randomHit == 1 {
			p1.score += 1
			game <- "stop"
		}
		game <- p2.name + " pong"
		randomHit = rand.Intn(4)
		if randomHit == 0 || randomHit == 1 {
			p2.score += 1
			game <- "stop"
		}
		if (p1.score == 11) || (p2.score == 11) {
			break
		}
		fmt.Println("________Next round________")
	}
	game <- p1.name + " " + strconv.Itoa(p1.score) + ":" + strconv.Itoa(p2.score) + " "

	close(game)
}

func main() {
	game := make(chan string)

	go gameplay(game)

	for val := range game {
		fmt.Println(val)
	}
}
