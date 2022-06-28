package materials

import (
	"go-raytracer/entities"
	"go-raytracer/lib"
)

type Metal struct {
	Albedo lib.Color
	Fuzz   float64
}

func (m Metal) Scatter(incidentRay entities.Ray, hitRec *entities.HitRecord, attenuation *lib.Color, scatteredRay *entities.Ray) bool {
	reflected := lib.Reflect(incidentRay.Dir.Unit(), hitRec.Normal)
	*attenuation = m.Albedo
	*scatteredRay = entities.Ray{
		Orig: hitRec.P,
		Dir:  reflected.Sum(lib.RandomUnitSphere().Scale(m.Fuzz))}

	return scatteredRay.Dir.Dot(hitRec.Normal) > 0
}
