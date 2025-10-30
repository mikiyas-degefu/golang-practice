package main

import (
	"fmt"
)

func add(a int, b int) int{
	return a + b
}


func main(){
	a := 12
	b := 43
	result := add(a,b)

	fmt.Println(result)
}