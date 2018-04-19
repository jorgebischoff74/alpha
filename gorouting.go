package main

import(
	"fmt"
	"strings"
	"time"
)

func nombre_lento(name string){

	letras := strings.Split(name,"")
	for _,letra := range (letras){
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)

	}
}

func main(){

	go nombre_lento("PepeJuanito")
	fmt.Println("Me quede esperando")

	var wait string
	fmt.Scanln(&wait)


}