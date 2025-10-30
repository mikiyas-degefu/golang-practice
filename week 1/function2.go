package main

import (
	"fmt"
)


func values()(int, int){
	return 1,3
}

func main(){
	a, b := values()

	fmt.Println(a,b)
}