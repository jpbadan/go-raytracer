package materials

import (
	"go-raytracer/entities"
	"go-raytracer/lib"
	"math"
)

type Dielectric struct {
	IndexOfRefraction float64
}

func (d Dielectric) Scatter(incidentRay entities.Ray, hitRec *entities.HitRecord, attenuation *lib.Color, scatteredRay *entities.Ray) bool {
	*attenuation = lib.Color{R: 1, G: 1, B: 1}

	var refraction_ratio float64

	if hitRec.FrontFace {
		refraction_ratio = 1.0 / d.IndexOfRefraction
	} else {
		refraction_ratio = d.IndexOfRefraction
	}

	unit_direction := incidentRay.Dir.Unit()

	cos_theta := math.Min(unit_direction.Neg().Dot(hitRec.Normal), 1.0)
	sin_theta := math.Sqrt(1.0 - cos_theta*cos_theta)

	cannotRefract := refraction_ratio*sin_theta > 1.0

	var direction lib.Vec
	if cannotRefract || reflectance(cos_theta, refraction_ratio) > lib.RandomFloat(0, 1) {
		direction = lib.Reflect(unit_direction, hitRec.Normal)
	} else {
		direction = lib.Refract(unit_direction, hitRec.Normal, refraction_ratio)
	}

	*scatteredRay = entities.Ray{Orig: hitRec.P, Dir: direction}

	return true
}

//Schlick's approximation for reflectance.
func reflectance(cosine, refIndex float64) float64 {
	r0 := (1.0 - refIndex) / (1 + refIndex)
	r0 = r0 * r0

	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}
