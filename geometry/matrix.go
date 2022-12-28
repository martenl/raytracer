package geometry

import (
	"math"
	"reflect"
)

type Matrix struct {
	height, width int
	data          [][]float64
}

func Create(height, width int) Matrix {
	var data = make([][]float64, height)
	for i := 0; i < height; i++ {
		data[i] = make([]float64, width)
	}
	return Matrix{height: height, width: width, data: data}
}

func CreateWithValues(height, width int, values ...float64) (m Matrix) {
	m = Create(height, width)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			index := i*width + j
			m.Set(j, i, values[index])
			//println(i, j, index, values[index], m.At(j, i))
		}
	}
	return
}

func Identity(n int) (m Matrix) {
	m = Create(n, n)
	for i := 0; i < n; i++ {
		m.Set(i, i, 1.0)
	}
	return
}

func (m Matrix) Equals(m2 Matrix) bool {
	if m.width != m2.width || m.height != m2.height {
		return false
	}
	return reflect.DeepEqual(m.data, m2.data)
}

func (m Matrix) Height() (height int) {
	return m.height
}

func (m Matrix) Width() (width int) {
	return m.width
}

func (m Matrix) Dims() (height int, width int) {
	return m.height, m.width
}

func (m Matrix) At(x, y int) float64 {
	if x >= 0 && x < m.width && y >= 0 && y < m.height {
		return m.data[y][x]
	}
	return 0
}

func (m Matrix) Set(x, y int, value float64) Matrix {
	if x >= 0 && x < m.width && y >= 0 && y < m.height {
		m.data[y][x] = value
	}
	return m
}

func (m Matrix) Add(m2 Matrix) Matrix {
	for i := 0; i < m.width; i++ {
		for j := 0; j < m.height; j++ {
			m.Set(i, j, m.At(i, j)+m2.At(i, j))
		}
	}
	return m
}

func (m Matrix) Subtract(m2 Matrix) Matrix {
	for i := 0; i < m.width; i++ {
		for j := 0; j < m.height; j++ {
			m.Set(i, j, m.At(i, j)-m2.At(i, j))
		}
	}
	return m
}

func (m Matrix) ScalarMultiply(scalar float64) Matrix {
	for i := 0; i < m.width; i++ {
		for j := 0; j < m.height; j++ {
			m.Set(i, j, scalar*m.At(i, j))
		}
	}
	return m
}

func (m Matrix) Multiply(m2 Matrix) (m3 Matrix) {
	if m.width != m2.height {
		panic("Sizes don't match")
	}
	m3 = Create(m.height, m2.width)

	for i := 0; i < m.height; i++ {
		for j := 0; j < m2.width; j++ {
			var value = 0.0
			for k := 0; k < m.width; k++ {
				value += m.At(k, i) * m2.At(j, k)
			}
			m3.Set(j, i, value)
		}
	}
	return
}

func (m Matrix) Transpose() (m2 Matrix) {
	m2 = Create(m.width, m.height)
	for i := 0; i < m.width; i++ {
		for j := 0; j < m.height; j++ {
			m2.Set(j, i, m.At(i, j))
		}
	}
	return
}

func (m Matrix) Submatrix(row, column int) (submatrix Matrix) {
	submatrix = Create(m.height-1, m.width-1)
	for i, i2 := 0, 0; i < m.width; i++ {
		if i == column {
			continue
		}
		for j, j2 := 0, 0; j < m.height; j++ {
			if j == row {
				continue
			}
			submatrix.Set(i2, j2, m.At(i, j))
			j2++
		}
		i2++
	}
	return
}

func (m Matrix) Determinant() (determinant float64) {
	if m.width != m.height {
		return 0
	}
	if m.height == 1 {
		return m.data[0][0]
	}
	if m.height == 2 {
		return m.data[0][0]*m.data[1][1] - m.data[1][0]*m.data[0][1]
	}
	for i := 0; i < m.height; i++ {
		determinant += m.At(0, i) * m.Cofactor(i, 0)
	}
	return
}

