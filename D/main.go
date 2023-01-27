package main

import (
	"fmt"
	"math"
)

func main() {
	initial := []int{12, 54, 89, 21, 66, 47, 14, 285, 96}

	pipeline := make(chan int)

	go worker(pipeline)
	go worker(pipeline)

	for _, n := range initial {
		pipeline <- n
	}

	close(pipeline)
}

func worker(pipeline <-chan int) {
	for n := range pipeline {
		result := math.Pow(float64(n), 2)
		fmt.Println(result)
	}
}
