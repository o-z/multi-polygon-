package main

import (
	"fmt"

	"math"
)

func main() {
	var edge int
	fmt.Scanf("%d", &edge)
	matrix := [51][51]int{}
	var x2, y2 int
	x := 25.0
	y := 0.0
	var x1 float64
	var y1 float64

	for i := 0; i < edge; i++ {
		x1 = x*math.Cos((2*math.Pi)/float64(edge)) - y*math.Sin((2*math.Pi)/float64(edge)) // source: https://en.wikipedia.org/wiki/Cartesian_coordinate_system#Translation
		y1 = x*math.Sin((2*math.Pi)/float64(edge)) + y*math.Cos((2*math.Pi)/float64(edge))

		x = x1
		y = y1

		if x1 > float64(int(x1))+0.5 {
			x2 = int(x1+1) + 25
		} else {
			x2 = int(x1) + 25
		}
		if y1 > float64(int(y1))+0.5 {
			y2 = int(y1+1) + 25
		} else {
			y2 = int(y1) + 25
		}
		matrix[x2][y2] = 1

	}
	for i := 0; i <= 50; i++ {
		for j := 0; j <= 50; j++ {
			if matrix[i][j] == 1 {
				fmt.Print("x")

			} else {
				fmt.Print(" ")
			}
			fmt.Print(" ")

		}
		fmt.Println()
	}
}
