package main

import (
	"fmt"
	"math"
)

func main() {

	// Classical for loop
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}

	// While loop
	fmt.Println()
	j := 0
	for j < 10 {
		if j%3 == 0 {
			fmt.Printf("%d ", j)
		}
		j++
	}

	// Infinite (breakable) loop
	fmt.Println()
	arr := []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	it := 0
	for {
		if it >= len(arr) {
			break
		} else {
			fmt.Printf("%d ", arr[it])
			it = int(math.Pow(math.Sqrt(float64(it))+1, 2))
		}
	}

}
