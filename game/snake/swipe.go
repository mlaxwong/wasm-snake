package snake

import (
	"syscall/js"
)

type swipe int

const (
	SWIPE_UP    swipe = 1
	SWIPE_DOWN  swipe = 2
	SWIPE_LEFT  swipe = 3
	SWIPE_RIGHT swipe = 4
)

type swipeEvent struct {
	swipe int
}

func listenToSwipe(event chan swipeEvent) {
	window := js.Global()
	callback := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]
		e.Call("preventDefault")
		event <- swipeEvent{
			swipe: e.Get("detail").Int(),
		}
		return nil
	})
	window.Call("addEventListener", "swipe", callback)
}
