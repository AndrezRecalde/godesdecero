package main

import (
	"fmt"
	"godesdecero/variables"
)

func main() {
	//variables.RestoVariables()
	estado, texto := variables.ConvierteaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
}
