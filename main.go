package main

import (
	"raytracer/geometry"
	"raytracer/raytracing"
)

func main() {
	mySphere := raytracing.CreateSphere(
		geometry.CreatePoint(0, 0, 0),
		geometry.CreateVector(1, 0, 0),
		raytracing.Red())

	println(mySphere.Intersect(
		raytracing.CreateRay(
			geometry.CreatePoint(0,0,0),
			geometry.CreateVector(0,0,1))))

	println(mySphere.Intersect(
		raytracing.CreateRay(
			geometry.CreatePoint(0,1,-5),
			geometry.CreateVector(0,0,1))))

	println(mySphere.Intersect(
		raytracing.CreateRay(
			geometry.CreatePoint(0,0,5),
			geometry.CreateVector(0,0,1))))

	p := geometry.CreatePoint(-3, 4, 5).Multiply(geometry.Identity(4).Translate(5,-3,2))
	println(p.X(),p.Y(),p.Z())
}
