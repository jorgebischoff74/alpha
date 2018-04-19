package main

import (
	"fmt"
	"io/ioutil"
)

func main(){

	filedata,err := ioutil.ReadFile("./hola1.txt")

	if (err != nil){
		fmt.Println(err)
	}else {
		fmt.Println("El archivo contiene: " + string(filedata))
	}

}