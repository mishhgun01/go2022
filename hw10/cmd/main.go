package main

import (
	"fmt"
	"math/rand"
	"sync"
)

//gameplay - структура игры с каналами и группами ожидания, а также счётом
type gameplay struct {
	waitgroup *sync.WaitGroup
	scoreOf   map[string]int
	channel   chan string
}

func main() {
	game := gameplay{
		waitgroup: new(sync.WaitGroup),
		scoreOf:   make(map[string]int),
		channel:   make(chan string),
	}
	game.init()
}

//init инициализирует игру
func (g *gameplay) init() {
	g.waitgroup.Add(2)
	go g.play("Max")
	go g.play("Mike")
	g.channel <- "begin"
	g.waitgroup.Wait()
	fmt.Println("Finally, Max:", g.scoreOf["Max"], " Mike:", g.scoreOf["Mike"])
}

//play отвечает за основную логику игры
func (g *gameplay) play(name string) {
	var action string
	for val := range g.channel {
		if g.gotWinner() {
			close(g.channel)
			g.waitgroup.Done()
			return
		}
		switch val {
		case "ping":
			action = "pong"
		default:
			action = "ping"
		}
		fmt.Println(name, action)
		if rand.Intn(5) == 0 {
			g.scoreOf[name] += 1
			fmt.Println(name, "won serve")
			g.channel <- "stop"
		} else {
			g.channel <- action
		}
	}
	g.waitgroup.Done()
}

//gotWinner проверяет выиграна ли игра
func (g *gameplay) gotWinner() bool {
	for i := range g.scoreOf {
		if g.scoreOf[i] == 11 {
			return true
		}
	}
	return false
}
