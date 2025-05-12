package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func squareNumber(num int) int {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return num * num
}

func main() {
	// Generate a slice of integers
	numbers := make([]int, 10)
	for i := 0; i < 10; i++ {
		numbers[i] = rand.Intn(10) + 1
	}

	fmt.Println("Input numbers:", numbers)

	inputChan := make(chan int)
	outputChan := make(chan int)

	var workerWg sync.WaitGroup
	var collectorWg sync.WaitGroup

	numWorkers := 3
	workerWg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		workerID := i + 1
		go func(id int) {
			defer workerWg.Done()

			for num := range inputChan {
				fmt.Printf("Worker %d processing: %d\n", id, num)
				result := squareNumber(num)
				outputChan <- result
			}
			fmt.Printf("Worker %d finished\n", id)
		}(workerID)
	}

	// Start collector goroutine to collect results
	collectorWg.Add(1)
	results := make([]int, 0, len(numbers))
	go func() {
		defer collectorWg.Done()
		for result := range outputChan {
			results = append(results, result)
		}
	}()

	// Send numbers to workers
	for _, num := range numbers {
		inputChan <- num
	}
	close(inputChan)

	workerWg.Wait()
	close(outputChan)

	collectorWg.Wait()

	fmt.Println("Results:", results)
	fmt.Println("Processing complete!")
}
