package geometry

import "math"

type Vector struct {
	data Matrix
}

func CreateVector(x, y, z float64) Vector {
	return Vector{
		data: CreateWithValues(4, 1, x, y, z, 0),
	}
}

func (v Vector) X() float64 {
	return v.data.At(0, 0)
}

func (v Vector) Y() float64 {
	return v.data.At(0, 1)
}

func (v Vector) Z() float64 {
	return v.data.At(0, 2)
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.X()*v.X() + v.Y()*v.Y() + v.Z()*v.Z())
}

func (v Vector) Normalize() Vector {
	length := v.Length()
	return CreateVector(v.X()/length, v.Y()/length, v.Z()/length)
}

func (v Vector) Dot(v2 Vector) float64 {
	return v.X()*v2.X() + v.Y()*v2.Y() + v.Z()*v2.Z()
}

func (v Vector) Cross(v2 Vector) Vector {
	return CreateVector(v.Y()*v2.Z()-v.Z()*v2.Y(), v.Z()*v2.X()-v.X()*v2.Z(), v.X()*v2.Y()-v.Y()*v2.X())
}
