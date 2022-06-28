package lib

import (
	"image"
	"image/png"
	"math"
	"math/rand"
	"os"
	"path/filepath"
)

//		=== MATH UTILS ===

// Converts degrees to radians
func Deg2rad(deg float64) float64 {
	return deg * math.Pi / 180.0
}

// Returns a random floar on the specified interval
func RandomFloat(min, max float64) float64 {
	return min + (max-min)*rand.Float64()
}

// Returns x if its on interval min max. Else returns either min or max.
func Clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

// Returns a random unit vector
func RandomUnitVector() Vec {
	return RandomUnitSphere().Unit()
}

// Returns a random vector in an unit sphere
func RandomUnitSphere() Vec {
	var p Vec
	p = p.Random(-1, 1)

	for p.LengthSquared() >= 1 {
		p = p.Random(-1, 1)
	}

	return p.Unit()
}

// Returns a random vector in an unit sphere in the same hemisphere as 'normal Vec'
func RandomInHemisphere(normal Vec) Vec {
	inUnitSphere := RandomUnitSphere()

	// inthe same hemisphere as the normal
	if inUnitSphere.Dot(normal) > 0.0 {
		return inUnitSphere
	} else {
		return inUnitSphere.Neg()
	}
}

// Reflects a Vector Vec relative to the Normal Vector
func Reflect(vec, normal Vec) Vec {
	return vec.Sub(normal.Scale(2 * vec.Dot(normal)))
}

// Refracts a vector UV relative to a normal vector
func Refract(uv, n Vec, etai_over_etat float64) Vec {
	cos_theta := math.Min(uv.Neg().Dot(n), 1.0)
	r_out_perp := uv.Sum(n.Scale(cos_theta)).Scale(etai_over_etat)
	r_out_parallel := n.Scale(-math.Sqrt(math.Abs(1.0 - r_out_perp.LengthSquared())))

	return r_out_perp.Sum(r_out_parallel)
}

// 		=== FILE UTILS ===

// Saves image to savePath
func SaveImage(img image.Image, name string, savePath string) {

	outputPath := filepath.Join(savePath, name+".png")
	f, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}

}
