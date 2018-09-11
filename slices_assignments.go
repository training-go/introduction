package main

import (
	"fmt"
)

func main() {
	primes := []int{2, 3, 5, 7, 11, 13, 15}
	smallestNumber := getMinElement(primes, 0, 7)
	fmt.Println("Smallest number is", smallestNumber)
	summOfEls := getSum(primes)
	fmt.Println("Sum number is", summOfEls)
	printSlice(primes)
	fmt.Printf("Ints %v reversed: %v\n", primes, reverseInts(primes))
}

func getMinElement(s []int, min, max int) int {
	var result = max
	if len(s) >= max {
		sliced := s[min:max]
		for _, e := range sliced {
			if e < result {
				result = e
			}
		}
	} else {
		fmt.Println("max exceeds the length of primes")
	}
	return result
}

func getSum(s []int) int {
	var result int = 0
	for _, e := range s {
		result += e
	}
	return result
}

func printSlice(s []int) {
	fmt.Println(s)
}

func reverseInts(s []int) []int {
	if len(s) == 0 {
		return s
	}
	return append(reverseInts(s[1:]), s[0])
}
