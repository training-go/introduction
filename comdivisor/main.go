package main

import "fmt"

func greatestDivisor(a, b int) int {
	if a == 0 {
		return b
	}

	for b != 0 {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}

func main() {
	fmt.Print(greatestDivisor(63, 9))
}



//algorithm quicksort(A, lo, hi) is
//if lo < hi then
//p := partition(A, lo, hi)
//quicksort(A, lo, p - 1)
//quicksort(A, p + 1, hi)
//
//algorithm partition(A, lo, hi) is
//pivot := A[hi]
//i := lo
//for j := lo to hi - 1 do
//if A[j] < pivot then
//swap A[i] with A[j]
//i := i + 1
//swap A[i] with A[hi]
//return i



//I am adding this comment to test git commit


//adding this comment to fix the conflict
