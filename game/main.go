package main

import "github.com/mlaxwong/wasm-snake/game/snake"

var signal = make(chan int)

func keepAlive() {
	for {
		<-signal
	}
}

func main() {
	snake.NewGame("game", 30, 30).Start()
	keepAlive()
}
