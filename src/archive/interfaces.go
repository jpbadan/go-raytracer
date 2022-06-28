package main

type Hittable interface {
	// HitRecord
	hit(ray Ray, t_min float32, t_max float32, record *HitRecord) bool
}
