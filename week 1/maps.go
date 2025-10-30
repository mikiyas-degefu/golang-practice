package main

import (
	"fmt"
)


func main(){
	m := make(map[string]int) //make(map[key-type]val-type).

	fmt.Println(m)

	m["k1"] = 10
	m["k2"] = 20
	m["32"] = 20

	fmt.Println(m)

	delete(m, "k2")
	fmt.Println(m)
	

	_, val := m["k1"] // check value is exist
	fmt.Println(val)
}