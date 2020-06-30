package main

import "fmt"

func main() {
	someStr := "Hello olo eHe LEEe"
	charFreq := map[byte]int{}

	// Creating/Updating (key, value) pairs
	for i := 0; i < len(someStr); i++ {
		charFreq[someStr[i]]++
	}

	// Using range to iterate over (key, value) pairs
	for k, v := range charFreq {
		fmt.Printf("{%c: %d} ", k, v)
	}
	fmt.Println()

	// Checking if (key, value) pairs are present
	stats := map[string]float32{"strength": 2.1, "intelligence": 1.5, "charisma": 0.7, "agility": 1.123}

	luckStat, prs := stats["luck"]
	fmt.Printf("Luck present: %v; value: %v\n", prs, luckStat)

	fmt.Printf("Present stats: ")
	for k := range stats {
		fmt.Printf("%v ", k)
	}
	fmt.Println()
}
