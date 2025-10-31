package main

import "fmt"


type person struct{
	name string
	age int
}



func main(){
	s := person{name: "Mikiyas", age: 25}

	fmt.Println(s.name)
	fmt.Println(s.age)

	fmt.Println(&person{name: "Ann", age: 40})
}