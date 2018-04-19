package main

import (
	"fmt"
)

type User struct{
	edad int
	nombre, apellido string

}

func (usuario User) nombre_completo() string{
	return usuario.nombre + " " + usuario.apellido
}

func (this *User) change_name(n string){
	this.nombre = n
}

func main(){
	/*Usuario := User{nombre: "Jorge", apellido: "Niño"}*/
	user := new(User)
	
	user.nombre = "Pepe"
	user.apellido = "Goico"

	user.change_name("Marco")

	fmt.Println(user.nombre)
	
	
}