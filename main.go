package main

import (
	"fmt"
)

func main() {

	var fname string
	fmt.Print("Ingresar nombre: ")
	fmt.Scan(&fname)
	fmt.Print("Ingresar direccion : ")
	var lname string
	fmt.Scan(&lname)
	fmt.Print("tus datos son: nombre:  " + fname + " Direccion: " + lname + ".")
}
