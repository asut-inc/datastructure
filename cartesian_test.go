package datastructure

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type setValTestCase struct {
	inputData Coordinates
	expected  int
}

var cartesian = make(CartesianImpl)

var dataCartesian = []setValTestCase{
	{
		inputData: Coordinates{x: 0, y: 2},
		expected:  13,
	},
	{
		inputData: Coordinates{x: 0, y: 2},
		expected:  21,
	},
	{
		inputData: Coordinates{x: 0, y: 2},
		expected:  34,
	},
	{
		inputData: Coordinates{x: 0, y: 2},
		expected:  55,
	},
	{
		inputData: Coordinates{x: 0, y: 2},
		expected:  89,
	},
	{
		inputData: Coordinates{x: 0, y: 2},
		expected:  144,
	},
	{
		inputData: Coordinates{x: 0, y: 2},
		expected:  233,
	},
}

func TestCartesianImpl(t *testing.T) {
	for _, data := range dataCartesian {
		val := data.expected
		t.Run("SetVal", func(t *testing.T) {
			res, err := cartesian.SetVal(data.inputData, val)
			if err != nil {
				t.Error(
					"error:", err,
					"input:", data.inputData,
					"result:", res,
				)
			}

			assert.Equal(t, data.expected, res, fmt.Sprintf("Expected value (#{data.expected}) does not match with setted value (#{val})"))
		})
		t.Run("GetVal", func(t *testing.T) {
			res, err := cartesian.GetVal(data.inputData.x, data.inputData.y)
			if err != nil {
				t.Error(
					"error:", err,
					"input:", data.inputData,
					"result:", res,
				)
			}
			assert.Equal(t, data.expected, res, fmt.Sprintf("Expected value (#{data.expected}) does not match with got value (#{val})"))
		})
	}
}
