package geometries

import (
	"go-raytracer/entities"
	"go-raytracer/lib"
	"math"
)

type Sphere struct {
	Center   lib.Point
	Radius   float64
	Material entities.Material
}

func (s Sphere) Hit(ray entities.Ray, t_min, t_max float64, record *entities.HitRecord) bool {
	oc := ray.Orig.Sub(s.Center)
	a := ray.Dir.LengthSquared()
	half_b := oc.Dot(ray.Dir)
	c := oc.LengthSquared() - s.Radius*s.Radius

	discriminant := half_b*half_b - a*c
	if discriminant < 0 {
		return false
	}

	sqrtd := math.Sqrt(discriminant)

	// Find the nearest root in the acceptable range:
	root := (-half_b - sqrtd) / a
	if root < t_min || t_max < root {
		root = (-half_b + sqrtd) / a
		if root < t_min || t_max < root {
			return false
		}
	}

	record.T = root
	record.P = ray.At(record.T)
	record.Normal = record.P.Sub(s.Center).Scale(1 / s.Radius)
	record.Material = s.Material

	// Implements normal determination
	outwardNormal := record.P.Sub(s.Center).Scale(1 / s.Radius)
	record.SetFaceNormal(ray, outwardNormal)

	return true
}
