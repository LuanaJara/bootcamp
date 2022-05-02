package main

import (
	"fmt"
)

func main() {

	var fname string
	fmt.Print("name: ")
	fmt.Scan(&fname)
	fmt.Print("city: ")
	var lname string
	fmt.Scan(&lname)
	fmt.Print("data: name:  " + fname + " city: " + lname + ".")
}
