package entities

import "go-raytracer/lib"

type Ray struct {
	Orig lib.Point
	Dir  lib.Vec
}

func (r Ray) Origin() lib.Point {
	return r.Orig
}

func (r Ray) Direction() lib.Vec {
	return r.Dir
}

func (r Ray) At(t float64) lib.Point {
	return r.Orig.Sum(r.Dir.Scale(t))
}

// === UTILS ===

//Recursive function to calculate the color of a ray
func RayColor(ray Ray, world Geometry, iterationDepth int, reflectionThreshold float64) lib.Color {
	hitRec := HitRecord{}

	// Limits the number of iterations of raycolor
	if iterationDepth <= 0 {
		return lib.Color{R: 0, G: 0, B: 0}
	}

	if world.Hit(ray, reflectionThreshold, lib.INF, &hitRec) {
		//Simple difuse reflection:
		// target := hitRec.p.sum(hitRec.normal).sum(randomUnitSphere())

		//True Lambertian reflection:
		// target := hitRec.p.sum(hitRec.normal).sum(randomUnitVector())

		// Alternate reflection formulation:
		//target := hitRec.p.sum(randomInHemisphere(hitRec.normal))

		//return rayColor(Ray{hitRec.p, target.sub(hitRec.p)}, world, iterationDepth-1, defaults).scale(0.5)

		var scaterred Ray
		var attenuation lib.Color

		if hitRec.Material.Scatter(ray, &hitRec, &attenuation, &scaterred) {
			return attenuation.Mult(RayColor(scaterred, world, iterationDepth-1, reflectionThreshold))
		}
		return lib.Color{R: 0, G: 0, B: 0}
	}
	t := 0.5 * (ray.Dir.Unit().Y + 1.0)

	// return a Gradient background
	return lib.Color{R: 1.0, G: 1.0, B: 1.0}.Scale(1 - t).Sum(lib.Color{R: 0.5, G: 0.7, B: 1.0}.Scale(t))
}
