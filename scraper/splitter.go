package main

import (
	"image"
	"image/png"
	"os"
	"strconv"
)

type subImager interface {
	SubImage(r image.Rectangle) image.Image
}

func splitFromTop(img image.Image, size int) []image.Image {
	result := []image.Image{
		img.(subImager).SubImage(image.Rect(0, 0, size, size)),
		img.(subImager).SubImage(image.Rect(0, size, size, 2*size)),
		img.(subImager).SubImage(image.Rect(size, 0, 2*size, size)),
		img.(subImager).SubImage(image.Rect(size, size, 2*size, 2*size)),
	}
	return result
}

func saveSplitImages(path string, location string, splitImages []image.Image) {
	for i, img := range splitImages {
		file, err := os.Create(path + location + strconv.Itoa(i) + ".png")
		handleErr(err)
		png.Encode(file, img)
		defer file.Close()
	}
}
