package main

type Material interface {
	scatter(incidentRay Ray, hitRec *HitRecord, attenuation *Color, scatteredRay *Ray) bool
}

// === TEST MATERIAL ===

type Lambertian struct {
	albedo Color
}

func (l Lambertian) scatter(incidentRay Ray, hitRec *HitRecord, attenuation *Color, scatteredRay *Ray) bool {
	scatterDirection := hitRec.normal.sum(randomUnitVector())

	//Catch degenerate scatter direction
	if scatterDirection.nearZero() {
		scatterDirection = hitRec.normal
	}

	*scatteredRay = Ray{hitRec.p, scatterDirection}
	*attenuation = l.albedo
	return true
}

type Metal struct {
	albedo Color
	fuzz   float32
}

func (m Metal) scatter(incidentRay Ray, hitRec *HitRecord, attenuation *Color, scatteredRay *Ray) bool {
	reflected := vecReflect(incidentRay.dir.unit(), hitRec.normal)
	*scatteredRay = Ray{hitRec.p, reflected.sum(randomUnitSphere().scale(m.fuzz))}
	*attenuation = m.albedo

	return scatteredRay.dir.dot(hitRec.normal) > 0
}
