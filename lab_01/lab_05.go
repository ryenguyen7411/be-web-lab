package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func sum_and_product(a, b int) (int, int) {
	return a + b, a * b
}

func greet(name string) {
	fmt.Println("Hello,", name, "1 + 1 =", add(1, 1))
}

func main() {
	greet("Go Learner")

	a, b := sum_and_product(1, 2)
	fmt.Println("1 + 2 =", a, "and 1 * 2 =", b)
}
