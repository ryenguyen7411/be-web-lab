package main

import "fmt"

func getStudentStatus(isStudent bool) string {
	if isStudent {
		return "a student"
	}
	return "not a student"
}

func main() {
	var age int = 25
	var name string = "Alex"
	height := 1.75 // shorthand
	isStudent := false

	fmt.Println(name, "is", age, "years old and", height, "meters tall. He is", getStudentStatus(isStudent))
	fmt.Printf("age: %T\n", age)
	fmt.Printf("name: %T\n", name)
	fmt.Printf("height: %T\n", height)
	fmt.Printf("isStudent: %T\n", isStudent)
}
