package materials

import (
	"go-raytracer/entities"
	"go-raytracer/lib"
)

type Diffuse struct {
	Albedo lib.Color
}

func (l Diffuse) Scatter(incidentRay entities.Ray, hitRec *entities.HitRecord, attenuation *lib.Color, scatteredRay *entities.Ray) bool {
	scatterDirection := hitRec.Normal.Sum(lib.RandomUnitVector())

	//Catch degenerate scatter direction
	if scatterDirection.NearZero() {
		scatterDirection = hitRec.Normal
	}

	*attenuation = l.Albedo
	*scatteredRay = entities.Ray{Orig: hitRec.P, Dir: scatterDirection}

	return true
}
