package main

import "fmt"

func average(nums []int) float64 {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return float64(sum) / float64(len(nums))
}

func main() {
	nums := []int{10, 20, 30, 40, 50}
	nums = append(nums, 60)

	for i, v := range nums {
		fmt.Printf("Index %d: Value %d\n", i, v)
	}

	fmt.Println("Average:", average(nums))
}
