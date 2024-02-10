package structs

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
