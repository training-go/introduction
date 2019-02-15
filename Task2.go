package main

import "fmt"

func getMinElement(s []int) int {
	var help int = s[0]
	for i := 0; i < len(s); i++ {
		if help > s[i] {
			help = s[i]
		}
	}
	return help
}

func getSum(s []int) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum = sum + s[i]
	}
	return sum
}

func printSlice(s []int) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("Element %v is %v\n", i, s[i])
	}
}

func main() {
	slice := []int{4, 3, 6, -7, 2, 8}

	fmt.Println(getMinElement(slice))
	fmt.Println(getSum(slice))
	printSlice(slice)
}
