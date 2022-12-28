package raytracing

type RenderObject interface {
	Intersect(r Ray) []Intersection
	Color() Color
}
