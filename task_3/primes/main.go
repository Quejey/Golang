package main

import "fmt"

func main() {
	var n int
	fmt.Println("Введите максимальное число:")
	fmt.Scanln(&n)
	fmt.Println("Все простые числа до", n)
	for i := 0; i <= n; i++ {
		if IsSimple(i) {
			fmt.Println(i)
		}
	}
}

func IsSimple(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
