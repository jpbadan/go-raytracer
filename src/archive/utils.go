package main

import (
	"image"
	"image/png"
	"math"
	"math/rand"
	"os"
	"path/filepath"
)

// --- Utilitary functions ---

func saveImage(img image.Image, name string) {
	conf := getDefaults()
	// Saves image to default output path
	outputPath := filepath.Join(conf.output_path, name+".png")
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

// === Utilities ===

const INFINITY = math.MaxFloat32

func deg2rad(deg float32) float32 {
	return deg * math.Pi / 180.0
}

func randomFloat(min, max float32) float32 {
	return min + (max-min)*rand.Float32()
}

func clamp(x, min, max float32) float32 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

//Picks a random point in a unit radius sphere
func randomUnitSphere() Vec3 {
	var p Vec3
	p = p.random(-1, 1)

	for p.lengthSquared() >= 1 {
		p = p.random(-1, 1)
	}

	return p.unit()
}

func randomUnitVector() Vec3 {

	return randomUnitSphere().unit()
}

func randomInHemisphere(normal Vec3) Vec3 {
	inUnitSphere := randomUnitSphere()

	// inthe same hemisphere as the normal
	if inUnitSphere.dot(normal) > 0.0 {
		return inUnitSphere
	} else {
		return inUnitSphere.neg()
	}
}
