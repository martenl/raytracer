package scene

type ThreeDVector struct {
	X, Y, Z int
}

type Camera struct {
	Position  ThreeDVector
	Direction ThreeDVector
}

type Light struct {
	Position ThreeDVector
}

type World struct {
	Camera Camera
}
