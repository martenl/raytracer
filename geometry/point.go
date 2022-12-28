package geometry

type Point struct {
	data Matrix
}

func CreatePoint(x, y, z float64) (p Point) {
	data := CreateWithValues(4, 1, x, y, z, 1)
	return Point{data: data}
}

func (p Point) X() float64 {
	return p.data.At(0, 0)
}

func (p Point) Y() float64 {
	return p.data.At(0, 1)
}

func (p Point) Z() float64 {
	return p.data.At(0, 2)
}

func (p Point) Multiply(m Matrix) Point {
	return Point{data: p.data.Multiply(m)}
}

func (p Point) Add(v Vector) Point {
	return Point{p.data.Add(v.data)}
}

func (p Point) Subtract(p2 Point) Vector {
	return CreateVector(p.X()-p2.X(), p.Y()-p2.Y(), p.Z()-p2.Z())
}
