package snake

type food struct {
	coordinate coordinate
	point      int
}

func newFood(x int, y int) *food {
	return &food{
		coordinate: coordinate{x, y},
		point:      10,
	}
}
