package main

import "fmt"

func main() {
	var base float64
	var altura float64
	fmt.Println("Valor de la base: ")
	fmt.Scan(&base)
	fmt.Println("Valor de la altura: ")
	fmt.Scan(&altura)
	area := (base * altura) / 2
	fmt.Println("El valor del area es: ", area)

}
