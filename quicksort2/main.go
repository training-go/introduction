package main

import (
	"fmt"
)

//Arranges the numbers in a slice from the lowest to the highest
func QuickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	biggestNum := a[len(a)/2]

	var rightSlice []int
	for _, v := range a {
		if v > biggestNum {
			biggestNum = v
		}
	}

	for _, v := range a {
		if v < biggestNum {
			rightSlice = append(rightSlice, v)
		}
	}
	a = append(append([]int{}, QuickSort(rightSlice)...))
	a = append(a, biggestNum)
	return a
}

//Prints slice
func PrintArray(a []int) {
	fmt.Println(a)
}

func main() {
	slice := []int{9, 3, -4, 6, -5, 4}

	s := QuickSort(slice)
	PrintArray(s)
}
