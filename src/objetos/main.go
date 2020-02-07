package main

import (
	"fmt"

	"github.com/jdcd/12_masters_golang/src/objetos/user"
)

//go se basa en las estrucutras

type pepe struct {
	user.Usuario
}

func main() {
	u := new(pepe)
	u.Id = 10
	fmt.Println(u)
}
