package main

import (
	"fmt"
)

func main() {

	var arr [5]int32
	fmt.Printf("Empty array declaration: %v\n", arr)

	n := [][]int32{{1, 2, 3}, {4, 5, 6}}
	fmt.Printf("Initialised matrix: %v\n", n)

	var m [3][3]int32
	for i := int32(0); i < 3; i++ {
		for j := int32(1); j <= 3; j++ {
			m[i][j-1] = i*3 + j
		}
	}
	fmt.Printf("Constructed matrix: %v\n", m)
}
