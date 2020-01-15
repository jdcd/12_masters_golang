package main

import "fmt"

func main() {
	//los slice son vectores dinamicos, pueden ampliarse en tiempo de ejecución
	matriz := []int{1, 2, 3}
	fmt.Println(matriz)
	variante()
	variante2()
}

func variante() {
	items := [5]int{1, 2, 3, 4, 5}
	porcion := items[3:]

	fmt.Println(porcion)
}

func variante2() {
	items := make([]int, 4, 20)
	porcion := items[3:]

	fmt.Println(porcion)
}
