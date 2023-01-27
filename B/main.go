package main

import "fmt"

func main() {
	val, found := FindUnique([]int{3, 3, 4, 4, 5})
	fmt.Println(val, found)

	val, found = FindUnique([]int{2, 2, 5, 6, 5})

	fmt.Println(val, found)
}

func FindUnique(numbers []int) (unique int, found bool) {
	for _, num := range numbers {
		unique ^= num
	}

	return unique, unique != 0
}
