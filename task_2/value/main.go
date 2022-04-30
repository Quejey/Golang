package main

import (
	"fmt"
)

func main() {
	var number int
	var a int
	var b int
	var c int

	fmt.Println("Enter the number:")
	fmt.Scanln(&number)

	a = number / 100
	b = (number % 100) / 10
	c = number % 10

	fmt.Println(a, "hundred")
	fmt.Println(b, "dozen")
	fmt.Println(c, "unit")
}
