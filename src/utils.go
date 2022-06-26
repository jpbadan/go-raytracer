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

//Recursive function to calculate the color of a ray
func rayColor(ray Ray, world Hittable, iterationDepth int, defaults Defaults) Color {
	hitRec := HitRecord{}

	// Limits the number of iterations of raycolor
	if iterationDepth <= 0 {
		return Color{0, 0, 0}
	}

	if world.hit(ray, defaults.ray.reflectionThreshold, INFINITY, &hitRec) {
		target := hitRec.p.sum(hitRec.normal).sum(randomUnitVector())

		return rayColor(Ray{hitRec.p, target.sub(hitRec.p)}, world, iterationDepth-1, defaults).scale(0.5)
		// // Normal map:
		// return hitRec.normal.sum(Color{1, 1, 1}).scale(0.5)
	}

	t := 0.5 * (ray.dir.unit().y + 1.0)

	// return a Gradient background
	return Color{1.0, 1.0, 1.0}.scale(1 - t).sum(Color{0.5, 0.7, 1.0}.scale(t))
}

// func sphere_hit(center Point3, radius float32, ray Ray) float32 {
// 	oc := ray.orig.sub(center)
// 	a := ray.dir.lengthSquared()
// 	half_b := oc.dot(ray.dir)
// 	c := oc.lengthSquared() - radius*radius
// 	discriminant := half_b*half_b - a*c

// 	if discriminant < 0 {
// 		return -1.0
// 	} else {
// 		return (-half_b - float32(math.Sqrt(float64(discriminant)))) / a
// 	}
// }

// === Hittable interface and type ===

type Hittable interface {
	// HitRecord
	hit(ray Ray, t_min float32, t_max float32, record *HitRecord) bool
}

type HitRecord struct {
	p         Point3
	normal    Vec3
	t         float32
	frontFace bool
}

func (h *HitRecord) setFaceNormal(ray Ray, outwardNormal Vec3) {
	h.frontFace = ray.dir.dot(outwardNormal) < 0
	if h.frontFace {
		h.normal = outwardNormal
	} else {
		h.normal = outwardNormal.neg()
	}
}

// === Definition of the Sphere type ===

type Sphere struct {
	center Point3
	radius float32
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

	// Implements normal determination
	outwardNormal := record.p.sub(s.center).scale(1 / s.radius)
	record.setFaceNormal(ray, outwardNormal)

	return true
}

// === Creation of a list of hittable objects ===

type HittableList struct {
	objects []interface{ Hittable }
}

func (h *HittableList) add(obj Hittable) {
	h.objects = append(h.objects, obj)
}

func (h *HittableList) clear() {
	h.objects = nil
}

func (h *HittableList) hit(ray Ray, t_min float32, t_max float32, record *HitRecord) bool {
	var temp_rec HitRecord
	hit_anything := false
	closest_so_far := t_max

	for _, obj := range h.objects {
		if obj.hit(ray, t_min, closest_so_far, &temp_rec) {
			// println("HitRec: ", temp_rec.t)
			hit_anything = true
			closest_so_far = temp_rec.t
			*record = temp_rec
		}
	}

	return hit_anything
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
