package main

import (
	"fmt"
)

func qsort(a []int) []int {
	if len(a) < 2 { return a }

	left, right := 0, len(a) - 1

	pivotIndex := len(a)/2


	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	qsort(a[:left])
	qsort(a[left + 1:])

	return a
}

func printArray(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d ", a[i])
	}
}


func main(){
	slice := []int{9, 3, -4, 6, -5, 3}

	qsort(slice)
	printArray(slice)

}
