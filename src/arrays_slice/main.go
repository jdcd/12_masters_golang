package main

import "fmt"

func main() {
	tabla := [5]int{1, 2, 3, 54, 65}

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

}
