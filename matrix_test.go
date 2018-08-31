package math

import (
	"testing"
)

var algebra ByteAlgebra = new(ByteAlgebraImpl)

func TestDetZeroMatrix(t *testing.T) {
	zeroMatrix := [][]byte{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	det3ZeroMatrix := Det(zeroMatrix, algebra)
	if det3ZeroMatrix != 0 {
		t.Errorf("Det of zero matrix must be 0 but it's %d", det3ZeroMatrix)
	}
}

func TestDetOneMatrix(t *testing.T) {
	oneMatrix := [][]byte{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}

	det3OneMatrix := Det(oneMatrix, algebra)
	if det3OneMatrix != 1 {
		t.Errorf("Det of one matrix must be 1 but it's %d", det3OneMatrix)
	}
}

func TestDetMatrix1(t *testing.T) {
	m := [][]byte{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	det := Det(m, algebra)
	if det != 0 {
		t.Errorf("Det of matrix must be 0 but it's %d", det)
	}
}

func TestDetMatrix2(t *testing.T) {
	m := [][]byte{
		{5, 5, 2},
		{5, 3, 2},
		{5, 5, 1},
	}
	det := Det(m, algebra)
	if det != 10 {
		t.Errorf("Det of matrix must be 10 but it's %d", det)
	}
}

func TestDetMatrix3(t *testing.T) {
	m := [][]byte{
		{1, 5, 2, 4, 3},
		{2, 5, 7, 8, 5},
		{2, 6, 4, 8, 7},
		{3, 5, 4, 1, 2},
		{3, 5, 2, 4, 5},
	}
	det := Det(m, algebra)
	if det != 198 {
		t.Errorf("Det3 of matrix must be 198 but it's %d", det)
	}
}

func TestCross(t *testing.T) {
	matrix := [][]byte{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	cross := cross(matrix, 1, 1)
	if cross[0][0] != 1 || cross[0][1] != 3 || cross[1][0] != 7 || cross[1][1] != 9 {
		t.Error("Cross function returns wrong result")
	}
}