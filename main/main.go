package main

import (
	"github.com/mlaxwong/wasm-snake/game"
)

var signal = make(chan int)

func keepAlive() {
	for {
		<-signal
	}
}

func main() {
	game.NewGame("game", 30, 30).Start()
	keepAlive()
}
