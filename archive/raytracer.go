package main

import (
	"image"
	"math"
	"math/rand"
)

func raytracer(frameName string, defaults Defaults) {

	// Image
	aspect_ratio := defaults.image.aspect_ratio
	image_width := defaults.image.width
	image_height := int(float32(image_width) / aspect_ratio)

	//World
	var world HittableList

	// world.clear()
	materialGround := Lambertian{Color{0.8, 0.8, 0.0}}
	materialCenter := Lambertian{Color{0.3, 0.4, 0.5}}
	materialLeft := Metal{Color{0.8, 0.8, 0.8}, 0.05}
	materialRight := Metal{Color{0.8, 0.6, 0.2}, 1.0}

	world.add(Sphere{Point3{0, -100.5, -1}, 100, materialGround})
	world.add(Sphere{Point3{0, 0, -1}, 0.5, materialCenter})
	world.add(Sphere{Point3{-1, 0, -1}, 0.5, materialLeft})
	world.add(Sphere{Point3{1, 0, -1}, 0.5, materialRight})

	// for i := float32(-0.5); i <= 0.5; i += 0.01 {
	// 	world.add(Sphere{Point3{i, 0, -1}, float32(math.Abs(float64(i * 0.6))), materialLeft})
	// }

	// Camera
	cam := newCamera(defaults)

	// Render
	newRgba := image.NewRGBA(image.Rect(0, 0, image_width, image_height))

	for j := image_height - 1; j >= 0; j-- {
		for i := 0; i < image_width; i++ {

			pixelColor := Color{0, 0, 0}
			for s := 0; s < defaults.image.samples_per_pixel; s++ {

				u := (float32(i) + rand.Float32()) / float32(image_width-1)
				v := (float32(j) + rand.Float32()) / float32(image_height-1)

				ray := cam.getRay(u, v)
				pixelColor = pixelColor.sum(rayColor(ray, &world, defaults.ray.iterationDepth, defaults))
			}

			newRgba.SetRGBA(i, int(math.Abs(float64(image_height-j))), pixelColor.correct(defaults).toRGBA())

		}
		saveImage(newRgba, frameName)

	}
}
