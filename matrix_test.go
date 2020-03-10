package datastructure

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var dataMatrix [][][]int = [][][]int{
	{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	},
	{
		[]int{1000, 2000, 3000},
		[]int{4000, 5000, 6000},
		[]int{7000, 8000, 9000},
	},
	{
		[]int{-10000, 2, 3},
		[]int{4, -10000, 6},
		[]int{7, 8, -10000},
	},
}

func TestMatrix(t *testing.T) {
	for _, data := range dataMatrix {
		matrix := NewMatrixImpl(data)
		t.Run("test poured data", func(t *testing.T) {
			for rowIdx := 0; rowIdx < len(dataMatrix)-1; rowIdx++ {
				for columnIdx := 0; columnIdx < len(dataMatrix[0])-1; columnIdx++ {
					t.Run("assert value by index", func(t *testing.T) {
						res, err := matrix.GetVal(columnIdx, rowIdx)
						if err != nil {
							t.Error(
								"error:", err,
								"input:", data,
								"result:", res,
							)
						}
						expectedVal := data[rowIdx][columnIdx]
						assertTxtFormat := "Expected value (#{expectedVal}) does not match with got value (#{res})"
						assert.Equal(t, expectedVal, res, fmt.Sprintf(assertTxtFormat))
					})
				}
			}
		})
	}
}

func TestMatrixInterface(t *testing.T) {
	var matrix Matrix

	matrix = NewMatrixImpl(dataMatrix[0])

	_ = matrix
}
