package main

import (
	"image/color"
	"math"
	"math/rand"
)

type Vec3 struct{ x, y, z float32 }
type Color = Vec3
type Point3 = Vec3

func (v Vec3) sum(v2 Vec3) Vec3 {
	return Vec3{x: v.x + v2.x, y: v.y + v2.y, z: v.z + v2.z}
}

func (v Vec3) scale(t float32) Vec3 {
	return Vec3{x: v.x * t, y: v.y * t, z: v.z * t}
}

func (v Vec3) div(t float32) Vec3 {
	return Vec3{x: v.x / t, y: v.y / t, z: v.z / t}
}

func (v Vec3) neg() Vec3 {
	return Vec3{x: -v.x, y: -v.y, z: -v.z}
}

func (v Vec3) lengthSquared() float32 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

func (v Vec3) length() float32 {
	return float32(math.Sqrt(float64(v.lengthSquared())))
}

func (v Vec3) mult(v2 Vec3) Vec3 {
	return Vec3{x: v.x * v2.x, y: v.y * v2.y, z: v.z * v2.z}
}

func (v Vec3) sub(v2 Vec3) Vec3 {
	return v.sum(v2.neg())
}

func (v Vec3) dot(v2 Vec3) float32 {
	return v.x*v2.x + v.y*v2.y + v.z*v2.z
}

func (v Vec3) cross(v2 Vec3) Vec3 {
	return Vec3{
		x: v.y*v2.z - v.z*v2.y,
		y: v.z*v2.x - v.x*v2.z,
		z: v.x*v2.y - v.y*v2.x}
}

func (v Vec3) unit() Vec3 {
	return v.div(v.length())
}

func (v Vec3) randomUnit() Vec3 {
	return Vec3{
		x: rand.Float32(),
		y: rand.Float32(),
		z: rand.Float32()}
}

func (v Vec3) random(min, max float32) Vec3 {
	return Vec3{
		x: randomFloat(min, max),
		y: randomFloat(min, max),
		z: randomFloat(min, max)}
}

func (c Color) toRGBA() color.RGBA {
	red := uint8(255.999 * c.x)
	green := uint8(255.999 * c.y)
	blue := uint8(255.999 * c.z)

	return color.RGBA{R: red, G: green, B: blue, A: 255}
}

//Divide color by the number of samples and gamma correct
func (c Color) correct(defaults Defaults) Color {
	r := math.Pow(float64(c.x*(1.0/float32(defaults.image.samples_per_pixel))), float64(defaults.image.gamma_correction))
	g := math.Pow(float64(c.y*(1.0/float32(defaults.image.samples_per_pixel))), float64(defaults.image.gamma_correction))
	b := math.Pow(float64(c.z*(1.0/float32(defaults.image.samples_per_pixel))), float64(defaults.image.gamma_correction))

	return Color{float32(r), float32(g), float32(b)}

}
