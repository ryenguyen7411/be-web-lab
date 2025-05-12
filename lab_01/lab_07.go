package main

import "fmt"

func main() {
	scores := map[string]int{
		"Alice": 90,
		"Bob":   85,
		"Rye":   110,
	}

	for name, score := range scores {
		fmt.Printf("%s: %d\n", name, score)
	}

	val, ok := scores["Thanh"]
	if !ok {
		fmt.Println("Thanh not found")
	} else {
		fmt.Println("Thanh's score:", val)
	}
}
