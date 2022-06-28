package lib

import (
	"math"
	"math/rand"
)

type Vec struct{ X, Y, Z float64 }

func (a Vec) Sum(b Vec) Vec {
	return Vec{X: a.X + b.X, Y: a.Y + b.Y, Z: a.Z + b.Z}
}

func (a Vec) Scale(t float64) Vec {
	return Vec{X: a.X * t, Y: a.Y * t, Z: a.Z * t}
}

func (a Vec) Div(t float64) Vec {
	return Vec{X: a.X / t, Y: a.Y / t, Z: a.Z / t}
}

func (a Vec) Neg() Vec {
	return Vec{X: -a.X, Y: -a.Y, Z: -a.Z}
}

func (a Vec) LengthSquared() float64 {
	return a.X*a.X + a.Y*a.Y + a.Z*a.Z
}

func (a Vec) Length() float64 {
	return float64(math.Sqrt(float64(a.LengthSquared())))
}

func (a Vec) Mult(b Vec) Vec {
	return Vec{X: a.X * b.X, Y: a.Y * b.Y, Z: a.Z * b.Z}
}

//subtracts vector b
func (a Vec) Sub(b Vec) Vec {
	return a.Sum(b.Neg())
}

func (a Vec) Dot(b Vec) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (a Vec) Cross(b Vec) Vec {
	return Vec{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X}
}

func (a Vec) Unit() Vec {
	return a.Div(a.Length())
}

func (a Vec) RandomUnit() Vec {
	return Vec{
		X: rand.Float64(),
		Y: rand.Float64(),
		Z: rand.Float64()}
}

func (a Vec) Random(min, maX float64) Vec {
	return Vec{
		X: RandomFloat(min, maX),
		Y: RandomFloat(min, maX),
		Z: RandomFloat(min, maX)}
}

func (a Vec) NearZero() bool {
	// Returns true if the vector is close to zero in all dimensions
	s := 1e-8
	return math.Abs(a.X) < s && math.Abs(a.Y) < s && math.Abs(a.Z) < s
}
