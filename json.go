package main

import (
	"net/http"
	"encoding/json"
)

type Curso struct{
	Title string `json:"titulo"`
	NumeroVideos int `json:"numero de video"`
}

type Cursos []Curso

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		cursos := Cursos{
			Curso{"Curso de Go", 10},
			Curso{"Curso de Rudy", 20},
			Curso{"Curso de NodeJS", 30},
			Curso{"Curso de Java", 40},
			}
		json.NewEncoder(w).Encode(cursos)
	})

	http.ListenAndServe(":8081",nil)
}
