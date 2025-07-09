package main

import (
	"fmt"
)

func findMax(numbers []int) int {
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			temp := numbers[i]
			numbers[i] = numbers[i+1]
			numbers[i+1] = temp
		}
	}
	max := numbers[len(numbers)-1]

	return max
}

func main() {

	numbers := []int{1, 3, 2, 8}
	max := findMax(numbers)
	fmt.Println(max)
}
