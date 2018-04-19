package main

import(
	"fmt"
	"strconv"
)

func main(){
	edad := "44"
	str_edad,_ := strconv.Atoi(edad)

	fmt.Println(str_edad + 10)

}
