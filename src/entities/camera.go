package entities

import "go-raytracer/lib"

type Camera struct {
	Viewport struct {
		Height, Width, Aspect_ratio float64
	}

	Focal_length float64
	Origin       lib.Point

	Horizontal,
	Vertical,
	Lower_left_corner lib.Vec
}

func (c Camera) GetRay(u, v float64) Ray {
	rayDirection := c.Lower_left_corner.Sum(c.Horizontal.Scale(u)).Sum(c.Vertical.Scale(v)).Sub(c.Origin)

	return Ray{Orig: c.Origin, Dir: rayDirection}
}
