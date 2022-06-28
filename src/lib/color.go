package lib

import (
	"image/color"
	"math"
)

type Color struct{ R, G, B float64 }

func (a Color) Sum(b Color) Color {
	return Color{R: a.R + b.R, G: a.G + b.G, B: a.B + b.B}
}

func (a Color) Scale(t float64) Color {
	return Color{R: a.R * t, G: a.G * t, B: a.B * t}

}

func (a Color) Neg() Color {
	return Color{R: -a.R, G: -a.G, B: -a.B}
}

func (a Color) Mult(b Color) Color {
	return Color{R: a.R * b.R, G: a.G * b.G, B: a.B * b.B}
}

//subtracts color b
func (a Color) Sub(b Color) Color {
	return a.Sum(b.Neg())
}

func (c Color) ToRGBA() color.RGBA {
	R := uint8(255.999 * c.R)
	G := uint8(255.999 * c.G)
	B := uint8(255.999 * c.B)

	return color.RGBA{R: R, G: G, B: B, A: 255}
}

//Divide color by the number of samples and gamma correct
func (c Color) Correct(samplesPerPixel int, gammaCorrection float64) Color {
	R := math.Pow(c.R*(1.0/float64(samplesPerPixel)), gammaCorrection)
	G := math.Pow(c.G*(1.0/float64(samplesPerPixel)), gammaCorrection)
	B := math.Pow(c.B*(1.0/float64(samplesPerPixel)), gammaCorrection)

	return Color{R, G, B}

}
