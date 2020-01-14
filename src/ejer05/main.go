package main

import "fmt"

func main() {

	num := 0
	for {
		fmt.Println("Continue...")
		fmt.Println("Digite el número secreto....")
		fmt.Scanln(&num)
		if num == 100 {
			fmt.Println("Número correcto, vuelva pronto !!")
			break
		}
	}

	i := 0

	for i < 10 {
		fmt.Printf("\n Valor de i: %d", i)
		if i == 5 {
			fmt.Printf(" multiplicamos por 2 \n")
			i = i * 2
			continue
		}
		fmt.Printf("	paso por aquí")
		i++
	}

}
