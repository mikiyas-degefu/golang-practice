package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	reg, _ := regexp.Compile(`[^\w]`)
	s = reg.ReplaceAllString(s, "")

	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right]{
			return false
		}

		left++
		right --
	}

	return true

}

func main(){
	t1 := "wow"
	fmt.Println(isPalindrome(t1))

	t2 := "noon"
	fmt.Println(isPalindrome(t2))

	t3 := "noan"
	fmt.Println(isPalindrome(t3))

	t4 := "wwa"
	fmt.Println(isPalindrome(t4))
}