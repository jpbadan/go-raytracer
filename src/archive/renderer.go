package main

//Recursive function to calculate the color of a ray
func rayColor(ray Ray, world Hittable, iterationDepth int, defaults Defaults) Color {
	hitRec := HitRecord{}

	// Limits the number of iterations of raycolor
	if iterationDepth <= 0 {
		return Color{0, 0, 0}
	}

	if world.hit(ray, defaults.ray.reflectionThreshold, INFINITY, &hitRec) {
		//Simple difuse reflection:
		// target := hitRec.p.sum(hitRec.normal).sum(randomUnitSphere())

		//True Lambertian reflection:
		// target := hitRec.p.sum(hitRec.normal).sum(randomUnitVector())

		// Alternate reflection formulation:
		//target := hitRec.p.sum(randomInHemisphere(hitRec.normal))

		//return rayColor(Ray{hitRec.p, target.sub(hitRec.p)}, world, iterationDepth-1, defaults).scale(0.5)

		var scaterred Ray
		var attenuation Color

		if hitRec.material.scatter(ray, &hitRec, &attenuation, &scaterred) {
			return attenuation.mult(rayColor(scaterred, world, iterationDepth-1, defaults))
		}
		return Color{0, 0, 0}
	}
	t := 0.5 * (ray.dir.unit().y + 1.0)

	// return a Gradient background
	return Color{1.0, 1.0, 1.0}.scale(1 - t).sum(Color{0.5, 0.7, 1.0}.scale(t))
}
