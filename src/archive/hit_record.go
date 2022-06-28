package main

// === Hittable interface and type ===

type HitRecord struct {
	p         Point3
	normal    Vec3
	material  Material
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
