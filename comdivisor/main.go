package main

import "fmt"

//Finds the greatest divisor between two numbers
func GreatestDivisor(a, b int) int {
	if a == 0 {
		return b
	}

	for b != 0 {
		if a > b {
			a = a - b
			continue
		}
		b = b - a
	}
	return a
}

func main() {
	fmt.Print(GreatestDivisor(21, 9))
}