func (m Matrix) submatrix(row, column int) (submatrix Matrix) {
	submatrix = Create(m.height-1, m.width-1)
	for i, i2 := 0, 0; i < m.width; i++ {
		if i == column {
			continue
		}
		for j, j2 := 0, 0; j < m.height; j++ {
			if j == row {
				continue
			}
			submatrix.Set(i2, j2, m.At(i, j))
			j2++
		}
		i2++
	}
	return
}

func (m Matrix) Minor(row, column int) float64 {
	return m.submatrix(row, column).Determinant()
}

func (m Matrix) Cofactor(row, column int) (cofactor float64) {
	cofactor = m.Minor(row, column)
	if (column+row)%2 == 1 {
		cofactor *= -1
	}
	return
}

func (m Matrix) Inverse() Matrix {
	determinant := m.Determinant()
	if determinant == 0 {
		panic("Matrix is not invertible")
	}
	println("Determinant is ", determinant)
	inverseMatrix := Create(m.height, m.width)
	for i := 0; i < m.height; i++ {
		for j := 0; j < m.width; j++ {
			value := m.Cofactor(j, i) / determinant
			inverseMatrix.Set(j, i, value)
		}
	}
	return inverseMatrix
}

func (m Matrix) Translate(x, y, z float64) Matrix {
	translationMatrix := Identity(4)
	translationMatrix.Set(3, 0, x).Set(3, 1, y).Set(3, 2, z)
	translationMatrix.Print()
	return translationMatrix.Multiply(m)
}

func (m Matrix) Scale(x, y, z float64) Matrix {
	scalingMatrix := Identity(4)
	scalingMatrix.Set(0, 0, x).Set(1, 1, y).Set(2, 2, z)
	return scalingMatrix.Multiply(m)
}

func (m Matrix) rotate(rad float64, axis rune) Matrix {
	createRotationMatrix := func() (rm Matrix) {
		rm = Identity(4)
		switch axis {
		case 'x':
			rm.Set(1, 1, math.Cos(rad))
			rm.Set(2, 1, -math.Sin(rad))
			rm.Set(1, 2, math.Sin(rad))
			rm.Set(2, 2, math.Cos(rad))
			rm.Print()
			return
		case 'y':
			rm.Set(0, 0, math.Cos(rad))
			rm.Set(2, 0, math.Sin(rad))
			rm.Set(0, 2, -math.Sin(rad))
			rm.Set(2, 2, math.Cos(rad))
			return
		case 'z':
			rm.Set(0, 0, math.Cos(rad))
			rm.Set(1, 0, -math.Sin(rad))
			rm.Set(0, 1, math.Sin(rad))
			rm.Set(1, 1, math.Cos(rad))
			return
		default:
			return
		}
	}
	return createRotationMatrix().Multiply(m)
}

func (m Matrix) RotateXaxis(rad float64) Matrix {
	return m.rotate(rad, 'x')
}

func (m Matrix) RotateYaxis(rad float64) Matrix {
	return m.rotate(rad, 'y')
}

func (m Matrix) RotateZaxis(rad float64) Matrix {
	return m.rotate(rad, 'z')
}

func (m Matrix) Shear(xy, xz, yx, yz, zx, zy float64) Matrix {
	shearMatrix := Identity(4)
	shearMatrix.Set(1, 0, xy).Set(2, 0, xz)
	shearMatrix.Set(0, 1, yx).Set(2, 1, yz)
	shearMatrix.Set(0, 2, zx).Set(1, 2, zy)
	return shearMatrix.Multiply(m)
}

func (m Matrix) Print() {
	for i := 0; i < m.height; i++ {
		for j := 0; j < m.width; j++ {
			print(m.At(j, i))
		}
		println()
	}
}
