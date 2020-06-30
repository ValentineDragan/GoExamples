package main

import (
	"errors"
	"fmt"
	"math"
)

// Function that calculates square root (of positive numbers only!)
// Returns an error if input is negative
func getSquareRoot(x int) (float64, error) {
	if x < 0 {
		return -1, errors.New("Can't calculate root for negative number")
	}
	return math.Sqrt(float64(x)), nil
}

// Custom Error Type
type negArgError struct {
	arg     int
	message string
}

func (e *negArgError) Error() string {
	return fmt.Sprintf("error on input %d: %s", e.arg, e.message)
}

// Implementation of the same function, but with custom error type
func getSquareRoot2(x int) (float64, error) {
	if x < 0 {
		return -1, &negArgError{x, "Negative number"}
	}
	return math.Sqrt(float64(x)), nil
}

func main() {
	nums := []int{10, -3, 48, 0}

	// using method 1
	for _, x := range nums {
		if r, err := getSquareRoot(x); err == nil {
			fmt.Printf("Root of %d is: %f\n", x, r)
		} else {
			fmt.Printf("Root for %d failed: %v\n", x, err)
		}
	}

	// using method 2
	fmt.Println()
	for _, x := range nums {
		if r, err := getSquareRoot2(x); err == nil {
			fmt.Printf("Root of %d is: %f\n", x, r)
		} else {
			fmt.Printf("Root for %d failed: %v\n", x, err)
		}
	}

}
