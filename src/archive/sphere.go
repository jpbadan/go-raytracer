package main

import "math"

// === Definition of the Sphere type ===

type Sphere struct {
	center   Point3
	radius   float32
	material Material
}

func (s Sphere) hit(ray Ray, t_min float32, t_max float32, record *HitRecord) bool {
	oc := ray.orig.sub(s.center)
	a := ray.dir.lengthSquared()
	half_b := oc.dot(ray.dir)
	c := oc.lengthSquared() - s.radius*s.radius

	discriminant := half_b*half_b - a*c
	if discriminant < 0 {
		return false
	}

	sqrtd := float32(math.Sqrt(float64(discriminant)))

	// Find the nearest root in the acceptable range:
	root := (-half_b - sqrtd) / a
	if root < t_min || t_max < root {
		root = (-half_b + sqrtd) / a
		if root < t_min || t_max < root {
			return false
		}
	}

	record.t = root
	record.p = ray.at(record.t)
	record.normal = record.p.sub(s.center).scale(1 / s.radius)
	record.material = s.material

	// Implements normal determination
	outwardNormal := record.p.sub(s.center).scale(1 / s.radius)
	record.setFaceNormal(ray, outwardNormal)

	return true
}
