package main

type Ray struct {
	orig Point3
	dir  Vec3
}

func (r Ray) origin() Point3 {
	return r.orig
}

func (r Ray) direction() Vec3 {
	return r.dir
}

func (r Ray) at(t float32) Point3 {
	return r.orig.sum(r.dir.scale(t))
}
