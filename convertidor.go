package main

import "fmt"

func main() {
	var f float64
	fmt.Println("Temperatura en grados Farenheit: ")
	fmt.Scan(&f)
	Celcius := ((f - 32) * 5 / 9)
	fmt.Println("Temperatura en Celcius: ", Celcius)

}
