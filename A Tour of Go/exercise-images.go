package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	x, y int
}

func (m Image) ColorModel() color.Model {
	return color.NRGBA64Model

}
func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.x, m.y)

}
func (m Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{x: 256, y: 256}
	pic.ShowImage(m)

}
