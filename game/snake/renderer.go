package snake

import "fmt"

func (game *Game) render() {
	game.canvas.clear()
	renderHeader(game)
	renderArena(game)
	renderFood(game)
	renderSnake(game)
}

func renderHeader(game *Game) {
	padding := 5
	fontSize := (float64(game.arena.height) * .2) - float64(padding)

	scoreText := fmt.Sprintf("Score: %d", game.score)
	bestScoreText := fmt.Sprintf("Best: %d", game.bestScore)
	scoreFont := fmt.Sprintf("%dpx serif", fontSize)

	game.canvas.ctx.Set("font", scoreFont)
	game.canvas.ctx.Set("fillStyle", "#000")
	game.canvas.ctx.Call("fillText", scoreText, getArenaOffsetX(game), getArenaOffsetY(game)-float64(padding))
	game.canvas.ctx.Call("fillText", bestScoreText, getArenaOffsetX(game)+100, getArenaOffsetY(game)-float64(padding))
}

func renderArena(game *Game) {
	game.canvas.ctx.Call("beginPath")
	game.canvas.ctx.Set("strokeStyle", "#666")
	game.canvas.ctx.Call("rect", getArenaOffsetX(game), getArenaOffsetY(game), getArenaResizedWidth(game), getArenaResizedHeight(game))
	game.canvas.ctx.Call("stroke")

	game.canvas.ctx.Call("clearRect", getArenaOffsetX(game), getArenaOffsetY(game), getArenaResizedWidth(game), getArenaResizedHeight(game))
}

func renderArenaPixel(game *Game, x int, y int, color string) {
	scaledX := (float64(x-1) * getArenaScaledWidth(game)) + getArenaOffsetX(game)
	scaledY := (float64(y-1) * getArenaScaledHeight(game)) + getArenaOffsetY(game)

	game.canvas.ctx.Set("fillStyle", color)
	game.canvas.ctx.Call("fillRect", scaledX, scaledY, getArenaScaledWidth(game), getArenaScaledHeight(game))
}

func renderSnake(game *Game) {
	for _, c := range game.arena.snake.body {
		renderArenaPixel(game, c.x, c.y, "red")
	}
}

func renderFood(game *Game) {
	renderArenaPixel(game, game.arena.food.coordinate.x, game.arena.food.coordinate.y, "orange")
}

func getArenaResizedWidth(game *Game) float64 {
	percent := .95 // 90% of screen
	return game.arena.width * percent
}

func getArenaResizedHeight(game *Game) float64 {
	percent := .95 // 90% of screen
	return game.arena.height * percent
}

func getArenaScaledWidth(game *Game) float64 {
	return getArenaResizedWidth(game) / float64(game.arena.xGrid)
}

func getArenaScaledHeight(game *Game) float64 {
	return getArenaResizedHeight(game) / float64(game.arena.yGrid)
}

func getArenaOffsetX(game *Game) float64 {
	offsetPercent := .5
	return (float64(game.canvas.width) * offsetPercent) - (getArenaResizedWidth(game) * offsetPercent)
}

func getArenaOffsetY(game *Game) float64 {
	offsetPercent := .8
	return (float64(game.canvas.height) * offsetPercent) - (getArenaResizedHeight(game) * offsetPercent)
}
