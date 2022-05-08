package main

import (
	"fmt"
)
type persona struct{
	nombre string
	edad int
}

func main() {
fmt.Println(persona {"Luana", 18})
fmt.Println(persona {"Santi", 19})
fmt.Println(persona {"Ori", 21})

s := persona{nombre: "abuela", edad: 50}
fmt.Println(s.nombre)

sp := &s
fmt.Println(sp.edad)


}
