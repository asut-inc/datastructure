package datastructure

type MatrixImpl struct {
	CartesianImpl
}

func (m *MatrixImpl) PouringData(matrixData [][]int) error {
	columnsCount := len(matrixData[0])
	rowsCount := len(matrixData)
	coords := NewCoord(0, 0)

	for rowIdx := 0; rowIdx < rowsCount-1; rowIdx++ {
		for columnIdx := 0; columnIdx < columnsCount-1; columnIdx++ {
			coords.x = columnIdx
			coords.y = rowIdx
			_, err := m.SetVal(*coords, matrixData[rowIdx][columnIdx])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func NewMatrixImpl(matrixData [][]int) *MatrixImpl {
	matrix := MatrixImpl{make(CartesianImpl, len(matrixData))}
	_ = matrix.PouringData(matrixData)

	return &matrix
}
