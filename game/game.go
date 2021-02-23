package game

import (
	"time"
)

type gamestatus int

const (
	GAME_IDE   gamestatus = 1
	GAME_START gamestatus = 2
	GAME_PAUSE gamestatus = 3
	GAME_OVER  gamestatus = 4
)

var (
	keyboardEventsChan = make(chan keyboardEvent)
	swipeEventsChan    = make(chan swipeEvent)
	scoreChan          = make(chan int)
	gameOverChan       = make(chan bool)
)

type Game struct {
	status    gamestatus
	canvas    *canvas
	arena     *arena
	fps       time.Duration
	score     int
	bestScore int
}

func NewGame(elementId string, xGrid int, yGrid int) *Game {

	go listenToKeyboard(keyboardEventsChan)
	go listenToSwipe(swipeEventsChan)

	canvas := NewCanvas("game")

	arena := newArena(
		canvas,
		xGrid,
		yGrid,
		scoreChan,
		gameOverChan,
	)

	fps := time.Duration(50) * time.Millisecond

	return &Game{
		status:    GAME_IDE,
		canvas:    canvas,
		arena:     arena,
		fps:       fps,
		score:     0,
		bestScore: 0,
	}
}

func (self *Game) Start() {
	// go listenToKeyboard(keyboardEventsChan)

	done := make(chan bool)
	self.loop()
	t := ticker2(self.loop, []interface{}{}, self.fps, done)

	// ticker2(self.test, self.fps, done, []interface{}{1})

	for {
		select {
		case <-t.C:
			self.loop()
		case gameOver := <-gameOverChan:
			if gameOver {
				self.gameOveer()
			}
		case s := <-scoreChan:
			self.addScore(s)
		case e := <-keyboardEventsChan:
			switch e.key {
			case int(KEY_UP):
				self.changeSnakeDirection(UP)
			case int(KEY_DOWN):
				self.changeSnakeDirection(DOWN)
			case int(KEY_LEFT):
				self.changeSnakeDirection(LEFT)
			case int(KEY_RIGHT):
				self.changeSnakeDirection(RIGHT)
			case int(KEY_SPACE):
				self.togglePauseResume()
			}
		case e := <-swipeEventsChan:
			switch e.swipe {
			case int(SWIPE_UP):
				self.changeSnakeDirection(UP)
			case int(SWIPE_DOWN):
				self.changeSnakeDirection(DOWN)
			case int(SWIPE_LEFT):
				self.changeSnakeDirection(LEFT)
			case int(SWIPE_RIGHT):
				self.changeSnakeDirection(RIGHT)
			}
		}
	}
}

func (self *Game) reset() {
	self.status = GAME_IDE
	self.score = 0
	self.arena.reset()
}

func (self *Game) togglePauseResume() {
	switch self.status {
	case GAME_START:
		self.pause()
	case GAME_PAUSE:
		self.resume()
	}
}

func (self *Game) pause() {
	self.arena.pause()
	self.status = GAME_PAUSE
}

func (self *Game) resume() {
	self.arena.resume()
	self.status = GAME_START
}

func (self *Game) loop(args ...interface{}) {
	self.render()
	// self.arena.moveSnake()
}

func (self *Game) addScore(score int) {
	self.score += score
	if self.score > self.bestScore {
		self.bestScore = self.score
	}
}

func (self *Game) gameOveer() {
	self.status = GAME_OVER
}

func (self *Game) changeSnakeDirection(direction direction) {
	if self.status == GAME_OVER {
		self.reset()
		return
	}

	if !self.arena.isStarted() && direction != NONE { // move to start arena
		self.status = GAME_START
		self.arena.start()
	}
	self.arena.snake.changeDirection(direction)
}
