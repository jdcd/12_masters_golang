package main

//mapas son estructuras clave valor
import "fmt"

type perro struct {
	raza   string
	nombre string
}

func main() {
	paises := make(map[string]string)
	paises["Mexico"] = "D.F"
	paises["Colombia"] = "Bogotà"
	fmt.Println(paises)
	fmt.Println(paises["Colombia"])

	campeonato := map[string]int{
		"Arsenal":  89,
		"Chelsea":  85,
		"Man City": 84,
	}

	campeonato["Liverpool"] = 80
	delete(campeonato, "Man City")

	for equipo, puntaje := range campeonato {
		fmt.Println(equipo, "	", puntaje)
	}

	puntaje, existe := campeonato["Chelseaaa"]
	fmt.Printf("El puntaje capturado de Chelsea es %d, existe ?  = %t \n", puntaje, existe)

	fmt.Println(campeonato)

	fmt.Println("**************************************")
	perros := make(map[string]perro)
	millos := perros["millos"]

	if (perro{}) == millos {
		fmt.Println("vacio !")
	} else {

	}

}
