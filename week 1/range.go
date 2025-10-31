package main

import "fmt"


func main(){
	nums := []int{1,2,3,4}
	sum := 0

	for _,num := range nums {
		sum += num
	}

	fmt.Println(sum)

	for i,num := range nums { //index, val
		fmt.Println(i, num)
	}

	items := map[string]string{"a": "apple", "b" : "banana"}
	
	for key, val := range items {
		fmt.Println(key, val)
	}

	for i,c := range "go"{
		fmt.Println(i,c)
	}
}