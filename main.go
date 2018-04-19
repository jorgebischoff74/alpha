package main


import (
	"net/http"
	"log"
	"fmt"
)

func main(){

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hola Mundo")
	})
	http.HandleFunc("/dos", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hola Mundo, dos")
	})
	log.Fatal(http.ListenAndServe("localhost:3000", nil))

}
