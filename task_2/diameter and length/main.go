package main

import (
	"fmt"
	"math"
)

func main() {
	pi := math.Pi
	var radius float64
	var area_of_circle float64
	var diameter float64
	var length_of_circle float64

	fmt.Println("Enter the area of circle:")
	fmt.Scanln(&area_of_circle)

	radius = math.Sqrt(area_of_circle / pi)
	diameter = 2 * radius
	length_of_circle = 2 * pi * radius

	fmt.Println("The diameter equal", diameter)
	fmt.Println("The length of circle equal", length_of_circle)
}
