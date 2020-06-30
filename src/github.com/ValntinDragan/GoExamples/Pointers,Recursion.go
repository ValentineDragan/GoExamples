package main

import "fmt"

// Reverses the array (passed through reference)
func reverseArr(arr *[]int) {
	n := len(*arr)
	for i := 0; i < n/2; i++ {
		(*arr)[i], (*arr)[n-i-1] = (*arr)[n-i-1], (*arr)[i]
	}
}

// Finds the path between the current/starting city and the destination, recursively.
// Also marks the visited cities in the cities map (passed by reference) to avoid infinite loops
func findPathRecursively(cities *map[string]string, current string, destination string) (bool, string) {
	if current == destination {
		return true, destination
	}

	nextCity, prs := (*cities)[current]
	if prs {
		if nextCity == "visited" {
			return false, "Loop. City already visited."
		}
		(*cities)[current] = "visited"
		found, path := findPathRecursively(cities, nextCity, destination)
		return found, current + " -> " + path
	}
	return false, "Path not found"

}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Printf("original array: %v\n", a)
	reverseArr(&a)
	fmt.Printf("updated array: %v\n", a)

	cities := map[string]string{"London": "Paris", "Barcelona": "Edinburgh", "Paris": "Barcelona", "Edinburgh": "New York", "Berlin": "Barcelona"}

	found, path := findPathRecursively(&cities, "London", "New York")
	fmt.Println("Searching path between London and New York...")
	if found {
		fmt.Printf("Path found: %v\n", path)
	} else {
		fmt.Println("Path not found.")
	}

}
