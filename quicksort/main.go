package main

import (
	"fmt"
)

//Implements the quicksort function
func QuickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivotIndex := len(a) / 2

	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	QuickSort(a[:left])
	QuickSort(a[left+1:])

	return a
}

//Prints the array
func PrintArray(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d ", a[i])
	}
}
