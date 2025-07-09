package main

import "fmt"

func twoSum(nums []int, target int) []int {
	mat := make(map[int]int)
	result := []int{}
	for i, num := range nums {
		if val, found := mat[target-num]; found {
			result = append(result, val, i)
			return result
		}
		mat[num] = i
	}
	return result
}

func main() {
	numbers := []int{2, 3, 4, 5}
	target := 5
	result := twoSum(numbers, target)
	fmt.Println(result)
}
