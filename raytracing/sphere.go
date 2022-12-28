package raytracing

import (
	"math"
	"raytracer/geometry"
)

type Sphere struct {
	center geometry.Point
	radius geometry.Vector
	color  Color
}

func CreateSphere(p geometry.Point, v geometry.Vector, c Color) Sphere {
	return Sphere{center: p, radius: v, color: c}
}
func (s Sphere) Intersect(r Ray) []Intersection {
	sphereToRay := r.origin.Subtract(s.center)
	a := r.direction.Dot(r.direction)
	b := 2 * r.direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay)-1
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return []Intersection{}
	}
	sqrRootDiscriminant := math.Sqrt(discriminant)
	t1 := (-b - sqrRootDiscriminant) / (2 * a)
	t2 := (-b + sqrRootDiscriminant) / (2 * a)
	println(a, b, c, discriminant, sqrRootDiscriminant, t1, t2)
	return []Intersection{{T: t1, Object: s}, {T: t2, Object: s}}
}

func (s Sphere) Color() Color {
	return s.color
}
