package main

import "fmt"

func getMinElement(s []int) int {
    smallest := 9223372036854775807

    for _, j := range s {
        if smallest > j {
          smallest = j
        }
    }
    
    return smallest    
}

func main() {
 s := []int{4, 5, 23, 10}

 fmt.Println(getMinElement(s))
}
