package main

import "image"

func TestImageOutput() {
	width, height := 512, 256

	newRgba := image.NewRGBA(image.Rect(0, 0, width, height)) //new image

	for j := height - 1; j >= 0; j-- {
		for i := 0; i < width; i++ {

			pixelColor := Color{float32(i) / float32(width-1), float32(j) / float32(height-1), 0.255}

			newRgba.SetRGBA(i, j, pixelColor.toRGBA())
		}
	}

	saveImage(newRgba, "test-Image")
	println("Exported test image to output folder")
}
