package main

import "fmt"

func main() {
	age := 18

	if age > 18 {
		fmt.Println("You are an adult.")
	} else if age == 18 {
		fmt.Println("You are 18.")
	} else {
		fmt.Println("You are a minor.")
	}
}
