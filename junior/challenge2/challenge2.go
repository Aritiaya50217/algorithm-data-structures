package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	str := strings.ToLower(s)
	word := strings.ReplaceAll(str, " ", "")

	runes := []rune(word)
	length := len(runes)

	for i := 0; i < length/2; i++ {
		if runes[i] != runes[length-1-i] {
			return false
		}
	}
	return true
}

func main() {
	word := "racecar"
	res := isPalindrome(word)
	fmt.Println(res)
}
