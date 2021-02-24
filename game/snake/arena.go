package snake

import (
	"time"

	"github.com/mlaxwong/wasm-snake/game/utils"
)

type arena struct {
	width        float64
	height       float64
	xGrid        int
	yGrid        int
	scaledWidth  float64
	scaledHeight float64
	done         chan bool
	ticker       *time.Ticker
	speed        float64
	started      bool
	snake        *snake
	food         *food
	scoreChan    chan (int)
	gameOverChan chan (bool)
}

func newArena(canvas *canvas, xGrid int, yGrid int, scoreChan chan (int), gameOverChan chan (bool)) *arena {
	scaledWidth := float64(canvas.width) / float64(xGrid)
	scaledHeight := float64(canvas.height) / float64(yGrid)

	width := scaledWidth * float64(xGrid)
	height := scaledHeight * float64(yGrid)

	a := &arena{
		width:        width,
		height:       height,
		xGrid:        xGrid,
		yGrid:        yGrid,
		scaledWidth:  scaledWidth,
		scaledHeight: scaledHeight,
		scoreChan:    scoreChan,
		gameOverChan: gameOverChan,
	}

	a.initArena()

	return a
}

func (self *arena) initArena() {
	self.started = false
	self.speed = 1.0
	self.placeSnake()
	self.placeFood()
}

func (self *arena) start() {
	if !self.isStarted() {
		self.done = make(chan bool)
		self.ticker = ticker2(self.loop, []interface{}{}, self.speedToDuration(self.speed), self.done)
		self.started = true
		// fmt.Println("Arena started")
	}
}

func (self *arena) pause() {
	self.ticker.Stop()
	close(self.done)
	self.started = false
}

func (self *arena) resume() {
	self.start()
}

func (self *arena) reset() {
	self.initArena()
}

func (self *arena) isStarted() bool {
	return self.started
}

func (self *arena) placeSnake() {
	self.snake = newSnake(self, self.xGrid/2, self.yGrid/2)
}

func (self *arena) placeFood() {
	var x, y int
	for {
		x = utils.RandInt(1, self.xGrid)
		y = utils.RandInt(1, self.yGrid)

		if !self.isOccupired(coordinate{x, y}) {
			break
		}
	}

	self.food = newFood(x, y)
}

func (self *arena) addScore(score int) {
	self.scoreChan <- score
}

func (self *arena) increaseSpeed() {
	self.ticker.Stop()
	self.speed += .1

	self.ticker = ticker2(self.loop, []interface{}{}, self.speedToDuration(self.speed), self.done)
}

func (self *arena) loop(args ...interface{}) {
	self.moveSnake()
	self.snakeHaveFood()
	self.snakeDie()
}

func (self *arena) moveSnake() {
	self.snake.move()
}

func (self *arena) snakeHaveFood() {
	if self.snake.isCollision(self.food.coordinate) {
		// fmt.Println("ate!")
		self.snake.eat(*self.food)
		self.addScore(self.food.point)
		self.increaseSpeed()
		self.placeFood()
	}
}

func (self *arena) snakeDie() {
	if self.snake.isBodyCollision(self.snake.head()) {
		self.snake.die()
		self.pause()
		self.gameOverChan <- true
	}
}

func (self *arena) isOccupired(c coordinate) bool {
	return self.snake.isCollision(c)
}

func (self *arena) speedToDuration(speed float64) time.Duration {
	// fmt.Println(speed)
	// fmt.Println(300 / speed)
	return time.Duration(300/speed) * time.Millisecond
}
