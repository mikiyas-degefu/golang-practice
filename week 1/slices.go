package main

import (
	"fmt"
)


func main(){
	var s []string

	fmt.Println(len(s), s == nil)

	s = make([]string, 3)
	fmt.Println(len(s), s, cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)

	s = append(s, "d")
	fmt.Println(s)

	fmt.Println(s[:3])
}