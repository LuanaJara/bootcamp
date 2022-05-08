package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
}

func (per *Persona) cargar() {
	fmt.Print("Ingrese el nombre:")
	fmt.Scan(&per.nombre)
	fmt.Print("Ingrese la edad:")
	fmt.Scan(&per.edad)
}

func (per Persona) imprimir() {
	fmt.Println("Datos Personales")
	fmt.Println("***** **********")
	fmt.Println("Nombre:", per.nombre)
	fmt.Println("Edad:", per.edad)
}

func (per Persona) mayorEdad() {
	if per.edad >= 18 {
		fmt.Println(per.nombre, "es mayor de edad")
	} else {
		fmt.Println(per.nombre, "no es mayor de edad")
	}
}

type Empleado struct {
	Persona
	sueldo int
}

func (emp *Empleado) cargar() {
	fmt.Println("Ingreso de datos del empleado")
	emp.Persona.cargar()
	fmt.Print("Ingrese el sueldo:")
	fmt.Scan(&emp.sueldo)
}

func (emp Empleado) imprimir() {
	emp.Persona.imprimir()
	fmt.Println("Sueldo:", emp.sueldo)
}

func (emp Empleado) pagaImpuestos() {
	if emp.sueldo >= 300 {
		fmt.Println(emp.nombre, "debe pagar impuestos")
	} else {
		fmt.Println(emp.nombre, "no debe pagar impuestos")
	}
}

func main() {
	var persona1 Persona
	persona1.cargar()
	persona1.imprimir()
	persona1.mayorEdad()
	var empleado1 Empleado
	empleado1.cargar()
	empleado1.imprimir()
	empleado1.pagaImpuestos()
}
