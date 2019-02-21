package main

import (
	"fmt"
)

func qsort(a []int) []int {
	if len(a) < 2 { return a }
	biggestnum := a[len(a)/2]

	var rightSlice []int
	for _, v := range a {
		if v > biggestnum {
			biggestnum = v
		}
	}

	for i := range a {
		if a[i] < biggestnum {
			rightSlice = append(rightSlice,a[i])
		}
	}
	a = append(append([]int{}, qsort(rightSlice)...))
	a = append(a,biggestnum)
	return a
}

func printArray(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d ", a[i])
	}
}

func main(){
	slice := []int{9, 3, -4, 6, -5, 4}

	s :=qsort(slice)
	printArray(s)

}
