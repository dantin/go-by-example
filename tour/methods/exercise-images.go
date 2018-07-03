// version 1.0 2018-07-03
package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	w, h int
	c    uint8
}

func (img *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)

}

func (img *Image) At(x, y int) color.Color {
	return color.RGBA{img.c + uint8(x*y+x*y), img.c + uint8(x*y+x*y), 255, 255}
}

func main() {
	m := Image{100, 100, 128}
	pic.ShowImage(&m)
}
