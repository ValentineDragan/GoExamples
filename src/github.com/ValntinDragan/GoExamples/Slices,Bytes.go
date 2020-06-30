package main

import (
	"fmt"
)

func main() {

	s1 := make([]int32, 5)
	fmt.Printf("empty: %v len: %v cap: %v\n", s1, len(s1), cap(s1))

	s2 := []int32{}
	s2 = append(s2, 1, 2, 3, 4, 5)
	fmt.Printf("appended: %v len: %v cap: %v\n", s2, len(s2), cap(s2))

	// Copying
	s3 := make([]int32, 3)
	fmt.Printf("before copy: %v len: %v cap: %v\n", s3, len(s3), cap(s3))
	copy(s3, s2)
	fmt.Printf("after copy: %v len: %v cap: %v\n", s3, len(s3), cap(s3))

	// Using Bytes for char
	s4 := []byte{'h', 'i', ' ', 'm', 'y', ' ', 'n', 'a', 'm', 'e', ' ', 'i', 's'}
	fmt.Printf("full array: %c\n", s4)
	fmt.Printf("sliced: %c\n", s4[3:10])

	// Const
	const maxSize = 10
	s5 := make([]bool, maxSize)
	fmt.Printf("%v\n", s5)

}
