package main

import (
	"go-raytracer/entities"
	"go-raytracer/geometries"
	"go-raytracer/lib"
	"go-raytracer/materials"
	"image"
	"math"
	"math/rand"
)

func Render(frameName string, defaults Defaults) {

	// Image
	aspect_ratio := defaults.image.aspect_ratio
	image_width := defaults.image.width
	image_height := int(float64(image_width) / aspect_ratio)

	//World
	var world entities.GeometryList

	// world.clear()
	materialGround := materials.Diffuse{Albedo: lib.Color{R: 0.8, G: 0.8, B: 0.0}}
	materialCenter := materials.Diffuse{Albedo: lib.Color{R: 0.3, G: 0.4, B: 0.5}}
	materialLeft := materials.Metal{Albedo: lib.Color{R: 0.8, G: 0.8, B: 0.8}, Fuzz: 0.05}
	materialRight := materials.Metal{Albedo: lib.Color{R: 0.8, G: 0.6, B: 0.2}, Fuzz: 1.0}

	world.Add(geometries.Sphere{Center: lib.Point{X: 0, Y: -100.5, Z: -1}, Radius: 100, Material: materialGround})
	world.Add(geometries.Sphere{Center: lib.Point{X: 0, Y: 0, Z: -1}, Radius: 0.5, Material: materialCenter})
	world.Add(geometries.Sphere{Center: lib.Point{X: -1, Y: 0, Z: -1}, Radius: 0.5, Material: materialLeft})
	world.Add(geometries.Sphere{Center: lib.Point{X: 1, Y: 0, Z: -1}, Radius: 0.5, Material: materialRight})

	// for i := float64(-0.5); i <= 0.5; i += 0.01 {
	// 	world.add(Sphere{lib.Point{i, 0, -1}, float64(math.Abs(float64(i * 0.6))), materialLeft})
	// }

	// Camera
	cam := entities.Camera{}

	cam.Viewport.Height = defaults.camera.viewport_height
	cam.Viewport.Width = defaults.camera.aspect_ratio * defaults.camera.viewport_height
	cam.Viewport.Aspect_ratio = defaults.camera.aspect_ratio
	cam.Focal_length = defaults.camera.focal_length
	cam.Origin = lib.Point{X: 0, Y: 0, Z: 0}
	cam.Horizontal = lib.Vec{X: cam.Viewport.Width, Y: 0, Z: 0}
	cam.Vertical = lib.Vec{X: 0, Y: cam.Viewport.Height, Z: 0}
	cam.Lower_left_corner = cam.Origin.Sub(cam.Horizontal.Div(2.0)).Sub(cam.Vertical.Div(2.0)).Sub(lib.Vec{X: 0, Y: 0, Z: cam.Focal_length})

	// Render
	newRgba := image.NewRGBA(image.Rect(0, 0, image_width, image_height))

	for j := image_height - 1; j >= 0; j-- {
		for i := 0; i < image_width; i++ {

			pixelColor := lib.Color{R: 0, G: 0, B: 0}
			for s := 0; s < defaults.image.samples_per_pixel; s++ {

				u := (float64(i) + rand.Float64()) / float64(image_width-1)
				v := (float64(j) + rand.Float64()) / float64(image_height-1)

				ray := cam.GetRay(u, v)
				pixelColor = pixelColor.Sum(entities.RayColor(ray, &world, defaults.ray.iterationDepth, defaults.ray.reflectionThreshold))
			}

			newRgba.SetRGBA(i, int(math.Abs(float64(image_height-j))), pixelColor.Correct(defaults.image.samples_per_pixel, defaults.image.gamma_correction).ToRGBA())

		}
		lib.SaveImage(newRgba, frameName, defaults.output_path)

	}
}
