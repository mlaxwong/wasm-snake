package snake

import "syscall/js"

type canvas struct {
	ctx         js.Value
	htmlElement js.Value
	width       int
	height      int
}

func NewCanvas(elementId string) *canvas {
	window := js.Global()
	document := window.Get("document")
	htmlElement := document.Call("getElementById", elementId)
	ctx := htmlElement.Call("getContext", "2d")

	width := htmlElement.Get("width").Int()
	height := htmlElement.Get("height").Int()

	return &canvas{
		ctx,
		htmlElement,
		width,
		height,
	}
}

// func (self *canvas) Draw(x int, y int, color string) {
// 	scaledX := float64(x) * self.ScaledWidth
// 	scaledY := float64(y) * self.scaledHeight

// 	self.ctx.Set("fillStyle", color)
// 	self.Ctx.Call("fillRect", scaledX, scaledY, self.ScaledWidth, self.scaledHeight)
// }

func (self *canvas) clear() {
	self.ctx.Call("clearRect", 0, 0, self.width, self.height)
}
