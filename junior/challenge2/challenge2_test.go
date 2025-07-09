package main

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	word := "A man a plan"
	res := isPalindrome(word)
	fmt.Println(res)
}

func BenchmarkIsPalindrome(b *testing.B) {
	word := "racecar"
	for i := 0; i < b.N; i++ {
		_ = isPalindrome(word)
	}

}
