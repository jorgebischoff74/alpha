package main

import(
 "fmt"
)

func main(){
	x:= 15
	y:=15

	if (x < y){
		fmt.Printf("%d es menor que %d \n", x , y)
	} else if ( x > y){
		fmt.Println(" Entre al elseif")
	} else{
		fmt.Println("Entre al else")
	}
}

