package matrix

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

type Matrix struct {
	M    [][]float64
	X, Y int
}

func HelloMatrix() {
	fmt.Println("Hello, matrix")
}

func NewMatrix(x, y int) (*Matrix, error) {
	if x < 0 || y < 0 {
		return nil, errors.New("negative matrix sizes")
	}
	m := make([][]float64, x)
	for i := 0; i < x; i++ {
		m[i] = make([]float64, y)
	}
	ret := &Matrix{M: m, X: x, Y: y}
	return ret, nil
}

func (m *Matrix) String() string {
	var ret bytes.Buffer
	for i := 0; i < m.Y; i++ {
		ret.WriteString("[")
		for j := 0; j < m.X-1; j++ {
			ret.WriteString(ftoa(m.M[j][i]))
			ret.WriteString(", ")
		}
		if m.X > 0 {
			ret.WriteString(ftoa(m.M[m.X-1][i]))
		}
		if i == m.Y-1 {
			ret.WriteString("]")
		} else {
			ret.WriteString("]\n")
		}
	}
	return ret.String()
}

func MatrixFromSlice(arr [][]float64) (*Matrix, error) {
	x := len(arr)
	if x == 0 {
		ret, _ := NewMatrix(0, 0)
		return ret, nil
	}
	y := len(arr[0])
	for i := 0; i < x; i++ {
		if len(arr[i]) != y {
			return nil, errors.New("rows sizes mismatch")
		}
	}
	ret, _ := NewMatrix(x, y)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			ret.M[i][j] = arr[i][j]
		}
	}
	return ret, nil
}

func (this *Matrix) Add(other *Matrix) *Matrix {
	return nil
}

func (this *Matrix) Mul(other *Matrix) *Matrix {
	return nil
}

func (m *Matrix) TriUp() *Matrix {
	return nil
}

func (m *Matrix) Diag() *Matrix {
	return nil
}

func (m *Matrix) Trans() *Matrix {
	return nil
}

func Copy(m *Matrix) *Matrix {
	return nil
}

func ftoa(f float64) string {
	return strconv.FormatFloat(f, 'g', -1, 64)
}
