package main

import "golang.org/x/tour/pic"

func createPic(dx, dy int) [][]uint8 {
	array := make([][]uint8, dy)
	for y := range array {
		array[y] = make([]uint8, dx)
	}

	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			array[y][x] = uint8(x ^ y)
		}
	}

	return array
}

func main() {
	pic.Show(createPic)
}
