package main

import (
	"fmt"
)

func countChars(strings ...string) {
	fmt.Printf("%s :", strings)
	counter := 0
	for _, s := range strings {
		counter += len(s)
	}
	fmt.Printf("%d\n", counter)
}

func main() {
	countChars("the", "quick", "brown", "fox")

	strings := []string{"jumps", "over", "the", "lazy", "dog"}
	countChars(strings...)
}
