package entities

import (
	"go-raytracer/lib"
)

type Material interface {
	Scatter(incidentRay Ray, hitRec *HitRecord, attenuation *lib.Color, scatteredRay *Ray) bool
}
