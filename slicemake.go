package main

import (
	"fmt"
)

func main(){
	
	slice:= []int{1,2,3,4,5}
	copia:= make([]int,len(slice), cap(slice)*2)

	copy(copia,slice)

	fmt.Println(slice)
	fmt.Println(copia)

}