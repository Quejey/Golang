package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, n float64
	var f int
	for {
		var method int
		fmt.Print("Введите арифметическую операцию:\n 1 - простой калькулятор;\n 2 - возведения числа в степень;\n 3 - факториал числа;\n 4 - выход;\n")
		fmt.Scanln(&method)
		if method == 1 {
			simpleCalculator(a, b)
		} else if method == 2 {
			numberToPowerN(a, n)
		} else if method == 3 {
			factorial(f)
		} else if method == 4 {
			break
		} else {
			fmt.Println("Ошибка")
		}
	}
}

func simpleCalculator(a, b float64) {
	var op string
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)
	fmt.Print("Введите арифметическую операцию (+, -, *, /): ")
	fmt.Scanln(&op)
	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("\nОшибка, На ноль делить нельзя!\n")
		} else {
			fmt.Printf("%.2f", a/b)
		}
	}
}

func numberToPowerN(a, n float64) {
	fmt.Print("Введите число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите степень: ")
	fmt.Scanln(&n)

	power := math.Pow(a, n)
	fmt.Println(a, "в степени", n, "=", power)
}

func factorial(f int) {
	factorial := 1
	fmt.Print("Введите число: ")
	fmt.Scanln(&f)
	for i := 1; i <= f; i++ {
		factorial = factorial * i
	}
	fmt.Println("Факториал ", f, " = ", factorial)
}
