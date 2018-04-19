package main

import (
	"fmt"
)

func main(){
	var x,y *int
	valor := 5

	x = &valor
	y = &valor

	fmt.Println(x)
	fmt.Println(*x)
	fmt.Println(y)
	fmt.Println(*y)

	*x = 6
	fmt.Println("---------------------")
	fmt.Println(x)
	fmt.Println(*x)
	fmt.Println(y)
	fmt.Println(*y)	


}