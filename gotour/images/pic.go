package main

import "golang.org/x/tour/pic"
import "image"
import "image/color"

type Image struct {
	width  int
	height int
}

func (im Image) ColorModel() color.Model {
	return color.RGBAModel
}
func (im Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, im.width, im.height)

}
func (im Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
	m := Image{250, 200}
	pic.ShowImage(m)
}
