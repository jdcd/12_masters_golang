package main

//Interfaces definen conductos
//structuras grabar valores

type humano interface {
	repirar()
	pensar()
	comer()
	sexo() string
}

type animal interface {
	repirar()
	comer()
	EsCarnivoro() bool
}

type vegetal interface {
	ClasificacionVegetal() string
}

type hombre struct {
	edad      int
	altura    float32
	peso      float32
	repirando bool
	pensando  bool
	comiendo  bool
}

type mujer struct {
	edad      int
	altura    float32
	peso      float32
	repirando bool
	pensando  bool
	comiendo  bool
}

func (h *hombre) repirar() { h.repirando = true }
func (h *hombre) pensar()  { h.repirando = true }
func (h *hombre) comer()   { h.repirando = true }
func (h *hombre) sexo()    { h.repirando = true }

func main() {}
