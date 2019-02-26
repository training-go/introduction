package main

import "fmt"

// QuickSort arranges the numbers in a slice from the lowest to the highest using two slices
func QuickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	pivot := a[len(a)-1]

	var rightSlice []int
	var leftSlice []int
	for i := 0; i < len(a)-1; i++ {
		if a[i] < pivot {
			leftSlice = append(leftSlice, a[i])
			continue
		}
		rightSlice = append(rightSlice, a[i])
	}
	a = QuickSort(leftSlice)
	a = append(a, pivot)
	a = append(a, QuickSort(rightSlice)...)
	return a
}

func main() {
	slice := []int{91, -3, -4, 6, -5, 4, 4, 91, 4}
	s := QuickSort(slice)
	fmt.Print(s)
}
