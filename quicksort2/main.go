package main

import (
	"fmt"
)

func qsort(a []int) []int {
	if len(a) < 2 { return a }


	biggestnum := a[len(a)/2]




	var rightslice []int
	for i := range a {
		if a[i] > biggestnum {
			biggestnum = a[i]
		}
	}

	for i := range a {
		if a[i] < biggestnum {
			rightslice = append(rightslice,a[i])
		}
	}



	a = append(append([]int{}, qsort(rightslice)...))
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
