package datastructure

type Coordinates struct {
	x int
	y int
}

func NewCoord(x int, y int) *Coordinates {
	return &Coordinates{x, y}
}

type Cartesian interface {
	SetVal(coord Coordinates, val int) (value int, err error)
	GetVal(x int, y int) (value int, err error)
}
