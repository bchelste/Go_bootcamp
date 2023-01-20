package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {

	width := 300
	height := 300
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	rect := image.Rectangle{upLeft, lowRight}

	myImage := image.NewNRGBA(rect)

	cyan := color.RGBA{100, 200, 200, 0xff}
	red := color.RGBA{255, 0, 0, 0xff}
	blue := color.RGBA{0, 0, 255, 0xff}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if ((i < int(float64(height) * 0.25)) || (i > int(float64(height) * 0.75))) ||
				((j < int(float64(height) * 0.25)) || (j > int(float64(height) * 0.75))) {
				myImage.Set(i, j, color.RGBA{
					R: uint8((i + j) & 255),
					G: uint8((i + j) << 1 & 255),
					B: uint8((i + j) << 2 & 255),
					A: 255,
				})
			} else if ((i > int(float64(height) * 0.33)) && (i < int(float64(height) * 0.66))) {
				if ((j > int(float64(width) * 0.33)) && (j < int(float64(width) * 0.44))) {
					myImage.Set(i, j, color.White)
				} else if ((j > int(float64(width) * 0.44)) && (j < int(float64(width) * 0.55))) {
					myImage.Set(i, j, blue)
				} else if ((j > int(float64(width) * 0.55)) && (j < int(float64(width) * 0.66))) {
					myImage.Set(i, j, red)
				} else {
					myImage.Set(i, j, cyan)
				}
			} else {
				myImage.Set(i, j, cyan)
			}
		}
	}
	f, err := os.Create("amazing_logo.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = png.Encode(f, myImage)
	if (err != nil) {
		panic(err)
	}
}