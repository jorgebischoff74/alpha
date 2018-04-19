package main

import(
	"fmt"
	"os"
	"bufio"

)

func main(){
	execreadfile()
	
}

func execreadfile(){
	ver := readfile()
	fmt.Println(ver)

}

func readfile() bool {

	archivo, err := os.Open("./hola.txt")

	defer func(){
		archivo.Close()
		fmt.Println("Defer")
		recover()
	}()	

	if (err != nil){
		panic(err)
		fmt.Println("Error leyendo el archio")
	}
	
	scanner := bufio.NewScanner(archivo)

	var i int

	for scanner.Scan(){
		i++
		linea := scanner.Text()
		fmt.Println(i)
		fmt.Println(linea)
	}

	
	return true

}