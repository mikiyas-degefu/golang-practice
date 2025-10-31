package main

import "fmt"

func byValue(ival int) {
	ival = 0
}

func byRef(iptr *int){
	*iptr = 0
}


func main(){
	i := 1

	byValue(i)
	fmt.Println(i)

	byRef(&i)
	fmt.Println(i)
	
}