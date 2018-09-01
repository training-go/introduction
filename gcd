package main

import (
	"fmt"
)

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
		fmt.Println("first iter === ", a)
	}
	return a
}

func main() {

	var x, y int
	fmt.Println("Enter int one, then pres enter and enter int two!")
	fmt.Scan(&x, &y)
	result := GCD(x, y)
	fmt.Println(" GCD is ==== ", result)
}
