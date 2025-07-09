package main

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	target := 6
	result := twoSum(nums, target)
	fmt.Println(result)
}

func BenchmarkTwoSum(b *testing.B) {
	nums := []int{2, 3, 0, 5}
	target := 0
	for i := 0; i < b.N; i++ {
		_ = twoSum(nums, target)
	}
}
