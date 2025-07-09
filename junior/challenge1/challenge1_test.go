package main

import (
	"fmt"
	"testing"
)

func TestFindMax(t *testing.T) {
	numbers := []int{-1, 0, -2, -3}
	max := findMax(numbers)
	fmt.Println(max)
}

func BenchmarkFindMax(b *testing.B) {
	numbers := []int{-1, 0, -2, -3}
	for i := 0; i < b.N; i++ {
		_ = findMax(numbers)
	}
}
