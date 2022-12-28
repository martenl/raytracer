package raytracing

type Color struct {
	r, g, b, a float64
}

func CreateColor(r, g, b, a float64) Color {
	return Color{
		r: r,
		g: g,
		b: b,
		a: a,
	}
}

func Red() Color {
	return CreateColor(1.0, 0, 0, 0)
}

func (c Color) Equals(c2 Color) bool {
	return c.r == c2.r && c.g == c2.g && c.b == c2.b && c.a == c2.a
}

func (c Color) Add(c2 Color) Color {
	return Color{
		r: c.r + c2.r,
		b: c.b + c2.b,
		g: c.g + c2.g,
		a: c.a + c2.a}
}

func (c Color) Subtract(c2 Color) Color {
	return Color{
		r: c.r - c2.r,
		b: c.b - c2.b,
		g: c.g - c2.g,
		a: c.a - c2.a}
}

func (c Color) Multiply(c2 Color) Color {
	return Color{
		r: c.r * c2.r,
		b: c.b * c2.b,
		g: c.g * c2.g,
		a: c.a * c2.a}
}

func (c Color) HadamardProduct(c2 Color) Color {
	return c.Multiply(c2)
}

func (c Color) MultiplyWithScalar(s float64) Color {
	return Color{
		r: c.r * s,
		b: c.b * s,
		g: c.g * s,
		a: c.a * s}
}
