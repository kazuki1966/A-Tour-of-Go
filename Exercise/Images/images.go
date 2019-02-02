package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image :: インターフェイス定義
type Image struct{}

// ColorModel :: イメージのカラーモデル
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds :: イメージの大きさ
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 256, 256)
}

// At :: イメージのピクセル単位の色
func (i Image) At(x, y int) color.Color {
	return color.RGB{uint8(x), uint8(y), 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
