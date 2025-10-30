package main

import "fmt"


func main(){
	i := 1

	fmt.Println("Loop 1")
	for i <= 3{
		fmt.Println(i)
		i++
	}

	fmt.Println("\n\n\nLoop 2")
	for j := 0; j < 3; j++{
		fmt.Println(j)
	}
	
	fmt.Println("\n\n\nLoop 3")
	for i := range 3{
		fmt.Println("range", i)
	}

	fmt.Println("\n\n\nLoop4")
	for {
		fmt.Println("loop")
		break
	}

	fmt.Println("\n\n\nLoop5")
	for n := range 6{
		if n % 2 == 0{
			continue
		}
		fmt.Println(n)
	}

}