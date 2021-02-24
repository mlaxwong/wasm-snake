package snake

type snake struct {
	arena           *arena
	alive           bool
	direction       direction
	directionLocked bool
	body            []coordinate
	poop            int
}

func newSnake(arena *arena, x int, y int) *snake {
	return &snake{
		arena:           arena,
		alive:           true,
		direction:       NONE,
		directionLocked: false,
		body: []coordinate{
			coordinate{x, y},
		},
		poop: 0,
	}
}

func (self *snake) head() coordinate {
	return self.body[len(self.body)-1]
}

func (self *snake) bodyOnly() []coordinate {
	return self.body[:len(self.body)-1]
}

func (self *snake) move() {

	if self.direction == NONE {
		return
	}

	maxX := self.arena.xGrid
	maxY := self.arena.yGrid

	head := self.head()
	next := coordinate{x: head.x, y: head.y}

	switch self.direction {
	case UP:
		if next.y > 1 {
			next.y--
		} else {
			next.y = maxY
		}
	case DOWN:
		if next.y < maxY {
			next.y++
		} else {
			next.y = 1
		}
	case LEFT:
		if next.x > 1 {
			next.x--
		} else {
			next.x = maxX
		}
	case RIGHT:
		if next.x < maxX {
			next.x++
		} else {
			next.x = 1
		}
	}

	if self.poop > 0 {
		self.body = append(self.body, next)
		self.poop--
	} else {
		self.body = append(self.body[1:], next)
	}

	self.directionLocked = false
}

func (self *snake) changeDirection(d direction) {
	if !isOpposite(self.direction, d) && !self.directionLocked {
		self.direction = d
		self.directionLocked = true
	}
}

func (self *snake) eat(food food) {
	self.poop++
}

func (self *snake) die() {
	self.alive = false
}

func (self *snake) isCollision(c coordinate) bool {
	for _, b := range self.body {
		if b.x == c.x && b.y == c.y {
			return true
		}
	}
	return false
}

func (self *snake) isBodyCollision(c coordinate) bool {
	for _, b := range self.bodyOnly() {
		if b.x == c.x && b.y == c.y {
			return true
		}
	}
	return false
}

func (self *snake) isHeadCollision(c coordinate) bool {
	if self.head().x == c.x && self.head().y == c.y {
		return true
	} else {
		return false
	}
}
