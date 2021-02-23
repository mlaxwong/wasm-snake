package game

import (
	"syscall/js"
)

type key int

const (
	KEY_UP    key = 38
	KEY_DOWN  key = 40
	KEY_LEFT  key = 37
	KEY_RIGHT key = 39
	KEY_SPACE key = 32
)

type keyboardEvent struct {
	key int
}

func listenToKeyboard(event chan keyboardEvent) {
	window := js.Global()
	callback := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]
		e.Call("preventDefault")
		// fmt.Println(e.Get("keyCode").Int())
		event <- keyboardEvent{
			key: e.Get("keyCode").Int(),
		}
		return nil
	})
	window.Call("addEventListener", "keyup", callback)
}
