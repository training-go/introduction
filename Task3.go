package main

import "fmt"

func main() {
	s := []int{4, 3, 6, -7, 2, 8}
	help := 0
	for i := 0; i < len(s)/2; i++ {
		help = s[i]
		s[i] = s[len(s)-i-1]
		s[len(s)-i-1] = help
	}

	fmt.Print(s)
}
