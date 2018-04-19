package main

import(
	"fmt"
)

func main(){

	fmt.Println("Esto es ciclos")
	i :=0 
	for  {
		
		if ( i == 3){
			i++
			continue
			
		}
		
		fmt.Printf("Ciclo: %d \n",i)
		
		i++
		
		if ( i == 10){
			break
		}
	}

}