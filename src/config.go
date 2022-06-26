package main

type Defaults struct {
	output_path string

	image struct {
		aspect_ratio      float32
		width             int
		samples_per_pixel int
		gamma_correction  float32
	}

	camera struct {
		focal_length    float32
		viewport_height float32
		aspect_ratio    float32
	}

	ray struct {
		iterationDepth      int
		reflectionThreshold float32
	}
}

// Constructor
func getDefaults() Defaults {
	d := new(Defaults)

	//File
	d.output_path = "../out"

	//Image
	d.image.aspect_ratio = 16.0 / 10.0
	d.image.width = 150
	d.image.samples_per_pixel = 50 //100
	d.image.gamma_correction = 0.5 //0.5

	//Camera
	d.camera.focal_length = 1.25   //1.0
	d.camera.viewport_height = 2.0 //2.0
	d.camera.aspect_ratio = d.image.aspect_ratio

	//Ray
	d.ray.iterationDepth = 10        //50
	d.ray.reflectionThreshold = 0.01 //0.001

	return *d
}
