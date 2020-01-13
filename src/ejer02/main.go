package main

import (
	"fmt"
	"strconv"
)

// Usado dentro de todas las funciones del packete
var texto string
var cantidad int
var disponible bool

func main() {
	// variables unicamente disponible en el scope de la funcion
	var numero int
	var cadena string
	var status bool
	numeroX := 3 //Autoasignación del tipo de variable

	// usar variable para
	fmt.Println(numeroX)

	//Puedo hacer definciones con asignación, de manera multiple
	nombre, edad, feliz := "Julian", 30, false

	// Imprimiendo variables no definidas
	fmt.Println(numero, cadena, status)

	// Impirmiendo variables definidas
	fmt.Println(nombre, edad, feliz)

	showDisponible()

	newText := strconv.Itoa(edad)

	fmt.Println("Cast text: ", newText)

}

func showDisponible() {
	fmt.Println(disponible)
}
