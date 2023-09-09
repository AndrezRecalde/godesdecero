package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero_uno int
var numero_dos int
var leyenda string
var err error

func IngresoNumeros() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese numero uno: ")
	if scanner.Scan() {
		numero_uno, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto " + err.Error())
		}
	}

	fmt.Println("Ingrese numero dos: ")
	if scanner.Scan() {
		numero_dos, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto " + err.Error())
		}
	}

	fmt.Println("Ingrese leyenda: ")
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, numero_uno*numero_dos)
}
