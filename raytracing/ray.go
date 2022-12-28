package raytracing

import "raytracer/geometry"

type Ray struct {
	origin    geometry.Point
	direction geometry.Vector
}

func CreateRay(o geometry.Point, v geometry.Vector) Ray {
	return Ray{origin: o, direction: v}
}
