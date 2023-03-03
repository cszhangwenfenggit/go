package main

import (
	"fmt"
	"image/color"
)

type Point struct {
	x, y float64
}

type ColorPoint struct {
	Point
	Color color.RGBA
}

func main() {
	var cp ColorPoint
	cp.Point.x = 1
	fmt.Println(cp.x)
}
