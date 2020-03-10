package datastructure

type Coordinates struct {
	x int
	y int
}

type Cartesian interface {
	SetVal(coord Coordinates, val int) (value int, err error)
	GetVal(x int, y int) (value int, err error)
}
