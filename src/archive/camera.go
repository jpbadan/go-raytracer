package main

type Camera struct {
	viewport struct {
		height, width, aspect_ratio float32
	}

	focal_length float32
	origin       Point3

	horizontal,
	vertical,
	lower_left_corner Vec3
}

// Contructor
func newCamera(defaults Defaults) *Camera {
	c := new(Camera)

	//viewPort
	c.viewport.height = defaults.camera.viewport_height
	c.viewport.width = defaults.camera.aspect_ratio * defaults.camera.viewport_height
	c.viewport.aspect_ratio = defaults.camera.aspect_ratio

	c.focal_length = defaults.camera.focal_length

	c.origin = Point3{0, 0, 0}
	c.horizontal = Vec3{c.viewport.width, 0, 0}
	c.vertical = Vec3{0, c.viewport.height, 0}
	c.lower_left_corner = c.origin.sub(c.horizontal.div(2.0)).sub(c.vertical.div(2.0)).sub(Vec3{0, 0, c.focal_length})

	return c
}

func (c Camera) getRay(u, v float32) Ray {
	rayDirection := c.lower_left_corner.sum(c.horizontal.scale(u)).sum(c.vertical.scale(v)).sub(c.origin)

	return Ray{orig: c.origin, dir: rayDirection}
}
