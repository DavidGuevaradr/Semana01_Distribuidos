package main

import "fmt"

func main() {
	var lado float64
	fmt.Println("valor del lado: ")
	fmt.Scanf("%f", &lado)
	area := lado * lado
	fmt.Println("El valor del area es: ", area)

}
