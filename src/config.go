package main

type Defaults struct {
	output_path string

	image struct {
		aspect_ratio      float64
		width             int
		samples_per_pixel int
		gamma_correction  float64
	}

	camera struct {
		focal_length    float64
		viewport_height float64
		aspect_ratio    float64
	}

	ray struct {
		iterationDepth      int
		reflectionThreshold float64
	}
}

// Constructor
func getDefaults() Defaults {
	d := new(Defaults)

	//File settings
	d.output_path = "../out/"

	//Image
	d.image.aspect_ratio = 16.0 / 10.0
	d.image.width = 250
	d.image.samples_per_pixel = 25 //100
	d.image.gamma_correction = 0.5 //0.5

	//Camera
	d.camera.focal_length = 1.25   //1.0
	d.camera.viewport_height = 2.0 //2.0
	d.camera.aspect_ratio = d.image.aspect_ratio

	//Ray
	d.ray.iterationDepth = 5         //50
	d.ray.reflectionThreshold = 0.01 //0.001

	return *d
}
