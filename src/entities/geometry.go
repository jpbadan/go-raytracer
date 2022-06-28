package entities

type Geometry interface {
	// HitRecord
	Hit(ray Ray, t_min, t_max float64, record *HitRecord) bool
}

type GeometryList struct {
	objects []interface{ Geometry }
}

func (h *GeometryList) Add(obj Geometry) {
	h.objects = append(h.objects, obj)
}

func (h *GeometryList) Clear() {
	h.objects = nil
}

func (h *GeometryList) Hit(ray Ray, t_min, t_max float64, record *HitRecord) bool {
	var temp_rec HitRecord
	hit_anything := false
	closest_so_far := t_max

	for _, obj := range h.objects {
		if obj.Hit(ray, t_min, closest_so_far, &temp_rec) {
			// println("HitRec: ", temp_rec.t)
			hit_anything = true
			closest_so_far = temp_rec.T
			*record = temp_rec
		}
	}

	return hit_anything
}
