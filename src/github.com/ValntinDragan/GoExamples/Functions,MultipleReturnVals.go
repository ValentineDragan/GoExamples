package main

import (
	"fmt"
	"math"
)

func calcSum(arr []int32) int32 {
	sum := int32(0)
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func isSumSquare(arr []int32) (bool, int32) {
	sum := calcSum(arr)
	intSqrt := int32(math.Sqrt(float64(sum)))
	if intSqrt*intSqrt == sum {
		return true, intSqrt
	}
	return false, 0
}

func printSumSquare(arr []int32) {
	isSquare, root := isSumSquare(arr)
	if isSquare {
		fmt.Printf("%v has the sum %v, which is a perfect square of root %v!\n", arr, root*root, root)
	} else {
		fmt.Printf("%v does not have a perfect square sum of its elements.\n", arr)
	}
}

func main() {
	a := []int32{1, 3, 5, 7, 9, 11, 13, 15}
	fmt.Printf("The sum of %v is: %v\n", a, calcSum(a))

	printSumSquare(a)

	printSumSquare([]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}
