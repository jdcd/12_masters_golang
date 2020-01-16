package main

import (
	"fmt"
	"time"
)

func main() {
	//Los condicionales son básicos, se debe poner la llave { en la misma linea
	if 2 > 1 {
		fmt.Println("All is fine")
	} else {
		fmt.Println("The world is crazy !")
	}

	//Como novedad podemos asignar varibales en la misma liea el if, aqí un ejemplo algo redundante
	//Soporta else if
	if day := getDay(); day == "Tuesday" {
		fmt.Println("Today is Monday")
	} else {
		fmt.Println("Today is not Tuesday, today is: ", day)
	}
}

func getDay() string {
	return time.Now().Weekday().String()
}
