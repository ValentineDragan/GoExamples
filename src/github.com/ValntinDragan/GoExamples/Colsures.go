package main

import (
	"fmt"
)

// intSeq is returning a function (which in turn returns an int)
func intSeq() func() int {
	i := 0
	// the returned function encloses over the variable i
	return func() int {
		i++
		return i
	}
}

func main() {

	seq := intSeq()

	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())

	newSeq := intSeq()
	fmt.Println(newSeq())
}
