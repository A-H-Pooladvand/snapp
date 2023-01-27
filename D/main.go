package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{12, 54, 89, 21, 66, 47, 14, 285, 96}

	// Create a channel to send numbers through the pipeline
	pipeline := make(chan int)

	// Start 2 concurrent worker functions
	for i := 0; i < 2; i++ {
		go worker(pipeline)
	}

	// Send the numbers through the pipeline
	for _, num := range nums {
		pipeline <- num
	}
	close(pipeline)

	// Wait for all workers to finish
	var input string
	fmt.Scanln(&input)
}

func worker(pipeline <-chan int) {
	for num := range pipeline {
		// Do the calculation
		result := math.Pow(float64(num), 2)
		fmt.Println(result)
	}
}
