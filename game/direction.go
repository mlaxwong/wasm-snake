package game

type direction int

const (
	NONE direction = 1 + iota
	UP
	DOWN
	LEFT
	RIGHT
)

func isOpposite(from direction, to direction) bool {
	opposites := map[direction]direction{
		UP:    DOWN,
		DOWN:  UP,
		LEFT:  RIGHT,
		RIGHT: LEFT,
	}
	return opposites[from] == to
}
