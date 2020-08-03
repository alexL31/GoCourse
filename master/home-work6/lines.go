package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	green := color.RGBA{0, 255, 0, 255}
	white := color.RGBA{255, 255, 255, 1}
	rectImg := image.NewRGBA(image.Rect(0, 0, 200, 200))

	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)

	for x := 0; x < 200; x++ {
		rectImg.Set(x, 50, white)
		rectImg.Set(x, 150, white)
		rectImg.Set(50, x, white)
		rectImg.Set(150, x, white)
	}

	file, err := os.Create("rectangle.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()
	png.Encode(file, rectImg)
}
