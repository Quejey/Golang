package main

import "fmt"

func main() {
	var width int
	var length int
	var area int

	fmt.Println("Enter the width:")
	fmt.Scanln(&width)
	fmt.Println("Enter the length:")
	fmt.Scanln(&length)
	area = width * length

	fmt.Println("The area of rectangle equal", area)
}
