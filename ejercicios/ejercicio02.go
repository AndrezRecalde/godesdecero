package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

func Ejercicio02() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Ingrese un número: ")

		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 0; i <= 10; i++ {
		resultado := numero * i
		fmt.Println(numero, "x", i, "= ", resultado)
	}
}
