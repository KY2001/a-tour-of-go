package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type I interface {
	ColorModel() color.Model
	Bounds() image.Rectangle
	At(x, y int) color.Color
}

type Image struct{}

func (a Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (a Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 500, 500)
}

func (a Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x) * uint8(y), uint8(x) ^ (uint8(x) ^ uint8(y)), 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
