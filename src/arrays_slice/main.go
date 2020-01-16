package main

import "fmt"

func main() {
	//los slice son vectores dinamicos, pueden ampliarse en tiempo de ejecución
	matriz := []int{1, 2, 3}
	fmt.Println(matriz)
	variante()
	variante2()
	variante3()
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
	fmt.Printf("largo %d, capacidad %d \n", len(items), cap(items))

}

func variante3(){
	nums := make([]int,0,0) 
	for i := 0; i <= 99 ; i++{
		nums=append(nums,i)
	}
	fmt.Printf("Largo %d, Capacidad %d \n", len(nums), cap(nums))
}
