package datastructure

import "fmt"

type CartesianImpl map[Coordinates]int

func (c *CartesianImpl) SetVal(coords Coordinates, val int) (value int, err error) {
	if c == nil {
		return 0, fmt.Errorf("cartesian nil")
	}

	(*c)[coords] = val

	return val, nil
}

func (c *CartesianImpl) GetVal(x int, y int) (value int, err error) {
	if c == nil {
		return 0, fmt.Errorf("cartesian nil")
	}

	return (*c)[Coordinates{x, y}], nil
}
