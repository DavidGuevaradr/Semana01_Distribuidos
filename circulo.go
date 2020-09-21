package main

import "fmt"

func main() {
	var radio float64
	const pi float64 = 3.1416
	fmt.Println("Valor del radio: ")
	fmt.Scan(&radio)
	area := pi * radio * radio
	fmt.Println("El valor del area es: ", area)

}
