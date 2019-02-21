package main

import "fmt"

//Arranges the numbers in a slice from the lowest to the highest
func QuickSort(a []int) []int {
	if len(a) < 2 { return a }
	pivot := a[len(a)/2]

	var rightSlice []int
	var leftSlice []int
	for _, v := range a {
		if v < pivot {
			leftSlice = append(leftSlice,v)
		} else if v >= pivot{
			rightSlice = append(rightSlice,v)
		}
	}
	//leftSlice = append(leftSlice,pivot)

	a = append(append([]int{}, QuickSort(leftSlice)...))
	a = append(append([]int{}, QuickSort(rightSlice)...))

	return a
}

func printArray(a []int) {
	fmt.Println(a)
}

func main(){
	slice := []int{9, 3, -4, 6, -5, 4}

	s :=QuickSort(slice)
	printArray(s)
}