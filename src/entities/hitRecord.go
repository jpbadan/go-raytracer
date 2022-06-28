package entities

import (
	"go-raytracer/lib"
)

type HitRecord struct {
	P         lib.Point
	Normal    lib.Vec
	Material  Material
	T         float64
	FrontFace bool
}

func (h *HitRecord) SetFaceNormal(ray Ray, outwardNormal lib.Vec) {
	h.FrontFace = ray.Dir.Dot(outwardNormal) < 0
	if h.FrontFace {
		h.Normal = outwardNormal
	} else {
		h.Normal = outwardNormal.Neg()
	}
}
