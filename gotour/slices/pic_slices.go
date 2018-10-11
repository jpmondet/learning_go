package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var pixels [][]uint8
	for i := 0; i < dy; i++ {
		element := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			element[j] = uint8((j ^ i) ^ (i^j)*2)
		}
		pixels = append(pixels, element)
	}

	return pixels
}

func main() {
	pic.Show(Pic)
}
