package datastructure

type CartesianImpl struct {
	CartesianBasic
}

type CartesianBasic map[Coordinates]int

func (c *CartesianImpl) SetVal(coords Coordinates, val int) (value int, err error) {
	_ = coords
	_ = val

	return 1, nil
}

func (c *CartesianImpl) GetVal(x int, y int) (value int, err error) {
	_, _ = x, y

	return 1, nil
}
