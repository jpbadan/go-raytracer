package main

// Golang raytracer - João BADAN @ 2022

// displayImage renders an image to the playground's console by
// base64-encoding the encoded image and printing it to stdout
// with the prefix "IMAGE:".

func main() {
	//TestImageOutput()
	defaultSettings := getDefaults()

	raytracer("test-render", defaultSettings)
}
