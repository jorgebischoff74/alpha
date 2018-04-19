package main

import(
	"fmt"
)

type Human struct{
	name string
}

type Tutor struct{
	Human
	Dummy
}

type Dummy struct{
	Human
}

func(this Human) hablar() string{

	return " Ble Ble"
}

func(this Tutor) hablar() string{
	this.Human.hablar()
	return this.Human.hablar() + " Bla Bla Bliiiii"
}

func main(){

	tutor := Tutor{Human{"Jorge"}, Dummy{Human{"Pepe"}}}

	fmt.Println(tutor.hablar())

}
